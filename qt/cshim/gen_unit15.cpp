#include <QtCore/qobjectdefs.h>
#include <QtCore/qobject.h>
#include <QtCore/qstring.h>
#include <QtCore/qmetatype.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qnamespace.h>
#include "gen_qmetaobject.h"

extern "C" {

void* QMetaObject_new() {
	return new QMetaObject();
}
void* QMetaObject_new2(void* arg1) {
	return new QMetaObject(*(QMetaObject*)arg1);
}
const char* QMetaObject_className(void* self) {
	auto&& __r = ((QMetaObject*)self)->className();
	return __r;
}
void* QMetaObject_superClass(void* self) {
	auto&& __r = ((QMetaObject*)self)->superClass();
	return (void*)(__r);
}
int QMetaObject_inherits(void* self, void* metaObject) {
	auto&& __r = ((QMetaObject*)self)->inherits((QMetaObject*)metaObject);
	return __r;
}
void* QMetaObject_cast(void* self, void* obj) {
	auto&& __r = ((QMetaObject*)self)->cast((QObject*)obj);
	return (void*)(__r);
}
void* QMetaObject_cast2(void* self, void* obj) {
	auto&& __r = ((QMetaObject*)self)->cast((QObject*)obj);
	return (void*)(__r);
}
char* QMetaObject_tr(void* self, const char* s, const char* c, long long* __out_len) {
	auto&& __r = ((QMetaObject*)self)->tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QMetaObject_metaType(void* self) {
	auto&& __r = ((QMetaObject*)self)->metaType();
	return new QMetaType(__r);
}
int QMetaObject_methodOffset(void* self) {
	auto&& __r = ((QMetaObject*)self)->methodOffset();
	return __r;
}
int QMetaObject_enumeratorOffset(void* self) {
	auto&& __r = ((QMetaObject*)self)->enumeratorOffset();
	return __r;
}
int QMetaObject_propertyOffset(void* self) {
	auto&& __r = ((QMetaObject*)self)->propertyOffset();
	return __r;
}
int QMetaObject_classInfoOffset(void* self) {
	auto&& __r = ((QMetaObject*)self)->classInfoOffset();
	return __r;
}
int QMetaObject_constructorCount(void* self) {
	auto&& __r = ((QMetaObject*)self)->constructorCount();
	return __r;
}
int QMetaObject_methodCount(void* self) {
	auto&& __r = ((QMetaObject*)self)->methodCount();
	return __r;
}
int QMetaObject_enumeratorCount(void* self) {
	auto&& __r = ((QMetaObject*)self)->enumeratorCount();
	return __r;
}
int QMetaObject_propertyCount(void* self) {
	auto&& __r = ((QMetaObject*)self)->propertyCount();
	return __r;
}
int QMetaObject_classInfoCount(void* self) {
	auto&& __r = ((QMetaObject*)self)->classInfoCount();
	return __r;
}
int QMetaObject_indexOfConstructor(void* self, const char* constructor) {
	auto&& __r = ((QMetaObject*)self)->indexOfConstructor(constructor);
	return __r;
}
int QMetaObject_indexOfMethod(void* self, const char* method) {
	auto&& __r = ((QMetaObject*)self)->indexOfMethod(method);
	return __r;
}
int QMetaObject_indexOfSignal(void* self, const char* signal) {
	auto&& __r = ((QMetaObject*)self)->indexOfSignal(signal);
	return __r;
}
int QMetaObject_indexOfSlot(void* self, const char* slot) {
	auto&& __r = ((QMetaObject*)self)->indexOfSlot(slot);
	return __r;
}
int QMetaObject_indexOfEnumerator(void* self, const char* name) {
	auto&& __r = ((QMetaObject*)self)->indexOfEnumerator(name);
	return __r;
}
int QMetaObject_indexOfProperty(void* self, const char* name) {
	auto&& __r = ((QMetaObject*)self)->indexOfProperty(name);
	return __r;
}
int QMetaObject_indexOfClassInfo(void* self, const char* name) {
	auto&& __r = ((QMetaObject*)self)->indexOfClassInfo(name);
	return __r;
}
void* QMetaObject_constructor(void* self, int index) {
	auto&& __r = ((QMetaObject*)self)->constructor(index);
	return new QMetaMethod(__r);
}
void* QMetaObject_method(void* self, int index) {
	auto&& __r = ((QMetaObject*)self)->method(index);
	return new QMetaMethod(__r);
}
void* QMetaObject_enumerator(void* self, int index) {
	auto&& __r = ((QMetaObject*)self)->enumerator(index);
	return new QMetaEnum(__r);
}
void* QMetaObject_property(void* self, int index) {
	auto&& __r = ((QMetaObject*)self)->property(index);
	return new QMetaProperty(__r);
}
void* QMetaObject_classInfo(void* self, int index) {
	auto&& __r = ((QMetaObject*)self)->classInfo(index);
	return new QMetaClassInfo(__r);
}
void* QMetaObject_userProperty(void* self) {
	auto&& __r = ((QMetaObject*)self)->userProperty();
	return new QMetaProperty(__r);
}
int QMetaObject_checkConnectArgs(const char* signal, const char* method) {
	auto&& __r = QMetaObject::checkConnectArgs(signal, method);
	return __r;
}
int QMetaObject_checkConnectArgs2(void* signal, void* method) {
	auto&& __r = QMetaObject::checkConnectArgs(*(QMetaMethod*)signal, *(QMetaMethod*)method);
	return __r;
}
void* QMetaObject_normalizedSignature(const char* method) {
	auto&& __r = QMetaObject::normalizedSignature(method);
	return new QByteArray(__r);
}
void* QMetaObject_normalizedType(const char* type_) {
	auto&& __r = QMetaObject::normalizedType(type_);
	return new QByteArray(__r);
}
void* QMetaObject_connect(void* sender, int signal_index, void* receiver, int method_index) {
	auto&& __r = QMetaObject::connect((QObject*)sender, signal_index, (QObject*)receiver, method_index);
	return new QMetaObject::Connection(__r);
}
int QMetaObject_disconnectOne(void* sender, int signal_index, void* receiver, int method_index) {
	auto&& __r = QMetaObject::disconnectOne((QObject*)sender, signal_index, (QObject*)receiver, method_index);
	return __r;
}
void QMetaObject_connectSlotsByName(void* o) {
	QMetaObject::connectSlotsByName((QObject*)o);
}
int QMetaObject_invokeMethod(void* obj, const char* member, int arg3, void* ret) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(arg3), *(QGenericReturnArgument*)ret);
	return __r;
}
int QMetaObject_invokeMethod2(void* obj, const char* member, void* ret) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericReturnArgument*)ret);
	return __r;
}
int QMetaObject_invokeMethod3(void* obj, const char* member, int type_, void* val0) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(type_), *(QGenericArgument*)val0);
	return __r;
}
int QMetaObject_invokeMethod4(void* obj, const char* member, void* val0) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericArgument*)val0);
	return __r;
}
void* QMetaObject_newInstance(void* self, void* val0) {
	auto&& __r = ((QMetaObject*)self)->newInstance(*(QGenericArgument*)val0);
	return (void*)(__r);
}
void* QMetaObject_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QMetaObject*)self)->operator=(*(QMetaObject*)arg1);
	return (void*)&(__r);
}
char* QMetaObject_tr2(void* self, const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = ((QMetaObject*)self)->tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QMetaObject_connect2(void* sender, int signal_index, void* receiver, int method_index, int type_) {
	auto&& __r = QMetaObject::connect((QObject*)sender, signal_index, (QObject*)receiver, method_index, type_);
	return new QMetaObject::Connection(__r);
}
void* QMetaObject_connect3(void* sender, int signal_index, void* receiver, int method_index, int type_, int* types) {
	int types_local{};
	auto&& __r = QMetaObject::connect((QObject*)sender, signal_index, (QObject*)receiver, method_index, type_, &types_local);
	*types = types_local;
	return new QMetaObject::Connection(__r);
}
int QMetaObject_invokeMethod5(void* obj, const char* member, int arg3, void* ret, void* val0) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(arg3), *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0);
	return __r;
}
int QMetaObject_invokeMethod6(void* obj, const char* member, int arg3, void* ret, void* val0, void* val1) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(arg3), *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0, *(QGenericArgument*)val1);
	return __r;
}
int QMetaObject_invokeMethod7(void* obj, const char* member, int arg3, void* ret, void* val0, void* val1, void* val2) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(arg3), *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2);
	return __r;
}
int QMetaObject_invokeMethod8(void* obj, const char* member, int arg3, void* ret, void* val0, void* val1, void* val2, void* val3) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(arg3), *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3);
	return __r;
}
int QMetaObject_invokeMethod9(void* obj, const char* member, int arg3, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(arg3), *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4);
	return __r;
}
int QMetaObject_invokeMethod10(void* obj, const char* member, int arg3, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(arg3), *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5);
	return __r;
}
int QMetaObject_invokeMethod11(void* obj, const char* member, int arg3, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(arg3), *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6);
	return __r;
}
int QMetaObject_invokeMethod12(void* obj, const char* member, int arg3, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(arg3), *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7);
	return __r;
}
int QMetaObject_invokeMethod13(void* obj, const char* member, int arg3, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(arg3), *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7, *(QGenericArgument*)val8);
	return __r;
}
int QMetaObject_invokeMethod14(void* obj, const char* member, int arg3, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(arg3), *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7, *(QGenericArgument*)val8, *(QGenericArgument*)val9);
	return __r;
}
int QMetaObject_invokeMethod15(void* obj, const char* member, void* ret, void* val0) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0);
	return __r;
}
int QMetaObject_invokeMethod16(void* obj, const char* member, void* ret, void* val0, void* val1) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0, *(QGenericArgument*)val1);
	return __r;
}
int QMetaObject_invokeMethod17(void* obj, const char* member, void* ret, void* val0, void* val1, void* val2) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2);
	return __r;
}
int QMetaObject_invokeMethod18(void* obj, const char* member, void* ret, void* val0, void* val1, void* val2, void* val3) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3);
	return __r;
}
int QMetaObject_invokeMethod19(void* obj, const char* member, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4);
	return __r;
}
int QMetaObject_invokeMethod20(void* obj, const char* member, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5);
	return __r;
}
int QMetaObject_invokeMethod21(void* obj, const char* member, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6);
	return __r;
}
int QMetaObject_invokeMethod22(void* obj, const char* member, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7);
	return __r;
}
int QMetaObject_invokeMethod23(void* obj, const char* member, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7, *(QGenericArgument*)val8);
	return __r;
}
int QMetaObject_invokeMethod24(void* obj, const char* member, void* ret, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericReturnArgument*)ret, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7, *(QGenericArgument*)val8, *(QGenericArgument*)val9);
	return __r;
}
int QMetaObject_invokeMethod25(void* obj, const char* member, int type_, void* val0, void* val1) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(type_), *(QGenericArgument*)val0, *(QGenericArgument*)val1);
	return __r;
}
int QMetaObject_invokeMethod26(void* obj, const char* member, int type_, void* val0, void* val1, void* val2) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(type_), *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2);
	return __r;
}
int QMetaObject_invokeMethod27(void* obj, const char* member, int type_, void* val0, void* val1, void* val2, void* val3) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(type_), *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3);
	return __r;
}
int QMetaObject_invokeMethod28(void* obj, const char* member, int type_, void* val0, void* val1, void* val2, void* val3, void* val4) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(type_), *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4);
	return __r;
}
int QMetaObject_invokeMethod29(void* obj, const char* member, int type_, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(type_), *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5);
	return __r;
}
int QMetaObject_invokeMethod30(void* obj, const char* member, int type_, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(type_), *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6);
	return __r;
}
int QMetaObject_invokeMethod31(void* obj, const char* member, int type_, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(type_), *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7);
	return __r;
}
int QMetaObject_invokeMethod32(void* obj, const char* member, int type_, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(type_), *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7, *(QGenericArgument*)val8);
	return __r;
}
int QMetaObject_invokeMethod33(void* obj, const char* member, int type_, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, (Qt::ConnectionType)(type_), *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7, *(QGenericArgument*)val8, *(QGenericArgument*)val9);
	return __r;
}
int QMetaObject_invokeMethod34(void* obj, const char* member, void* val0, void* val1) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericArgument*)val0, *(QGenericArgument*)val1);
	return __r;
}
int QMetaObject_invokeMethod35(void* obj, const char* member, void* val0, void* val1, void* val2) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2);
	return __r;
}
int QMetaObject_invokeMethod36(void* obj, const char* member, void* val0, void* val1, void* val2, void* val3) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3);
	return __r;
}
int QMetaObject_invokeMethod37(void* obj, const char* member, void* val0, void* val1, void* val2, void* val3, void* val4) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4);
	return __r;
}
int QMetaObject_invokeMethod38(void* obj, const char* member, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5);
	return __r;
}
int QMetaObject_invokeMethod39(void* obj, const char* member, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6);
	return __r;
}
int QMetaObject_invokeMethod40(void* obj, const char* member, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7);
	return __r;
}
int QMetaObject_invokeMethod41(void* obj, const char* member, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7, *(QGenericArgument*)val8);
	return __r;
}
int QMetaObject_invokeMethod42(void* obj, const char* member, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9) {
	auto&& __r = QMetaObject::invokeMethod((QObject*)obj, member, *(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7, *(QGenericArgument*)val8, *(QGenericArgument*)val9);
	return __r;
}
void* QMetaObject_newInstance2(void* self, void* val0, void* val1) {
	auto&& __r = ((QMetaObject*)self)->newInstance(*(QGenericArgument*)val0, *(QGenericArgument*)val1);
	return (void*)(__r);
}
void* QMetaObject_newInstance3(void* self, void* val0, void* val1, void* val2) {
	auto&& __r = ((QMetaObject*)self)->newInstance(*(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2);
	return (void*)(__r);
}
void* QMetaObject_newInstance4(void* self, void* val0, void* val1, void* val2, void* val3) {
	auto&& __r = ((QMetaObject*)self)->newInstance(*(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3);
	return (void*)(__r);
}
void* QMetaObject_newInstance5(void* self, void* val0, void* val1, void* val2, void* val3, void* val4) {
	auto&& __r = ((QMetaObject*)self)->newInstance(*(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4);
	return (void*)(__r);
}
void* QMetaObject_newInstance6(void* self, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5) {
	auto&& __r = ((QMetaObject*)self)->newInstance(*(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5);
	return (void*)(__r);
}
void* QMetaObject_newInstance7(void* self, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6) {
	auto&& __r = ((QMetaObject*)self)->newInstance(*(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6);
	return (void*)(__r);
}
void* QMetaObject_newInstance8(void* self, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7) {
	auto&& __r = ((QMetaObject*)self)->newInstance(*(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7);
	return (void*)(__r);
}
void* QMetaObject_newInstance9(void* self, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8) {
	auto&& __r = ((QMetaObject*)self)->newInstance(*(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7, *(QGenericArgument*)val8);
	return (void*)(__r);
}
void* QMetaObject_newInstance10(void* self, void* val0, void* val1, void* val2, void* val3, void* val4, void* val5, void* val6, void* val7, void* val8, void* val9) {
	auto&& __r = ((QMetaObject*)self)->newInstance(*(QGenericArgument*)val0, *(QGenericArgument*)val1, *(QGenericArgument*)val2, *(QGenericArgument*)val3, *(QGenericArgument*)val4, *(QGenericArgument*)val5, *(QGenericArgument*)val6, *(QGenericArgument*)val7, *(QGenericArgument*)val8, *(QGenericArgument*)val9);
	return (void*)(__r);
}
void* QMetaObject_d_get(void* self) {
	auto __r = ((QMetaObject*)self)->d;
	return new QMetaObject::Data(__r);
}
void QMetaObject_d_set(void* self, void* d) {
	((QMetaObject*)self)->d = *(QMetaObject::Data*)d;
}
void QMetaObject_delete(void* self) { delete (QMetaObject*)self; }

}

#include <QtCore/qfloat16.h>
#include <QtCore/qnamespace.h>
#include "gen_qfloat16.h"

extern "C" {

void* qfloat16_new() {
	return new qfloat16();
}
void* qfloat16_new2(int arg1) {
	return new qfloat16((Qt::Initialization)(arg1));
}
void* qfloat16_new3(float f) {
	return new qfloat16(f);
}
void* qfloat16_new4(void* arg1) {
	return new qfloat16(*(qfloat16*)arg1);
}
float qfloat16_ToFloat(void* self) {
	auto&& __r = static_cast<float>(*(qfloat16*)self);
	return __r;
}
int qfloat16_isInf(void* self) {
	auto&& __r = ((qfloat16*)self)->isInf();
	return __r;
}
int qfloat16_isNaN(void* self) {
	auto&& __r = ((qfloat16*)self)->isNaN();
	return __r;
}
int qfloat16_isFinite(void* self) {
	auto&& __r = ((qfloat16*)self)->isFinite();
	return __r;
}
int qfloat16_fpClassify(void* self) {
	auto&& __r = ((qfloat16*)self)->fpClassify();
	return __r;
}
void* qfloat16_copySign(void* self, void* sign) {
	auto&& __r = ((qfloat16*)self)->copySign(*(qfloat16*)sign);
	return new qfloat16(__r);
}
void* qfloat16__limit_epsilon() {
	auto&& __r = qfloat16::_limit_epsilon();
	return new qfloat16(__r);
}
void* qfloat16__limit_min() {
	auto&& __r = qfloat16::_limit_min();
	return new qfloat16(__r);
}
void* qfloat16__limit_denorm_min() {
	auto&& __r = qfloat16::_limit_denorm_min();
	return new qfloat16(__r);
}
void* qfloat16__limit_max() {
	auto&& __r = qfloat16::_limit_max();
	return new qfloat16(__r);
}
void* qfloat16__limit_lowest() {
	auto&& __r = qfloat16::_limit_lowest();
	return new qfloat16(__r);
}
void* qfloat16__limit_infinity() {
	auto&& __r = qfloat16::_limit_infinity();
	return new qfloat16(__r);
}
void* qfloat16__limit_quiet_NaN() {
	auto&& __r = qfloat16::_limit_quiet_NaN();
	return new qfloat16(__r);
}
void* qfloat16__limit_signaling_NaN() {
	auto&& __r = qfloat16::_limit_signaling_NaN();
	return new qfloat16(__r);
}
int qfloat16_isNormal(void* self) {
	auto&& __r = ((qfloat16*)self)->isNormal();
	return __r;
}
void* qfloat16_OperatorMinus(void* self) {
	auto&& __r = -(*(qfloat16*)self);
	return new qfloat16(__r);
}
void* qfloat16_OperatorPlus(void* self, void* b) {
	auto&& __r = (*(qfloat16*)self) + (*(qfloat16*)b);
	return new qfloat16(__r);
}
void* qfloat16_OperatorMinus2(void* self, void* b) {
	auto&& __r = (*(qfloat16*)self) - (*(qfloat16*)b);
	return new qfloat16(__r);
}
void* qfloat16_OperatorMultiply(void* self, void* b) {
	auto&& __r = (*(qfloat16*)self) * (*(qfloat16*)b);
	return new qfloat16(__r);
}
void* qfloat16_OperatorDivide(void* self, void* b) {
	auto&& __r = (*(qfloat16*)self) / (*(qfloat16*)b);
	return new qfloat16(__r);
}
double qfloat16_OperatorPlus2(void* self, double rhs) {
	auto&& __r = (*(qfloat16*)self) + (rhs);
	return __r;
}
double qfloat16_OperatorMinus3(void* self, double rhs) {
	auto&& __r = (*(qfloat16*)self) - (rhs);
	return __r;
}
double qfloat16_OperatorMultiply2(void* self, double rhs) {
	auto&& __r = (*(qfloat16*)self) * (rhs);
	return __r;
}
double qfloat16_OperatorDivide2(void* self, double rhs) {
	auto&& __r = (*(qfloat16*)self) / (rhs);
	return __r;
}
void* qfloat16_OperatorPlusAssign(void* self, double rhs) {
	auto&& __r = (*(qfloat16*)self) += (rhs);
	return (void*)&(__r);
}
void* qfloat16_OperatorMinusAssign(void* self, double rhs) {
	auto&& __r = (*(qfloat16*)self) -= (rhs);
	return (void*)&(__r);
}
void* qfloat16_OperatorMultiplyAssign(void* self, double rhs) {
	auto&& __r = (*(qfloat16*)self) *= (rhs);
	return (void*)&(__r);
}
void* qfloat16_OperatorDivideAssign(void* self, double rhs) {
	auto&& __r = (*(qfloat16*)self) /= (rhs);
	return (void*)&(__r);
}
float qfloat16_OperatorPlus3(void* self, float rhs) {
	auto&& __r = (*(qfloat16*)self) + (rhs);
	return __r;
}
float qfloat16_OperatorMinus4(void* self, float rhs) {
	auto&& __r = (*(qfloat16*)self) - (rhs);
	return __r;
}
float qfloat16_OperatorMultiply3(void* self, float rhs) {
	auto&& __r = (*(qfloat16*)self) * (rhs);
	return __r;
}
float qfloat16_OperatorDivide3(void* self, float rhs) {
	auto&& __r = (*(qfloat16*)self) / (rhs);
	return __r;
}
void* qfloat16_OperatorPlusAssign2(void* self, float rhs) {
	auto&& __r = (*(qfloat16*)self) += (rhs);
	return (void*)&(__r);
}
void* qfloat16_OperatorMinusAssign2(void* self, float rhs) {
	auto&& __r = (*(qfloat16*)self) -= (rhs);
	return (void*)&(__r);
}
void* qfloat16_OperatorMultiplyAssign2(void* self, float rhs) {
	auto&& __r = (*(qfloat16*)self) *= (rhs);
	return (void*)&(__r);
}
void* qfloat16_OperatorDivideAssign2(void* self, float rhs) {
	auto&& __r = (*(qfloat16*)self) /= (rhs);
	return (void*)&(__r);
}
double qfloat16_OperatorPlus4(void* self, int rhs) {
	auto&& __r = (*(qfloat16*)self) + (rhs);
	return __r;
}
double qfloat16_OperatorMinus5(void* self, int rhs) {
	auto&& __r = (*(qfloat16*)self) - (rhs);
	return __r;
}
double qfloat16_OperatorMultiply4(void* self, int rhs) {
	auto&& __r = (*(qfloat16*)self) * (rhs);
	return __r;
}
double qfloat16_OperatorDivide4(void* self, int rhs) {
	auto&& __r = (*(qfloat16*)self) / (rhs);
	return __r;
}
int qfloat16_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(qfloat16*)self) == (*(qfloat16*)rhs);
	return __r;
}
int qfloat16_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(qfloat16*)self) != (*(qfloat16*)rhs);
	return __r;
}
int qfloat16_OperatorLesser(void* self, void* rhs) {
	auto&& __r = (*(qfloat16*)self) < (*(qfloat16*)rhs);
	return __r;
}
int qfloat16_OperatorGreater(void* self, void* rhs) {
	auto&& __r = (*(qfloat16*)self) > (*(qfloat16*)rhs);
	return __r;
}
int qfloat16_OperatorLesserOrEqual(void* self, void* rhs) {
	auto&& __r = (*(qfloat16*)self) <= (*(qfloat16*)rhs);
	return __r;
}
int qfloat16_OperatorGreaterOrEqual(void* self, void* rhs) {
	auto&& __r = (*(qfloat16*)self) >= (*(qfloat16*)rhs);
	return __r;
}
void qfloat16_delete(void* self) { delete (qfloat16*)self; }

}

#include <QtCore/qshareddata.h>
#include "gen_qadoptshareddatatag.h"

extern "C" {

void* QAdoptSharedDataTag_new() {
	return new QAdoptSharedDataTag();
}
void QAdoptSharedDataTag_delete(void* self) { delete (QAdoptSharedDataTag*)self; }

}

#include <QtCore/qcoreevent.h>
#include <QtCore/qbytearray.h>
#include "gen_qdynamicpropertychangeevent.h"

extern "C" {

void* QDynamicPropertyChangeEvent_base(void* self) { return (void*)static_cast<QEvent*>((QDynamicPropertyChangeEvent*)self); }
void* QDynamicPropertyChangeEvent_new(void* name) {
	return new QDynamicPropertyChangeEvent(*(QByteArray*)name);
}
void* QDynamicPropertyChangeEvent_propertyName(void* self) {
	auto&& __r = ((QDynamicPropertyChangeEvent*)self)->propertyName();
	return new QByteArray(__r);
}
int QDynamicPropertyChangeEvent_registerEventType() {
	auto&& __r = QDynamicPropertyChangeEvent::registerEventType();
	return __r;
}
int QDynamicPropertyChangeEvent_registerEventType2(int hint) {
	auto&& __r = QDynamicPropertyChangeEvent::registerEventType(hint);
	return __r;
}
void QDynamicPropertyChangeEvent_delete(void* self) { delete (QDynamicPropertyChangeEvent*)self; }

}

#include <QtCore/quuid.h>
#include <QtCore/qbytearrayview.h>
#include "gen_quuid__id128bytes.h"

extern "C" {

void* QUuid__Id128Bytes_new() {
	return new QUuid::Id128Bytes();
}
void* QUuid__Id128Bytes_new2(void* arg1) {
	return new QUuid::Id128Bytes(*(QUuid::Id128Bytes*)arg1);
}
void* QUuid__Id128Bytes_ToQByteArrayView(void* self) {
	auto&& __r = static_cast<QByteArrayView>(*(QUuid::Id128Bytes*)self);
	return new QByteArrayView(__r);
}
void QUuid__Id128Bytes_delete(void* self) { delete (QUuid::Id128Bytes*)self; }

}

#include <QtCore/qcommandlineparser.h>
#include <QtCore/qstring.h>
#include <QtCore/qcommandlineoption.h>
#include <QtCore/qcoreapplication.h>
#include "gen_qcommandlineparser.h"

extern "C" {

void* QCommandLineParser_new() {
	return new QCommandLineParser();
}
char* QCommandLineParser_tr(const char* sourceText, long long* __out_len) {
	auto&& __r = QCommandLineParser::tr(sourceText);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCommandLineParser_setSingleDashWordOptionMode(void* self, int parsingMode) {
	((QCommandLineParser*)self)->setSingleDashWordOptionMode((QCommandLineParser::SingleDashWordOptionMode)(parsingMode));
}
void QCommandLineParser_setOptionsAfterPositionalArgumentsMode(void* self, int mode) {
	((QCommandLineParser*)self)->setOptionsAfterPositionalArgumentsMode((QCommandLineParser::OptionsAfterPositionalArgumentsMode)(mode));
}
int QCommandLineParser_addOption(void* self, void* commandLineOption) {
	auto&& __r = ((QCommandLineParser*)self)->addOption(*(QCommandLineOption*)commandLineOption);
	return __r;
}
int QCommandLineParser_addOptions(void* self, void** options, long long options_len) {
	QList<QCommandLineOption> options_list; options_list.reserve(options_len);
	for (long long __i = 0; __i < options_len; __i++) options_list.append(*(QCommandLineOption*)options[__i]);
	auto&& __r = ((QCommandLineParser*)self)->addOptions(options_list);
	return __r;
}
void* QCommandLineParser_addVersionOption(void* self) {
	auto&& __r = ((QCommandLineParser*)self)->addVersionOption();
	return new QCommandLineOption(__r);
}
void* QCommandLineParser_addHelpOption(void* self) {
	auto&& __r = ((QCommandLineParser*)self)->addHelpOption();
	return new QCommandLineOption(__r);
}
void QCommandLineParser_setApplicationDescription(void* self, const char* description, long long description_len) {
	QString description_qstr = QString::fromUtf8(description, description_len);
	((QCommandLineParser*)self)->setApplicationDescription(description_qstr);
}
char* QCommandLineParser_applicationDescription(void* self, long long* __out_len) {
	auto&& __r = ((QCommandLineParser*)self)->applicationDescription();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCommandLineParser_addPositionalArgument(void* self, const char* name, long long name_len, const char* description, long long description_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString description_qstr = QString::fromUtf8(description, description_len);
	((QCommandLineParser*)self)->addPositionalArgument(name_qstr, description_qstr);
}
void QCommandLineParser_clearPositionalArguments(void* self) {
	((QCommandLineParser*)self)->clearPositionalArguments();
}
void QCommandLineParser_process(void* self, char** arguments, long long arguments_len) {
	QList<QString> arguments_list; arguments_list.reserve(arguments_len);
	for (long long __i = 0; __i < arguments_len; __i++) arguments_list.append(QString::fromUtf8(arguments[__i]));
	((QCommandLineParser*)self)->process(arguments_list);
}
void QCommandLineParser_process2(void* self, void* app) {
	((QCommandLineParser*)self)->process(*(QCoreApplication*)app);
}
int QCommandLineParser_parse(void* self, char** arguments, long long arguments_len) {
	QList<QString> arguments_list; arguments_list.reserve(arguments_len);
	for (long long __i = 0; __i < arguments_len; __i++) arguments_list.append(QString::fromUtf8(arguments[__i]));
	auto&& __r = ((QCommandLineParser*)self)->parse(arguments_list);
	return __r;
}
char* QCommandLineParser_errorText(void* self, long long* __out_len) {
	auto&& __r = ((QCommandLineParser*)self)->errorText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QCommandLineParser_isSet(void* self, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = ((QCommandLineParser*)self)->isSet(name_qstr);
	return __r;
}
char* QCommandLineParser_value(void* self, const char* name, long long name_len, long long* __out_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = ((QCommandLineParser*)self)->value(name_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QCommandLineParser_values(void* self, const char* name, long long name_len, long long* __out_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = ((QCommandLineParser*)self)->values(name_qstr);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int QCommandLineParser_isSet2(void* self, void* option) {
	auto&& __r = ((QCommandLineParser*)self)->isSet(*(QCommandLineOption*)option);
	return __r;
}
char* QCommandLineParser_value2(void* self, void* option, long long* __out_len) {
	auto&& __r = ((QCommandLineParser*)self)->value(*(QCommandLineOption*)option);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QCommandLineParser_values2(void* self, void* option, long long* __out_len) {
	auto&& __r = ((QCommandLineParser*)self)->values(*(QCommandLineOption*)option);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QCommandLineParser_positionalArguments(void* self, long long* __out_len) {
	auto&& __r = ((QCommandLineParser*)self)->positionalArguments();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QCommandLineParser_optionNames(void* self, long long* __out_len) {
	auto&& __r = ((QCommandLineParser*)self)->optionNames();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char** QCommandLineParser_unknownOptionNames(void* self, long long* __out_len) {
	auto&& __r = ((QCommandLineParser*)self)->unknownOptionNames();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QCommandLineParser_showVersion(void* self) {
	((QCommandLineParser*)self)->showVersion();
}
void QCommandLineParser_showHelp(void* self) {
	((QCommandLineParser*)self)->showHelp();
}
char* QCommandLineParser_helpText(void* self, long long* __out_len) {
	auto&& __r = ((QCommandLineParser*)self)->helpText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCommandLineParser_showMessageAndExit(int type_, const char* message, long long message_len) {
	QString message_qstr = QString::fromUtf8(message, message_len);
	QCommandLineParser::showMessageAndExit((QCommandLineParser::MessageType)(type_), message_qstr);
}
char* QCommandLineParser_tr2(const char* sourceText, const char* disambiguation, long long* __out_len) {
	auto&& __r = QCommandLineParser::tr(sourceText, disambiguation);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCommandLineParser_tr3(const char* sourceText, const char* disambiguation, int n, long long* __out_len) {
	auto&& __r = QCommandLineParser::tr(sourceText, disambiguation, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCommandLineParser_addPositionalArgument2(void* self, const char* name, long long name_len, const char* description, long long description_len, const char* syntax, long long syntax_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	QString description_qstr = QString::fromUtf8(description, description_len);
	QString syntax_qstr = QString::fromUtf8(syntax, syntax_len);
	((QCommandLineParser*)self)->addPositionalArgument(name_qstr, description_qstr, syntax_qstr);
}
void QCommandLineParser_showHelp2(void* self, int exitCode) {
	((QCommandLineParser*)self)->showHelp(exitCode);
}
void QCommandLineParser_showMessageAndExit2(int type_, const char* message, long long message_len, int exitCode) {
	QString message_qstr = QString::fromUtf8(message, message_len);
	QCommandLineParser::showMessageAndExit((QCommandLineParser::MessageType)(type_), message_qstr, exitCode);
}
void QCommandLineParser_delete(void* self) { delete (QCommandLineParser*)self; }

}

#include <QtCore/qthreadpool.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qrunnable.h>
#include <QtCore/qthread.h>
#include <QtCore/qdeadlinetimer.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qthreadpool.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QThreadPool__Go : public QThreadPool {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QThreadPool__Go(long long id, A&&... a) : QThreadPool(std::forward<A>(a)...), __goId(id) {}
	using QThreadPool::sender;
	using QThreadPool::senderSignalIndex;
	using QThreadPool::receivers;
	using QThreadPool::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QThreadPool::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QThreadPool::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QThreadPool::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QThreadPool::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QThreadPool::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QThreadPool::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QThreadPool::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QThreadPool::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QThreadPool::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QThreadPool::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QThreadPool::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QThreadPool::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QThreadPool::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QThreadPool;
static std::mutex __connsMu_QThreadPool;

extern "C" {

void* QThreadPool_base(void* self) { return (void*)static_cast<QObject*>((QThreadPool*)self); }
void QThreadPool__set_override(void* self, int idx) { ((QThreadPool__Go*)self)->__goOn[idx] = 1; }
void* QThreadPool_fromQObject(void* o) { return (void*)qobject_cast<QThreadPool*>((QObject*)o); }
void* QThreadPool_new(long long __goid) {
	return new QThreadPool__Go(__goid);
}
void* QThreadPool_new2(long long __goid, void* parent) {
	return new QThreadPool__Go(__goid, (QObject*)parent);
}
char* QThreadPool_tr(const char* s, long long* __out_len) {
	auto&& __r = QThreadPool::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QThreadPool_globalInstance() {
	auto&& __r = QThreadPool::globalInstance();
	return (void*)(__r);
}
void QThreadPool_start(void* self, void* runnable) {
	((QThreadPool*)self)->start((QRunnable*)runnable);
}
int QThreadPool_tryStart(void* self, void* runnable) {
	auto&& __r = ((QThreadPool*)self)->tryStart((QRunnable*)runnable);
	return __r;
}
void QThreadPool_startOnReservedThread(void* self, void* runnable) {
	((QThreadPool*)self)->startOnReservedThread((QRunnable*)runnable);
}
int QThreadPool_expiryTimeout(void* self) {
	auto&& __r = ((QThreadPool*)self)->expiryTimeout();
	return __r;
}
void QThreadPool_setExpiryTimeout(void* self, int expiryTimeout) {
	((QThreadPool*)self)->setExpiryTimeout(expiryTimeout);
}
int QThreadPool_maxThreadCount(void* self) {
	auto&& __r = ((QThreadPool*)self)->maxThreadCount();
	return __r;
}
void QThreadPool_setMaxThreadCount(void* self, int maxThreadCount) {
	((QThreadPool*)self)->setMaxThreadCount(maxThreadCount);
}
int QThreadPool_activeThreadCount(void* self) {
	auto&& __r = ((QThreadPool*)self)->activeThreadCount();
	return __r;
}
void QThreadPool_setStackSize(void* self, unsigned int stackSize) {
	((QThreadPool*)self)->setStackSize(stackSize);
}
unsigned int QThreadPool_stackSize(void* self) {
	auto&& __r = ((QThreadPool*)self)->stackSize();
	return __r;
}
void QThreadPool_setThreadPriority(void* self, int priority) {
	((QThreadPool*)self)->setThreadPriority((QThread::Priority)(priority));
}
int QThreadPool_threadPriority(void* self) {
	auto&& __r = ((QThreadPool*)self)->threadPriority();
	return (int)(__r);
}
void QThreadPool_reserveThread(void* self) {
	((QThreadPool*)self)->reserveThread();
}
void QThreadPool_releaseThread(void* self) {
	((QThreadPool*)self)->releaseThread();
}
void QThreadPool_setServiceLevel(void* self, int serviceLevel) {
	((QThreadPool*)self)->setServiceLevel((QThread::QualityOfService)(serviceLevel));
}
int QThreadPool_serviceLevel(void* self) {
	auto&& __r = ((QThreadPool*)self)->serviceLevel();
	return (int)(__r);
}
int QThreadPool_waitForDone(void* self, int msecs) {
	auto&& __r = ((QThreadPool*)self)->waitForDone(msecs);
	return __r;
}
int QThreadPool_waitForDone2(void* self) {
	auto&& __r = ((QThreadPool*)self)->waitForDone();
	return __r;
}
void QThreadPool_clear(void* self) {
	((QThreadPool*)self)->clear();
}
int QThreadPool_contains(void* self, void* thread) {
	auto&& __r = ((QThreadPool*)self)->contains((QThread*)thread);
	return __r;
}
int QThreadPool_tryTake(void* self, void* runnable) {
	auto&& __r = ((QThreadPool*)self)->tryTake((QRunnable*)runnable);
	return __r;
}
void* QThreadPool_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QThreadPool::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QThreadPool_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QThreadPool::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QThreadPool_sender(void* self) {
	auto&& __r = ((QThreadPool__Go*)self)->sender();
	return (void*)(__r);
}
int QThreadPool_senderSignalIndex(void* self) {
	auto&& __r = ((QThreadPool__Go*)self)->senderSignalIndex();
	return __r;
}
int QThreadPool_receivers(void* self, const char* signal) {
	auto&& __r = ((QThreadPool__Go*)self)->receivers(signal);
	return __r;
}
int QThreadPool_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QThreadPool__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QThreadPool_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QThreadPool::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QThreadPool_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QThreadPool::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QThreadPool_start2(void* self, void* runnable, int priority) {
	((QThreadPool*)self)->start((QRunnable*)runnable, priority);
}
int QThreadPool_waitForDone3(void* self, void* deadline) {
	auto&& __r = ((QThreadPool*)self)->waitForDone(*(QDeadlineTimer*)deadline);
	return __r;
}
void* QThreadPool_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QThreadPool::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QThreadPool_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QThreadPool::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QThreadPool_metaObject_default(void* self) {
	auto&& __r = ((QThreadPool*)self)->QThreadPool::metaObject();
	return (void*)(__r);
}
int QThreadPool_event_default(void* self, void* event) {
	auto&& __r = ((QThreadPool*)self)->QThreadPool::event((QEvent*)event);
	return __r;
}
int QThreadPool_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QThreadPool*)self)->QThreadPool::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QThreadPool_timerEvent_default(void* self, void* event) {
	((QThreadPool__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QThreadPool_childEvent_default(void* self, void* event) {
	((QThreadPool__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QThreadPool_customEvent_default(void* self, void* event) {
	((QThreadPool__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QThreadPool_connectNotify_default(void* self, void* signal) {
	((QThreadPool__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QThreadPool_disconnectNotify_default(void* self, void* signal) {
	((QThreadPool__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QThreadPool_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QThreadPool*)self, &QThreadPool::destroyed, (QThreadPool*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QThreadPool);
	__conns_QThreadPool[__id] = __c;
}
void QThreadPool_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QThreadPool*)self, &QThreadPool::destroyed, (QThreadPool*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QThreadPool);
	__conns_QThreadPool[__id] = __c;
}
void QThreadPool_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QThreadPool*)self, &QThreadPool::objectNameChanged, (QThreadPool*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QThreadPool);
	__conns_QThreadPool[__id] = __c;
}
void QThreadPool__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QThreadPool);
		auto __it = __conns_QThreadPool.find(id);
		if (__it == __conns_QThreadPool.end()) return;
		__c = __it->second;
		__conns_QThreadPool.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QThreadPool__Go_new(long long id, void* parent) { return new QThreadPool__Go(id, (QObject*)parent); }
void QThreadPool_delete(void* self) { delete (QThreadPool*)self; }

}

#include <QtCore/qmargins.h>
#include <QtCore/qrect.h>
#include "gen_qmargins.h"

extern "C" {

void* QMargins_new() {
	return new QMargins();
}
void* QMargins_new2(int left, int top, int right, int bottom) {
	return new QMargins(left, top, right, bottom);
}
void* QMargins_new3(void* arg1) {
	return new QMargins(*(QMargins*)arg1);
}
int QMargins_isNull(void* self) {
	auto&& __r = ((QMargins*)self)->isNull();
	return __r;
}
int QMargins_left(void* self) {
	auto&& __r = ((QMargins*)self)->left();
	return __r;
}
int QMargins_top(void* self) {
	auto&& __r = ((QMargins*)self)->top();
	return __r;
}
int QMargins_right(void* self) {
	auto&& __r = ((QMargins*)self)->right();
	return __r;
}
int QMargins_bottom(void* self) {
	auto&& __r = ((QMargins*)self)->bottom();
	return __r;
}
void QMargins_setLeft(void* self, int left) {
	((QMargins*)self)->setLeft(left);
}
void QMargins_setTop(void* self, int top) {
	((QMargins*)self)->setTop(top);
}
void QMargins_setRight(void* self, int right) {
	((QMargins*)self)->setRight(right);
}
void QMargins_setBottom(void* self, int bottom) {
	((QMargins*)self)->setBottom(bottom);
}
void* QMargins_OperatorPlusAssign(void* self, void* margins) {
	auto&& __r = ((QMargins*)self)->operator+=(*(QMargins*)margins);
	return (void*)&(__r);
}
void* QMargins_OperatorMinusAssign(void* self, void* margins) {
	auto&& __r = ((QMargins*)self)->operator-=(*(QMargins*)margins);
	return (void*)&(__r);
}
void* QMargins_OperatorPlusAssign2(void* self, int arg1) {
	auto&& __r = ((QMargins*)self)->operator+=(arg1);
	return (void*)&(__r);
}
void* QMargins_OperatorMinusAssign2(void* self, int arg1) {
	auto&& __r = ((QMargins*)self)->operator-=(arg1);
	return (void*)&(__r);
}
void* QMargins_OperatorMultiplyAssign(void* self, int arg1) {
	auto&& __r = ((QMargins*)self)->operator*=(arg1);
	return (void*)&(__r);
}
void* QMargins_OperatorDivideAssign(void* self, int arg1) {
	auto&& __r = ((QMargins*)self)->operator/=(arg1);
	return (void*)&(__r);
}
void* QMargins_OperatorMultiplyAssign2(void* self, double arg1) {
	auto&& __r = ((QMargins*)self)->operator*=(arg1);
	return (void*)&(__r);
}
void* QMargins_OperatorDivideAssign2(void* self, double arg1) {
	auto&& __r = ((QMargins*)self)->operator/=(arg1);
	return (void*)&(__r);
}
void* QMargins_toMarginsF(void* self) {
	auto&& __r = ((QMargins*)self)->toMarginsF();
	return new QMarginsF(__r);
}
void* QMargins_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QMargins*)self)->operator=(*(QMargins*)arg1);
	return (void*)&(__r);
}
void* QMargins_OperatorPlus(void* self, void* m2) {
	auto&& __r = (*(QMargins*)self) + (*(QMargins*)m2);
	return new QMargins(__r);
}
void* QMargins_OperatorMinus(void* self, void* m2) {
	auto&& __r = (*(QMargins*)self) - (*(QMargins*)m2);
	return new QMargins(__r);
}
void* QMargins_OperatorPlus2(void* self, int rhs) {
	auto&& __r = (*(QMargins*)self) + (rhs);
	return new QMargins(__r);
}
void* QMargins_OperatorMinus2(void* self, int rhs) {
	auto&& __r = (*(QMargins*)self) - (rhs);
	return new QMargins(__r);
}
void* QMargins_OperatorMultiply(void* self, int factor) {
	auto&& __r = (*(QMargins*)self) * (factor);
	return new QMargins(__r);
}
void* QMargins_OperatorMultiply2(void* self, double factor) {
	auto&& __r = (*(QMargins*)self) * (factor);
	return new QMargins(__r);
}
void* QMargins_OperatorDivide(void* self, int divisor) {
	auto&& __r = (*(QMargins*)self) / (divisor);
	return new QMargins(__r);
}
void* QMargins_OperatorDivide2(void* self, double divisor) {
	auto&& __r = (*(QMargins*)self) / (divisor);
	return new QMargins(__r);
}
void* QMargins_OperatorBitwiseOr(void* self, void* m2) {
	auto&& __r = (*(QMargins*)self) | (*(QMargins*)m2);
	return new QMargins(__r);
}
int QMargins_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QMargins*)self) == (*(QMargins*)rhs);
	return __r;
}
int QMargins_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QMargins*)self) != (*(QMargins*)rhs);
	return __r;
}
void* QMargins_OperatorPlus3(void* self) {
	auto&& __r = +(*(QMargins*)self);
	return new QMargins(__r);
}
void* QMargins_OperatorMinus3(void* self) {
	auto&& __r = -(*(QMargins*)self);
	return new QMargins(__r);
}
int QMargins_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QMargins*)self) == (*(QMarginsF*)rhs);
	return __r;
}
int QMargins_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QMargins*)self) != (*(QMarginsF*)rhs);
	return __r;
}
void* QMargins_OperatorPlus4(void* self, void* rectangle) {
	auto&& __r = (*(QMargins*)self) + (*(QRect*)rectangle);
	return new QRect(__r);
}
void QMargins_delete(void* self) { delete (QMargins*)self; }

}

#include <QtCore/qplugin.h>
#include "gen_qpluginmetadata__header.h"

extern "C" {

void* QPluginMetaData__Header_new(void* arg1) {
	return new QPluginMetaData::Header(*(QPluginMetaData::Header*)arg1);
}
unsigned char QPluginMetaData__Header_version_get(void* self) {
	auto __r = ((QPluginMetaData::Header*)self)->version;
	return __r;
}
void QPluginMetaData__Header_version_set(void* self, unsigned char version) {
	((QPluginMetaData::Header*)self)->version = version;
}
unsigned char QPluginMetaData__Header_qt_major_version_get(void* self) {
	auto __r = ((QPluginMetaData::Header*)self)->qt_major_version;
	return __r;
}
void QPluginMetaData__Header_qt_major_version_set(void* self, unsigned char qt_major_version) {
	((QPluginMetaData::Header*)self)->qt_major_version = qt_major_version;
}
unsigned char QPluginMetaData__Header_qt_minor_version_get(void* self) {
	auto __r = ((QPluginMetaData::Header*)self)->qt_minor_version;
	return __r;
}
void QPluginMetaData__Header_qt_minor_version_set(void* self, unsigned char qt_minor_version) {
	((QPluginMetaData::Header*)self)->qt_minor_version = qt_minor_version;
}
unsigned char QPluginMetaData__Header_plugin_arch_requirements_get(void* self) {
	auto __r = ((QPluginMetaData::Header*)self)->plugin_arch_requirements;
	return __r;
}
void QPluginMetaData__Header_plugin_arch_requirements_set(void* self, unsigned char plugin_arch_requirements) {
	((QPluginMetaData::Header*)self)->plugin_arch_requirements = plugin_arch_requirements;
}
void QPluginMetaData__Header_delete(void* self) { delete (QPluginMetaData::Header*)self; }

}

#include <QtCore/qsemaphore.h>
#include "gen_qsemaphorereleaser.h"

extern "C" {

void* QSemaphoreReleaser_new() {
	return new QSemaphoreReleaser();
}
void* QSemaphoreReleaser_new2(void* sem) {
	return new QSemaphoreReleaser(*(QSemaphore*)sem);
}
void* QSemaphoreReleaser_new3(void* sem) {
	return new QSemaphoreReleaser((QSemaphore*)sem);
}
void* QSemaphoreReleaser_new4(void* sem, int n) {
	return new QSemaphoreReleaser(*(QSemaphore*)sem, n);
}
void* QSemaphoreReleaser_new5(void* sem, int n) {
	return new QSemaphoreReleaser((QSemaphore*)sem, n);
}
void QSemaphoreReleaser_swap(void* self, void* other) {
	((QSemaphoreReleaser*)self)->swap(*(QSemaphoreReleaser*)other);
}
void* QSemaphoreReleaser_semaphore(void* self) {
	auto&& __r = ((QSemaphoreReleaser*)self)->semaphore();
	return (void*)(__r);
}
void* QSemaphoreReleaser_cancel(void* self) {
	auto&& __r = ((QSemaphoreReleaser*)self)->cancel();
	return (void*)(__r);
}
void QSemaphoreReleaser_delete(void* self) { delete (QSemaphoreReleaser*)self; }

}

#include <QtCore/qwaitcondition.h>
#include <QtCore/qmutex.h>
#include <QtCore/qdeadlinetimer.h>
#include <QtCore/qreadwritelock.h>
#include "gen_qwaitcondition.h"

extern "C" {

void* QWaitCondition_new() {
	return new QWaitCondition();
}
int QWaitCondition_wait(void* self, void* lockedMutex) {
	auto&& __r = ((QWaitCondition*)self)->wait((QMutex*)lockedMutex);
	return __r;
}
int QWaitCondition_wait2(void* self, void* lockedMutex, unsigned long time) {
	auto&& __r = ((QWaitCondition*)self)->wait((QMutex*)lockedMutex, time);
	return __r;
}
int QWaitCondition_wait3(void* self, void* lockedReadWriteLock) {
	auto&& __r = ((QWaitCondition*)self)->wait((QReadWriteLock*)lockedReadWriteLock);
	return __r;
}
int QWaitCondition_wait4(void* self, void* lockedReadWriteLock, unsigned long time) {
	auto&& __r = ((QWaitCondition*)self)->wait((QReadWriteLock*)lockedReadWriteLock, time);
	return __r;
}
void QWaitCondition_wakeOne(void* self) {
	((QWaitCondition*)self)->wakeOne();
}
void QWaitCondition_wakeAll(void* self) {
	((QWaitCondition*)self)->wakeAll();
}
void QWaitCondition_notify_one(void* self) {
	((QWaitCondition*)self)->notify_one();
}
void QWaitCondition_notify_all(void* self) {
	((QWaitCondition*)self)->notify_all();
}
int QWaitCondition_wait5(void* self, void* lockedMutex, void* deadline) {
	auto&& __r = ((QWaitCondition*)self)->wait((QMutex*)lockedMutex, *(QDeadlineTimer*)deadline);
	return __r;
}
int QWaitCondition_wait6(void* self, void* lockedReadWriteLock, void* deadline) {
	auto&& __r = ((QWaitCondition*)self)->wait((QReadWriteLock*)lockedReadWriteLock, *(QDeadlineTimer*)deadline);
	return __r;
}
void QWaitCondition_delete(void* self) { delete (QWaitCondition*)self; }

}

#include <QtGui/qvectornd.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qpoint.h>
#include <QtCore/qvariant.h>
#include <QtGui/qmatrix4x4.h>
#include "gen_qvector4d.h"

extern "C" {

void* QVector4D_new() {
	return new QVector4D();
}
void* QVector4D_new2(int arg1) {
	return new QVector4D((Qt::Initialization)(arg1));
}
void* QVector4D_new3(float xpos, float ypos, float zpos, float wpos) {
	return new QVector4D(xpos, ypos, zpos, wpos);
}
void* QVector4D_new4(void* point) {
	return new QVector4D(*(QPoint*)point);
}
void* QVector4D_new5(void* point) {
	return new QVector4D(*(QPointF*)point);
}
void* QVector4D_new6(void* vector) {
	return new QVector4D(*(QVector2D*)vector);
}
void* QVector4D_new7(void* vector, float zpos, float wpos) {
	return new QVector4D(*(QVector2D*)vector, zpos, wpos);
}
void* QVector4D_new8(void* vector) {
	return new QVector4D(*(QVector3D*)vector);
}
void* QVector4D_new9(void* vector, float wpos) {
	return new QVector4D(*(QVector3D*)vector, wpos);
}
void* QVector4D_new10(void* arg1) {
	return new QVector4D(*(QVector4D*)arg1);
}
int QVector4D_isNull(void* self) {
	auto&& __r = ((QVector4D*)self)->isNull();
	return __r;
}
float QVector4D_x(void* self) {
	auto&& __r = ((QVector4D*)self)->x();
	return __r;
}
float QVector4D_y(void* self) {
	auto&& __r = ((QVector4D*)self)->y();
	return __r;
}
float QVector4D_z(void* self) {
	auto&& __r = ((QVector4D*)self)->z();
	return __r;
}
float QVector4D_w(void* self) {
	auto&& __r = ((QVector4D*)self)->w();
	return __r;
}
void QVector4D_setX(void* self, float x) {
	((QVector4D*)self)->setX(x);
}
void QVector4D_setY(void* self, float y) {
	((QVector4D*)self)->setY(y);
}
void QVector4D_setZ(void* self, float z) {
	((QVector4D*)self)->setZ(z);
}
void QVector4D_setW(void* self, float w) {
	((QVector4D*)self)->setW(w);
}
float QVector4D_OperatorSubscript(void* self, int i) {
	auto&& __r = ((QVector4D*)self)->operator[](i);
	return __r;
}
float QVector4D_length(void* self) {
	auto&& __r = ((QVector4D*)self)->length();
	return __r;
}
float QVector4D_lengthSquared(void* self) {
	auto&& __r = ((QVector4D*)self)->lengthSquared();
	return __r;
}
void* QVector4D_normalized(void* self) {
	auto&& __r = ((QVector4D*)self)->normalized();
	return new QVector4D(__r);
}
void QVector4D_normalize(void* self) {
	((QVector4D*)self)->normalize();
}
void* QVector4D_OperatorPlusAssign(void* self, void* vector) {
	auto&& __r = ((QVector4D*)self)->operator+=(*(QVector4D*)vector);
	return (void*)&(__r);
}
void* QVector4D_OperatorMinusAssign(void* self, void* vector) {
	auto&& __r = ((QVector4D*)self)->operator-=(*(QVector4D*)vector);
	return (void*)&(__r);
}
void* QVector4D_OperatorMultiplyAssign(void* self, float factor) {
	auto&& __r = ((QVector4D*)self)->operator*=(factor);
	return (void*)&(__r);
}
void* QVector4D_OperatorMultiplyAssign2(void* self, void* vector) {
	auto&& __r = ((QVector4D*)self)->operator*=(*(QVector4D*)vector);
	return (void*)&(__r);
}
void* QVector4D_OperatorDivideAssign(void* self, float divisor) {
	auto&& __r = ((QVector4D*)self)->operator/=(divisor);
	return (void*)&(__r);
}
void* QVector4D_OperatorDivideAssign2(void* self, void* vector) {
	auto&& __r = ((QVector4D*)self)->operator/=(*(QVector4D*)vector);
	return (void*)&(__r);
}
float QVector4D_dotProduct(void* v1, void* v2) {
	auto&& __r = QVector4D::dotProduct(*(QVector4D*)v1, *(QVector4D*)v2);
	return __r;
}
void* QVector4D_toVector2D(void* self) {
	auto&& __r = ((QVector4D*)self)->toVector2D();
	return new QVector2D(__r);
}
void* QVector4D_toVector2DAffine(void* self) {
	auto&& __r = ((QVector4D*)self)->toVector2DAffine();
	return new QVector2D(__r);
}
void* QVector4D_toVector3D(void* self) {
	auto&& __r = ((QVector4D*)self)->toVector3D();
	return new QVector3D(__r);
}
void* QVector4D_toVector3DAffine(void* self) {
	auto&& __r = ((QVector4D*)self)->toVector3DAffine();
	return new QVector3D(__r);
}
void* QVector4D_toPoint(void* self) {
	auto&& __r = ((QVector4D*)self)->toPoint();
	return new QPoint(__r);
}
void* QVector4D_toPointF(void* self) {
	auto&& __r = ((QVector4D*)self)->toPointF();
	return new QPointF(__r);
}
void* QVector4D_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QVector4D*)self);
	return new QVariant(__r);
}
int QVector4D_OperatorEqual(void* self, void* v2) {
	auto&& __r = (*(QVector4D*)self) == (*(QVector4D*)v2);
	return __r;
}
int QVector4D_OperatorNotEqual(void* self, void* v2) {
	auto&& __r = (*(QVector4D*)self) != (*(QVector4D*)v2);
	return __r;
}
void* QVector4D_OperatorPlus(void* self, void* v2) {
	auto&& __r = (*(QVector4D*)self) + (*(QVector4D*)v2);
	return new QVector4D(__r);
}
void* QVector4D_OperatorMinus(void* self, void* v2) {
	auto&& __r = (*(QVector4D*)self) - (*(QVector4D*)v2);
	return new QVector4D(__r);
}
void* QVector4D_OperatorMultiply(void* self, float factor) {
	auto&& __r = (*(QVector4D*)self) * (factor);
	return new QVector4D(__r);
}
void* QVector4D_OperatorMultiply2(void* self, void* v2) {
	auto&& __r = (*(QVector4D*)self) * (*(QVector4D*)v2);
	return new QVector4D(__r);
}
void* QVector4D_OperatorMinus2(void* self) {
	auto&& __r = -(*(QVector4D*)self);
	return new QVector4D(__r);
}
void* QVector4D_OperatorDivide(void* self, float divisor) {
	auto&& __r = (*(QVector4D*)self) / (divisor);
	return new QVector4D(__r);
}
void* QVector4D_OperatorDivide2(void* self, void* divisor) {
	auto&& __r = (*(QVector4D*)self) / (*(QVector4D*)divisor);
	return new QVector4D(__r);
}
void* QVector4D_OperatorMultiply3(void* self, void* matrix) {
	auto&& __r = (*(QVector4D*)self) * (*(QMatrix4x4*)matrix);
	return new QVector4D(__r);
}
void QVector4D_delete(void* self) { delete (QVector4D*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qsize.h>
#include "gen_qresizeevent.h"

extern "C" {

void* QResizeEvent_base(void* self) { return (void*)static_cast<QEvent*>((QResizeEvent*)self); }
void* QResizeEvent_new(void* size, void* oldSize) {
	return new QResizeEvent(*(QSize*)size, *(QSize*)oldSize);
}
void* QResizeEvent_size(void* self) {
	auto&& __r = ((QResizeEvent*)self)->size();
	return (void*)&(__r);
}
void* QResizeEvent_oldSize(void* self) {
	auto&& __r = ((QResizeEvent*)self)->oldSize();
	return (void*)&(__r);
}
int QResizeEvent_registerEventType() {
	auto&& __r = QResizeEvent::registerEventType();
	return __r;
}
int QResizeEvent_registerEventType2(int hint) {
	auto&& __r = QResizeEvent::registerEventType(hint);
	return __r;
}
void QResizeEvent_delete(void* self) { delete (QResizeEvent*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qscreen.h>
#include <QtCore/qnamespace.h>
#include "gen_qscreenorientationchangeevent.h"

extern "C" {

void* QScreenOrientationChangeEvent_base(void* self) { return (void*)static_cast<QEvent*>((QScreenOrientationChangeEvent*)self); }
void* QScreenOrientationChangeEvent_new(void* screen, int orientation) {
	return new QScreenOrientationChangeEvent((QScreen*)screen, (Qt::ScreenOrientation)(orientation));
}
void* QScreenOrientationChangeEvent_screen(void* self) {
	auto&& __r = ((QScreenOrientationChangeEvent*)self)->screen();
	return (void*)(__r);
}
int QScreenOrientationChangeEvent_orientation(void* self) {
	auto&& __r = ((QScreenOrientationChangeEvent*)self)->orientation();
	return (int)(__r);
}
int QScreenOrientationChangeEvent_registerEventType() {
	auto&& __r = QScreenOrientationChangeEvent::registerEventType();
	return __r;
}
int QScreenOrientationChangeEvent_registerEventType2(int hint) {
	auto&& __r = QScreenOrientationChangeEvent::registerEventType(hint);
	return __r;
}
void QScreenOrientationChangeEvent_delete(void* self) { delete (QScreenOrientationChangeEvent*)self; }

}

#include <QtGui/qtextformat.h>
#include <QtCore/qstring.h>
#include <QtCore/qvariant.h>
#include <QtGui/qcolor.h>
#include <QtGui/qpen.h>
#include <QtGui/qbrush.h>
#include <QtCore/qnamespace.h>
#include "gen_qtextlistformat.h"

extern "C" {

void* QTextListFormat_base(void* self) { return (void*)static_cast<QTextFormat*>((QTextListFormat*)self); }
void* QTextListFormat_new() {
	return new QTextListFormat();
}
void* QTextListFormat_new2(void* arg1) {
	return new QTextListFormat(*(QTextListFormat*)arg1);
}
int QTextListFormat_isValid(void* self) {
	auto&& __r = ((QTextListFormat*)self)->isValid();
	return __r;
}
void QTextListFormat_setStyle(void* self, int style) {
	((QTextListFormat*)self)->setStyle((QTextListFormat::Style)(style));
}
int QTextListFormat_style(void* self) {
	auto&& __r = ((QTextListFormat*)self)->style();
	return (int)(__r);
}
void QTextListFormat_setIndent(void* self, int indent) {
	((QTextListFormat*)self)->setIndent(indent);
}
int QTextListFormat_indent(void* self) {
	auto&& __r = ((QTextListFormat*)self)->indent();
	return __r;
}
void QTextListFormat_setNumberPrefix(void* self, const char* numberPrefix, long long numberPrefix_len) {
	QString numberPrefix_qstr = QString::fromUtf8(numberPrefix, numberPrefix_len);
	((QTextListFormat*)self)->setNumberPrefix(numberPrefix_qstr);
}
char* QTextListFormat_numberPrefix(void* self, long long* __out_len) {
	auto&& __r = ((QTextListFormat*)self)->numberPrefix();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextListFormat_setNumberSuffix(void* self, const char* numberSuffix, long long numberSuffix_len) {
	QString numberSuffix_qstr = QString::fromUtf8(numberSuffix, numberSuffix_len);
	((QTextListFormat*)self)->setNumberSuffix(numberSuffix_qstr);
}
char* QTextListFormat_numberSuffix(void* self, long long* __out_len) {
	auto&& __r = ((QTextListFormat*)self)->numberSuffix();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextListFormat_setStart(void* self, int indent) {
	((QTextListFormat*)self)->setStart(indent);
}
int QTextListFormat_start(void* self) {
	auto&& __r = ((QTextListFormat*)self)->start();
	return __r;
}
void* QTextListFormat_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QTextListFormat*)self)->operator=(*(QTextListFormat*)arg1);
	return (void*)&(__r);
}
void QTextListFormat_delete(void* self) { delete (QTextListFormat*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtCore/qstring.h>
#include "gen_qaccessibletableinterface.h"
#include "gen_purego_cb.h"

class QAccessibleTableInterface__Go : public QAccessibleTableInterface {
public:
	long long __goId;
	unsigned char __goOn[20] = {};
	template<typename... A> QAccessibleTableInterface__Go(long long id, A&&... a) : QAccessibleTableInterface(std::forward<A>(a)...), __goId(id) {}
	QAccessibleInterface* caption() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6749251983603663525LL, &__h); if (__h) return (QAccessibleInterface*)__r; }
		return nullptr;
	}
	QAccessibleInterface* summary() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -6646608006195182317LL, &__h); if (__h) return (QAccessibleInterface*)__r; }
		return nullptr;
	}
	QAccessibleInterface* cellAt(int row, int column) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int*))cppgo_virt_voidp_int_int)(__goId, 6973530125813990328LL, row, column, &__h); if (__h) return (QAccessibleInterface*)__r; }
		return nullptr;
	}
	int selectedCellCount() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1985908008123410713LL, &__h); if (__h) return __r; }
		return int{};
	}
	QList<QAccessibleInterface*> selectedCells() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, -2099517733663994687LL, &__len, &__h);
			if (__h) {
				QList<QAccessibleInterface*> __r; __r.reserve(__len);
				void** __aa = (void**)__a;
				for (long long __i = 0; __i < __len; __i++) __r.append((QAccessibleInterface*)__aa[__i]);
				if (__a) free(__a);
				return __r;
			}
		}
		return QList<QAccessibleInterface*>{};
	}
	QString columnDescription(int column) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virtp_charp_int) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, int, long long*, int*))cppgo_virtp_charp_int)(__goId, -4866595225830892309LL, column, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QString{};
	}
	QString rowDescription(int row) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virtp_charp_int) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, int, long long*, int*))cppgo_virtp_charp_int)(__goId, -2563375781797925217LL, row, &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QString{};
	}
	int selectedColumnCount() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -4569576858401950571LL, &__h); if (__h) return __r; }
		return int{};
	}
	int selectedRowCount() const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -2566786583858001057LL, &__h); if (__h) return __r; }
		return int{};
	}
	int columnCount() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8256963024669584834LL, &__h); if (__h) return __r; }
		return int{};
	}
	int rowCount() const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 1960960171956884010LL, &__h); if (__h) return __r; }
		return int{};
	}
	QList<int> selectedColumns() const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, -278488305062277297LL, &__len, &__h);
			if (__h) {
				QList<int> __r; __r.reserve(__len);
				int* __aa = (int*)__a;
				for (long long __i = 0; __i < __len; __i++) __r.append(__aa[__i]);
				if (__a) free(__a);
				return __r;
			}
		}
		return QList<int>{};
	}
	QList<int> selectedRows() const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, -5252052890467341575LL, &__len, &__h);
			if (__h) {
				QList<int> __r; __r.reserve(__len);
				int* __aa = (int*)__a;
				for (long long __i = 0; __i < __len; __i++) __r.append(__aa[__i]);
				if (__a) free(__a);
				return __r;
			}
		}
		return QList<int>{};
	}
	bool isColumnSelected(int column) const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 1772671393755888886LL, column, &__h); if (__h) return __r; }
		return bool{};
	}
	bool isRowSelected(int row) const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -8255619831741278870LL, row, &__h); if (__h) return __r; }
		return bool{};
	}
	bool selectRow(int row) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 2033665621955588835LL, row, &__h); if (__h) return __r; }
		return bool{};
	}
	bool selectColumn(int column) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 3930440648437166287LL, column, &__h); if (__h) return __r; }
		return bool{};
	}
	bool unselectRow(int row) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -7465067139058793362LL, row, &__h); if (__h) return __r; }
		return bool{};
	}
	bool unselectColumn(int column) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5189646775272561424LL, column, &__h); if (__h) return __r; }
		return bool{};
	}
	void modelChange(QAccessibleTableModelChangeEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -8156005454275955122LL, (void*)event, &__h); if (__h) return; }
		return;
	}
};

extern "C" {

void QAccessibleTableInterface__set_override(void* self, int idx) { ((QAccessibleTableInterface__Go*)self)->__goOn[idx] = 1; }
void* QAccessibleTableInterface_caption(void* self) {
	auto&& __r = ((QAccessibleTableInterface*)self)->caption();
	return (void*)(__r);
}
void* QAccessibleTableInterface_summary(void* self) {
	auto&& __r = ((QAccessibleTableInterface*)self)->summary();
	return (void*)(__r);
}
void* QAccessibleTableInterface_cellAt(void* self, int row, int column) {
	auto&& __r = ((QAccessibleTableInterface*)self)->cellAt(row, column);
	return (void*)(__r);
}
int QAccessibleTableInterface_selectedCellCount(void* self) {
	auto&& __r = ((QAccessibleTableInterface*)self)->selectedCellCount();
	return __r;
}
void** QAccessibleTableInterface_selectedCells(void* self, long long* __out_len) {
	auto&& __r = ((QAccessibleTableInterface*)self)->selectedCells();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
char* QAccessibleTableInterface_columnDescription(void* self, int column, long long* __out_len) {
	auto&& __r = ((QAccessibleTableInterface*)self)->columnDescription(column);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleTableInterface_rowDescription(void* self, int row, long long* __out_len) {
	auto&& __r = ((QAccessibleTableInterface*)self)->rowDescription(row);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QAccessibleTableInterface_selectedColumnCount(void* self) {
	auto&& __r = ((QAccessibleTableInterface*)self)->selectedColumnCount();
	return __r;
}
int QAccessibleTableInterface_selectedRowCount(void* self) {
	auto&& __r = ((QAccessibleTableInterface*)self)->selectedRowCount();
	return __r;
}
int QAccessibleTableInterface_columnCount(void* self) {
	auto&& __r = ((QAccessibleTableInterface*)self)->columnCount();
	return __r;
}
int QAccessibleTableInterface_rowCount(void* self) {
	auto&& __r = ((QAccessibleTableInterface*)self)->rowCount();
	return __r;
}
int* QAccessibleTableInterface_selectedColumns(void* self, long long* __out_len) {
	auto&& __r = ((QAccessibleTableInterface*)self)->selectedColumns();
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
int* QAccessibleTableInterface_selectedRows(void* self, long long* __out_len) {
	auto&& __r = ((QAccessibleTableInterface*)self)->selectedRows();
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
int QAccessibleTableInterface_isColumnSelected(void* self, int column) {
	auto&& __r = ((QAccessibleTableInterface*)self)->isColumnSelected(column);
	return __r;
}
int QAccessibleTableInterface_isRowSelected(void* self, int row) {
	auto&& __r = ((QAccessibleTableInterface*)self)->isRowSelected(row);
	return __r;
}
int QAccessibleTableInterface_selectRow(void* self, int row) {
	auto&& __r = ((QAccessibleTableInterface*)self)->selectRow(row);
	return __r;
}
int QAccessibleTableInterface_selectColumn(void* self, int column) {
	auto&& __r = ((QAccessibleTableInterface*)self)->selectColumn(column);
	return __r;
}
int QAccessibleTableInterface_unselectRow(void* self, int row) {
	auto&& __r = ((QAccessibleTableInterface*)self)->unselectRow(row);
	return __r;
}
int QAccessibleTableInterface_unselectColumn(void* self, int column) {
	auto&& __r = ((QAccessibleTableInterface*)self)->unselectColumn(column);
	return __r;
}
void QAccessibleTableInterface_modelChange(void* self, void* event) {
	((QAccessibleTableInterface*)self)->modelChange((QAccessibleTableModelChangeEvent*)event);
}
void* QAccessibleTableInterface_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QAccessibleTableInterface*)self)->operator=(*(QAccessibleTableInterface*)arg1);
	return (void*)&(__r);
}
void* QAccessibleTableInterface__Go_new(long long id) { return new QAccessibleTableInterface__Go(id); }
void QAccessibleTableInterface_delete(void* self) { delete (QAccessibleTableInterface*)self; }

}

#include <QtGui/qsurface.h>
#include <QtGui/qsurfaceformat.h>
#include <QtCore/qsize.h>
#include "gen_qsurface.h"

extern "C" {

int QSurface_surfaceClass(void* self) {
	auto&& __r = ((QSurface*)self)->surfaceClass();
	return (int)(__r);
}
void* QSurface_format(void* self) {
	auto&& __r = ((QSurface*)self)->format();
	return new QSurfaceFormat(__r);
}
int QSurface_surfaceType(void* self) {
	auto&& __r = ((QSurface*)self)->surfaceType();
	return (int)(__r);
}
int QSurface_supportsOpenGL(void* self) {
	auto&& __r = ((QSurface*)self)->supportsOpenGL();
	return __r;
}
void* QSurface_size(void* self) {
	auto&& __r = ((QSurface*)self)->size();
	return new QSize(__r);
}
void* QSurface_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QSurface*)self)->operator=(*(QSurface*)arg1);
	return (void*)&(__r);
}
void QSurface_delete(void* self) { delete (QSurface*)self; }

}

#include <QtGui/qimagewriter.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qstring.h>
#include <QtGui/qimageiohandler.h>
#include <QtGui/qimage.h>
#include "gen_qimagewriter.h"

extern "C" {

void* QImageWriter_new() {
	return new QImageWriter();
}
void* QImageWriter_new2(void* device, void* format) {
	return new QImageWriter((QIODevice*)device, *(QByteArray*)format);
}
void* QImageWriter_new3(const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QImageWriter(fileName_qstr);
}
void* QImageWriter_new4(const char* fileName, long long fileName_len, void* format) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QImageWriter(fileName_qstr, *(QByteArray*)format);
}
char* QImageWriter_tr(const char* sourceText, long long* __out_len) {
	auto&& __r = QImageWriter::tr(sourceText);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QImageWriter_setFormat(void* self, void* format) {
	((QImageWriter*)self)->setFormat(*(QByteArray*)format);
}
void* QImageWriter_format(void* self) {
	auto&& __r = ((QImageWriter*)self)->format();
	return new QByteArray(__r);
}
void QImageWriter_setDevice(void* self, void* device) {
	((QImageWriter*)self)->setDevice((QIODevice*)device);
}
void* QImageWriter_device(void* self) {
	auto&& __r = ((QImageWriter*)self)->device();
	return (void*)(__r);
}
void QImageWriter_setFileName(void* self, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QImageWriter*)self)->setFileName(fileName_qstr);
}
char* QImageWriter_fileName(void* self, long long* __out_len) {
	auto&& __r = ((QImageWriter*)self)->fileName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QImageWriter_setQuality(void* self, int quality) {
	((QImageWriter*)self)->setQuality(quality);
}
int QImageWriter_quality(void* self) {
	auto&& __r = ((QImageWriter*)self)->quality();
	return __r;
}
void QImageWriter_setCompression(void* self, int compression) {
	((QImageWriter*)self)->setCompression(compression);
}
int QImageWriter_compression(void* self) {
	auto&& __r = ((QImageWriter*)self)->compression();
	return __r;
}
void QImageWriter_setSubType(void* self, void* type_) {
	((QImageWriter*)self)->setSubType(*(QByteArray*)type_);
}
void* QImageWriter_subType(void* self) {
	auto&& __r = ((QImageWriter*)self)->subType();
	return new QByteArray(__r);
}
void** QImageWriter_supportedSubTypes(void* self, long long* __out_len) {
	auto&& __r = ((QImageWriter*)self)->supportedSubTypes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
void QImageWriter_setOptimizedWrite(void* self, int optimize) {
	((QImageWriter*)self)->setOptimizedWrite(optimize);
}
int QImageWriter_optimizedWrite(void* self) {
	auto&& __r = ((QImageWriter*)self)->optimizedWrite();
	return __r;
}
void QImageWriter_setProgressiveScanWrite(void* self, int progressive) {
	((QImageWriter*)self)->setProgressiveScanWrite(progressive);
}
int QImageWriter_progressiveScanWrite(void* self) {
	auto&& __r = ((QImageWriter*)self)->progressiveScanWrite();
	return __r;
}
int QImageWriter_transformation(void* self) {
	auto&& __r = ((QImageWriter*)self)->transformation();
	return (int)(__r);
}
void QImageWriter_setTransformation(void* self, int orientation) {
	((QImageWriter*)self)->setTransformation(QFlags<QImageIOHandler::Transformation>(orientation));
}
void QImageWriter_setText(void* self, const char* key, long long key_len, const char* text, long long text_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QImageWriter*)self)->setText(key_qstr, text_qstr);
}
int QImageWriter_canWrite(void* self) {
	auto&& __r = ((QImageWriter*)self)->canWrite();
	return __r;
}
int QImageWriter_write(void* self, void* image) {
	auto&& __r = ((QImageWriter*)self)->write(*(QImage*)image);
	return __r;
}
int QImageWriter_error(void* self) {
	auto&& __r = ((QImageWriter*)self)->error();
	return (int)(__r);
}
char* QImageWriter_errorString(void* self, long long* __out_len) {
	auto&& __r = ((QImageWriter*)self)->errorString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QImageWriter_supportsOption(void* self, int option) {
	auto&& __r = ((QImageWriter*)self)->supportsOption((QImageIOHandler::ImageOption)(option));
	return __r;
}
void** QImageWriter_supportedImageFormats(long long* __out_len) {
	auto&& __r = QImageWriter::supportedImageFormats();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
void** QImageWriter_supportedMimeTypes(long long* __out_len) {
	auto&& __r = QImageWriter::supportedMimeTypes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
void** QImageWriter_imageFormatsForMimeType(void* mimeType, long long* __out_len) {
	auto&& __r = QImageWriter::imageFormatsForMimeType(*(QByteArray*)mimeType);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QByteArray(__e);
	return __arr;
}
char* QImageWriter_tr2(const char* sourceText, const char* disambiguation, long long* __out_len) {
	auto&& __r = QImageWriter::tr(sourceText, disambiguation);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QImageWriter_tr3(const char* sourceText, const char* disambiguation, int n, long long* __out_len) {
	auto&& __r = QImageWriter::tr(sourceText, disambiguation, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QImageWriter_delete(void* self) { delete (QImageWriter*)self; }

}

#include <QtGui/qpdfwriter.h>
#include <QtCore/qobject.h>
#include <QtGui/qpagedpaintdevice.h>
#include <QtCore/qstring.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/quuid.h>
#include <QtCore/qbytearray.h>
#include <QtGui/qpdfoutputintent.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qpaintdevice.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbindingstorage.h>
#include <QtGui/qpagelayout.h>
#include <QtGui/qpagesize.h>
#include <QtCore/qmargins.h>
#include <QtGui/qpageranges.h>
#include <QtGui/qpainter.h>
#include <QtCore/qpoint.h>
#include "gen_qpdfwriter.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QPdfWriter;
static std::mutex __connsMu_QPdfWriter;

extern "C" {

void* QPdfWriter_base(void* self) { return (void*)static_cast<QObject*>((QPdfWriter*)self); }
void* QPdfWriter_fromQObject(void* o) { return (void*)qobject_cast<QPdfWriter*>((QObject*)o); }
void* QPdfWriter_new(const char* filename, long long filename_len) {
	QString filename_qstr = QString::fromUtf8(filename, filename_len);
	return new QPdfWriter(filename_qstr);
}
void* QPdfWriter_new2(void* device) {
	return new QPdfWriter((QIODevice*)device);
}
char* QPdfWriter_tr(const char* s, long long* __out_len) {
	auto&& __r = QPdfWriter::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QPdfWriter_setPdfVersion(void* self, int version) {
	((QPdfWriter*)self)->setPdfVersion((QPagedPaintDevice::PdfVersion)(version));
}
int QPdfWriter_pdfVersion(void* self) {
	auto&& __r = ((QPdfWriter*)self)->pdfVersion();
	return (int)(__r);
}
char* QPdfWriter_title(void* self, long long* __out_len) {
	auto&& __r = ((QPdfWriter*)self)->title();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QPdfWriter_setTitle(void* self, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	((QPdfWriter*)self)->setTitle(title_qstr);
}
char* QPdfWriter_creator(void* self, long long* __out_len) {
	auto&& __r = ((QPdfWriter*)self)->creator();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QPdfWriter_setCreator(void* self, const char* creator, long long creator_len) {
	QString creator_qstr = QString::fromUtf8(creator, creator_len);
	((QPdfWriter*)self)->setCreator(creator_qstr);
}
void* QPdfWriter_documentId(void* self) {
	auto&& __r = ((QPdfWriter*)self)->documentId();
	return new QUuid(__r);
}
void QPdfWriter_setDocumentId(void* self, void* documentId) {
	((QPdfWriter*)self)->setDocumentId(*(QUuid*)documentId);
}
char* QPdfWriter_author(void* self, long long* __out_len) {
	auto&& __r = ((QPdfWriter*)self)->author();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QPdfWriter_setAuthor(void* self, const char* author, long long author_len) {
	QString author_qstr = QString::fromUtf8(author, author_len);
	((QPdfWriter*)self)->setAuthor(author_qstr);
}
int QPdfWriter_newPage(void* self) {
	auto&& __r = ((QPdfWriter*)self)->newPage();
	return __r;
}
void QPdfWriter_setResolution(void* self, int resolution) {
	((QPdfWriter*)self)->setResolution(resolution);
}
int QPdfWriter_resolution(void* self) {
	auto&& __r = ((QPdfWriter*)self)->resolution();
	return __r;
}
void QPdfWriter_setDocumentXmpMetadata(void* self, void* xmpMetadata) {
	((QPdfWriter*)self)->setDocumentXmpMetadata(*(QByteArray*)xmpMetadata);
}
void* QPdfWriter_documentXmpMetadata(void* self) {
	auto&& __r = ((QPdfWriter*)self)->documentXmpMetadata();
	return new QByteArray(__r);
}
void QPdfWriter_addFileAttachment(void* self, const char* fileName, long long fileName_len, void* data) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	((QPdfWriter*)self)->addFileAttachment(fileName_qstr, *(QByteArray*)data);
}
int QPdfWriter_colorModel(void* self) {
	auto&& __r = ((QPdfWriter*)self)->colorModel();
	return (int)(__r);
}
void QPdfWriter_setColorModel(void* self, int model) {
	((QPdfWriter*)self)->setColorModel((QPdfWriter::ColorModel)(model));
}
void* QPdfWriter_outputIntent(void* self) {
	auto&& __r = ((QPdfWriter*)self)->outputIntent();
	return new QPdfOutputIntent(__r);
}
void QPdfWriter_setOutputIntent(void* self, void* intent) {
	((QPdfWriter*)self)->setOutputIntent(*(QPdfOutputIntent*)intent);
}
void* QPdfWriter_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QPdfWriter::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QPdfWriter_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QPdfWriter::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
int QPdfWriter_setPageLayout(void* self, void* pageLayout) {
	auto&& __r = ((QPdfWriter*)self)->setPageLayout(*(QPageLayout*)pageLayout);
	return __r;
}
int QPdfWriter_setPageSize(void* self, void* pageSize) {
	auto&& __r = ((QPdfWriter*)self)->setPageSize(*(QPageSize*)pageSize);
	return __r;
}
int QPdfWriter_setPageOrientation(void* self, int orientation) {
	auto&& __r = ((QPdfWriter*)self)->setPageOrientation((QPageLayout::Orientation)(orientation));
	return __r;
}
int QPdfWriter_setPageMargins(void* self, void* margins, int units) {
	auto&& __r = ((QPdfWriter*)self)->setPageMargins(*(QMarginsF*)margins, (QPageLayout::Unit)(units));
	return __r;
}
void* QPdfWriter_pageLayout(void* self) {
	auto&& __r = ((QPdfWriter*)self)->pageLayout();
	return new QPageLayout(__r);
}
void QPdfWriter_setPageRanges(void* self, void* ranges) {
	((QPdfWriter*)self)->setPageRanges(*(QPageRanges*)ranges);
}
void* QPdfWriter_pageRanges(void* self) {
	auto&& __r = ((QPdfWriter*)self)->pageRanges();
	return new QPageRanges(__r);
}
int QPdfWriter_devType(void* self) {
	auto&& __r = ((QPdfWriter*)self)->devType();
	return __r;
}
int QPdfWriter_paintingActive(void* self) {
	auto&& __r = ((QPdfWriter*)self)->paintingActive();
	return __r;
}
int QPdfWriter_width(void* self) {
	auto&& __r = ((QPdfWriter*)self)->width();
	return __r;
}
int QPdfWriter_height(void* self) {
	auto&& __r = ((QPdfWriter*)self)->height();
	return __r;
}
int QPdfWriter_widthMM(void* self) {
	auto&& __r = ((QPdfWriter*)self)->widthMM();
	return __r;
}
int QPdfWriter_heightMM(void* self) {
	auto&& __r = ((QPdfWriter*)self)->heightMM();
	return __r;
}
int QPdfWriter_logicalDpiX(void* self) {
	auto&& __r = ((QPdfWriter*)self)->logicalDpiX();
	return __r;
}
int QPdfWriter_logicalDpiY(void* self) {
	auto&& __r = ((QPdfWriter*)self)->logicalDpiY();
	return __r;
}
int QPdfWriter_physicalDpiX(void* self) {
	auto&& __r = ((QPdfWriter*)self)->physicalDpiX();
	return __r;
}
int QPdfWriter_physicalDpiY(void* self) {
	auto&& __r = ((QPdfWriter*)self)->physicalDpiY();
	return __r;
}
double QPdfWriter_devicePixelRatio(void* self) {
	auto&& __r = ((QPdfWriter*)self)->devicePixelRatio();
	return __r;
}
double QPdfWriter_devicePixelRatioF(void* self) {
	auto&& __r = ((QPdfWriter*)self)->devicePixelRatioF();
	return __r;
}
int QPdfWriter_colorCount(void* self) {
	auto&& __r = ((QPdfWriter*)self)->colorCount();
	return __r;
}
int QPdfWriter_depth(void* self) {
	auto&& __r = ((QPdfWriter*)self)->depth();
	return __r;
}
double QPdfWriter_devicePixelRatioFScale() {
	auto&& __r = QPdfWriter::devicePixelRatioFScale();
	return __r;
}
int QPdfWriter_encodeMetricF(int metric, double value) {
	auto&& __r = QPdfWriter::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
char* QPdfWriter_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QPdfWriter::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QPdfWriter_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QPdfWriter::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QPdfWriter_addFileAttachment2(void* self, const char* fileName, long long fileName_len, void* data, const char* mimeType, long long mimeType_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	QString mimeType_qstr = QString::fromUtf8(mimeType, mimeType_len);
	((QPdfWriter*)self)->addFileAttachment(fileName_qstr, *(QByteArray*)data, mimeType_qstr);
}
void* QPdfWriter_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QPdfWriter::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QPdfWriter_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QPdfWriter::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QPdfWriter_asQPagedPaintDevice(void* self) {
	auto&& __r = static_cast<QPagedPaintDevice*>((QPdfWriter*)self);
	return (void*)(__r);
}
void QPdfWriter_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPdfWriter*)self, &QPdfWriter::destroyed, (QPdfWriter*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPdfWriter);
	__conns_QPdfWriter[__id] = __c;
}
void QPdfWriter_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPdfWriter*)self, &QPdfWriter::destroyed, (QPdfWriter*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPdfWriter);
	__conns_QPdfWriter[__id] = __c;
}
void QPdfWriter_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPdfWriter*)self, &QPdfWriter::objectNameChanged, (QPdfWriter*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPdfWriter);
	__conns_QPdfWriter[__id] = __c;
}
void QPdfWriter__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QPdfWriter);
		auto __it = __conns_QPdfWriter.find(id);
		if (__it == __conns_QPdfWriter.end()) return;
		__c = __it->second;
		__conns_QPdfWriter.erase(__it);
	}
	QObject::disconnect(__c);
}
void QPdfWriter_delete(void* self) { delete (QPdfWriter*)self; }

}

#include <QtGui/qtextlist.h>
#include <QtGui/qtextobject.h>
#include <QtGui/qtextdocument.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qtextformat.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qobject.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qtextlist.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTextList;
static std::mutex __connsMu_QTextList;

extern "C" {

void* QTextList_base(void* self) { return (void*)static_cast<QTextBlockGroup*>((QTextList*)self); }
void QTextList_chain(void* self, void** out) {
	QTextList* __p = (QTextList*)self;
	QTextBlockGroup* __b0 = static_cast<QTextBlockGroup*>(__p); out[0] = (void*)__b0;
	QTextObject* __b1 = static_cast<QTextObject*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void* QTextList_fromQObject(void* o) { return (void*)qobject_cast<QTextList*>((QObject*)o); }
void* QTextList_new(void* doc) {
	return new QTextList((QTextDocument*)doc);
}
char* QTextList_tr(const char* s, long long* __out_len) {
	auto&& __r = QTextList::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QTextList_count(void* self) {
	auto&& __r = ((QTextList*)self)->count();
	return __r;
}
void* QTextList_item(void* self, int i) {
	auto&& __r = ((QTextList*)self)->item(i);
	return new QTextBlock(__r);
}
int QTextList_itemNumber(void* self, void* arg1) {
	auto&& __r = ((QTextList*)self)->itemNumber(*(QTextBlock*)arg1);
	return __r;
}
char* QTextList_itemText(void* self, void* arg1, long long* __out_len) {
	auto&& __r = ((QTextList*)self)->itemText(*(QTextBlock*)arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextList_removeItem(void* self, int i) {
	((QTextList*)self)->removeItem(i);
}
void QTextList_remove(void* self, void* arg1) {
	((QTextList*)self)->remove(*(QTextBlock*)arg1);
}
void QTextList_add(void* self, void* block) {
	((QTextList*)self)->add(*(QTextBlock*)block);
}
void QTextList_setFormat(void* self, void* format) {
	((QTextList*)self)->setFormat(*(QTextListFormat*)format);
}
void* QTextList_format(void* self) {
	auto&& __r = ((QTextList*)self)->format();
	return new QTextListFormat(__r);
}
void* QTextList_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTextList::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTextList_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTextList::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QTextList_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTextList::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTextList_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTextList::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTextList_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTextList::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTextList_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTextList::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QTextList_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextList*)self, &QTextList::destroyed, (QTextList*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextList);
	__conns_QTextList[__id] = __c;
}
void QTextList_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextList*)self, &QTextList::destroyed, (QTextList*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextList);
	__conns_QTextList[__id] = __c;
}
void QTextList_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextList*)self, &QTextList::objectNameChanged, (QTextList*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextList);
	__conns_QTextList[__id] = __c;
}
void QTextList__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTextList);
		auto __it = __conns_QTextList.find(id);
		if (__it == __conns_QTextList.end()) return;
		__c = __it->second;
		__conns_QTextList.erase(__it);
	}
	QObject::disconnect(__c);
}
void QTextList_delete(void* self) { delete (QTextList*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtGui/qcolor.h>
#include "gen_qstyleoptionfocusrect.h"

extern "C" {

void* QStyleOptionFocusRect_base(void* self) { return (void*)static_cast<QStyleOption*>((QStyleOptionFocusRect*)self); }
void* QStyleOptionFocusRect_new() {
	return new QStyleOptionFocusRect();
}
void* QStyleOptionFocusRect_new2(void* other) {
	return new QStyleOptionFocusRect(*(QStyleOptionFocusRect*)other);
}
void* QStyleOptionFocusRect_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionFocusRect*)self)->operator=(*(QStyleOptionFocusRect*)arg1);
	return (void*)&(__r);
}
void* QStyleOptionFocusRect_backgroundColor_get(void* self) {
	auto __r = ((QStyleOptionFocusRect*)self)->backgroundColor;
	return new QColor(__r);
}
void QStyleOptionFocusRect_backgroundColor_set(void* self, void* backgroundColor) {
	((QStyleOptionFocusRect*)self)->backgroundColor = *(QColor*)backgroundColor;
}
void QStyleOptionFocusRect_delete(void* self) { delete (QStyleOptionFocusRect*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtGui/qtransform.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qrect.h>
#include "gen_qstyleoptiongraphicsitem.h"

extern "C" {

void* QStyleOptionGraphicsItem_base(void* self) { return (void*)static_cast<QStyleOption*>((QStyleOptionGraphicsItem*)self); }
void* QStyleOptionGraphicsItem_new() {
	return new QStyleOptionGraphicsItem();
}
void* QStyleOptionGraphicsItem_new2(void* other) {
	return new QStyleOptionGraphicsItem(*(QStyleOptionGraphicsItem*)other);
}
void* QStyleOptionGraphicsItem_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionGraphicsItem*)self)->operator=(*(QStyleOptionGraphicsItem*)arg1);
	return (void*)&(__r);
}
double QStyleOptionGraphicsItem_levelOfDetailFromTransform(void* worldTransform) {
	auto&& __r = QStyleOptionGraphicsItem::levelOfDetailFromTransform(*(QTransform*)worldTransform);
	return __r;
}
void* QStyleOptionGraphicsItem_exposedRect_get(void* self) {
	auto __r = ((QStyleOptionGraphicsItem*)self)->exposedRect;
	return new QRectF(__r);
}
void QStyleOptionGraphicsItem_exposedRect_set(void* self, void* exposedRect) {
	((QStyleOptionGraphicsItem*)self)->exposedRect = *(QRectF*)exposedRect;
}
void QStyleOptionGraphicsItem_delete(void* self) { delete (QStyleOptionGraphicsItem*)self; }

}

#include <QtWidgets/qcolormap.h>
#include <QtGui/qcolor.h>
#include "gen_qcolormap.h"

extern "C" {

void* QColormap_new(void* colormap) {
	return new QColormap(*(QColormap*)colormap);
}
void QColormap_initialize() {
	QColormap::initialize();
}
void QColormap_cleanup() {
	QColormap::cleanup();
}
void* QColormap_instance() {
	auto&& __r = QColormap::instance();
	return new QColormap(__r);
}
void* QColormap_OperatorAssign(void* self, void* colormap) {
	auto&& __r = ((QColormap*)self)->operator=(*(QColormap*)colormap);
	return (void*)&(__r);
}
int QColormap_mode(void* self) {
	auto&& __r = ((QColormap*)self)->mode();
	return (int)(__r);
}
int QColormap_depth(void* self) {
	auto&& __r = ((QColormap*)self)->depth();
	return __r;
}
int QColormap_size(void* self) {
	auto&& __r = ((QColormap*)self)->size();
	return __r;
}
unsigned int QColormap_pixel(void* self, void* color) {
	auto&& __r = ((QColormap*)self)->pixel(*(QColor*)color);
	return __r;
}
void* QColormap_colorAt(void* self, unsigned int pixel) {
	auto&& __r = ((QColormap*)self)->colorAt(pixel);
	return new QColor(__r);
}
void** QColormap_colormap(void* self, long long* __out_len) {
	auto&& __r = ((QColormap*)self)->colormap();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QColor(__e);
	return __arr;
}
void* QColormap_instance2(int screen) {
	auto&& __r = QColormap::instance(screen);
	return new QColormap(__r);
}
void QColormap_delete(void* self) { delete (QColormap*)self; }

}

#include <QtWidgets/qgesture.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qpangesture.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QPanGesture__Go : public QPanGesture {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QPanGesture__Go(long long id, A&&... a) : QPanGesture(std::forward<A>(a)...), __goId(id) {}
	using QPanGesture::sender;
	using QPanGesture::senderSignalIndex;
	using QPanGesture::receivers;
	using QPanGesture::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QPanGesture::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QPanGesture::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QPanGesture::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QPanGesture::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QPanGesture::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QPanGesture::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QPanGesture::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QPanGesture::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QPanGesture::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QPanGesture::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QPanGesture::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QPanGesture::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QPanGesture::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QPanGesture;
static std::mutex __connsMu_QPanGesture;

extern "C" {

void* QPanGesture_base(void* self) { return (void*)static_cast<QGesture*>((QPanGesture*)self); }
void QPanGesture_chain(void* self, void** out) {
	QPanGesture* __p = (QPanGesture*)self;
	QGesture* __b0 = static_cast<QGesture*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QPanGesture__set_override(void* self, int idx) { ((QPanGesture__Go*)self)->__goOn[idx] = 1; }
void* QPanGesture_fromQObject(void* o) { return (void*)qobject_cast<QPanGesture*>((QObject*)o); }
void* QPanGesture_new(long long __goid) {
	return new QPanGesture__Go(__goid);
}
void* QPanGesture_new2(long long __goid, void* parent) {
	return new QPanGesture__Go(__goid, (QObject*)parent);
}
char* QPanGesture_tr(const char* s, long long* __out_len) {
	auto&& __r = QPanGesture::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QPanGesture_lastOffset(void* self) {
	auto&& __r = ((QPanGesture*)self)->lastOffset();
	return new QPointF(__r);
}
void* QPanGesture_offset(void* self) {
	auto&& __r = ((QPanGesture*)self)->offset();
	return new QPointF(__r);
}
void* QPanGesture_delta(void* self) {
	auto&& __r = ((QPanGesture*)self)->delta();
	return new QPointF(__r);
}
double QPanGesture_acceleration(void* self) {
	auto&& __r = ((QPanGesture*)self)->acceleration();
	return __r;
}
void QPanGesture_setLastOffset(void* self, void* value) {
	((QPanGesture*)self)->setLastOffset(*(QPointF*)value);
}
void QPanGesture_setOffset(void* self, void* value) {
	((QPanGesture*)self)->setOffset(*(QPointF*)value);
}
void QPanGesture_setAcceleration(void* self, double value) {
	((QPanGesture*)self)->setAcceleration(value);
}
void* QPanGesture_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QPanGesture::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QPanGesture_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QPanGesture::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QPanGesture_sender(void* self) {
	auto&& __r = ((QPanGesture__Go*)self)->sender();
	return (void*)(__r);
}
int QPanGesture_senderSignalIndex(void* self) {
	auto&& __r = ((QPanGesture__Go*)self)->senderSignalIndex();
	return __r;
}
int QPanGesture_receivers(void* self, const char* signal) {
	auto&& __r = ((QPanGesture__Go*)self)->receivers(signal);
	return __r;
}
int QPanGesture_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QPanGesture__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QPanGesture_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QPanGesture::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QPanGesture_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QPanGesture::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QPanGesture_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QPanGesture::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QPanGesture_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QPanGesture::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QPanGesture_metaObject_default(void* self) {
	auto&& __r = ((QPanGesture*)self)->QPanGesture::metaObject();
	return (void*)(__r);
}
int QPanGesture_event_default(void* self, void* event) {
	auto&& __r = ((QPanGesture*)self)->QPanGesture::event((QEvent*)event);
	return __r;
}
int QPanGesture_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QPanGesture*)self)->QPanGesture::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QPanGesture_timerEvent_default(void* self, void* event) {
	((QPanGesture__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QPanGesture_childEvent_default(void* self, void* event) {
	((QPanGesture__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QPanGesture_customEvent_default(void* self, void* event) {
	((QPanGesture__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QPanGesture_connectNotify_default(void* self, void* signal) {
	((QPanGesture__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QPanGesture_disconnectNotify_default(void* self, void* signal) {
	((QPanGesture__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QPanGesture_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPanGesture*)self, &QPanGesture::destroyed, (QPanGesture*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPanGesture);
	__conns_QPanGesture[__id] = __c;
}
void QPanGesture_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPanGesture*)self, &QPanGesture::destroyed, (QPanGesture*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPanGesture);
	__conns_QPanGesture[__id] = __c;
}
void QPanGesture_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPanGesture*)self, &QPanGesture::objectNameChanged, (QPanGesture*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPanGesture);
	__conns_QPanGesture[__id] = __c;
}
void QPanGesture__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QPanGesture);
		auto __it = __conns_QPanGesture.find(id);
		if (__it == __conns_QPanGesture.end()) return;
		__c = __it->second;
		__conns_QPanGesture.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QPanGesture__Go_new(long long id, void* parent) { return new QPanGesture__Go(id, (QObject*)parent); }
void QPanGesture_delete(void* self) { delete (QPanGesture*)self; }

}

#include <QtWidgets/qgraphicseffect.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtGui/qcolor.h>
#include <QtGui/qpainter.h>
#include <QtCore/qrect.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qpixmap.h>
#include <QtCore/qpoint.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qgraphicscolorizeeffect.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGraphicsColorizeEffect__Go : public QGraphicsColorizeEffect {
public:
	long long __goId;
	unsigned char __goOn[11] = {};
	template<typename... A> QGraphicsColorizeEffect__Go(long long id, A&&... a) : QGraphicsColorizeEffect(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsColorizeEffect::updateBoundingRect;
	using QGraphicsColorizeEffect::sourceIsPixmap;
	using QGraphicsColorizeEffect::sourceBoundingRect;
	using QGraphicsColorizeEffect::drawSource;
	using QGraphicsColorizeEffect::sourcePixmap;
	using QGraphicsColorizeEffect::sender;
	using QGraphicsColorizeEffect::senderSignalIndex;
	using QGraphicsColorizeEffect::receivers;
	using QGraphicsColorizeEffect::isSignalConnected;
	void cppgoBase_draw(QPainter* __a0) { QGraphicsColorizeEffect::draw(__a0); }
	void cppgoBase_sourceChanged(QFlags<QGraphicsEffect::ChangeFlag> __a0) { QGraphicsColorizeEffect::sourceChanged(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGraphicsColorizeEffect::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QGraphicsColorizeEffect::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGraphicsColorizeEffect::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGraphicsColorizeEffect::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGraphicsColorizeEffect::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGraphicsColorizeEffect::metaObject();
	}
	void draw(QPainter* painter) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1044652152830004333LL, (void*)painter, &__h); if (__h) return; }
		QGraphicsColorizeEffect::draw(painter);
	}
	QRectF boundingRectFor(const QRectF& sourceRect) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 5831744872130794138LL, (void*)&sourceRect, &__h); if (__h) return *(QRectF*)__r; }
		return QGraphicsColorizeEffect::boundingRectFor(sourceRect);
	}
	void sourceChanged(QFlags<QGraphicsEffect::ChangeFlag> flags) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4520368087989481956LL, (int)(flags), &__h); if (__h) return; }
		QGraphicsColorizeEffect::sourceChanged(flags);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsColorizeEffect::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsColorizeEffect::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGraphicsColorizeEffect::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QGraphicsColorizeEffect::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGraphicsColorizeEffect::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsColorizeEffect::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsColorizeEffect::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGraphicsColorizeEffect;
static std::mutex __connsMu_QGraphicsColorizeEffect;

extern "C" {

void* QGraphicsColorizeEffect_base(void* self) { return (void*)static_cast<QGraphicsEffect*>((QGraphicsColorizeEffect*)self); }
void QGraphicsColorizeEffect_chain(void* self, void** out) {
	QGraphicsColorizeEffect* __p = (QGraphicsColorizeEffect*)self;
	QGraphicsEffect* __b0 = static_cast<QGraphicsEffect*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsColorizeEffect__set_override(void* self, int idx) { ((QGraphicsColorizeEffect__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsColorizeEffect_fromQObject(void* o) { return (void*)qobject_cast<QGraphicsColorizeEffect*>((QObject*)o); }
void* QGraphicsColorizeEffect_new(long long __goid) {
	return new QGraphicsColorizeEffect__Go(__goid);
}
void* QGraphicsColorizeEffect_new2(long long __goid, void* parent) {
	return new QGraphicsColorizeEffect__Go(__goid, (QObject*)parent);
}
char* QGraphicsColorizeEffect_tr(const char* s, long long* __out_len) {
	auto&& __r = QGraphicsColorizeEffect::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsColorizeEffect_color(void* self) {
	auto&& __r = ((QGraphicsColorizeEffect*)self)->color();
	return new QColor(__r);
}
double QGraphicsColorizeEffect_strength(void* self) {
	auto&& __r = ((QGraphicsColorizeEffect*)self)->strength();
	return __r;
}
void QGraphicsColorizeEffect_setColor(void* self, void* c) {
	((QGraphicsColorizeEffect*)self)->setColor(*(QColor*)c);
}
void QGraphicsColorizeEffect_setStrength(void* self, double strength) {
	((QGraphicsColorizeEffect*)self)->setStrength(strength);
}
void QGraphicsColorizeEffect_colorChanged(void* self, void* color) {
	((QGraphicsColorizeEffect*)self)->colorChanged(*(QColor*)color);
}
void QGraphicsColorizeEffect_strengthChanged(void* self, double strength) {
	((QGraphicsColorizeEffect*)self)->strengthChanged(strength);
}
void QGraphicsColorizeEffect_updateBoundingRect(void* self) {
	((QGraphicsColorizeEffect__Go*)self)->updateBoundingRect();
}
int QGraphicsColorizeEffect_sourceIsPixmap(void* self) {
	auto&& __r = ((QGraphicsColorizeEffect__Go*)self)->sourceIsPixmap();
	return __r;
}
void* QGraphicsColorizeEffect_sourceBoundingRect(void* self) {
	auto&& __r = ((QGraphicsColorizeEffect__Go*)self)->sourceBoundingRect();
	return new QRectF(__r);
}
void QGraphicsColorizeEffect_drawSource(void* self, void* painter) {
	((QGraphicsColorizeEffect__Go*)self)->drawSource((QPainter*)painter);
}
void* QGraphicsColorizeEffect_sourcePixmap(void* self) {
	auto&& __r = ((QGraphicsColorizeEffect__Go*)self)->sourcePixmap();
	return new QPixmap(__r);
}
void* QGraphicsColorizeEffect_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGraphicsColorizeEffect::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsColorizeEffect_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGraphicsColorizeEffect::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsColorizeEffect_sender(void* self) {
	auto&& __r = ((QGraphicsColorizeEffect__Go*)self)->sender();
	return (void*)(__r);
}
int QGraphicsColorizeEffect_senderSignalIndex(void* self) {
	auto&& __r = ((QGraphicsColorizeEffect__Go*)self)->senderSignalIndex();
	return __r;
}
int QGraphicsColorizeEffect_receivers(void* self, const char* signal) {
	auto&& __r = ((QGraphicsColorizeEffect__Go*)self)->receivers(signal);
	return __r;
}
int QGraphicsColorizeEffect_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGraphicsColorizeEffect__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QGraphicsColorizeEffect_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGraphicsColorizeEffect::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGraphicsColorizeEffect_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGraphicsColorizeEffect::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsColorizeEffect_sourceBoundingRect2(void* self, int system) {
	auto&& __r = ((QGraphicsColorizeEffect__Go*)self)->sourceBoundingRect((Qt::CoordinateSystem)(system));
	return new QRectF(__r);
}
void* QGraphicsColorizeEffect_sourcePixmap2(void* self, int system) {
	auto&& __r = ((QGraphicsColorizeEffect__Go*)self)->sourcePixmap((Qt::CoordinateSystem)(system));
	return new QPixmap(__r);
}
void* QGraphicsColorizeEffect_sourcePixmap3(void* self, int system, void* offset) {
	auto&& __r = ((QGraphicsColorizeEffect__Go*)self)->sourcePixmap((Qt::CoordinateSystem)(system), (QPoint*)offset);
	return new QPixmap(__r);
}
void* QGraphicsColorizeEffect_sourcePixmap4(void* self, int system, void* offset, int mode) {
	auto&& __r = ((QGraphicsColorizeEffect__Go*)self)->sourcePixmap((Qt::CoordinateSystem)(system), (QPoint*)offset, (QGraphicsEffect::PixmapPadMode)(mode));
	return new QPixmap(__r);
}
void* QGraphicsColorizeEffect_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGraphicsColorizeEffect::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsColorizeEffect_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGraphicsColorizeEffect::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsColorizeEffect_metaObject_default(void* self) {
	auto&& __r = ((QGraphicsColorizeEffect*)self)->QGraphicsColorizeEffect::metaObject();
	return (void*)(__r);
}
void QGraphicsColorizeEffect_draw_default(void* self, void* painter) {
	((QGraphicsColorizeEffect__Go*)self)->cppgoBase_draw((QPainter*)painter);
}
void* QGraphicsColorizeEffect_boundingRectFor_default(void* self, void* sourceRect) {
	auto&& __r = ((QGraphicsColorizeEffect*)self)->QGraphicsColorizeEffect::boundingRectFor(*(QRectF*)sourceRect);
	return new QRectF(__r);
}
void QGraphicsColorizeEffect_sourceChanged_default(void* self, int flags) {
	((QGraphicsColorizeEffect__Go*)self)->cppgoBase_sourceChanged(QFlags<QGraphicsEffect::ChangeFlag>(flags));
}
int QGraphicsColorizeEffect_event_default(void* self, void* event) {
	auto&& __r = ((QGraphicsColorizeEffect*)self)->QGraphicsColorizeEffect::event((QEvent*)event);
	return __r;
}
int QGraphicsColorizeEffect_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsColorizeEffect*)self)->QGraphicsColorizeEffect::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QGraphicsColorizeEffect_timerEvent_default(void* self, void* event) {
	((QGraphicsColorizeEffect__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGraphicsColorizeEffect_childEvent_default(void* self, void* event) {
	((QGraphicsColorizeEffect__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QGraphicsColorizeEffect_customEvent_default(void* self, void* event) {
	((QGraphicsColorizeEffect__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGraphicsColorizeEffect_connectNotify_default(void* self, void* signal) {
	((QGraphicsColorizeEffect__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGraphicsColorizeEffect_disconnectNotify_default(void* self, void* signal) {
	((QGraphicsColorizeEffect__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QGraphicsColorizeEffect_connect_colorChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsColorizeEffect*)self, &QGraphicsColorizeEffect::colorChanged, (QGraphicsColorizeEffect*)self, [__id](const QColor& color){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&color);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsColorizeEffect);
	__conns_QGraphicsColorizeEffect[__id] = __c;
}
void QGraphicsColorizeEffect_connect_strengthChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsColorizeEffect*)self, &QGraphicsColorizeEffect::strengthChanged, (QGraphicsColorizeEffect*)self, [__id](qreal strength){
		if (cppgo_sig_void_unsignedlonglong) ((void(*)(long long, unsigned long long))cppgo_sig_void_unsignedlonglong)(__id, __builtin_bit_cast(unsigned long long, (double)(strength)));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsColorizeEffect);
	__conns_QGraphicsColorizeEffect[__id] = __c;
}
void QGraphicsColorizeEffect_connect_enabledChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsColorizeEffect*)self, &QGraphicsColorizeEffect::enabledChanged, (QGraphicsColorizeEffect*)self, [__id](bool enabled){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, enabled);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsColorizeEffect);
	__conns_QGraphicsColorizeEffect[__id] = __c;
}
void QGraphicsColorizeEffect_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsColorizeEffect*)self, &QGraphicsColorizeEffect::destroyed, (QGraphicsColorizeEffect*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsColorizeEffect);
	__conns_QGraphicsColorizeEffect[__id] = __c;
}
void QGraphicsColorizeEffect_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsColorizeEffect*)self, &QGraphicsColorizeEffect::destroyed, (QGraphicsColorizeEffect*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsColorizeEffect);
	__conns_QGraphicsColorizeEffect[__id] = __c;
}
void QGraphicsColorizeEffect_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsColorizeEffect*)self, &QGraphicsColorizeEffect::objectNameChanged, (QGraphicsColorizeEffect*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsColorizeEffect);
	__conns_QGraphicsColorizeEffect[__id] = __c;
}
void QGraphicsColorizeEffect__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsColorizeEffect);
		auto __it = __conns_QGraphicsColorizeEffect.find(id);
		if (__it == __conns_QGraphicsColorizeEffect.end()) return;
		__c = __it->second;
		__conns_QGraphicsColorizeEffect.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGraphicsColorizeEffect__Go_new(long long id, void* parent) { return new QGraphicsColorizeEffect__Go(id, (QObject*)parent); }
void QGraphicsColorizeEffect_delete(void* self) { delete (QGraphicsColorizeEffect*)self; }

}

#include <QtWidgets/qgroupbox.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qsize.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtGui/qregion.h>
#include <QtGui/qpalette.h>
#include <QtGui/qfont.h>
#include <QtGui/qfontmetrics.h>
#include <QtGui/qfontinfo.h>
#include <QtGui/qcursor.h>
#include <QtGui/qbitmap.h>
#include <QtGui/qpaintdevice.h>
#include <QtGui/qpainter.h>
#include <QtGui/qpixmap.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtGui/qicon.h>
#include <QtCore/qlocale.h>
#include <QtGui/qkeysequence.h>
#include <QtWidgets/qgraphicsproxywidget.h>
#include <QtCore/qbytearray.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qmargins.h>
#include <QtWidgets/qlayout.h>
#include <QtGui/qaction.h>
#include <QtCore/qobject.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtCore/qvariant.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qgroupbox.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGroupBox__Go : public QGroupBox {
public:
	long long __goId;
	unsigned char __goOn[49] = {};
	template<typename... A> QGroupBox__Go(long long id, A&&... a) : QGroupBox(std::forward<A>(a)...), __goId(id) {}
	using QGroupBox::updateMicroFocus;
	using QGroupBox::create;
	using QGroupBox::destroy;
	using QGroupBox::focusNextChild;
	using QGroupBox::focusPreviousChild;
	using QGroupBox::sender;
	using QGroupBox::senderSignalIndex;
	using QGroupBox::receivers;
	using QGroupBox::isSignalConnected;
	using QGroupBox::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QGroupBox::event(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QGroupBox::childEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QGroupBox::resizeEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QGroupBox::paintEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QGroupBox::focusInEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QGroupBox::changeEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QGroupBox::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QGroupBox::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QGroupBox::mouseReleaseEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionGroupBox* __a0) { QGroupBox::initStyleOption(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QGroupBox::mouseDoubleClickEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QGroupBox::wheelEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QGroupBox::keyPressEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QGroupBox::keyReleaseEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QGroupBox::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QGroupBox::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QGroupBox::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QGroupBox::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QGroupBox::closeEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QGroupBox::contextMenuEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QGroupBox::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QGroupBox::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QGroupBox::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QGroupBox::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QGroupBox::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QGroupBox::dropEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QGroupBox::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QGroupBox::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QGroupBox::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QGroupBox::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QGroupBox::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QGroupBox::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QGroupBox::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QGroupBox::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGroupBox::timerEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGroupBox::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGroupBox::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGroupBox::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGroupBox::metaObject();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QGroupBox::minimumSizeHint();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QGroupBox::event(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::childEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::resizeEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::paintEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::focusInEvent(event);
	}
	void changeEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::changeEvent(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::mousePressEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::mouseReleaseEvent(event);
	}
	void initStyleOption(QStyleOptionGroupBox* option) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QGroupBox::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QGroupBox::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QGroupBox::setVisible(visible);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QGroupBox::sizeHint();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QGroupBox::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QGroupBox::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QGroupBox::paintEngine();
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::mouseDoubleClickEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::wheelEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::keyPressEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::keyReleaseEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::closeEvent(event);
	}
	void contextMenuEvent(QContextMenuEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::contextMenuEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::dropEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QGroupBox::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QGroupBox::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QGroupBox::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QGroupBox::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QGroupBox::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QGroupBox::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QGroupBox::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QGroupBox::focusNextPrevChild(next);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGroupBox::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::timerEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGroupBox::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGroupBox::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGroupBox::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGroupBox;
static std::mutex __connsMu_QGroupBox;

extern "C" {

void* QGroupBox_base(void* self) { return (void*)static_cast<QWidget*>((QGroupBox*)self); }
void QGroupBox_chain(void* self, void** out) {
	QGroupBox* __p = (QGroupBox*)self;
	QWidget* __b0 = static_cast<QWidget*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QGroupBox__set_override(void* self, int idx) { ((QGroupBox__Go*)self)->__goOn[idx] = 1; }
void* QGroupBox_fromQObject(void* o) { return (void*)qobject_cast<QGroupBox*>((QObject*)o); }
void* QGroupBox_new(long long __goid, void* parent) {
	return new QGroupBox__Go(__goid, (QWidget*)parent);
}
void* QGroupBox_new2(long long __goid) {
	return new QGroupBox__Go(__goid);
}
void* QGroupBox_new3(long long __goid, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	return new QGroupBox__Go(__goid, title_qstr);
}
void* QGroupBox_new4(long long __goid, const char* title, long long title_len, void* parent) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	return new QGroupBox__Go(__goid, title_qstr, (QWidget*)parent);
}
char* QGroupBox_tr(const char* s, long long* __out_len) {
	auto&& __r = QGroupBox::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGroupBox_title(void* self, long long* __out_len) {
	auto&& __r = ((QGroupBox*)self)->title();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGroupBox_setTitle(void* self, const char* title, long long title_len) {
	QString title_qstr = QString::fromUtf8(title, title_len);
	((QGroupBox*)self)->setTitle(title_qstr);
}
int QGroupBox_alignment(void* self) {
	auto&& __r = ((QGroupBox*)self)->alignment();
	return (int)(__r);
}
void QGroupBox_setAlignment(void* self, int alignment) {
	((QGroupBox*)self)->setAlignment(alignment);
}
int QGroupBox_isFlat(void* self) {
	auto&& __r = ((QGroupBox*)self)->isFlat();
	return __r;
}
void QGroupBox_setFlat(void* self, int flat) {
	((QGroupBox*)self)->setFlat(flat);
}
int QGroupBox_isCheckable(void* self) {
	auto&& __r = ((QGroupBox*)self)->isCheckable();
	return __r;
}
void QGroupBox_setCheckable(void* self, int checkable) {
	((QGroupBox*)self)->setCheckable(checkable);
}
int QGroupBox_isChecked(void* self) {
	auto&& __r = ((QGroupBox*)self)->isChecked();
	return __r;
}
void QGroupBox_setChecked(void* self, int checked) {
	((QGroupBox*)self)->setChecked(checked);
}
void QGroupBox_clicked(void* self) {
	((QGroupBox*)self)->clicked();
}
void QGroupBox_toggled(void* self, int arg1) {
	((QGroupBox*)self)->toggled(arg1);
}
void QGroupBox_setTabOrder(void* arg1, void* arg2) {
	QGroupBox::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QGroupBox_mouseGrabber() {
	auto&& __r = QGroupBox::mouseGrabber();
	return (void*)(__r);
}
void* QGroupBox_keyboardGrabber() {
	auto&& __r = QGroupBox::keyboardGrabber();
	return (void*)(__r);
}
void* QGroupBox_find(unsigned long long arg1) {
	auto&& __r = QGroupBox::find(arg1);
	return (void*)(__r);
}
void* QGroupBox_createWindowContainer(void* window) {
	auto&& __r = QGroupBox::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QGroupBox_updateMicroFocus(void* self) {
	((QGroupBox__Go*)self)->updateMicroFocus();
}
void QGroupBox_create(void* self) {
	((QGroupBox__Go*)self)->create();
}
void QGroupBox_destroy(void* self) {
	((QGroupBox__Go*)self)->destroy();
}
int QGroupBox_focusNextChild(void* self) {
	auto&& __r = ((QGroupBox__Go*)self)->focusNextChild();
	return __r;
}
int QGroupBox_focusPreviousChild(void* self) {
	auto&& __r = ((QGroupBox__Go*)self)->focusPreviousChild();
	return __r;
}
void* QGroupBox_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGroupBox::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGroupBox_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGroupBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGroupBox_sender(void* self) {
	auto&& __r = ((QGroupBox__Go*)self)->sender();
	return (void*)(__r);
}
int QGroupBox_senderSignalIndex(void* self) {
	auto&& __r = ((QGroupBox__Go*)self)->senderSignalIndex();
	return __r;
}
int QGroupBox_receivers(void* self, const char* signal) {
	auto&& __r = ((QGroupBox__Go*)self)->receivers(signal);
	return __r;
}
int QGroupBox_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGroupBox__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QGroupBox_devicePixelRatioFScale() {
	auto&& __r = QGroupBox::devicePixelRatioFScale();
	return __r;
}
int QGroupBox_encodeMetricF(int metric, double value) {
	auto&& __r = QGroupBox::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QGroupBox_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QGroupBox__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QGroupBox_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGroupBox::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGroupBox_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGroupBox::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QGroupBox_clicked2(void* self, int checked) {
	((QGroupBox*)self)->clicked(checked);
}
void* QGroupBox_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QGroupBox::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QGroupBox_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QGroupBox::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QGroupBox_updateMicroFocus2(void* self, int query) {
	((QGroupBox__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QGroupBox_create2(void* self, unsigned long long arg1) {
	((QGroupBox__Go*)self)->create(arg1);
}
void QGroupBox_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QGroupBox__Go*)self)->create(arg1, initializeWindow);
}
void QGroupBox_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QGroupBox__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QGroupBox_destroy2(void* self, int destroyWindow) {
	((QGroupBox__Go*)self)->destroy(destroyWindow);
}
void QGroupBox_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QGroupBox__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QGroupBox_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGroupBox::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGroupBox_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGroupBox::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QGroupBox_metaObject_default(void* self) {
	auto&& __r = ((QGroupBox*)self)->QGroupBox::metaObject();
	return (void*)(__r);
}
void* QGroupBox_minimumSizeHint_default(void* self) {
	auto&& __r = ((QGroupBox*)self)->QGroupBox::minimumSizeHint();
	return new QSize(__r);
}
int QGroupBox_event_default(void* self, void* event) {
	auto&& __r = ((QGroupBox__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QGroupBox_childEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QGroupBox_resizeEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QGroupBox_paintEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QGroupBox_focusInEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QGroupBox_changeEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_changeEvent((QEvent*)event);
}
void QGroupBox_mousePressEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QGroupBox_mouseMoveEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QGroupBox_mouseReleaseEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QGroupBox_initStyleOption_default(void* self, void* option) {
	((QGroupBox__Go*)self)->cppgoBase_initStyleOption((QStyleOptionGroupBox*)option);
}
int QGroupBox_devType_default(void* self) {
	auto&& __r = ((QGroupBox*)self)->QGroupBox::devType();
	return __r;
}
void QGroupBox_setVisible_default(void* self, int visible) {
	((QGroupBox*)self)->QGroupBox::setVisible(visible);
}
void* QGroupBox_sizeHint_default(void* self) {
	auto&& __r = ((QGroupBox*)self)->QGroupBox::sizeHint();
	return new QSize(__r);
}
int QGroupBox_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QGroupBox*)self)->QGroupBox::heightForWidth(arg1);
	return __r;
}
int QGroupBox_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QGroupBox*)self)->QGroupBox::hasHeightForWidth();
	return __r;
}
void* QGroupBox_paintEngine_default(void* self) {
	auto&& __r = ((QGroupBox*)self)->QGroupBox::paintEngine();
	return (void*)(__r);
}
void QGroupBox_mouseDoubleClickEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QGroupBox_wheelEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QGroupBox_keyPressEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QGroupBox_keyReleaseEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QGroupBox_focusOutEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QGroupBox_enterEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QGroupBox_leaveEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QGroupBox_moveEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QGroupBox_closeEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QGroupBox_contextMenuEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)event);
}
void QGroupBox_tabletEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QGroupBox_actionEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QGroupBox_dragEnterEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QGroupBox_dragMoveEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QGroupBox_dragLeaveEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QGroupBox_dropEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QGroupBox_showEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QGroupBox_hideEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QGroupBox_metric_default(void* self, int arg1) {
	auto&& __r = ((QGroupBox__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QGroupBox_initPainter_default(void* self, void* painter) {
	((QGroupBox__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QGroupBox_redirected_default(void* self, void* offset) {
	auto&& __r = ((QGroupBox__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QGroupBox_sharedPainter_default(void* self) {
	auto&& __r = ((QGroupBox__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QGroupBox_inputMethodEvent_default(void* self, void* arg1) {
	((QGroupBox__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QGroupBox_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QGroupBox*)self)->QGroupBox::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QGroupBox_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QGroupBox__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QGroupBox_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGroupBox*)self)->QGroupBox::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QGroupBox_timerEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGroupBox_customEvent_default(void* self, void* event) {
	((QGroupBox__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGroupBox_connectNotify_default(void* self, void* signal) {
	((QGroupBox__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGroupBox_disconnectNotify_default(void* self, void* signal) {
	((QGroupBox__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QGroupBox_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGroupBox*)self, &QGroupBox::clicked, (QGroupBox*)self, [__id](bool checked){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGroupBox);
	__conns_QGroupBox[__id] = __c;
}
void QGroupBox_connect_clicked2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGroupBox*)self, &QGroupBox::clicked, (QGroupBox*)self, [__id](bool checked){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, checked);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGroupBox);
	__conns_QGroupBox[__id] = __c;
}
void QGroupBox_connect_toggled(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGroupBox*)self, &QGroupBox::toggled, (QGroupBox*)self, [__id](bool arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGroupBox);
	__conns_QGroupBox[__id] = __c;
}
void QGroupBox_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGroupBox*)self, &QGroupBox::windowTitleChanged, (QGroupBox*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGroupBox);
	__conns_QGroupBox[__id] = __c;
}
void QGroupBox_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGroupBox*)self, &QGroupBox::windowIconChanged, (QGroupBox*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGroupBox);
	__conns_QGroupBox[__id] = __c;
}
void QGroupBox_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGroupBox*)self, &QGroupBox::windowIconTextChanged, (QGroupBox*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGroupBox);
	__conns_QGroupBox[__id] = __c;
}
void QGroupBox_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGroupBox*)self, &QGroupBox::customContextMenuRequested, (QGroupBox*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGroupBox);
	__conns_QGroupBox[__id] = __c;
}
void QGroupBox_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGroupBox*)self, &QGroupBox::destroyed, (QGroupBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGroupBox);
	__conns_QGroupBox[__id] = __c;
}
void QGroupBox_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGroupBox*)self, &QGroupBox::destroyed, (QGroupBox*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGroupBox);
	__conns_QGroupBox[__id] = __c;
}
void QGroupBox_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGroupBox*)self, &QGroupBox::objectNameChanged, (QGroupBox*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGroupBox);
	__conns_QGroupBox[__id] = __c;
}
void QGroupBox__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGroupBox);
		auto __it = __conns_QGroupBox.find(id);
		if (__it == __conns_QGroupBox.end()) return;
		__c = __it->second;
		__conns_QGroupBox.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGroupBox__Go_new(long long id) { return new QGroupBox__Go(id); }
void QGroupBox_delete(void* self) { delete (QGroupBox*)self; }

}

#include <QtWidgets/qprogressdialog.h>
#include <QtWidgets/qdialog.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtWidgets/qlabel.h>
#include <QtWidgets/qpushbutton.h>
#include <QtWidgets/qprogressbar.h>
#include <QtCore/qsize.h>
#include <QtCore/qobject.h>
#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtGui/qregion.h>
#include <QtGui/qpalette.h>
#include <QtGui/qfont.h>
#include <QtGui/qfontmetrics.h>
#include <QtGui/qfontinfo.h>
#include <QtGui/qcursor.h>
#include <QtGui/qbitmap.h>
#include <QtGui/qpaintdevice.h>
#include <QtGui/qpainter.h>
#include <QtGui/qpixmap.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtGui/qicon.h>
#include <QtCore/qlocale.h>
#include <QtGui/qkeysequence.h>
#include <QtWidgets/qgraphicsproxywidget.h>
#include <QtCore/qbytearray.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qmargins.h>
#include <QtWidgets/qlayout.h>
#include <QtGui/qaction.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtCore/qvariant.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qprogressdialog.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QProgressDialog__Go : public QProgressDialog {
public:
	long long __goId;
	unsigned char __goOn[52] = {};
	template<typename... A> QProgressDialog__Go(long long id, A&&... a) : QProgressDialog(std::forward<A>(a)...), __goId(id) {}
	using QProgressDialog::forceShow;
	using QProgressDialog::adjustPosition;
	using QProgressDialog::updateMicroFocus;
	using QProgressDialog::create;
	using QProgressDialog::destroy;
	using QProgressDialog::focusNextChild;
	using QProgressDialog::focusPreviousChild;
	using QProgressDialog::sender;
	using QProgressDialog::senderSignalIndex;
	using QProgressDialog::receivers;
	using QProgressDialog::isSignalConnected;
	using QProgressDialog::getDecodedMetricF;
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QProgressDialog::resizeEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QProgressDialog::closeEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QProgressDialog::changeEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QProgressDialog::showEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QProgressDialog::keyPressEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QProgressDialog::contextMenuEvent(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QProgressDialog::eventFilter(__a0, __a1); }
	bool cppgoBase_event(QEvent* __a0) { return QProgressDialog::event(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QProgressDialog::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QProgressDialog::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QProgressDialog::mouseDoubleClickEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QProgressDialog::mouseMoveEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QProgressDialog::wheelEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QProgressDialog::keyReleaseEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QProgressDialog::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QProgressDialog::focusOutEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QProgressDialog::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QProgressDialog::leaveEvent(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QProgressDialog::paintEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QProgressDialog::moveEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QProgressDialog::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QProgressDialog::actionEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QProgressDialog::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QProgressDialog::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QProgressDialog::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QProgressDialog::dropEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QProgressDialog::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QProgressDialog::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QProgressDialog::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QProgressDialog::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QProgressDialog::sharedPainter(); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QProgressDialog::inputMethodEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QProgressDialog::focusNextPrevChild(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QProgressDialog::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QProgressDialog::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QProgressDialog::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QProgressDialog::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QProgressDialog::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QProgressDialog::metaObject();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QProgressDialog::sizeHint();
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::resizeEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::closeEvent(event);
	}
	void changeEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::changeEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::showEvent(event);
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QProgressDialog::setVisible(visible);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QProgressDialog::minimumSizeHint();
	}
	int exec() override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1745322561779819684LL, &__h); if (__h) return __r; }
		return QProgressDialog::exec();
	}
	void done(int arg1) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -2570997490807640431LL, arg1, &__h); if (__h) return; }
		QProgressDialog::done(arg1);
	}
	void accept() override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9106131928185084681LL, &__h); if (__h) return; }
		QProgressDialog::accept();
	}
	void reject() override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 2875220963396170150LL, &__h); if (__h) return; }
		QProgressDialog::reject();
	}
	void keyPressEvent(QKeyEvent* arg1) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)arg1, &__h); if (__h) return; }
		QProgressDialog::keyPressEvent(arg1);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QProgressDialog::contextMenuEvent(arg1);
	}
	bool eventFilter(QObject* arg1, QEvent* arg2) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)arg1, (void*)arg2, &__h); if (__h) return __r; }
		return QProgressDialog::eventFilter(arg1, arg2);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QProgressDialog::devType();
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QProgressDialog::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QProgressDialog::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QProgressDialog::paintEngine();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QProgressDialog::event(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::mousePressEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::mouseDoubleClickEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::mouseMoveEvent(event);
	}
	void wheelEvent(QWheelEvent* event) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::wheelEvent(event);
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::keyReleaseEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::focusOutEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::leaveEvent(event);
	}
	void paintEvent(QPaintEvent* event) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::paintEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::moveEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::actionEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::dropEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QProgressDialog::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QProgressDialog::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QProgressDialog::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QProgressDialog::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QProgressDialog::sharedPainter();
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QProgressDialog::inputMethodEvent(arg1);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery arg1) const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(arg1), &__h); if (__h) return *(QVariant*)__r; }
		return QProgressDialog::inputMethodQuery(arg1);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QProgressDialog::focusNextPrevChild(next);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QProgressDialog::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QProgressDialog::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QProgressDialog::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QProgressDialog;
static std::mutex __connsMu_QProgressDialog;

extern "C" {

void* QProgressDialog_base(void* self) { return (void*)static_cast<QDialog*>((QProgressDialog*)self); }
void QProgressDialog_chain(void* self, void** out) {
	QProgressDialog* __p = (QProgressDialog*)self;
	QDialog* __b0 = static_cast<QDialog*>(__p); out[0] = (void*)__b0;
	QWidget* __b1 = static_cast<QWidget*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QProgressDialog__set_override(void* self, int idx) { ((QProgressDialog__Go*)self)->__goOn[idx] = 1; }
void* QProgressDialog_fromQObject(void* o) { return (void*)qobject_cast<QProgressDialog*>((QObject*)o); }
void* QProgressDialog_new(long long __goid, void* parent) {
	return new QProgressDialog__Go(__goid, (QWidget*)parent);
}
void* QProgressDialog_new2(long long __goid) {
	return new QProgressDialog__Go(__goid);
}
void* QProgressDialog_new3(long long __goid, const char* labelText, long long labelText_len, const char* cancelButtonText, long long cancelButtonText_len, int minimum, int maximum) {
	QString labelText_qstr = QString::fromUtf8(labelText, labelText_len);
	QString cancelButtonText_qstr = QString::fromUtf8(cancelButtonText, cancelButtonText_len);
	return new QProgressDialog__Go(__goid, labelText_qstr, cancelButtonText_qstr, minimum, maximum);
}
void* QProgressDialog_new4(long long __goid, void* parent, int flags) {
	return new QProgressDialog__Go(__goid, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
}
void* QProgressDialog_new5(long long __goid, const char* labelText, long long labelText_len, const char* cancelButtonText, long long cancelButtonText_len, int minimum, int maximum, void* parent) {
	QString labelText_qstr = QString::fromUtf8(labelText, labelText_len);
	QString cancelButtonText_qstr = QString::fromUtf8(cancelButtonText, cancelButtonText_len);
	return new QProgressDialog__Go(__goid, labelText_qstr, cancelButtonText_qstr, minimum, maximum, (QWidget*)parent);
}
void* QProgressDialog_new6(long long __goid, const char* labelText, long long labelText_len, const char* cancelButtonText, long long cancelButtonText_len, int minimum, int maximum, void* parent, int flags) {
	QString labelText_qstr = QString::fromUtf8(labelText, labelText_len);
	QString cancelButtonText_qstr = QString::fromUtf8(cancelButtonText, cancelButtonText_len);
	return new QProgressDialog__Go(__goid, labelText_qstr, cancelButtonText_qstr, minimum, maximum, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
}
char* QProgressDialog_tr(const char* s, long long* __out_len) {
	auto&& __r = QProgressDialog::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QProgressDialog_setLabel(void* self, void* label) {
	((QProgressDialog*)self)->setLabel((QLabel*)label);
}
void QProgressDialog_setCancelButton(void* self, void* button) {
	((QProgressDialog*)self)->setCancelButton((QPushButton*)button);
}
void QProgressDialog_setBar(void* self, void* bar) {
	((QProgressDialog*)self)->setBar((QProgressBar*)bar);
}
int QProgressDialog_wasCanceled(void* self) {
	auto&& __r = ((QProgressDialog*)self)->wasCanceled();
	return __r;
}
int QProgressDialog_minimum(void* self) {
	auto&& __r = ((QProgressDialog*)self)->minimum();
	return __r;
}
int QProgressDialog_maximum(void* self) {
	auto&& __r = ((QProgressDialog*)self)->maximum();
	return __r;
}
int QProgressDialog_value(void* self) {
	auto&& __r = ((QProgressDialog*)self)->value();
	return __r;
}
char* QProgressDialog_labelText(void* self, long long* __out_len) {
	auto&& __r = ((QProgressDialog*)self)->labelText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QProgressDialog_minimumDuration(void* self) {
	auto&& __r = ((QProgressDialog*)self)->minimumDuration();
	return __r;
}
void QProgressDialog_setAutoReset(void* self, int reset) {
	((QProgressDialog*)self)->setAutoReset(reset);
}
int QProgressDialog_autoReset(void* self) {
	auto&& __r = ((QProgressDialog*)self)->autoReset();
	return __r;
}
void QProgressDialog_setAutoClose(void* self, int close_) {
	((QProgressDialog*)self)->setAutoClose(close_);
}
int QProgressDialog_autoClose(void* self) {
	auto&& __r = ((QProgressDialog*)self)->autoClose();
	return __r;
}
void QProgressDialog_open(void* self, void* receiver, const char* member) {
	((QProgressDialog*)self)->open((QObject*)receiver, member);
}
void QProgressDialog_cancel(void* self) {
	((QProgressDialog*)self)->cancel();
}
void QProgressDialog_reset(void* self) {
	((QProgressDialog*)self)->reset();
}
void QProgressDialog_setMaximum(void* self, int maximum) {
	((QProgressDialog*)self)->setMaximum(maximum);
}
void QProgressDialog_setMinimum(void* self, int minimum) {
	((QProgressDialog*)self)->setMinimum(minimum);
}
void QProgressDialog_setRange(void* self, int minimum, int maximum) {
	((QProgressDialog*)self)->setRange(minimum, maximum);
}
void QProgressDialog_setValue(void* self, int progress) {
	((QProgressDialog*)self)->setValue(progress);
}
void QProgressDialog_setLabelText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QProgressDialog*)self)->setLabelText(text_qstr);
}
void QProgressDialog_setCancelButtonText(void* self, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QProgressDialog*)self)->setCancelButtonText(text_qstr);
}
void QProgressDialog_setMinimumDuration(void* self, int ms) {
	((QProgressDialog*)self)->setMinimumDuration(ms);
}
void QProgressDialog_canceled(void* self) {
	((QProgressDialog*)self)->canceled();
}
void QProgressDialog_forceShow(void* self) {
	((QProgressDialog__Go*)self)->forceShow();
}
void QProgressDialog_adjustPosition(void* self, void* arg1) {
	((QProgressDialog__Go*)self)->adjustPosition((QWidget*)arg1);
}
void QProgressDialog_setTabOrder(void* arg1, void* arg2) {
	QProgressDialog::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QProgressDialog_mouseGrabber() {
	auto&& __r = QProgressDialog::mouseGrabber();
	return (void*)(__r);
}
void* QProgressDialog_keyboardGrabber() {
	auto&& __r = QProgressDialog::keyboardGrabber();
	return (void*)(__r);
}
void* QProgressDialog_find(unsigned long long arg1) {
	auto&& __r = QProgressDialog::find(arg1);
	return (void*)(__r);
}
void* QProgressDialog_createWindowContainer(void* window) {
	auto&& __r = QProgressDialog::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QProgressDialog_updateMicroFocus(void* self) {
	((QProgressDialog__Go*)self)->updateMicroFocus();
}
void QProgressDialog_create(void* self) {
	((QProgressDialog__Go*)self)->create();
}
void QProgressDialog_destroy(void* self) {
	((QProgressDialog__Go*)self)->destroy();
}
int QProgressDialog_focusNextChild(void* self) {
	auto&& __r = ((QProgressDialog__Go*)self)->focusNextChild();
	return __r;
}
int QProgressDialog_focusPreviousChild(void* self) {
	auto&& __r = ((QProgressDialog__Go*)self)->focusPreviousChild();
	return __r;
}
void* QProgressDialog_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QProgressDialog::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QProgressDialog_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QProgressDialog::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QProgressDialog_sender(void* self) {
	auto&& __r = ((QProgressDialog__Go*)self)->sender();
	return (void*)(__r);
}
int QProgressDialog_senderSignalIndex(void* self) {
	auto&& __r = ((QProgressDialog__Go*)self)->senderSignalIndex();
	return __r;
}
int QProgressDialog_receivers(void* self, const char* signal) {
	auto&& __r = ((QProgressDialog__Go*)self)->receivers(signal);
	return __r;
}
int QProgressDialog_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QProgressDialog__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QProgressDialog_devicePixelRatioFScale() {
	auto&& __r = QProgressDialog::devicePixelRatioFScale();
	return __r;
}
int QProgressDialog_encodeMetricF(int metric, double value) {
	auto&& __r = QProgressDialog::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QProgressDialog_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QProgressDialog__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QProgressDialog_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QProgressDialog::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QProgressDialog_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QProgressDialog::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QProgressDialog_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QProgressDialog::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QProgressDialog_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QProgressDialog::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QProgressDialog_updateMicroFocus2(void* self, int query) {
	((QProgressDialog__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QProgressDialog_create2(void* self, unsigned long long arg1) {
	((QProgressDialog__Go*)self)->create(arg1);
}
void QProgressDialog_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QProgressDialog__Go*)self)->create(arg1, initializeWindow);
}
void QProgressDialog_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QProgressDialog__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QProgressDialog_destroy2(void* self, int destroyWindow) {
	((QProgressDialog__Go*)self)->destroy(destroyWindow);
}
void QProgressDialog_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QProgressDialog__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QProgressDialog_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QProgressDialog::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QProgressDialog_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QProgressDialog::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QProgressDialog_metaObject_default(void* self) {
	auto&& __r = ((QProgressDialog*)self)->QProgressDialog::metaObject();
	return (void*)(__r);
}
void* QProgressDialog_sizeHint_default(void* self) {
	auto&& __r = ((QProgressDialog*)self)->QProgressDialog::sizeHint();
	return new QSize(__r);
}
void QProgressDialog_resizeEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QProgressDialog_closeEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QProgressDialog_changeEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_changeEvent((QEvent*)event);
}
void QProgressDialog_showEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QProgressDialog_setVisible_default(void* self, int visible) {
	((QProgressDialog*)self)->QProgressDialog::setVisible(visible);
}
void* QProgressDialog_minimumSizeHint_default(void* self) {
	auto&& __r = ((QProgressDialog*)self)->QProgressDialog::minimumSizeHint();
	return new QSize(__r);
}
int QProgressDialog_exec_default(void* self) {
	auto&& __r = ((QProgressDialog*)self)->QProgressDialog::exec();
	return __r;
}
void QProgressDialog_done_default(void* self, int arg1) {
	((QProgressDialog*)self)->QProgressDialog::done(arg1);
}
void QProgressDialog_accept_default(void* self) {
	((QProgressDialog*)self)->QProgressDialog::accept();
}
void QProgressDialog_reject_default(void* self) {
	((QProgressDialog*)self)->QProgressDialog::reject();
}
void QProgressDialog_keyPressEvent_default(void* self, void* arg1) {
	((QProgressDialog__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)arg1);
}
void QProgressDialog_contextMenuEvent_default(void* self, void* arg1) {
	((QProgressDialog__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
int QProgressDialog_eventFilter_default(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QProgressDialog__Go*)self)->cppgoBase_eventFilter((QObject*)arg1, (QEvent*)arg2);
	return __r;
}
int QProgressDialog_devType_default(void* self) {
	auto&& __r = ((QProgressDialog*)self)->QProgressDialog::devType();
	return __r;
}
int QProgressDialog_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QProgressDialog*)self)->QProgressDialog::heightForWidth(arg1);
	return __r;
}
int QProgressDialog_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QProgressDialog*)self)->QProgressDialog::hasHeightForWidth();
	return __r;
}
void* QProgressDialog_paintEngine_default(void* self) {
	auto&& __r = ((QProgressDialog*)self)->QProgressDialog::paintEngine();
	return (void*)(__r);
}
int QProgressDialog_event_default(void* self, void* event) {
	auto&& __r = ((QProgressDialog__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QProgressDialog_mousePressEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QProgressDialog_mouseReleaseEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QProgressDialog_mouseDoubleClickEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QProgressDialog_mouseMoveEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QProgressDialog_wheelEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)event);
}
void QProgressDialog_keyReleaseEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QProgressDialog_focusInEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QProgressDialog_focusOutEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QProgressDialog_enterEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QProgressDialog_leaveEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QProgressDialog_paintEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)event);
}
void QProgressDialog_moveEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QProgressDialog_tabletEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QProgressDialog_actionEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QProgressDialog_dragEnterEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QProgressDialog_dragMoveEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QProgressDialog_dragLeaveEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QProgressDialog_dropEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QProgressDialog_hideEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QProgressDialog_metric_default(void* self, int arg1) {
	auto&& __r = ((QProgressDialog__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QProgressDialog_initPainter_default(void* self, void* painter) {
	((QProgressDialog__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QProgressDialog_redirected_default(void* self, void* offset) {
	auto&& __r = ((QProgressDialog__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QProgressDialog_sharedPainter_default(void* self) {
	auto&& __r = ((QProgressDialog__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QProgressDialog_inputMethodEvent_default(void* self, void* arg1) {
	((QProgressDialog__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void* QProgressDialog_inputMethodQuery_default(void* self, int arg1) {
	auto&& __r = ((QProgressDialog*)self)->QProgressDialog::inputMethodQuery((Qt::InputMethodQuery)(arg1));
	return new QVariant(__r);
}
int QProgressDialog_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QProgressDialog__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QProgressDialog_timerEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QProgressDialog_childEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QProgressDialog_customEvent_default(void* self, void* event) {
	((QProgressDialog__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QProgressDialog_connectNotify_default(void* self, void* signal) {
	((QProgressDialog__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QProgressDialog_disconnectNotify_default(void* self, void* signal) {
	((QProgressDialog__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QProgressDialog_connect_canceled(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressDialog*)self, &QProgressDialog::canceled, (QProgressDialog*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressDialog);
	__conns_QProgressDialog[__id] = __c;
}
void QProgressDialog_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressDialog*)self, &QProgressDialog::finished, (QProgressDialog*)self, [__id](int result){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, result);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressDialog);
	__conns_QProgressDialog[__id] = __c;
}
void QProgressDialog_connect_accepted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressDialog*)self, &QProgressDialog::accepted, (QProgressDialog*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressDialog);
	__conns_QProgressDialog[__id] = __c;
}
void QProgressDialog_connect_rejected(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressDialog*)self, &QProgressDialog::rejected, (QProgressDialog*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressDialog);
	__conns_QProgressDialog[__id] = __c;
}
void QProgressDialog_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressDialog*)self, &QProgressDialog::windowTitleChanged, (QProgressDialog*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressDialog);
	__conns_QProgressDialog[__id] = __c;
}
void QProgressDialog_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressDialog*)self, &QProgressDialog::windowIconChanged, (QProgressDialog*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressDialog);
	__conns_QProgressDialog[__id] = __c;
}
void QProgressDialog_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressDialog*)self, &QProgressDialog::windowIconTextChanged, (QProgressDialog*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressDialog);
	__conns_QProgressDialog[__id] = __c;
}
void QProgressDialog_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressDialog*)self, &QProgressDialog::customContextMenuRequested, (QProgressDialog*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressDialog);
	__conns_QProgressDialog[__id] = __c;
}
void QProgressDialog_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressDialog*)self, &QProgressDialog::destroyed, (QProgressDialog*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressDialog);
	__conns_QProgressDialog[__id] = __c;
}
void QProgressDialog_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressDialog*)self, &QProgressDialog::destroyed, (QProgressDialog*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressDialog);
	__conns_QProgressDialog[__id] = __c;
}
void QProgressDialog_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProgressDialog*)self, &QProgressDialog::objectNameChanged, (QProgressDialog*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProgressDialog);
	__conns_QProgressDialog[__id] = __c;
}
void QProgressDialog__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QProgressDialog);
		auto __it = __conns_QProgressDialog.find(id);
		if (__it == __conns_QProgressDialog.end()) return;
		__c = __it->second;
		__conns_QProgressDialog.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QProgressDialog__Go_new(long long id) { return new QProgressDialog__Go(id); }
void QProgressDialog_delete(void* self) { delete (QProgressDialog*)self; }

}

#include <QtWidgets/qtablewidget.h>
#include <QtWidgets/qtableview.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qmimedata.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qitemselectionmodel.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qpoint.h>
#include <QtCore/qrect.h>
#include <QtWidgets/qabstractitemview.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtWidgets/qheaderview.h>
#include <QtWidgets/qstyleoption.h>
#include <QtGui/qregion.h>
#include <QtCore/qsize.h>
#include <QtWidgets/qabstractitemdelegate.h>
#include <QtCore/qvariant.h>
#include <QtCore/qobject.h>
#include <QtWidgets/qscrollbar.h>
#include <QtWidgets/qabstractscrollarea.h>
#include <QtCore/qmargins.h>
#include <QtWidgets/qframe.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyle.h>
#include <QtGui/qpalette.h>
#include <QtGui/qfont.h>
#include <QtGui/qfontmetrics.h>
#include <QtGui/qfontinfo.h>
#include <QtGui/qcursor.h>
#include <QtGui/qbitmap.h>
#include <QtGui/qpaintdevice.h>
#include <QtGui/qpixmap.h>
#include <QtWidgets/qgraphicseffect.h>
#include <QtGui/qicon.h>
#include <QtCore/qlocale.h>
#include <QtGui/qkeysequence.h>
#include <QtWidgets/qgraphicsproxywidget.h>
#include <QtCore/qbytearray.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtWidgets/qlayout.h>
#include <QtGui/qaction.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qbackingstore.h>
#include <QtGui/qwindow.h>
#include <QtGui/qscreen.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qtablewidget.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QTableWidget__Go : public QTableWidget {
public:
	long long __goId;
	unsigned char __goOn[93] = {};
	template<typename... A> QTableWidget__Go(long long id, A&&... a) : QTableWidget(std::forward<A>(a)...), __goId(id) {}
	using QTableWidget::rowMoved;
	using QTableWidget::columnMoved;
	using QTableWidget::rowResized;
	using QTableWidget::columnResized;
	using QTableWidget::rowCountChanged;
	using QTableWidget::columnCountChanged;
	using QTableWidget::state;
	using QTableWidget::scheduleDelayedItemsLayout;
	using QTableWidget::executeDelayedItemsLayout;
	using QTableWidget::setDirtyRegion;
	using QTableWidget::scrollDirtyRegion;
	using QTableWidget::dirtyRegionOffset;
	using QTableWidget::startAutoScroll;
	using QTableWidget::stopAutoScroll;
	using QTableWidget::doAutoScroll;
	using QTableWidget::dropIndicatorPosition;
	using QTableWidget::setViewportMargins;
	using QTableWidget::viewportMargins;
	using QTableWidget::drawFrame;
	using QTableWidget::updateMicroFocus;
	using QTableWidget::create;
	using QTableWidget::destroy;
	using QTableWidget::focusNextChild;
	using QTableWidget::focusPreviousChild;
	using QTableWidget::sender;
	using QTableWidget::senderSignalIndex;
	using QTableWidget::receivers;
	using QTableWidget::isSignalConnected;
	using QTableWidget::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QTableWidget::event(__a0); }
	QList<QString> cppgoBase_mimeTypes() { return QTableWidget::mimeTypes(); }
	bool cppgoBase_dropMimeData(int __a0, int __a1, const QMimeData* __a2, int __a3) { return QTableWidget::dropMimeData(__a0, __a1, __a2, (Qt::DropAction)__a3); }
	QFlags<Qt::DropAction> cppgoBase_supportedDropActions() { return QTableWidget::supportedDropActions(); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QTableWidget::dropEvent(__a0); }
	void cppgoBase_scrollContentsBy(int __a0, int __a1) { QTableWidget::scrollContentsBy(__a0, __a1); }
	void cppgoBase_initViewItemOption(QStyleOptionViewItem* __a0) { QTableWidget::initViewItemOption(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QTableWidget::paintEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QTableWidget::timerEvent(__a0); }
	int cppgoBase_horizontalOffset() { return QTableWidget::horizontalOffset(); }
	int cppgoBase_verticalOffset() { return QTableWidget::verticalOffset(); }
	QModelIndex cppgoBase_moveCursor(int __a0, QFlags<Qt::KeyboardModifier> __a1) { return QTableWidget::moveCursor((QAbstractItemView::CursorAction)__a0, __a1); }
	void cppgoBase_setSelection(const QRect& __a0, QFlags<QItemSelectionModel::SelectionFlag> __a1) { QTableWidget::setSelection(__a0, __a1); }
	QRegion cppgoBase_visualRegionForSelection(const QItemSelection& __a0) { return QTableWidget::visualRegionForSelection(__a0); }
	QList<QModelIndex> cppgoBase_selectedIndexes() { return QTableWidget::selectedIndexes(); }
	void cppgoBase_updateGeometries() { QTableWidget::updateGeometries(); }
	QSize cppgoBase_viewportSizeHint() { return QTableWidget::viewportSizeHint(); }
	int cppgoBase_sizeHintForRow(int __a0) { return QTableWidget::sizeHintForRow(__a0); }
	int cppgoBase_sizeHintForColumn(int __a0) { return QTableWidget::sizeHintForColumn(__a0); }
	void cppgoBase_verticalScrollbarAction(int __a0) { QTableWidget::verticalScrollbarAction(__a0); }
	void cppgoBase_horizontalScrollbarAction(int __a0) { QTableWidget::horizontalScrollbarAction(__a0); }
	bool cppgoBase_isIndexHidden(const QModelIndex& __a0) { return QTableWidget::isIndexHidden(__a0); }
	void cppgoBase_selectionChanged(const QItemSelection& __a0, const QItemSelection& __a1) { QTableWidget::selectionChanged(__a0, __a1); }
	void cppgoBase_currentChanged(const QModelIndex& __a0, const QModelIndex& __a1) { QTableWidget::currentChanged(__a0, __a1); }
	void cppgoBase_rowsInserted(const QModelIndex& __a0, int __a1, int __a2) { QTableWidget::rowsInserted(__a0, __a1, __a2); }
	void cppgoBase_rowsAboutToBeRemoved(const QModelIndex& __a0, int __a1, int __a2) { QTableWidget::rowsAboutToBeRemoved(__a0, __a1, __a2); }
	void cppgoBase_updateEditorData() { QTableWidget::updateEditorData(); }
	void cppgoBase_updateEditorGeometries() { QTableWidget::updateEditorGeometries(); }
	void cppgoBase_verticalScrollbarValueChanged(int __a0) { QTableWidget::verticalScrollbarValueChanged(__a0); }
	void cppgoBase_horizontalScrollbarValueChanged(int __a0) { QTableWidget::horizontalScrollbarValueChanged(__a0); }
	void cppgoBase_closeEditor(QWidget* __a0, int __a1) { QTableWidget::closeEditor(__a0, (QAbstractItemDelegate::EndEditHint)__a1); }
	void cppgoBase_commitData(QWidget* __a0) { QTableWidget::commitData(__a0); }
	void cppgoBase_editorDestroyed(QObject* __a0) { QTableWidget::editorDestroyed(__a0); }
	bool cppgoBase_edit(const QModelIndex& __a0, int __a1, QEvent* __a2) { return QTableWidget::edit(__a0, (QAbstractItemView::EditTrigger)__a1, __a2); }
	QFlags<QItemSelectionModel::SelectionFlag> cppgoBase_selectionCommand(const QModelIndex& __a0, const QEvent* __a1) { return QTableWidget::selectionCommand(__a0, __a1); }
	void cppgoBase_startDrag(QFlags<Qt::DropAction> __a0) { QTableWidget::startDrag(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QTableWidget::focusNextPrevChild(__a0); }
	bool cppgoBase_viewportEvent(QEvent* __a0) { return QTableWidget::viewportEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QTableWidget::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QTableWidget::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QTableWidget::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QTableWidget::mouseDoubleClickEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QTableWidget::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QTableWidget::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QTableWidget::dragLeaveEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QTableWidget::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QTableWidget::focusOutEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QTableWidget::keyPressEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QTableWidget::resizeEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QTableWidget::inputMethodEvent(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QTableWidget::eventFilter(__a0, __a1); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QTableWidget::wheelEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QTableWidget::contextMenuEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QTableWidget::changeEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QTableWidget::initStyleOption(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QTableWidget::keyReleaseEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QTableWidget::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QTableWidget::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QTableWidget::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QTableWidget::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QTableWidget::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QTableWidget::actionEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QTableWidget::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QTableWidget::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QTableWidget::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QTableWidget::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QTableWidget::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QTableWidget::sharedPainter(); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QTableWidget::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QTableWidget::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QTableWidget::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QTableWidget::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QTableWidget::metaObject();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QTableWidget::event(e);
	}
	QList<QString> mimeTypes() const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, 5464467312394730418LL, &__len, &__h);
			if (__h) {
				QList<QString> __r; __r.reserve(__len);
				char** __aa = (char**)__a;
				for (long long __i = 0; __i < __len; __i++) { __r.append(QString::fromUtf8(__aa[__i])); free(__aa[__i]); }
				if (__a) free(__a);
				return __r;
			}
		}
		return QTableWidget::mimeTypes();
	}
	bool dropMimeData(int row, int column, const QMimeData* data, Qt::DropAction action) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_int_int_voidp_int) { int __r = ((int(*)(long long, long long, int, int, void*, int, int*))cppgo_virt_int_int_int_voidp_int)(__goId, -9049994684305682212LL, row, column, (void*)data, (int)(action), &__h); if (__h) return __r; }
		return QTableWidget::dropMimeData(row, column, data, action);
	}
	QFlags<Qt::DropAction> supportedDropActions() const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -7784642176186821205LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QTableWidget::supportedDropActions();
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::dropEvent(event);
	}
	void setRootIndex(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1574077111192992763LL, (void*)&index, &__h); if (__h) return; }
		QTableWidget::setRootIndex(index);
	}
	void setSelectionModel(QItemSelectionModel* selectionModel) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4689322347236366780LL, (void*)selectionModel, &__h); if (__h) return; }
		QTableWidget::setSelectionModel(selectionModel);
	}
	void doItemsLayout() override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1888629738259754754LL, &__h); if (__h) return; }
		QTableWidget::doItemsLayout();
	}
	QRect visualRect(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -1267696851865372997LL, (void*)&index, &__h); if (__h) return *(QRect*)__r; }
		return QTableWidget::visualRect(index);
	}
	void scrollTo(const QModelIndex& index, QAbstractItemView::ScrollHint hint) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -4832340528568269947LL, (void*)&index, (int)(hint), &__h); if (__h) return; }
		QTableWidget::scrollTo(index, hint);
	}
	QModelIndex indexAt(const QPoint& p) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -7252823847285716802LL, (void*)&p, &__h); if (__h) return *(QModelIndex*)__r; }
		return QTableWidget::indexAt(p);
	}
	void scrollContentsBy(int dx, int dy) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, -6118001150126564745LL, dx, dy, &__h); if (__h) return; }
		QTableWidget::scrollContentsBy(dx, dy);
	}
	void initViewItemOption(QStyleOptionViewItem* option) const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -7450787277202094556LL, (void*)option, &__h); if (__h) return; }
		QTableWidget::initViewItemOption(option);
	}
	void paintEvent(QPaintEvent* e) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)e, &__h); if (__h) return; }
		QTableWidget::paintEvent(e);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::timerEvent(event);
	}
	int horizontalOffset() const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1500310024801917802LL, &__h); if (__h) return __r; }
		return QTableWidget::horizontalOffset();
	}
	int verticalOffset() const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 5218707810677921016LL, &__h); if (__h) return __r; }
		return QTableWidget::verticalOffset();
	}
	QModelIndex moveCursor(QAbstractItemView::CursorAction cursorAction, QFlags<Qt::KeyboardModifier> modifiers) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_voidp_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int*))cppgo_virt_voidp_int_int)(__goId, -158523882194220890LL, (int)(cursorAction), (int)(modifiers), &__h); if (__h) return *(QModelIndex*)__r; }
		return QTableWidget::moveCursor(cursorAction, modifiers);
	}
	void setSelection(const QRect& rect, QFlags<QItemSelectionModel::SelectionFlag> command) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -1539266838375037305LL, (void*)&rect, (int)(command), &__h); if (__h) return; }
		QTableWidget::setSelection(rect, command);
	}
	QRegion visualRegionForSelection(const QItemSelection& selection) const override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 7063795720810552074LL, (void*)&selection, &__h); if (__h) return *(QRegion*)__r; }
		return QTableWidget::visualRegionForSelection(selection);
	}
	QList<QModelIndex> selectedIndexes() const override {
		int __h = 0;
		if (__goOn[21] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, 740869982805721908LL, &__len, &__h);
			if (__h) {
				QList<QModelIndex> __r; __r.reserve(__len);
				void** __aa = (void**)__a;
				for (long long __i = 0; __i < __len; __i++) __r.append(*(QModelIndex*)__aa[__i]);
				if (__a) free(__a);
				return __r;
			}
		}
		return QTableWidget::selectedIndexes();
	}
	void updateGeometries() override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1022771470898460576LL, &__h); if (__h) return; }
		QTableWidget::updateGeometries();
	}
	QSize viewportSizeHint() const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 503579067252048239LL, &__h); if (__h) return *(QSize*)__r; }
		return QTableWidget::viewportSizeHint();
	}
	int sizeHintForRow(int row) const override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -4084029023708556876LL, row, &__h); if (__h) return __r; }
		return QTableWidget::sizeHintForRow(row);
	}
	int sizeHintForColumn(int column) const override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -3666792239577137638LL, column, &__h); if (__h) return __r; }
		return QTableWidget::sizeHintForColumn(column);
	}
	void verticalScrollbarAction(int action) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -1155555122391250623LL, action, &__h); if (__h) return; }
		QTableWidget::verticalScrollbarAction(action);
	}
	void horizontalScrollbarAction(int action) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -785551405282638793LL, action, &__h); if (__h) return; }
		QTableWidget::horizontalScrollbarAction(action);
	}
	bool isIndexHidden(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 6057989667089378859LL, (void*)&index, &__h); if (__h) return __r; }
		return QTableWidget::isIndexHidden(index);
	}
	void selectionChanged(const QItemSelection& selected, const QItemSelection& deselected) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -5892205462145366901LL, (void*)&selected, (void*)&deselected, &__h); if (__h) return; }
		QTableWidget::selectionChanged(selected, deselected);
	}
	void currentChanged(const QModelIndex& current, const QModelIndex& previous) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -6978583972112316034LL, (void*)&current, (void*)&previous, &__h); if (__h) return; }
		QTableWidget::currentChanged(current, previous);
	}
	void keyboardSearch(const QString& search) override {
		int __h = 0;
		QByteArray search_ba = search.toUtf8();
		if (__goOn[31] && cppgo_virt_void_constcharp_longlong) { ((void(*)(long long, long long, const char*, long long, int*))cppgo_virt_void_constcharp_longlong)(__goId, -5041452885022114200LL, (char*)search_ba.constData(), search_ba.size(), &__h); if (__h) return; }
		QTableWidget::keyboardSearch(search);
	}
	QAbstractItemDelegate* itemDelegateForIndex(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -3811438228839363824LL, (void*)&index, &__h); if (__h) return (QAbstractItemDelegate*)__r; }
		return QTableWidget::itemDelegateForIndex(index);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QTableWidget::inputMethodQuery(query);
	}
	void reset() override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1086335023529244512LL, &__h); if (__h) return; }
		QTableWidget::reset();
	}
	void selectAll() override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 8758113809152574300LL, &__h); if (__h) return; }
		QTableWidget::selectAll();
	}
	void rowsInserted(const QModelIndex& parent, int start, int end) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, 4472434241583495620LL, (void*)&parent, start, end, &__h); if (__h) return; }
		QTableWidget::rowsInserted(parent, start, end);
	}
	void rowsAboutToBeRemoved(const QModelIndex& parent, int start, int end) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, -8151049665126058157LL, (void*)&parent, start, end, &__h); if (__h) return; }
		QTableWidget::rowsAboutToBeRemoved(parent, start, end);
	}
	void updateEditorData() override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9029941737538200621LL, &__h); if (__h) return; }
		QTableWidget::updateEditorData();
	}
	void updateEditorGeometries() override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -8726300139121449691LL, &__h); if (__h) return; }
		QTableWidget::updateEditorGeometries();
	}
	void verticalScrollbarValueChanged(int value) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6366953527806415524LL, value, &__h); if (__h) return; }
		QTableWidget::verticalScrollbarValueChanged(value);
	}
	void horizontalScrollbarValueChanged(int value) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 1786430177580376614LL, value, &__h); if (__h) return; }
		QTableWidget::horizontalScrollbarValueChanged(value);
	}
	void closeEditor(QWidget* editor, QAbstractItemDelegate::EndEditHint hint) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, 8648484264628816806LL, (void*)editor, (int)(hint), &__h); if (__h) return; }
		QTableWidget::closeEditor(editor, hint);
	}
	void commitData(QWidget* editor) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6021383853323633948LL, (void*)editor, &__h); if (__h) return; }
		QTableWidget::commitData(editor);
	}
	void editorDestroyed(QObject* editor) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6979418566323745651LL, (void*)editor, &__h); if (__h) return; }
		QTableWidget::editorDestroyed(editor);
	}
	bool edit(const QModelIndex& index, QAbstractItemView::EditTrigger trigger, QEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_voidp_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, void*, int*))cppgo_virt_int_voidp_int_voidp)(__goId, 4206689019978932209LL, (void*)&index, (int)(trigger), (void*)event, &__h); if (__h) return __r; }
		return QTableWidget::edit(index, trigger, event);
	}
	QFlags<QItemSelectionModel::SelectionFlag> selectionCommand(const QModelIndex& index, const QEvent* event) const override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -2447490284072398302LL, (void*)&index, (void*)event, &__h); if (__h) return QFlags<QItemSelectionModel::SelectionFlag>(__r); }
		return QTableWidget::selectionCommand(index, event);
	}
	void startDrag(QFlags<Qt::DropAction> supportedActions) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -8112366604639385491LL, (int)(supportedActions), &__h); if (__h) return; }
		QTableWidget::startDrag(supportedActions);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QTableWidget::focusNextPrevChild(next);
	}
	bool viewportEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -1971465653061861519LL, (void*)event, &__h); if (__h) return __r; }
		return QTableWidget::viewportEvent(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::mousePressEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::mouseDoubleClickEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[56] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::dragLeaveEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[57] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[58] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::focusOutEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[59] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::keyPressEvent(event);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[60] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::resizeEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[61] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::inputMethodEvent(event);
	}
	bool eventFilter(QObject* object, QEvent* event) override {
		int __h = 0;
		if (__goOn[62] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)object, (void*)event, &__h); if (__h) return __r; }
		return QTableWidget::eventFilter(object, event);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[63] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QTableWidget::minimumSizeHint();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[64] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QTableWidget::sizeHint();
	}
	void setupViewport(QWidget* viewport) override {
		int __h = 0;
		if (__goOn[65] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4531996381736730688LL, (void*)viewport, &__h); if (__h) return; }
		QTableWidget::setupViewport(viewport);
	}
	void wheelEvent(QWheelEvent* arg1) override {
		int __h = 0;
		if (__goOn[66] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)arg1, &__h); if (__h) return; }
		QTableWidget::wheelEvent(arg1);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[67] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QTableWidget::contextMenuEvent(arg1);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[68] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QTableWidget::changeEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[69] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QTableWidget::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[70] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QTableWidget::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[71] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QTableWidget::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[72] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QTableWidget::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[73] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QTableWidget::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[74] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QTableWidget::paintEngine();
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[75] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::keyReleaseEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[76] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[77] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[78] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[79] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[80] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[81] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::actionEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[82] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[83] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[84] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QTableWidget::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[85] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QTableWidget::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[86] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QTableWidget::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[87] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QTableWidget::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[88] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QTableWidget::sharedPainter();
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[89] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[90] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QTableWidget::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[91] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QTableWidget::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[92] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QTableWidget::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTableWidget;
static std::mutex __connsMu_QTableWidget;

extern "C" {

void* QTableWidget_base(void* self) { return (void*)static_cast<QTableView*>((QTableWidget*)self); }
void QTableWidget_chain(void* self, void** out) {
	QTableWidget* __p = (QTableWidget*)self;
	QTableView* __b0 = static_cast<QTableView*>(__p); out[0] = (void*)__b0;
	QAbstractItemView* __b1 = static_cast<QAbstractItemView*>(__b0); out[1] = (void*)__b1;
	QAbstractScrollArea* __b2 = static_cast<QAbstractScrollArea*>(__b1); out[2] = (void*)__b2;
	QFrame* __b3 = static_cast<QFrame*>(__b2); out[3] = (void*)__b3;
	QWidget* __b4 = static_cast<QWidget*>(__b3); out[4] = (void*)__b4;
	QObject* __b5 = static_cast<QObject*>(__b4); out[5] = (void*)__b5;
}
void QTableWidget__set_override(void* self, int idx) { ((QTableWidget__Go*)self)->__goOn[idx] = 1; }
void* QTableWidget_fromQObject(void* o) { return (void*)qobject_cast<QTableWidget*>((QObject*)o); }
void* QTableWidget_new(long long __goid, void* parent) {
	return new QTableWidget__Go(__goid, (QWidget*)parent);
}
void* QTableWidget_new2(long long __goid) {
	return new QTableWidget__Go(__goid);
}
void* QTableWidget_new3(long long __goid, int rows, int columns) {
	return new QTableWidget__Go(__goid, rows, columns);
}
void* QTableWidget_new4(long long __goid, int rows, int columns, void* parent) {
	return new QTableWidget__Go(__goid, rows, columns, (QWidget*)parent);
}
char* QTableWidget_tr(const char* s, long long* __out_len) {
	auto&& __r = QTableWidget::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTableWidget_setRowCount(void* self, int rows) {
	((QTableWidget*)self)->setRowCount(rows);
}
int QTableWidget_rowCount(void* self) {
	auto&& __r = ((QTableWidget*)self)->rowCount();
	return __r;
}
void QTableWidget_setColumnCount(void* self, int columns) {
	((QTableWidget*)self)->setColumnCount(columns);
}
int QTableWidget_columnCount(void* self) {
	auto&& __r = ((QTableWidget*)self)->columnCount();
	return __r;
}
int QTableWidget_row(void* self, void* item) {
	auto&& __r = ((QTableWidget*)self)->row((QTableWidgetItem*)item);
	return __r;
}
int QTableWidget_column(void* self, void* item) {
	auto&& __r = ((QTableWidget*)self)->column((QTableWidgetItem*)item);
	return __r;
}
void* QTableWidget_item(void* self, int row, int column) {
	auto&& __r = ((QTableWidget*)self)->item(row, column);
	return (void*)(__r);
}
void QTableWidget_setItem(void* self, int row, int column, void* item) {
	((QTableWidget*)self)->setItem(row, column, (QTableWidgetItem*)item);
}
void* QTableWidget_takeItem(void* self, int row, int column) {
	auto&& __r = ((QTableWidget*)self)->takeItem(row, column);
	return (void*)(__r);
}
void** QTableWidget_items(void* self, void* data, long long* __out_len) {
	auto&& __r = ((QTableWidget*)self)->items((QMimeData*)data);
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QTableWidget_indexFromItem(void* self, void* item) {
	auto&& __r = ((QTableWidget*)self)->indexFromItem((QTableWidgetItem*)item);
	return new QModelIndex(__r);
}
void* QTableWidget_itemFromIndex(void* self, void* index) {
	auto&& __r = ((QTableWidget*)self)->itemFromIndex(*(QModelIndex*)index);
	return (void*)(__r);
}
void* QTableWidget_verticalHeaderItem(void* self, int row) {
	auto&& __r = ((QTableWidget*)self)->verticalHeaderItem(row);
	return (void*)(__r);
}
void QTableWidget_setVerticalHeaderItem(void* self, int row, void* item) {
	((QTableWidget*)self)->setVerticalHeaderItem(row, (QTableWidgetItem*)item);
}
void* QTableWidget_takeVerticalHeaderItem(void* self, int row) {
	auto&& __r = ((QTableWidget*)self)->takeVerticalHeaderItem(row);
	return (void*)(__r);
}
void* QTableWidget_horizontalHeaderItem(void* self, int column) {
	auto&& __r = ((QTableWidget*)self)->horizontalHeaderItem(column);
	return (void*)(__r);
}
void QTableWidget_setHorizontalHeaderItem(void* self, int column, void* item) {
	((QTableWidget*)self)->setHorizontalHeaderItem(column, (QTableWidgetItem*)item);
}
void* QTableWidget_takeHorizontalHeaderItem(void* self, int column) {
	auto&& __r = ((QTableWidget*)self)->takeHorizontalHeaderItem(column);
	return (void*)(__r);
}
void QTableWidget_setVerticalHeaderLabels(void* self, char** labels, long long labels_len) {
	QList<QString> labels_list; labels_list.reserve(labels_len);
	for (long long __i = 0; __i < labels_len; __i++) labels_list.append(QString::fromUtf8(labels[__i]));
	((QTableWidget*)self)->setVerticalHeaderLabels(labels_list);
}
void QTableWidget_setHorizontalHeaderLabels(void* self, char** labels, long long labels_len) {
	QList<QString> labels_list; labels_list.reserve(labels_len);
	for (long long __i = 0; __i < labels_len; __i++) labels_list.append(QString::fromUtf8(labels[__i]));
	((QTableWidget*)self)->setHorizontalHeaderLabels(labels_list);
}
int QTableWidget_currentRow(void* self) {
	auto&& __r = ((QTableWidget*)self)->currentRow();
	return __r;
}
int QTableWidget_currentColumn(void* self) {
	auto&& __r = ((QTableWidget*)self)->currentColumn();
	return __r;
}
void* QTableWidget_currentItem(void* self) {
	auto&& __r = ((QTableWidget*)self)->currentItem();
	return (void*)(__r);
}
void QTableWidget_setCurrentItem(void* self, void* item) {
	((QTableWidget*)self)->setCurrentItem((QTableWidgetItem*)item);
}
void QTableWidget_setCurrentItem2(void* self, void* item, int command) {
	((QTableWidget*)self)->setCurrentItem((QTableWidgetItem*)item, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void QTableWidget_setCurrentCell(void* self, int row, int column) {
	((QTableWidget*)self)->setCurrentCell(row, column);
}
void QTableWidget_setCurrentCell2(void* self, int row, int column, int command) {
	((QTableWidget*)self)->setCurrentCell(row, column, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void QTableWidget_sortItems(void* self, int column) {
	((QTableWidget*)self)->sortItems(column);
}
void QTableWidget_setSortingEnabled(void* self, int enable) {
	((QTableWidget*)self)->setSortingEnabled(enable);
}
int QTableWidget_isSortingEnabled(void* self) {
	auto&& __r = ((QTableWidget*)self)->isSortingEnabled();
	return __r;
}
void QTableWidget_editItem(void* self, void* item) {
	((QTableWidget*)self)->editItem((QTableWidgetItem*)item);
}
void QTableWidget_openPersistentEditor(void* self, void* item) {
	((QTableWidget*)self)->openPersistentEditor((QTableWidgetItem*)item);
}
void QTableWidget_closePersistentEditor(void* self, void* item) {
	((QTableWidget*)self)->closePersistentEditor((QTableWidgetItem*)item);
}
int QTableWidget_isPersistentEditorOpen(void* self, void* item) {
	auto&& __r = ((QTableWidget*)self)->isPersistentEditorOpen((QTableWidgetItem*)item);
	return __r;
}
void* QTableWidget_cellWidget(void* self, int row, int column) {
	auto&& __r = ((QTableWidget*)self)->cellWidget(row, column);
	return (void*)(__r);
}
void QTableWidget_setCellWidget(void* self, int row, int column, void* widget) {
	((QTableWidget*)self)->setCellWidget(row, column, (QWidget*)widget);
}
void QTableWidget_removeCellWidget(void* self, int row, int column) {
	((QTableWidget*)self)->removeCellWidget(row, column);
}
void QTableWidget_setRangeSelected(void* self, void* range_, int select_) {
	((QTableWidget*)self)->setRangeSelected(*(QTableWidgetSelectionRange*)range_, select_);
}
void** QTableWidget_selectedRanges(void* self, long long* __out_len) {
	auto&& __r = ((QTableWidget*)self)->selectedRanges();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QTableWidgetSelectionRange(__e);
	return __arr;
}
void** QTableWidget_selectedItems(void* self, long long* __out_len) {
	auto&& __r = ((QTableWidget*)self)->selectedItems();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void** QTableWidget_findItems(void* self, const char* text, long long text_len, int flags, long long* __out_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QTableWidget*)self)->findItems(text_qstr, QFlags<Qt::MatchFlag>(flags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
int QTableWidget_visualRow(void* self, int logicalRow) {
	auto&& __r = ((QTableWidget*)self)->visualRow(logicalRow);
	return __r;
}
int QTableWidget_visualColumn(void* self, int logicalColumn) {
	auto&& __r = ((QTableWidget*)self)->visualColumn(logicalColumn);
	return __r;
}
void* QTableWidget_itemAt(void* self, void* p) {
	auto&& __r = ((QTableWidget*)self)->itemAt(*(QPoint*)p);
	return (void*)(__r);
}
void* QTableWidget_itemAt2(void* self, int x, int y) {
	auto&& __r = ((QTableWidget*)self)->itemAt(x, y);
	return (void*)(__r);
}
void* QTableWidget_visualItemRect(void* self, void* item) {
	auto&& __r = ((QTableWidget*)self)->visualItemRect((QTableWidgetItem*)item);
	return new QRect(__r);
}
void* QTableWidget_itemPrototype(void* self) {
	auto&& __r = ((QTableWidget*)self)->itemPrototype();
	return (void*)(__r);
}
void QTableWidget_setItemPrototype(void* self, void* item) {
	((QTableWidget*)self)->setItemPrototype((QTableWidgetItem*)item);
}
int QTableWidget_supportedDragActions(void* self) {
	auto&& __r = ((QTableWidget*)self)->supportedDragActions();
	return (int)(__r);
}
void QTableWidget_setSupportedDragActions(void* self, int actions) {
	((QTableWidget*)self)->setSupportedDragActions(QFlags<Qt::DropAction>(actions));
}
void QTableWidget_scrollToItem(void* self, void* item) {
	((QTableWidget*)self)->scrollToItem((QTableWidgetItem*)item);
}
void QTableWidget_insertRow(void* self, int row) {
	((QTableWidget*)self)->insertRow(row);
}
void QTableWidget_insertColumn(void* self, int column) {
	((QTableWidget*)self)->insertColumn(column);
}
void QTableWidget_removeRow(void* self, int row) {
	((QTableWidget*)self)->removeRow(row);
}
void QTableWidget_removeColumn(void* self, int column) {
	((QTableWidget*)self)->removeColumn(column);
}
void QTableWidget_clear(void* self) {
	((QTableWidget*)self)->clear();
}
void QTableWidget_clearContents(void* self) {
	((QTableWidget*)self)->clearContents();
}
void QTableWidget_itemPressed(void* self, void* item) {
	((QTableWidget*)self)->itemPressed((QTableWidgetItem*)item);
}
void QTableWidget_itemClicked(void* self, void* item) {
	((QTableWidget*)self)->itemClicked((QTableWidgetItem*)item);
}
void QTableWidget_itemDoubleClicked(void* self, void* item) {
	((QTableWidget*)self)->itemDoubleClicked((QTableWidgetItem*)item);
}
void QTableWidget_itemActivated(void* self, void* item) {
	((QTableWidget*)self)->itemActivated((QTableWidgetItem*)item);
}
void QTableWidget_itemEntered(void* self, void* item) {
	((QTableWidget*)self)->itemEntered((QTableWidgetItem*)item);
}
void QTableWidget_itemChanged(void* self, void* item) {
	((QTableWidget*)self)->itemChanged((QTableWidgetItem*)item);
}
void QTableWidget_currentItemChanged(void* self, void* current, void* previous) {
	((QTableWidget*)self)->currentItemChanged((QTableWidgetItem*)current, (QTableWidgetItem*)previous);
}
void QTableWidget_itemSelectionChanged(void* self) {
	((QTableWidget*)self)->itemSelectionChanged();
}
void QTableWidget_cellPressed(void* self, int row, int column) {
	((QTableWidget*)self)->cellPressed(row, column);
}
void QTableWidget_cellClicked(void* self, int row, int column) {
	((QTableWidget*)self)->cellClicked(row, column);
}
void QTableWidget_cellDoubleClicked(void* self, int row, int column) {
	((QTableWidget*)self)->cellDoubleClicked(row, column);
}
void QTableWidget_cellActivated(void* self, int row, int column) {
	((QTableWidget*)self)->cellActivated(row, column);
}
void QTableWidget_cellEntered(void* self, int row, int column) {
	((QTableWidget*)self)->cellEntered(row, column);
}
void QTableWidget_cellChanged(void* self, int row, int column) {
	((QTableWidget*)self)->cellChanged(row, column);
}
void QTableWidget_currentCellChanged(void* self, int currentRow, int currentColumn, int previousRow, int previousColumn) {
	((QTableWidget*)self)->currentCellChanged(currentRow, currentColumn, previousRow, previousColumn);
}
void QTableWidget_rowMoved(void* self, int row, int oldIndex, int newIndex) {
	((QTableWidget__Go*)self)->rowMoved(row, oldIndex, newIndex);
}
void QTableWidget_columnMoved(void* self, int column, int oldIndex, int newIndex) {
	((QTableWidget__Go*)self)->columnMoved(column, oldIndex, newIndex);
}
void QTableWidget_rowResized(void* self, int row, int oldHeight, int newHeight) {
	((QTableWidget__Go*)self)->rowResized(row, oldHeight, newHeight);
}
void QTableWidget_columnResized(void* self, int column, int oldWidth, int newWidth) {
	((QTableWidget__Go*)self)->columnResized(column, oldWidth, newWidth);
}
void QTableWidget_rowCountChanged(void* self, int oldCount, int newCount) {
	((QTableWidget__Go*)self)->rowCountChanged(oldCount, newCount);
}
void QTableWidget_columnCountChanged(void* self, int oldCount, int newCount) {
	((QTableWidget__Go*)self)->columnCountChanged(oldCount, newCount);
}
int QTableWidget_state(void* self) {
	auto&& __r = ((QTableWidget__Go*)self)->state();
	return (int)(__r);
}
void QTableWidget_scheduleDelayedItemsLayout(void* self) {
	((QTableWidget__Go*)self)->scheduleDelayedItemsLayout();
}
void QTableWidget_executeDelayedItemsLayout(void* self) {
	((QTableWidget__Go*)self)->executeDelayedItemsLayout();
}
void QTableWidget_setDirtyRegion(void* self, void* region) {
	((QTableWidget__Go*)self)->setDirtyRegion(*(QRegion*)region);
}
void QTableWidget_scrollDirtyRegion(void* self, int dx, int dy) {
	((QTableWidget__Go*)self)->scrollDirtyRegion(dx, dy);
}
void* QTableWidget_dirtyRegionOffset(void* self) {
	auto&& __r = ((QTableWidget__Go*)self)->dirtyRegionOffset();
	return new QPoint(__r);
}
void QTableWidget_startAutoScroll(void* self) {
	((QTableWidget__Go*)self)->startAutoScroll();
}
void QTableWidget_stopAutoScroll(void* self) {
	((QTableWidget__Go*)self)->stopAutoScroll();
}
void QTableWidget_doAutoScroll(void* self) {
	((QTableWidget__Go*)self)->doAutoScroll();
}
int QTableWidget_dropIndicatorPosition(void* self) {
	auto&& __r = ((QTableWidget__Go*)self)->dropIndicatorPosition();
	return (int)(__r);
}
void QTableWidget_setViewportMargins(void* self, int left, int top, int right, int bottom) {
	((QTableWidget__Go*)self)->setViewportMargins(left, top, right, bottom);
}
void QTableWidget_setViewportMargins2(void* self, void* margins) {
	((QTableWidget__Go*)self)->setViewportMargins(*(QMargins*)margins);
}
void* QTableWidget_viewportMargins(void* self) {
	auto&& __r = ((QTableWidget__Go*)self)->viewportMargins();
	return new QMargins(__r);
}
void QTableWidget_drawFrame(void* self, void* arg1) {
	((QTableWidget__Go*)self)->drawFrame((QPainter*)arg1);
}
void QTableWidget_setTabOrder(void* arg1, void* arg2) {
	QTableWidget::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QTableWidget_mouseGrabber() {
	auto&& __r = QTableWidget::mouseGrabber();
	return (void*)(__r);
}
void* QTableWidget_keyboardGrabber() {
	auto&& __r = QTableWidget::keyboardGrabber();
	return (void*)(__r);
}
void* QTableWidget_find(unsigned long long arg1) {
	auto&& __r = QTableWidget::find(arg1);
	return (void*)(__r);
}
void* QTableWidget_createWindowContainer(void* window) {
	auto&& __r = QTableWidget::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QTableWidget_updateMicroFocus(void* self) {
	((QTableWidget__Go*)self)->updateMicroFocus();
}
void QTableWidget_create(void* self) {
	((QTableWidget__Go*)self)->create();
}
void QTableWidget_destroy(void* self) {
	((QTableWidget__Go*)self)->destroy();
}
int QTableWidget_focusNextChild(void* self) {
	auto&& __r = ((QTableWidget__Go*)self)->focusNextChild();
	return __r;
}
int QTableWidget_focusPreviousChild(void* self) {
	auto&& __r = ((QTableWidget__Go*)self)->focusPreviousChild();
	return __r;
}
void* QTableWidget_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTableWidget::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTableWidget_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTableWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QTableWidget_sender(void* self) {
	auto&& __r = ((QTableWidget__Go*)self)->sender();
	return (void*)(__r);
}
int QTableWidget_senderSignalIndex(void* self) {
	auto&& __r = ((QTableWidget__Go*)self)->senderSignalIndex();
	return __r;
}
int QTableWidget_receivers(void* self, const char* signal) {
	auto&& __r = ((QTableWidget__Go*)self)->receivers(signal);
	return __r;
}
int QTableWidget_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QTableWidget__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QTableWidget_devicePixelRatioFScale() {
	auto&& __r = QTableWidget::devicePixelRatioFScale();
	return __r;
}
int QTableWidget_encodeMetricF(int metric, double value) {
	auto&& __r = QTableWidget::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QTableWidget_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QTableWidget__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QTableWidget_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTableWidget::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTableWidget_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTableWidget::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTableWidget_sortItems2(void* self, int column, int order) {
	((QTableWidget*)self)->sortItems(column, (Qt::SortOrder)(order));
}
void QTableWidget_scrollToItem2(void* self, void* item, int hint) {
	((QTableWidget*)self)->scrollToItem((QTableWidgetItem*)item, (QAbstractItemView::ScrollHint)(hint));
}
void* QTableWidget_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QTableWidget::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QTableWidget_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QTableWidget::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QTableWidget_updateMicroFocus2(void* self, int query) {
	((QTableWidget__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QTableWidget_create2(void* self, unsigned long long arg1) {
	((QTableWidget__Go*)self)->create(arg1);
}
void QTableWidget_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QTableWidget__Go*)self)->create(arg1, initializeWindow);
}
void QTableWidget_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QTableWidget__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QTableWidget_destroy2(void* self, int destroyWindow) {
	((QTableWidget__Go*)self)->destroy(destroyWindow);
}
void QTableWidget_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QTableWidget__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QTableWidget_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTableWidget::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTableWidget_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTableWidget::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QTableWidget_metaObject_default(void* self) {
	auto&& __r = ((QTableWidget*)self)->QTableWidget::metaObject();
	return (void*)(__r);
}
int QTableWidget_event_default(void* self, void* e) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
char** QTableWidget_mimeTypes_default(void* self, long long* __out_len) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_mimeTypes();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
int QTableWidget_dropMimeData_default(void* self, int row, int column, void* data, int action) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_dropMimeData(row, column, (QMimeData*)data, action);
	return __r;
}
int QTableWidget_supportedDropActions_default(void* self) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_supportedDropActions();
	return (int)(__r);
}
void QTableWidget_dropEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QTableWidget_setRootIndex_default(void* self, void* index) {
	((QTableWidget*)self)->QTableWidget::setRootIndex(*(QModelIndex*)index);
}
void QTableWidget_setSelectionModel_default(void* self, void* selectionModel) {
	((QTableWidget*)self)->QTableWidget::setSelectionModel((QItemSelectionModel*)selectionModel);
}
void QTableWidget_doItemsLayout_default(void* self) {
	((QTableWidget*)self)->QTableWidget::doItemsLayout();
}
void* QTableWidget_visualRect_default(void* self, void* index) {
	auto&& __r = ((QTableWidget*)self)->QTableWidget::visualRect(*(QModelIndex*)index);
	return new QRect(__r);
}
void QTableWidget_scrollTo_default(void* self, void* index, int hint) {
	((QTableWidget*)self)->QTableWidget::scrollTo(*(QModelIndex*)index, (QAbstractItemView::ScrollHint)(hint));
}
void* QTableWidget_indexAt_default(void* self, void* p) {
	auto&& __r = ((QTableWidget*)self)->QTableWidget::indexAt(*(QPoint*)p);
	return new QModelIndex(__r);
}
void QTableWidget_scrollContentsBy_default(void* self, int dx, int dy) {
	((QTableWidget__Go*)self)->cppgoBase_scrollContentsBy(dx, dy);
}
void QTableWidget_initViewItemOption_default(void* self, void* option) {
	((QTableWidget__Go*)self)->cppgoBase_initViewItemOption((QStyleOptionViewItem*)option);
}
void QTableWidget_paintEvent_default(void* self, void* e) {
	((QTableWidget__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)e);
}
void QTableWidget_timerEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
int QTableWidget_horizontalOffset_default(void* self) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_horizontalOffset();
	return __r;
}
int QTableWidget_verticalOffset_default(void* self) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_verticalOffset();
	return __r;
}
void* QTableWidget_moveCursor_default(void* self, int cursorAction, int modifiers) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_moveCursor(cursorAction, QFlags<Qt::KeyboardModifier>(modifiers));
	return new QModelIndex(__r);
}
void QTableWidget_setSelection_default(void* self, void* rect, int command) {
	((QTableWidget__Go*)self)->cppgoBase_setSelection(*(QRect*)rect, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void* QTableWidget_visualRegionForSelection_default(void* self, void* selection) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_visualRegionForSelection(*(QItemSelection*)selection);
	return new QRegion(__r);
}
void** QTableWidget_selectedIndexes_default(void* self, long long* __out_len) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_selectedIndexes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void QTableWidget_updateGeometries_default(void* self) {
	((QTableWidget__Go*)self)->cppgoBase_updateGeometries();
}
void* QTableWidget_viewportSizeHint_default(void* self) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_viewportSizeHint();
	return new QSize(__r);
}
int QTableWidget_sizeHintForRow_default(void* self, int row) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_sizeHintForRow(row);
	return __r;
}
int QTableWidget_sizeHintForColumn_default(void* self, int column) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_sizeHintForColumn(column);
	return __r;
}
void QTableWidget_verticalScrollbarAction_default(void* self, int action) {
	((QTableWidget__Go*)self)->cppgoBase_verticalScrollbarAction(action);
}
void QTableWidget_horizontalScrollbarAction_default(void* self, int action) {
	((QTableWidget__Go*)self)->cppgoBase_horizontalScrollbarAction(action);
}
int QTableWidget_isIndexHidden_default(void* self, void* index) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_isIndexHidden(*(QModelIndex*)index);
	return __r;
}
void QTableWidget_selectionChanged_default(void* self, void* selected, void* deselected) {
	((QTableWidget__Go*)self)->cppgoBase_selectionChanged(*(QItemSelection*)selected, *(QItemSelection*)deselected);
}
void QTableWidget_currentChanged_default(void* self, void* current, void* previous) {
	((QTableWidget__Go*)self)->cppgoBase_currentChanged(*(QModelIndex*)current, *(QModelIndex*)previous);
}
void QTableWidget_keyboardSearch_default(void* self, const char* search, long long search_len) {
	QString search_qstr = QString::fromUtf8(search, search_len);
	((QTableWidget*)self)->QTableWidget::keyboardSearch(search_qstr);
}
void* QTableWidget_itemDelegateForIndex_default(void* self, void* index) {
	auto&& __r = ((QTableWidget*)self)->QTableWidget::itemDelegateForIndex(*(QModelIndex*)index);
	return (void*)(__r);
}
void* QTableWidget_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QTableWidget*)self)->QTableWidget::inputMethodQuery((Qt::InputMethodQuery)(query));
	return new QVariant(__r);
}
void QTableWidget_reset_default(void* self) {
	((QTableWidget*)self)->QTableWidget::reset();
}
void QTableWidget_selectAll_default(void* self) {
	((QTableWidget*)self)->QTableWidget::selectAll();
}
void QTableWidget_rowsInserted_default(void* self, void* parent, int start, int end) {
	((QTableWidget__Go*)self)->cppgoBase_rowsInserted(*(QModelIndex*)parent, start, end);
}
void QTableWidget_rowsAboutToBeRemoved_default(void* self, void* parent, int start, int end) {
	((QTableWidget__Go*)self)->cppgoBase_rowsAboutToBeRemoved(*(QModelIndex*)parent, start, end);
}
void QTableWidget_updateEditorData_default(void* self) {
	((QTableWidget__Go*)self)->cppgoBase_updateEditorData();
}
void QTableWidget_updateEditorGeometries_default(void* self) {
	((QTableWidget__Go*)self)->cppgoBase_updateEditorGeometries();
}
void QTableWidget_verticalScrollbarValueChanged_default(void* self, int value) {
	((QTableWidget__Go*)self)->cppgoBase_verticalScrollbarValueChanged(value);
}
void QTableWidget_horizontalScrollbarValueChanged_default(void* self, int value) {
	((QTableWidget__Go*)self)->cppgoBase_horizontalScrollbarValueChanged(value);
}
void QTableWidget_closeEditor_default(void* self, void* editor, int hint) {
	((QTableWidget__Go*)self)->cppgoBase_closeEditor((QWidget*)editor, hint);
}
void QTableWidget_commitData_default(void* self, void* editor) {
	((QTableWidget__Go*)self)->cppgoBase_commitData((QWidget*)editor);
}
void QTableWidget_editorDestroyed_default(void* self, void* editor) {
	((QTableWidget__Go*)self)->cppgoBase_editorDestroyed((QObject*)editor);
}
int QTableWidget_edit_default(void* self, void* index, int trigger, void* event) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_edit(*(QModelIndex*)index, trigger, (QEvent*)event);
	return __r;
}
int QTableWidget_selectionCommand_default(void* self, void* index, void* event) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_selectionCommand(*(QModelIndex*)index, (QEvent*)event);
	return (int)(__r);
}
void QTableWidget_startDrag_default(void* self, int supportedActions) {
	((QTableWidget__Go*)self)->cppgoBase_startDrag(QFlags<Qt::DropAction>(supportedActions));
}
int QTableWidget_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QTableWidget_viewportEvent_default(void* self, void* event) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_viewportEvent((QEvent*)event);
	return __r;
}
void QTableWidget_mousePressEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QTableWidget_mouseMoveEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QTableWidget_mouseReleaseEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QTableWidget_mouseDoubleClickEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QTableWidget_dragEnterEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QTableWidget_dragMoveEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QTableWidget_dragLeaveEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QTableWidget_focusInEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QTableWidget_focusOutEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QTableWidget_keyPressEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QTableWidget_resizeEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QTableWidget_inputMethodEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
int QTableWidget_eventFilter_default(void* self, void* object, void* event) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_eventFilter((QObject*)object, (QEvent*)event);
	return __r;
}
void* QTableWidget_minimumSizeHint_default(void* self) {
	auto&& __r = ((QTableWidget*)self)->QTableWidget::minimumSizeHint();
	return new QSize(__r);
}
void* QTableWidget_sizeHint_default(void* self) {
	auto&& __r = ((QTableWidget*)self)->QTableWidget::sizeHint();
	return new QSize(__r);
}
void QTableWidget_setupViewport_default(void* self, void* viewport) {
	((QTableWidget*)self)->QTableWidget::setupViewport((QWidget*)viewport);
}
void QTableWidget_wheelEvent_default(void* self, void* arg1) {
	((QTableWidget__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)arg1);
}
void QTableWidget_contextMenuEvent_default(void* self, void* arg1) {
	((QTableWidget__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
void QTableWidget_changeEvent_default(void* self, void* arg1) {
	((QTableWidget__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QTableWidget_initStyleOption_default(void* self, void* option) {
	((QTableWidget__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QTableWidget_devType_default(void* self) {
	auto&& __r = ((QTableWidget*)self)->QTableWidget::devType();
	return __r;
}
void QTableWidget_setVisible_default(void* self, int visible) {
	((QTableWidget*)self)->QTableWidget::setVisible(visible);
}
int QTableWidget_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QTableWidget*)self)->QTableWidget::heightForWidth(arg1);
	return __r;
}
int QTableWidget_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QTableWidget*)self)->QTableWidget::hasHeightForWidth();
	return __r;
}
void* QTableWidget_paintEngine_default(void* self) {
	auto&& __r = ((QTableWidget*)self)->QTableWidget::paintEngine();
	return (void*)(__r);
}
void QTableWidget_keyReleaseEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QTableWidget_enterEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QTableWidget_leaveEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QTableWidget_moveEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QTableWidget_closeEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QTableWidget_tabletEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QTableWidget_actionEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QTableWidget_showEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QTableWidget_hideEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QTableWidget_metric_default(void* self, int arg1) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QTableWidget_initPainter_default(void* self, void* painter) {
	((QTableWidget__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QTableWidget_redirected_default(void* self, void* offset) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QTableWidget_sharedPainter_default(void* self) {
	auto&& __r = ((QTableWidget__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QTableWidget_childEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QTableWidget_customEvent_default(void* self, void* event) {
	((QTableWidget__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QTableWidget_connectNotify_default(void* self, void* signal) {
	((QTableWidget__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QTableWidget_disconnectNotify_default(void* self, void* signal) {
	((QTableWidget__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QTableWidget_connect_itemPressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::itemPressed, (QTableWidget*)self, [__id](QTableWidgetItem* item){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)item);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_itemClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::itemClicked, (QTableWidget*)self, [__id](QTableWidgetItem* item){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)item);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_itemDoubleClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::itemDoubleClicked, (QTableWidget*)self, [__id](QTableWidgetItem* item){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)item);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_itemActivated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::itemActivated, (QTableWidget*)self, [__id](QTableWidgetItem* item){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)item);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_itemEntered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::itemEntered, (QTableWidget*)self, [__id](QTableWidgetItem* item){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)item);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_itemChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::itemChanged, (QTableWidget*)self, [__id](QTableWidgetItem* item){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)item);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_currentItemChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::currentItemChanged, (QTableWidget*)self, [__id](QTableWidgetItem* current, QTableWidgetItem* previous){
		if (cppgo_sig_void_voidp_voidp) ((void(*)(long long, void*, void*))cppgo_sig_void_voidp_voidp)(__id, (void*)current, (void*)previous);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_itemSelectionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::itemSelectionChanged, (QTableWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_cellPressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::cellPressed, (QTableWidget*)self, [__id](int row, int column){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, row, column);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_cellClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::cellClicked, (QTableWidget*)self, [__id](int row, int column){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, row, column);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_cellDoubleClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::cellDoubleClicked, (QTableWidget*)self, [__id](int row, int column){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, row, column);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_cellActivated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::cellActivated, (QTableWidget*)self, [__id](int row, int column){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, row, column);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_cellEntered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::cellEntered, (QTableWidget*)self, [__id](int row, int column){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, row, column);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_cellChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::cellChanged, (QTableWidget*)self, [__id](int row, int column){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, row, column);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_currentCellChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::currentCellChanged, (QTableWidget*)self, [__id](int currentRow, int currentColumn, int previousRow, int previousColumn){
		if (cppgo_sig_void_int_int_int_int) ((void(*)(long long, int, int, int, int))cppgo_sig_void_int_int_int_int)(__id, currentRow, currentColumn, previousRow, previousColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_pressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::pressed, (QTableWidget*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::clicked, (QTableWidget*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_doubleClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::doubleClicked, (QTableWidget*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_activated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::activated, (QTableWidget*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_entered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::entered, (QTableWidget*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_viewportEntered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::viewportEntered, (QTableWidget*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_iconSizeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::iconSizeChanged, (QTableWidget*)self, [__id](const QSize& size){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&size);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::windowTitleChanged, (QTableWidget*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::windowIconChanged, (QTableWidget*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::windowIconTextChanged, (QTableWidget*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::customContextMenuRequested, (QTableWidget*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::destroyed, (QTableWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::destroyed, (QTableWidget*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTableWidget*)self, &QTableWidget::objectNameChanged, (QTableWidget*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
	__conns_QTableWidget[__id] = __c;
}
void QTableWidget__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTableWidget);
		auto __it = __conns_QTableWidget.find(id);
		if (__it == __conns_QTableWidget.end()) return;
		__c = __it->second;
		__conns_QTableWidget.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QTableWidget__Go_new(long long id) { return new QTableWidget__Go(id); }
void QTableWidget_delete(void* self) { delete (QTableWidget*)self; }

}

