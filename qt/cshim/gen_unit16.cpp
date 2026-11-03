#include <QtCore/qobjectdefs.h>
#include "gen_qmetaobject__superdata.h"

extern "C" {

void* QMetaObject__SuperData_new() {
	return new QMetaObject::SuperData();
}
void* QMetaObject__SuperData_new2(void* mo) {
	return new QMetaObject::SuperData((QMetaObject*)mo);
}
void* QMetaObject__SuperData_new3(void* arg1) {
	return new QMetaObject::SuperData(*(QMetaObject::SuperData*)arg1);
}
void* QMetaObject__SuperData_OperatorMinusGreater(void* self) {
	auto&& __r = ((QMetaObject::SuperData*)self)->operator->();
	return (void*)(__r);
}
void* QMetaObject__SuperData_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QMetaObject::SuperData*)self)->operator=(*(QMetaObject::SuperData*)arg1);
	return (void*)&(__r);
}
void* QMetaObject__SuperData_direct_get(void* self) {
	auto __r = ((QMetaObject::SuperData*)self)->direct;
	return (void*)(__r);
}
void QMetaObject__SuperData_direct_set(void* self, void* direct) {
	((QMetaObject::SuperData*)self)->direct = (QMetaObject*)direct;
}
void QMetaObject__SuperData_delete(void* self) { delete (QMetaObject::SuperData*)self; }

}

#include <QtCore/qmetacontainer.h>
#include "gen_qmetacontainer.h"

extern "C" {

void* QMetaContainer_new() {
	return new QMetaContainer();
}
void* QMetaContainer_new2(void* arg1) {
	return new QMetaContainer(*(QMetaContainer*)arg1);
}
int QMetaContainer_hasInputIterator(void* self) {
	auto&& __r = ((QMetaContainer*)self)->hasInputIterator();
	return __r;
}
int QMetaContainer_hasForwardIterator(void* self) {
	auto&& __r = ((QMetaContainer*)self)->hasForwardIterator();
	return __r;
}
int QMetaContainer_hasBidirectionalIterator(void* self) {
	auto&& __r = ((QMetaContainer*)self)->hasBidirectionalIterator();
	return __r;
}
int QMetaContainer_hasRandomAccessIterator(void* self) {
	auto&& __r = ((QMetaContainer*)self)->hasRandomAccessIterator();
	return __r;
}
int QMetaContainer_hasSize(void* self) {
	auto&& __r = ((QMetaContainer*)self)->hasSize();
	return __r;
}
int QMetaContainer_canClear(void* self) {
	auto&& __r = ((QMetaContainer*)self)->canClear();
	return __r;
}
int QMetaContainer_hasIterator(void* self) {
	auto&& __r = ((QMetaContainer*)self)->hasIterator();
	return __r;
}
int QMetaContainer_hasConstIterator(void* self) {
	auto&& __r = ((QMetaContainer*)self)->hasConstIterator();
	return __r;
}
void QMetaContainer_delete(void* self) { delete (QMetaContainer*)self; }

}

#include <QtCore/qvarlengtharray.h>
#include "gen_qvlabasebase.h"

extern "C" {

long long QVLABaseBase_capacity(void* self) {
	auto&& __r = ((QVLABaseBase*)self)->capacity();
	return __r;
}
long long QVLABaseBase_size(void* self) {
	auto&& __r = ((QVLABaseBase*)self)->size();
	return __r;
}
int QVLABaseBase_empty(void* self) {
	auto&& __r = ((QVLABaseBase*)self)->empty();
	return __r;
}

}

#include <QtCore/qcoreapplication.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qeventloop.h>
#include <QtCore/qdeadlinetimer.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qabstracteventdispatcher.h>
#include <QtCore/qpermissions.h>
#include <QtCore/qtranslator.h>
#include <QtCore/qabstractnativeeventfilter.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qcoreapplication.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QCoreApplication;
static std::mutex __connsMu_QCoreApplication;

extern "C" {

void* QCoreApplication_base(void* self) { return (void*)static_cast<QObject*>((QCoreApplication*)self); }
void* QCoreApplication_fromQObject(void* o) { return (void*)qobject_cast<QCoreApplication*>((QObject*)o); }
char* QCoreApplication_tr(const char* s, long long* __out_len) {
	auto&& __r = QCoreApplication::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QCoreApplication_arguments(long long* __out_len) {
	auto&& __r = QCoreApplication::arguments();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QCoreApplication_setAttribute(int attribute) {
	QCoreApplication::setAttribute((Qt::ApplicationAttribute)(attribute));
}
int QCoreApplication_testAttribute(int attribute) {
	auto&& __r = QCoreApplication::testAttribute((Qt::ApplicationAttribute)(attribute));
	return __r;
}
void QCoreApplication_setOrganizationDomain(const char* orgDomain, long long orgDomain_len) {
	QString orgDomain_qstr = QString::fromUtf8(orgDomain, orgDomain_len);
	QCoreApplication::setOrganizationDomain(orgDomain_qstr);
}
char* QCoreApplication_organizationDomain(long long* __out_len) {
	auto&& __r = QCoreApplication::organizationDomain();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCoreApplication_setOrganizationName(const char* orgName, long long orgName_len) {
	QString orgName_qstr = QString::fromUtf8(orgName, orgName_len);
	QCoreApplication::setOrganizationName(orgName_qstr);
}
char* QCoreApplication_organizationName(long long* __out_len) {
	auto&& __r = QCoreApplication::organizationName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCoreApplication_setApplicationName(const char* application, long long application_len) {
	QString application_qstr = QString::fromUtf8(application, application_len);
	QCoreApplication::setApplicationName(application_qstr);
}
char* QCoreApplication_applicationName(long long* __out_len) {
	auto&& __r = QCoreApplication::applicationName();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCoreApplication_setApplicationVersion(const char* version, long long version_len) {
	QString version_qstr = QString::fromUtf8(version, version_len);
	QCoreApplication::setApplicationVersion(version_qstr);
}
char* QCoreApplication_applicationVersion(long long* __out_len) {
	auto&& __r = QCoreApplication::applicationVersion();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCoreApplication_setSetuidAllowed(int allow) {
	QCoreApplication::setSetuidAllowed(allow);
}
int QCoreApplication_isSetuidAllowed() {
	auto&& __r = QCoreApplication::isSetuidAllowed();
	return __r;
}
void* QCoreApplication_instance() {
	auto&& __r = QCoreApplication::instance();
	return (void*)(__r);
}
int QCoreApplication_instanceExists() {
	auto&& __r = QCoreApplication::instanceExists();
	return __r;
}
int QCoreApplication_exec() {
	auto&& __r = QCoreApplication::exec();
	return __r;
}
void QCoreApplication_processEvents() {
	QCoreApplication::processEvents();
}
void QCoreApplication_processEvents2(int flags, int maxtime) {
	QCoreApplication::processEvents(QFlags<QEventLoop::ProcessEventsFlag>(flags), maxtime);
}
void QCoreApplication_processEvents3(int flags, void* deadline) {
	QCoreApplication::processEvents(QFlags<QEventLoop::ProcessEventsFlag>(flags), *(QDeadlineTimer*)deadline);
}
int QCoreApplication_sendEvent(void* receiver, void* event) {
	auto&& __r = QCoreApplication::sendEvent((QObject*)receiver, (QEvent*)event);
	return __r;
}
void QCoreApplication_postEvent(void* receiver, void* event) {
	QCoreApplication::postEvent((QObject*)receiver, (QEvent*)event);
}
void QCoreApplication_sendPostedEvents() {
	QCoreApplication::sendPostedEvents();
}
void QCoreApplication_removePostedEvents(void* receiver) {
	QCoreApplication::removePostedEvents((QObject*)receiver);
}
void* QCoreApplication_eventDispatcher() {
	auto&& __r = QCoreApplication::eventDispatcher();
	return (void*)(__r);
}
void QCoreApplication_setEventDispatcher(void* eventDispatcher) {
	QCoreApplication::setEventDispatcher((QAbstractEventDispatcher*)eventDispatcher);
}
int QCoreApplication_notify(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QCoreApplication*)self)->notify((QObject*)arg1, (QEvent*)arg2);
	return __r;
}
int QCoreApplication_startingUp() {
	auto&& __r = QCoreApplication::startingUp();
	return __r;
}
int QCoreApplication_closingDown() {
	auto&& __r = QCoreApplication::closingDown();
	return __r;
}
char* QCoreApplication_applicationDirPath(long long* __out_len) {
	auto&& __r = QCoreApplication::applicationDirPath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCoreApplication_applicationFilePath(long long* __out_len) {
	auto&& __r = QCoreApplication::applicationFilePath();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
long long QCoreApplication_applicationPid() {
	auto&& __r = QCoreApplication::applicationPid();
	return __r;
}
int QCoreApplication_checkPermission(void* self, void* permission) {
	auto&& __r = ((QCoreApplication*)self)->checkPermission(*(QPermission*)permission);
	return (int)(__r);
}
void QCoreApplication_setLibraryPaths(char** arg1, long long arg1_len) {
	QList<QString> arg1_list; arg1_list.reserve(arg1_len);
	for (long long __i = 0; __i < arg1_len; __i++) arg1_list.append(QString::fromUtf8(arg1[__i]));
	QCoreApplication::setLibraryPaths(arg1_list);
}
char** QCoreApplication_libraryPaths(long long* __out_len) {
	auto&& __r = QCoreApplication::libraryPaths();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QCoreApplication_addLibraryPath(const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	QCoreApplication::addLibraryPath(arg1_qstr);
}
void QCoreApplication_removeLibraryPath(const char* arg1, long long arg1_len) {
	QString arg1_qstr = QString::fromUtf8(arg1, arg1_len);
	QCoreApplication::removeLibraryPath(arg1_qstr);
}
int QCoreApplication_installTranslator(void* messageFile) {
	auto&& __r = QCoreApplication::installTranslator((QTranslator*)messageFile);
	return __r;
}
int QCoreApplication_removeTranslator(void* messageFile) {
	auto&& __r = QCoreApplication::removeTranslator((QTranslator*)messageFile);
	return __r;
}
char* QCoreApplication_translate(const char* context, const char* key, long long* __out_len) {
	auto&& __r = QCoreApplication::translate(context, key);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCoreApplication_installNativeEventFilter(void* self, void* filterObj) {
	((QCoreApplication*)self)->installNativeEventFilter((QAbstractNativeEventFilter*)filterObj);
}
void QCoreApplication_removeNativeEventFilter(void* self, void* filterObj) {
	((QCoreApplication*)self)->removeNativeEventFilter((QAbstractNativeEventFilter*)filterObj);
}
int QCoreApplication_isQuitLockEnabled() {
	auto&& __r = QCoreApplication::isQuitLockEnabled();
	return __r;
}
void QCoreApplication_setQuitLockEnabled(int enabled) {
	QCoreApplication::setQuitLockEnabled(enabled);
}
void QCoreApplication_quit() {
	QCoreApplication::quit();
}
void QCoreApplication_exit() {
	QCoreApplication::exit();
}
void QCoreApplication_organizationNameChanged(void* self) {
	((QCoreApplication*)self)->organizationNameChanged();
}
void QCoreApplication_organizationDomainChanged(void* self) {
	((QCoreApplication*)self)->organizationDomainChanged();
}
void QCoreApplication_applicationNameChanged(void* self) {
	((QCoreApplication*)self)->applicationNameChanged();
}
void QCoreApplication_applicationVersionChanged(void* self) {
	((QCoreApplication*)self)->applicationVersionChanged();
}
void* QCoreApplication_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QCoreApplication::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QCoreApplication_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QCoreApplication::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QCoreApplication_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QCoreApplication::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCoreApplication_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QCoreApplication::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCoreApplication_setAttribute2(int attribute, int on) {
	QCoreApplication::setAttribute((Qt::ApplicationAttribute)(attribute), on);
}
void QCoreApplication_processEvents4(int flags) {
	QCoreApplication::processEvents(QFlags<QEventLoop::ProcessEventsFlag>(flags));
}
void QCoreApplication_postEvent2(void* receiver, void* event, int priority) {
	QCoreApplication::postEvent((QObject*)receiver, (QEvent*)event, priority);
}
void QCoreApplication_sendPostedEvents2(void* receiver) {
	QCoreApplication::sendPostedEvents((QObject*)receiver);
}
void QCoreApplication_sendPostedEvents3(void* receiver, int event_type) {
	QCoreApplication::sendPostedEvents((QObject*)receiver, event_type);
}
void QCoreApplication_removePostedEvents2(void* receiver, int eventType) {
	QCoreApplication::removePostedEvents((QObject*)receiver, eventType);
}
char* QCoreApplication_translate2(const char* context, const char* key, const char* disambiguation, long long* __out_len) {
	auto&& __r = QCoreApplication::translate(context, key, disambiguation);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QCoreApplication_translate3(const char* context, const char* key, const char* disambiguation, int n, long long* __out_len) {
	auto&& __r = QCoreApplication::translate(context, key, disambiguation, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QCoreApplication_exit2(int retcode) {
	QCoreApplication::exit(retcode);
}
void* QCoreApplication_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QCoreApplication::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QCoreApplication_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QCoreApplication::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QCoreApplication_connect_aboutToQuit(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCoreApplication*)self, &QCoreApplication::aboutToQuit, (QCoreApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCoreApplication);
	__conns_QCoreApplication[__id] = __c;
}
void QCoreApplication_connect_organizationNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCoreApplication*)self, &QCoreApplication::organizationNameChanged, (QCoreApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCoreApplication);
	__conns_QCoreApplication[__id] = __c;
}
void QCoreApplication_connect_organizationDomainChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCoreApplication*)self, &QCoreApplication::organizationDomainChanged, (QCoreApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCoreApplication);
	__conns_QCoreApplication[__id] = __c;
}
void QCoreApplication_connect_applicationNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCoreApplication*)self, &QCoreApplication::applicationNameChanged, (QCoreApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCoreApplication);
	__conns_QCoreApplication[__id] = __c;
}
void QCoreApplication_connect_applicationVersionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCoreApplication*)self, &QCoreApplication::applicationVersionChanged, (QCoreApplication*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCoreApplication);
	__conns_QCoreApplication[__id] = __c;
}
void QCoreApplication_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCoreApplication*)self, &QCoreApplication::destroyed, (QCoreApplication*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCoreApplication);
	__conns_QCoreApplication[__id] = __c;
}
void QCoreApplication_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCoreApplication*)self, &QCoreApplication::destroyed, (QCoreApplication*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCoreApplication);
	__conns_QCoreApplication[__id] = __c;
}
void QCoreApplication_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QCoreApplication*)self, &QCoreApplication::objectNameChanged, (QCoreApplication*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QCoreApplication);
	__conns_QCoreApplication[__id] = __c;
}
void QCoreApplication__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QCoreApplication);
		auto __it = __conns_QCoreApplication.find(id);
		if (__it == __conns_QCoreApplication.end()) return;
		__c = __it->second;
		__conns_QCoreApplication.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QCoreApplication_new_args(int* argc, char** argv) { return new QCoreApplication(*argc, argv); }
void QCoreApplication_delete(void* self) { delete (QCoreApplication*)self; }

}

#include <QtCore/qcborvalue.h>
#include <QtCore/qstring.h>
#include <QtCore/qcborcommon.h>
#include "gen_qcborparsererror.h"

extern "C" {

char* QCborParserError_errorString(void* self, long long* __out_len) {
	auto&& __r = ((QCborParserError*)self)->errorString();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
long long QCborParserError_offset_get(void* self) {
	auto __r = ((QCborParserError*)self)->offset;
	return __r;
}
void QCborParserError_offset_set(void* self, long long offset) {
	((QCborParserError*)self)->offset = offset;
}
void* QCborParserError_error_get(void* self) {
	auto __r = ((QCborParserError*)self)->error;
	return new QCborError(__r);
}
void QCborParserError_error_set(void* self, void* error_) {
	((QCborParserError*)self)->error = *(QCborError*)error_;
}
void QCborParserError_delete(void* self) { delete (QCborParserError*)self; }

}

#include <QtCore/qconcatenatetablesproxymodel.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qvariant.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmimedata.h>
#include <QtCore/qsize.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qdatastream.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qconcatenatetablesproxymodel.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QConcatenateTablesProxyModel__Go : public QConcatenateTablesProxyModel {
public:
	long long __goId;
	unsigned char __goOn[42] = {};
	template<typename... A> QConcatenateTablesProxyModel__Go(long long id, A&&... a) : QConcatenateTablesProxyModel(std::forward<A>(a)...), __goId(id) {}
	using QConcatenateTablesProxyModel::createIndex;
	using QConcatenateTablesProxyModel::encodeData;
	using QConcatenateTablesProxyModel::decodeData;
	using QConcatenateTablesProxyModel::beginInsertRows;
	using QConcatenateTablesProxyModel::endInsertRows;
	using QConcatenateTablesProxyModel::beginRemoveRows;
	using QConcatenateTablesProxyModel::endRemoveRows;
	using QConcatenateTablesProxyModel::beginMoveRows;
	using QConcatenateTablesProxyModel::endMoveRows;
	using QConcatenateTablesProxyModel::beginInsertColumns;
	using QConcatenateTablesProxyModel::endInsertColumns;
	using QConcatenateTablesProxyModel::beginRemoveColumns;
	using QConcatenateTablesProxyModel::endRemoveColumns;
	using QConcatenateTablesProxyModel::beginMoveColumns;
	using QConcatenateTablesProxyModel::endMoveColumns;
	using QConcatenateTablesProxyModel::beginResetModel;
	using QConcatenateTablesProxyModel::endResetModel;
	using QConcatenateTablesProxyModel::changePersistentIndex;
	using QConcatenateTablesProxyModel::changePersistentIndexList;
	using QConcatenateTablesProxyModel::persistentIndexList;
	using QConcatenateTablesProxyModel::sender;
	using QConcatenateTablesProxyModel::senderSignalIndex;
	using QConcatenateTablesProxyModel::receivers;
	using QConcatenateTablesProxyModel::isSignalConnected;
	void cppgoBase_resetInternalData() { QConcatenateTablesProxyModel::resetInternalData(); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QConcatenateTablesProxyModel::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QConcatenateTablesProxyModel::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QConcatenateTablesProxyModel::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QConcatenateTablesProxyModel::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QConcatenateTablesProxyModel::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QConcatenateTablesProxyModel::metaObject();
	}
	QVariant data(const QModelIndex& index, int role) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_voidp_int) { void* __r = ((void*(*)(long long, long long, void*, int, int*))cppgo_virt_voidp_voidp_int)(__goId, -8837375942098466299LL, (void*)&index, role, &__h); if (__h) return *(QVariant*)__r; }
		return QConcatenateTablesProxyModel::data(index, role);
	}
	bool setData(const QModelIndex& index, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp_int) { int __r = ((int(*)(long long, long long, void*, void*, int, int*))cppgo_virt_int_voidp_voidp_int)(__goId, 35191454575202219LL, (void*)&index, (void*)&value, role, &__h); if (__h) return __r; }
		return QConcatenateTablesProxyModel::setData(index, value, role);
	}
	QFlags<Qt::ItemFlag> flags(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1703382833545304812LL, (void*)&index, &__h); if (__h) return QFlags<Qt::ItemFlag>(__r); }
		return QConcatenateTablesProxyModel::flags(index);
	}
	QModelIndex index(int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, -8948777187306027381LL, row, column, (void*)&parent, &__h); if (__h) return *(QModelIndex*)__r; }
		return QConcatenateTablesProxyModel::index(row, column, parent);
	}
	QModelIndex parent(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 7371339085820260189LL, (void*)&index, &__h); if (__h) return *(QModelIndex*)__r; }
		return QConcatenateTablesProxyModel::parent(index);
	}
	int rowCount(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1960960171956884010LL, (void*)&parent, &__h); if (__h) return __r; }
		return QConcatenateTablesProxyModel::rowCount(parent);
	}
	QVariant headerData(int section, Qt::Orientation orientation, int role) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp_int_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int, int*))cppgo_virt_voidp_int_int_int)(__goId, -3297140021896693688LL, section, (int)(orientation), role, &__h); if (__h) return *(QVariant*)__r; }
		return QConcatenateTablesProxyModel::headerData(section, orientation, role);
	}
	int columnCount(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -8256963024669584834LL, (void*)&parent, &__h); if (__h) return __r; }
		return QConcatenateTablesProxyModel::columnCount(parent);
	}
	QList<QString> mimeTypes() const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virtp_voidp) {
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
		return QConcatenateTablesProxyModel::mimeTypes();
	}
	QMimeData* mimeData(const QList<QModelIndex>& indexes) const override {
		int __h = 0;
		std::vector<void*> indexes_v; indexes_v.reserve(indexes.size());
		for (long long __i = 0; __i < indexes.size(); __i++) indexes_v.push_back((void*)&indexes[__i]);
		if (__goOn[10] && cppgo_virt_voidp_voidpp_longlong) { void* __r = ((void*(*)(long long, long long, void**, long long, int*))cppgo_virt_voidp_voidpp_longlong)(__goId, 5175297091024974749LL, indexes_v.data(), (long long)indexes.size(), &__h); if (__h) return (QMimeData*)__r; }
		return QConcatenateTablesProxyModel::mimeData(indexes);
	}
	bool canDropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, 731539533644753052LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QConcatenateTablesProxyModel::canDropMimeData(data, action, row, column, parent);
	}
	bool dropMimeData(const QMimeData* data, Qt::DropAction action, int row, int column, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_voidp_int_int_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, int, int, void*, int*))cppgo_virt_int_voidp_int_int_int_voidp)(__goId, -9049994684305682212LL, (void*)data, (int)(action), row, column, (void*)&parent, &__h); if (__h) return __r; }
		return QConcatenateTablesProxyModel::dropMimeData(data, action, row, column, parent);
	}
	QSize span(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -8395342916587040543LL, (void*)&index, &__h); if (__h) return *(QSize*)__r; }
		return QConcatenateTablesProxyModel::span(index);
	}
	QModelIndex sibling(int row, int column, const QModelIndex& idx) const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_voidp_int_int_voidp) { void* __r = ((void*(*)(long long, long long, int, int, void*, int*))cppgo_virt_voidp_int_int_voidp)(__goId, 397329346813782521LL, row, column, (void*)&idx, &__h); if (__h) return *(QModelIndex*)__r; }
		return QConcatenateTablesProxyModel::sibling(row, column, idx);
	}
	bool hasChildren(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 3369452259144514380LL, (void*)&parent, &__h); if (__h) return __r; }
		return QConcatenateTablesProxyModel::hasChildren(parent);
	}
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant& value, int role) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_int_int_int_voidp_int) { int __r = ((int(*)(long long, long long, int, int, void*, int, int*))cppgo_virt_int_int_int_voidp_int)(__goId, 4852200306590054850LL, section, (int)(orientation), (void*)&value, role, &__h); if (__h) return __r; }
		return QConcatenateTablesProxyModel::setHeaderData(section, orientation, value, role);
	}
	bool clearItemData(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8914190500520542409LL, (void*)&index, &__h); if (__h) return __r; }
		return QConcatenateTablesProxyModel::clearItemData(index);
	}
	QFlags<Qt::DropAction> supportedDropActions() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -7784642176186821205LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QConcatenateTablesProxyModel::supportedDropActions();
	}
	QFlags<Qt::DropAction> supportedDragActions() const override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 3421601386088570762LL, &__h); if (__h) return QFlags<Qt::DropAction>(__r); }
		return QConcatenateTablesProxyModel::supportedDragActions();
	}
	bool insertRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -2775729825962395603LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QConcatenateTablesProxyModel::insertRows(row, count, parent);
	}
	bool insertColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, -8973293846230792237LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QConcatenateTablesProxyModel::insertColumns(column, count, parent);
	}
	bool removeRows(int row, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 5104072549182222656LL, row, count, (void*)&parent, &__h); if (__h) return __r; }
		return QConcatenateTablesProxyModel::removeRows(row, count, parent);
	}
	bool removeColumns(int column, int count, const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_int_int_int_voidp) { int __r = ((int(*)(long long, long long, int, int, void*, int*))cppgo_virt_int_int_int_voidp)(__goId, 2078548737056540420LL, column, count, (void*)&parent, &__h); if (__h) return __r; }
		return QConcatenateTablesProxyModel::removeColumns(column, count, parent);
	}
	bool moveRows(const QModelIndex& sourceParent, int sourceRow, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6074819356947405415LL, (void*)&sourceParent, sourceRow, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QConcatenateTablesProxyModel::moveRows(sourceParent, sourceRow, count, destinationParent, destinationChild);
	}
	bool moveColumns(const QModelIndex& sourceParent, int sourceColumn, int count, const QModelIndex& destinationParent, int destinationChild) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_voidp_int_int_voidp_int) { int __r = ((int(*)(long long, long long, void*, int, int, void*, int, int*))cppgo_virt_int_voidp_int_int_voidp_int)(__goId, -6321007591388309905LL, (void*)&sourceParent, sourceColumn, count, (void*)&destinationParent, destinationChild, &__h); if (__h) return __r; }
		return QConcatenateTablesProxyModel::moveColumns(sourceParent, sourceColumn, count, destinationParent, destinationChild);
	}
	void fetchMore(const QModelIndex& parent) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3553294593611130806LL, (void*)&parent, &__h); if (__h) return; }
		QConcatenateTablesProxyModel::fetchMore(parent);
	}
	bool canFetchMore(const QModelIndex& parent) const override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 8639262066473964150LL, (void*)&parent, &__h); if (__h) return __r; }
		return QConcatenateTablesProxyModel::canFetchMore(parent);
	}
	void sort(int column, Qt::SortOrder order) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 6912441694453382161LL, column, (int)(order), &__h); if (__h) return; }
		QConcatenateTablesProxyModel::sort(column, order);
	}
	QModelIndex buddy(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 6112860112248786499LL, (void*)&index, &__h); if (__h) return *(QModelIndex*)__r; }
		return QConcatenateTablesProxyModel::buddy(index);
	}
	QList<QModelIndex> match(const QModelIndex& start, int role, const QVariant& value, int hits, QFlags<Qt::MatchFlag> flags) const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virtp_voidp_voidp_int_voidp_int_int) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, void*, int, void*, int, int, long long*, int*))cppgo_virtp_voidp_voidp_int_voidp_int_int)(__goId, -4341501217981328394LL, (void*)&start, role, (void*)&value, hits, (int)(flags), &__len, &__h);
			if (__h) {
				QList<QModelIndex> __r; __r.reserve(__len);
				void** __aa = (void**)__a;
				for (long long __i = 0; __i < __len; __i++) __r.append(*(QModelIndex*)__aa[__i]);
				if (__a) free(__a);
				return __r;
			}
		}
		return QConcatenateTablesProxyModel::match(start, role, value, hits, flags);
	}
	void multiData(const QModelIndex& index, QModelRoleDataSpan roleDataSpan) const override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, 731311060944096642LL, (void*)&index, (void*)&roleDataSpan, &__h); if (__h) return; }
		QConcatenateTablesProxyModel::multiData(index, roleDataSpan);
	}
	bool submit() override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -3788331537442685179LL, &__h); if (__h) return __r; }
		return QConcatenateTablesProxyModel::submit();
	}
	void revert() override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1522951607536117149LL, &__h); if (__h) return; }
		QConcatenateTablesProxyModel::revert();
	}
	void resetInternalData() override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -4510675770784051183LL, &__h); if (__h) return; }
		QConcatenateTablesProxyModel::resetInternalData();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QConcatenateTablesProxyModel::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QConcatenateTablesProxyModel::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QConcatenateTablesProxyModel::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QConcatenateTablesProxyModel::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QConcatenateTablesProxyModel::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QConcatenateTablesProxyModel::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QConcatenateTablesProxyModel::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QConcatenateTablesProxyModel;
static std::mutex __connsMu_QConcatenateTablesProxyModel;

extern "C" {

void* QConcatenateTablesProxyModel_base(void* self) { return (void*)static_cast<QAbstractItemModel*>((QConcatenateTablesProxyModel*)self); }
void QConcatenateTablesProxyModel_chain(void* self, void** out) {
	QConcatenateTablesProxyModel* __p = (QConcatenateTablesProxyModel*)self;
	QAbstractItemModel* __b0 = static_cast<QAbstractItemModel*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QConcatenateTablesProxyModel__set_override(void* self, int idx) { ((QConcatenateTablesProxyModel__Go*)self)->__goOn[idx] = 1; }
void* QConcatenateTablesProxyModel_fromQObject(void* o) { return (void*)qobject_cast<QConcatenateTablesProxyModel*>((QObject*)o); }
void* QConcatenateTablesProxyModel_new(long long __goid) {
	return new QConcatenateTablesProxyModel__Go(__goid);
}
void* QConcatenateTablesProxyModel_new2(long long __goid, void* parent) {
	return new QConcatenateTablesProxyModel__Go(__goid, (QObject*)parent);
}
char* QConcatenateTablesProxyModel_tr(const char* s, long long* __out_len) {
	auto&& __r = QConcatenateTablesProxyModel::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void** QConcatenateTablesProxyModel_sourceModels(void* self, long long* __out_len) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->sourceModels();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void QConcatenateTablesProxyModel_addSourceModel(void* self, void* sourceModel) {
	((QConcatenateTablesProxyModel*)self)->addSourceModel((QAbstractItemModel*)sourceModel);
}
void QConcatenateTablesProxyModel_removeSourceModel(void* self, void* sourceModel) {
	((QConcatenateTablesProxyModel*)self)->removeSourceModel((QAbstractItemModel*)sourceModel);
}
void* QConcatenateTablesProxyModel_mapFromSource(void* self, void* sourceIndex) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->mapFromSource(*(QModelIndex*)sourceIndex);
	return new QModelIndex(__r);
}
void* QConcatenateTablesProxyModel_mapToSource(void* self, void* proxyIndex) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->mapToSource(*(QModelIndex*)proxyIndex);
	return new QModelIndex(__r);
}
void* QConcatenateTablesProxyModel_createIndex(void* self, int row, int column) {
	auto&& __r = ((QConcatenateTablesProxyModel__Go*)self)->createIndex(row, column);
	return new QModelIndex(__r);
}
void* QConcatenateTablesProxyModel_createIndex2(void* self, int row, int column, unsigned long long id) {
	auto&& __r = ((QConcatenateTablesProxyModel__Go*)self)->createIndex(row, column, id);
	return new QModelIndex(__r);
}
void QConcatenateTablesProxyModel_encodeData(void* self, void** indexes, long long indexes_len, void* stream) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	((QConcatenateTablesProxyModel__Go*)self)->encodeData(indexes_list, *(QDataStream*)stream);
}
int QConcatenateTablesProxyModel_decodeData(void* self, int row, int column, void* parent, void* stream) {
	auto&& __r = ((QConcatenateTablesProxyModel__Go*)self)->decodeData(row, column, *(QModelIndex*)parent, *(QDataStream*)stream);
	return __r;
}
void QConcatenateTablesProxyModel_beginInsertRows(void* self, void* parent, int first, int last) {
	((QConcatenateTablesProxyModel__Go*)self)->beginInsertRows(*(QModelIndex*)parent, first, last);
}
void QConcatenateTablesProxyModel_endInsertRows(void* self) {
	((QConcatenateTablesProxyModel__Go*)self)->endInsertRows();
}
void QConcatenateTablesProxyModel_beginRemoveRows(void* self, void* parent, int first, int last) {
	((QConcatenateTablesProxyModel__Go*)self)->beginRemoveRows(*(QModelIndex*)parent, first, last);
}
void QConcatenateTablesProxyModel_endRemoveRows(void* self) {
	((QConcatenateTablesProxyModel__Go*)self)->endRemoveRows();
}
int QConcatenateTablesProxyModel_beginMoveRows(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationRow) {
	auto&& __r = ((QConcatenateTablesProxyModel__Go*)self)->beginMoveRows(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationRow);
	return __r;
}
void QConcatenateTablesProxyModel_endMoveRows(void* self) {
	((QConcatenateTablesProxyModel__Go*)self)->endMoveRows();
}
void QConcatenateTablesProxyModel_beginInsertColumns(void* self, void* parent, int first, int last) {
	((QConcatenateTablesProxyModel__Go*)self)->beginInsertColumns(*(QModelIndex*)parent, first, last);
}
void QConcatenateTablesProxyModel_endInsertColumns(void* self) {
	((QConcatenateTablesProxyModel__Go*)self)->endInsertColumns();
}
void QConcatenateTablesProxyModel_beginRemoveColumns(void* self, void* parent, int first, int last) {
	((QConcatenateTablesProxyModel__Go*)self)->beginRemoveColumns(*(QModelIndex*)parent, first, last);
}
void QConcatenateTablesProxyModel_endRemoveColumns(void* self) {
	((QConcatenateTablesProxyModel__Go*)self)->endRemoveColumns();
}
int QConcatenateTablesProxyModel_beginMoveColumns(void* self, void* sourceParent, int sourceFirst, int sourceLast, void* destinationParent, int destinationColumn) {
	auto&& __r = ((QConcatenateTablesProxyModel__Go*)self)->beginMoveColumns(*(QModelIndex*)sourceParent, sourceFirst, sourceLast, *(QModelIndex*)destinationParent, destinationColumn);
	return __r;
}
void QConcatenateTablesProxyModel_endMoveColumns(void* self) {
	((QConcatenateTablesProxyModel__Go*)self)->endMoveColumns();
}
void QConcatenateTablesProxyModel_beginResetModel(void* self) {
	((QConcatenateTablesProxyModel__Go*)self)->beginResetModel();
}
void QConcatenateTablesProxyModel_endResetModel(void* self) {
	((QConcatenateTablesProxyModel__Go*)self)->endResetModel();
}
void QConcatenateTablesProxyModel_changePersistentIndex(void* self, void* from, void* to) {
	((QConcatenateTablesProxyModel__Go*)self)->changePersistentIndex(*(QModelIndex*)from, *(QModelIndex*)to);
}
void QConcatenateTablesProxyModel_changePersistentIndexList(void* self, void** from, long long from_len, void** to, long long to_len) {
	QList<QModelIndex> from_list; from_list.reserve(from_len);
	for (long long __i = 0; __i < from_len; __i++) from_list.append(*(QModelIndex*)from[__i]);
	QList<QModelIndex> to_list; to_list.reserve(to_len);
	for (long long __i = 0; __i < to_len; __i++) to_list.append(*(QModelIndex*)to[__i]);
	((QConcatenateTablesProxyModel__Go*)self)->changePersistentIndexList(from_list, to_list);
}
void** QConcatenateTablesProxyModel_persistentIndexList(void* self, long long* __out_len) {
	auto&& __r = ((QConcatenateTablesProxyModel__Go*)self)->persistentIndexList();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void* QConcatenateTablesProxyModel_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QConcatenateTablesProxyModel::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QConcatenateTablesProxyModel_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QConcatenateTablesProxyModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QConcatenateTablesProxyModel_sender(void* self) {
	auto&& __r = ((QConcatenateTablesProxyModel__Go*)self)->sender();
	return (void*)(__r);
}
int QConcatenateTablesProxyModel_senderSignalIndex(void* self) {
	auto&& __r = ((QConcatenateTablesProxyModel__Go*)self)->senderSignalIndex();
	return __r;
}
int QConcatenateTablesProxyModel_receivers(void* self, const char* signal) {
	auto&& __r = ((QConcatenateTablesProxyModel__Go*)self)->receivers(signal);
	return __r;
}
int QConcatenateTablesProxyModel_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QConcatenateTablesProxyModel__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QConcatenateTablesProxyModel_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QConcatenateTablesProxyModel::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QConcatenateTablesProxyModel_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QConcatenateTablesProxyModel::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QConcatenateTablesProxyModel_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QConcatenateTablesProxyModel::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QConcatenateTablesProxyModel_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QConcatenateTablesProxyModel::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QConcatenateTablesProxyModel_metaObject_default(void* self) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::metaObject();
	return (void*)(__r);
}
void* QConcatenateTablesProxyModel_data_default(void* self, void* index, int role) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::data(*(QModelIndex*)index, role);
	return new QVariant(__r);
}
int QConcatenateTablesProxyModel_setData_default(void* self, void* index, void* value, int role) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::setData(*(QModelIndex*)index, *(QVariant*)value, role);
	return __r;
}
int QConcatenateTablesProxyModel_flags_default(void* self, void* index) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::flags(*(QModelIndex*)index);
	return (int)(__r);
}
void* QConcatenateTablesProxyModel_index_default(void* self, int row, int column, void* parent) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::index(row, column, *(QModelIndex*)parent);
	return new QModelIndex(__r);
}
void* QConcatenateTablesProxyModel_parent_default(void* self, void* index) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::parent(*(QModelIndex*)index);
	return new QModelIndex(__r);
}
int QConcatenateTablesProxyModel_rowCount_default(void* self, void* parent) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::rowCount(*(QModelIndex*)parent);
	return __r;
}
void* QConcatenateTablesProxyModel_headerData_default(void* self, int section, int orientation, int role) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::headerData(section, (Qt::Orientation)(orientation), role);
	return new QVariant(__r);
}
int QConcatenateTablesProxyModel_columnCount_default(void* self, void* parent) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::columnCount(*(QModelIndex*)parent);
	return __r;
}
char** QConcatenateTablesProxyModel_mimeTypes_default(void* self, long long* __out_len) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::mimeTypes();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void* QConcatenateTablesProxyModel_mimeData_default(void* self, void** indexes, long long indexes_len) {
	QList<QModelIndex> indexes_list; indexes_list.reserve(indexes_len);
	for (long long __i = 0; __i < indexes_len; __i++) indexes_list.append(*(QModelIndex*)indexes[__i]);
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::mimeData(indexes_list);
	return (void*)(__r);
}
int QConcatenateTablesProxyModel_canDropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::canDropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
int QConcatenateTablesProxyModel_dropMimeData_default(void* self, void* data, int action, int row, int column, void* parent) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::dropMimeData((QMimeData*)data, (Qt::DropAction)(action), row, column, *(QModelIndex*)parent);
	return __r;
}
void* QConcatenateTablesProxyModel_span_default(void* self, void* index) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::span(*(QModelIndex*)index);
	return new QSize(__r);
}
void* QConcatenateTablesProxyModel_sibling_default(void* self, int row, int column, void* idx) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::sibling(row, column, *(QModelIndex*)idx);
	return new QModelIndex(__r);
}
int QConcatenateTablesProxyModel_hasChildren_default(void* self, void* parent) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::hasChildren(*(QModelIndex*)parent);
	return __r;
}
int QConcatenateTablesProxyModel_setHeaderData_default(void* self, int section, int orientation, void* value, int role) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::setHeaderData(section, (Qt::Orientation)(orientation), *(QVariant*)value, role);
	return __r;
}
int QConcatenateTablesProxyModel_clearItemData_default(void* self, void* index) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::clearItemData(*(QModelIndex*)index);
	return __r;
}
int QConcatenateTablesProxyModel_supportedDropActions_default(void* self) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::supportedDropActions();
	return (int)(__r);
}
int QConcatenateTablesProxyModel_supportedDragActions_default(void* self) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::supportedDragActions();
	return (int)(__r);
}
int QConcatenateTablesProxyModel_insertRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::insertRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QConcatenateTablesProxyModel_insertColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::insertColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QConcatenateTablesProxyModel_removeRows_default(void* self, int row, int count, void* parent) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::removeRows(row, count, *(QModelIndex*)parent);
	return __r;
}
int QConcatenateTablesProxyModel_removeColumns_default(void* self, int column, int count, void* parent) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::removeColumns(column, count, *(QModelIndex*)parent);
	return __r;
}
int QConcatenateTablesProxyModel_moveRows_default(void* self, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::moveRows(*(QModelIndex*)sourceParent, sourceRow, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
int QConcatenateTablesProxyModel_moveColumns_default(void* self, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::moveColumns(*(QModelIndex*)sourceParent, sourceColumn, count, *(QModelIndex*)destinationParent, destinationChild);
	return __r;
}
void QConcatenateTablesProxyModel_fetchMore_default(void* self, void* parent) {
	((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::fetchMore(*(QModelIndex*)parent);
}
int QConcatenateTablesProxyModel_canFetchMore_default(void* self, void* parent) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::canFetchMore(*(QModelIndex*)parent);
	return __r;
}
void QConcatenateTablesProxyModel_sort_default(void* self, int column, int order) {
	((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::sort(column, (Qt::SortOrder)(order));
}
void* QConcatenateTablesProxyModel_buddy_default(void* self, void* index) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::buddy(*(QModelIndex*)index);
	return new QModelIndex(__r);
}
void** QConcatenateTablesProxyModel_match_default(void* self, void* start, int role, void* value, int hits, int flags, long long* __out_len) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::match(*(QModelIndex*)start, role, *(QVariant*)value, hits, QFlags<Qt::MatchFlag>(flags));
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
void QConcatenateTablesProxyModel_multiData_default(void* self, void* index, void* roleDataSpan) {
	((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::multiData(*(QModelIndex*)index, *(QModelRoleDataSpan*)roleDataSpan);
}
int QConcatenateTablesProxyModel_submit_default(void* self) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::submit();
	return __r;
}
void QConcatenateTablesProxyModel_revert_default(void* self) {
	((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::revert();
}
void QConcatenateTablesProxyModel_resetInternalData_default(void* self) {
	((QConcatenateTablesProxyModel__Go*)self)->cppgoBase_resetInternalData();
}
int QConcatenateTablesProxyModel_event_default(void* self, void* event) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::event((QEvent*)event);
	return __r;
}
int QConcatenateTablesProxyModel_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QConcatenateTablesProxyModel*)self)->QConcatenateTablesProxyModel::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QConcatenateTablesProxyModel_timerEvent_default(void* self, void* event) {
	((QConcatenateTablesProxyModel__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QConcatenateTablesProxyModel_childEvent_default(void* self, void* event) {
	((QConcatenateTablesProxyModel__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QConcatenateTablesProxyModel_customEvent_default(void* self, void* event) {
	((QConcatenateTablesProxyModel__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QConcatenateTablesProxyModel_connectNotify_default(void* self, void* signal) {
	((QConcatenateTablesProxyModel__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QConcatenateTablesProxyModel_disconnectNotify_default(void* self, void* signal) {
	((QConcatenateTablesProxyModel__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QConcatenateTablesProxyModel_connect_headerDataChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::headerDataChanged, (QConcatenateTablesProxyModel*)self, [__id](Qt::Orientation orientation, int first, int last){
		if (cppgo_sig_void_int_int_int) ((void(*)(long long, int, int, int))cppgo_sig_void_int_int_int)(__id, (int)(orientation), first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_layoutChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::layoutChanged, (QConcatenateTablesProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_layoutChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::layoutChanged, (QConcatenateTablesProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_layoutAboutToBeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::layoutAboutToBeChanged, (QConcatenateTablesProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_layoutAboutToBeChanged2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::layoutAboutToBeChanged, (QConcatenateTablesProxyModel*)self, [__id](const QList<QPersistentModelIndex>& parents, QAbstractItemModel::LayoutChangeHint hint){
		std::vector<void*> parents_v; parents_v.reserve(parents.size());
		for (long long __i = 0; __i < parents.size(); __i++) parents_v.push_back((void*)&parents[__i]);
		if (cppgo_sig_void_voidpp_longlong_int) ((void(*)(long long, void**, long long, int))cppgo_sig_void_voidpp_longlong_int)(__id, parents_v.data(), (long long)parents.size(), (int)(hint));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_rowsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::rowsAboutToBeInserted, (QConcatenateTablesProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_rowsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::rowsInserted, (QConcatenateTablesProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_rowsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::rowsAboutToBeRemoved, (QConcatenateTablesProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_rowsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::rowsRemoved, (QConcatenateTablesProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_columnsAboutToBeInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::columnsAboutToBeInserted, (QConcatenateTablesProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_columnsInserted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::columnsInserted, (QConcatenateTablesProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_columnsAboutToBeRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::columnsAboutToBeRemoved, (QConcatenateTablesProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_columnsRemoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::columnsRemoved, (QConcatenateTablesProxyModel*)self, [__id](const QModelIndex& parent, int first, int last){
		if (cppgo_sig_void_voidp_int_int) ((void(*)(long long, void*, int, int))cppgo_sig_void_voidp_int_int)(__id, (void*)&parent, first, last);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_modelAboutToBeReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::modelAboutToBeReset, (QConcatenateTablesProxyModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_modelReset(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::modelReset, (QConcatenateTablesProxyModel*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_rowsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::rowsAboutToBeMoved, (QConcatenateTablesProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_rowsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::rowsMoved, (QConcatenateTablesProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationRow){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationRow);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_columnsAboutToBeMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::columnsAboutToBeMoved, (QConcatenateTablesProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_columnsMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::columnsMoved, (QConcatenateTablesProxyModel*)self, [__id](const QModelIndex& sourceParent, int sourceStart, int sourceEnd, const QModelIndex& destinationParent, int destinationColumn){
		if (cppgo_sig_void_voidp_int_int_voidp_int) ((void(*)(long long, void*, int, int, void*, int))cppgo_sig_void_voidp_int_int_voidp_int)(__id, (void*)&sourceParent, sourceStart, sourceEnd, (void*)&destinationParent, destinationColumn);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::destroyed, (QConcatenateTablesProxyModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::destroyed, (QConcatenateTablesProxyModel*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QConcatenateTablesProxyModel*)self, &QConcatenateTablesProxyModel::objectNameChanged, (QConcatenateTablesProxyModel*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
	__conns_QConcatenateTablesProxyModel[__id] = __c;
}
void QConcatenateTablesProxyModel__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QConcatenateTablesProxyModel);
		auto __it = __conns_QConcatenateTablesProxyModel.find(id);
		if (__it == __conns_QConcatenateTablesProxyModel.end()) return;
		__c = __it->second;
		__conns_QConcatenateTablesProxyModel.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QConcatenateTablesProxyModel__Go_new(long long id, void* parent) { return new QConcatenateTablesProxyModel__Go(id, (QObject*)parent); }
void QConcatenateTablesProxyModel_delete(void* self) { delete (QConcatenateTablesProxyModel*)self; }

}

#include <QtCore/qfuturewatcher.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include <QtCore/qfutureinterface.h>
#include "gen_qfuturewatcherbase.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QFutureWatcherBase;
static std::mutex __connsMu_QFutureWatcherBase;

extern "C" {

void* QFutureWatcherBase_base(void* self) { return (void*)static_cast<QObject*>((QFutureWatcherBase*)self); }
void* QFutureWatcherBase_fromQObject(void* o) { return (void*)qobject_cast<QFutureWatcherBase*>((QObject*)o); }
char* QFutureWatcherBase_tr(const char* s, long long* __out_len) {
	auto&& __r = QFutureWatcherBase::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QFutureWatcherBase_progressValue(void* self) {
	auto&& __r = ((QFutureWatcherBase*)self)->progressValue();
	return __r;
}
int QFutureWatcherBase_progressMinimum(void* self) {
	auto&& __r = ((QFutureWatcherBase*)self)->progressMinimum();
	return __r;
}
int QFutureWatcherBase_progressMaximum(void* self) {
	auto&& __r = ((QFutureWatcherBase*)self)->progressMaximum();
	return __r;
}
char* QFutureWatcherBase_progressText(void* self, long long* __out_len) {
	auto&& __r = ((QFutureWatcherBase*)self)->progressText();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QFutureWatcherBase_isStarted(void* self) {
	auto&& __r = ((QFutureWatcherBase*)self)->isStarted();
	return __r;
}
int QFutureWatcherBase_isFinished(void* self) {
	auto&& __r = ((QFutureWatcherBase*)self)->isFinished();
	return __r;
}
int QFutureWatcherBase_isRunning(void* self) {
	auto&& __r = ((QFutureWatcherBase*)self)->isRunning();
	return __r;
}
int QFutureWatcherBase_isCanceled(void* self) {
	auto&& __r = ((QFutureWatcherBase*)self)->isCanceled();
	return __r;
}
int QFutureWatcherBase_isPaused(void* self) {
	auto&& __r = ((QFutureWatcherBase*)self)->isPaused();
	return __r;
}
int QFutureWatcherBase_isSuspending(void* self) {
	auto&& __r = ((QFutureWatcherBase*)self)->isSuspending();
	return __r;
}
int QFutureWatcherBase_isSuspended(void* self) {
	auto&& __r = ((QFutureWatcherBase*)self)->isSuspended();
	return __r;
}
void QFutureWatcherBase_waitForFinished(void* self) {
	((QFutureWatcherBase*)self)->waitForFinished();
}
void QFutureWatcherBase_setPendingResultsLimit(void* self, int limit) {
	((QFutureWatcherBase*)self)->setPendingResultsLimit(limit);
}
void QFutureWatcherBase_started(void* self) {
	((QFutureWatcherBase*)self)->started();
}
void QFutureWatcherBase_finished(void* self) {
	((QFutureWatcherBase*)self)->finished();
}
void QFutureWatcherBase_canceled(void* self) {
	((QFutureWatcherBase*)self)->canceled();
}
void QFutureWatcherBase_paused(void* self) {
	((QFutureWatcherBase*)self)->paused();
}
void QFutureWatcherBase_suspending(void* self) {
	((QFutureWatcherBase*)self)->suspending();
}
void QFutureWatcherBase_suspended(void* self) {
	((QFutureWatcherBase*)self)->suspended();
}
void QFutureWatcherBase_resumed(void* self) {
	((QFutureWatcherBase*)self)->resumed();
}
void QFutureWatcherBase_resultReadyAt(void* self, int resultIndex) {
	((QFutureWatcherBase*)self)->resultReadyAt(resultIndex);
}
void QFutureWatcherBase_resultsReadyAt(void* self, int beginIndex, int endIndex) {
	((QFutureWatcherBase*)self)->resultsReadyAt(beginIndex, endIndex);
}
void QFutureWatcherBase_progressRangeChanged(void* self, int minimum, int maximum) {
	((QFutureWatcherBase*)self)->progressRangeChanged(minimum, maximum);
}
void QFutureWatcherBase_progressValueChanged(void* self, int progressValue) {
	((QFutureWatcherBase*)self)->progressValueChanged(progressValue);
}
void QFutureWatcherBase_progressTextChanged(void* self, const char* progressText, long long progressText_len) {
	QString progressText_qstr = QString::fromUtf8(progressText, progressText_len);
	((QFutureWatcherBase*)self)->progressTextChanged(progressText_qstr);
}
void QFutureWatcherBase_cancel(void* self) {
	((QFutureWatcherBase*)self)->cancel();
}
void QFutureWatcherBase_setSuspended(void* self, int suspend) {
	((QFutureWatcherBase*)self)->setSuspended(suspend);
}
void QFutureWatcherBase_suspend(void* self) {
	((QFutureWatcherBase*)self)->suspend();
}
void QFutureWatcherBase_resume(void* self) {
	((QFutureWatcherBase*)self)->resume();
}
void QFutureWatcherBase_toggleSuspended(void* self) {
	((QFutureWatcherBase*)self)->toggleSuspended();
}
void QFutureWatcherBase_setPaused(void* self, int paused) {
	((QFutureWatcherBase*)self)->setPaused(paused);
}
void QFutureWatcherBase_pause(void* self) {
	((QFutureWatcherBase*)self)->pause();
}
void QFutureWatcherBase_togglePaused(void* self) {
	((QFutureWatcherBase*)self)->togglePaused();
}
void* QFutureWatcherBase_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QFutureWatcherBase::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QFutureWatcherBase_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QFutureWatcherBase::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QFutureWatcherBase_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QFutureWatcherBase::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QFutureWatcherBase_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QFutureWatcherBase::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QFutureWatcherBase_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QFutureWatcherBase::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QFutureWatcherBase_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QFutureWatcherBase::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QFutureWatcherBase_connect_started(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFutureWatcherBase*)self, &QFutureWatcherBase::started, (QFutureWatcherBase*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFutureWatcherBase);
	__conns_QFutureWatcherBase[__id] = __c;
}
void QFutureWatcherBase_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFutureWatcherBase*)self, &QFutureWatcherBase::finished, (QFutureWatcherBase*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFutureWatcherBase);
	__conns_QFutureWatcherBase[__id] = __c;
}
void QFutureWatcherBase_connect_canceled(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFutureWatcherBase*)self, &QFutureWatcherBase::canceled, (QFutureWatcherBase*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFutureWatcherBase);
	__conns_QFutureWatcherBase[__id] = __c;
}
void QFutureWatcherBase_connect_paused(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFutureWatcherBase*)self, &QFutureWatcherBase::paused, (QFutureWatcherBase*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFutureWatcherBase);
	__conns_QFutureWatcherBase[__id] = __c;
}
void QFutureWatcherBase_connect_suspending(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFutureWatcherBase*)self, &QFutureWatcherBase::suspending, (QFutureWatcherBase*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFutureWatcherBase);
	__conns_QFutureWatcherBase[__id] = __c;
}
void QFutureWatcherBase_connect_suspended(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFutureWatcherBase*)self, &QFutureWatcherBase::suspended, (QFutureWatcherBase*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFutureWatcherBase);
	__conns_QFutureWatcherBase[__id] = __c;
}
void QFutureWatcherBase_connect_resumed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFutureWatcherBase*)self, &QFutureWatcherBase::resumed, (QFutureWatcherBase*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFutureWatcherBase);
	__conns_QFutureWatcherBase[__id] = __c;
}
void QFutureWatcherBase_connect_resultReadyAt(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFutureWatcherBase*)self, &QFutureWatcherBase::resultReadyAt, (QFutureWatcherBase*)self, [__id](int resultIndex){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, resultIndex);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFutureWatcherBase);
	__conns_QFutureWatcherBase[__id] = __c;
}
void QFutureWatcherBase_connect_resultsReadyAt(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFutureWatcherBase*)self, &QFutureWatcherBase::resultsReadyAt, (QFutureWatcherBase*)self, [__id](int beginIndex, int endIndex){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, beginIndex, endIndex);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFutureWatcherBase);
	__conns_QFutureWatcherBase[__id] = __c;
}
void QFutureWatcherBase_connect_progressRangeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFutureWatcherBase*)self, &QFutureWatcherBase::progressRangeChanged, (QFutureWatcherBase*)self, [__id](int minimum, int maximum){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, minimum, maximum);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFutureWatcherBase);
	__conns_QFutureWatcherBase[__id] = __c;
}
void QFutureWatcherBase_connect_progressValueChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFutureWatcherBase*)self, &QFutureWatcherBase::progressValueChanged, (QFutureWatcherBase*)self, [__id](int progressValue){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, progressValue);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFutureWatcherBase);
	__conns_QFutureWatcherBase[__id] = __c;
}
void QFutureWatcherBase_connect_progressTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFutureWatcherBase*)self, &QFutureWatcherBase::progressTextChanged, (QFutureWatcherBase*)self, [__id](const QString& progressText){
		QByteArray progressText_ba = progressText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)progressText_ba.constData(), progressText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFutureWatcherBase);
	__conns_QFutureWatcherBase[__id] = __c;
}
void QFutureWatcherBase_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFutureWatcherBase*)self, &QFutureWatcherBase::destroyed, (QFutureWatcherBase*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFutureWatcherBase);
	__conns_QFutureWatcherBase[__id] = __c;
}
void QFutureWatcherBase_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFutureWatcherBase*)self, &QFutureWatcherBase::destroyed, (QFutureWatcherBase*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFutureWatcherBase);
	__conns_QFutureWatcherBase[__id] = __c;
}
void QFutureWatcherBase_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QFutureWatcherBase*)self, &QFutureWatcherBase::objectNameChanged, (QFutureWatcherBase*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QFutureWatcherBase);
	__conns_QFutureWatcherBase[__id] = __c;
}
void QFutureWatcherBase__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QFutureWatcherBase);
		auto __it = __conns_QFutureWatcherBase.find(id);
		if (__it == __conns_QFutureWatcherBase.end()) return;
		__c = __it->second;
		__conns_QFutureWatcherBase.erase(__it);
	}
	QObject::disconnect(__c);
}
void QFutureWatcherBase_delete(void* self) { delete (QFutureWatcherBase*)self; }

}

#include <QtCore/qmargins.h>
#include <QtCore/qrect.h>
#include "gen_qmarginsf.h"

extern "C" {

void* QMarginsF_new() {
	return new QMarginsF();
}
void* QMarginsF_new2(double left, double top, double right, double bottom) {
	return new QMarginsF(left, top, right, bottom);
}
void* QMarginsF_new3(void* margins) {
	return new QMarginsF(*(QMargins*)margins);
}
void* QMarginsF_new4(void* arg1) {
	return new QMarginsF(*(QMarginsF*)arg1);
}
int QMarginsF_isNull(void* self) {
	auto&& __r = ((QMarginsF*)self)->isNull();
	return __r;
}
double QMarginsF_left(void* self) {
	auto&& __r = ((QMarginsF*)self)->left();
	return __r;
}
double QMarginsF_top(void* self) {
	auto&& __r = ((QMarginsF*)self)->top();
	return __r;
}
double QMarginsF_right(void* self) {
	auto&& __r = ((QMarginsF*)self)->right();
	return __r;
}
double QMarginsF_bottom(void* self) {
	auto&& __r = ((QMarginsF*)self)->bottom();
	return __r;
}
void QMarginsF_setLeft(void* self, double aleft) {
	((QMarginsF*)self)->setLeft(aleft);
}
void QMarginsF_setTop(void* self, double atop) {
	((QMarginsF*)self)->setTop(atop);
}
void QMarginsF_setRight(void* self, double aright) {
	((QMarginsF*)self)->setRight(aright);
}
void QMarginsF_setBottom(void* self, double abottom) {
	((QMarginsF*)self)->setBottom(abottom);
}
void* QMarginsF_OperatorPlusAssign(void* self, void* margins) {
	auto&& __r = ((QMarginsF*)self)->operator+=(*(QMarginsF*)margins);
	return (void*)&(__r);
}
void* QMarginsF_OperatorMinusAssign(void* self, void* margins) {
	auto&& __r = ((QMarginsF*)self)->operator-=(*(QMarginsF*)margins);
	return (void*)&(__r);
}
void* QMarginsF_OperatorPlusAssign2(void* self, double addend) {
	auto&& __r = ((QMarginsF*)self)->operator+=(addend);
	return (void*)&(__r);
}
void* QMarginsF_OperatorMinusAssign2(void* self, double subtrahend) {
	auto&& __r = ((QMarginsF*)self)->operator-=(subtrahend);
	return (void*)&(__r);
}
void* QMarginsF_OperatorMultiplyAssign(void* self, double factor) {
	auto&& __r = ((QMarginsF*)self)->operator*=(factor);
	return (void*)&(__r);
}
void* QMarginsF_OperatorDivideAssign(void* self, double divisor) {
	auto&& __r = ((QMarginsF*)self)->operator/=(divisor);
	return (void*)&(__r);
}
void* QMarginsF_toMargins(void* self) {
	auto&& __r = ((QMarginsF*)self)->toMargins();
	return new QMargins(__r);
}
void* QMarginsF_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QMarginsF*)self)->operator=(*(QMarginsF*)arg1);
	return (void*)&(__r);
}
int QMarginsF_OperatorEqual(void* self, void* rhs) {
	auto&& __r = (*(QMarginsF*)self) == (*(QMarginsF*)rhs);
	return __r;
}
int QMarginsF_OperatorNotEqual(void* self, void* rhs) {
	auto&& __r = (*(QMarginsF*)self) != (*(QMarginsF*)rhs);
	return __r;
}
int QMarginsF_OperatorEqual2(void* self, void* rhs) {
	auto&& __r = (*(QMarginsF*)self) == (*(QMargins*)rhs);
	return __r;
}
int QMarginsF_OperatorNotEqual2(void* self, void* rhs) {
	auto&& __r = (*(QMarginsF*)self) != (*(QMargins*)rhs);
	return __r;
}
void* QMarginsF_OperatorPlus(void* self, void* rhs) {
	auto&& __r = (*(QMarginsF*)self) + (*(QMarginsF*)rhs);
	return new QMarginsF(__r);
}
void* QMarginsF_OperatorMinus(void* self, void* rhs) {
	auto&& __r = (*(QMarginsF*)self) - (*(QMarginsF*)rhs);
	return new QMarginsF(__r);
}
void* QMarginsF_OperatorPlus2(void* self, double rhs) {
	auto&& __r = (*(QMarginsF*)self) + (rhs);
	return new QMarginsF(__r);
}
void* QMarginsF_OperatorMinus2(void* self, double rhs) {
	auto&& __r = (*(QMarginsF*)self) - (rhs);
	return new QMarginsF(__r);
}
void* QMarginsF_OperatorMultiply(void* self, double rhs) {
	auto&& __r = (*(QMarginsF*)self) * (rhs);
	return new QMarginsF(__r);
}
void* QMarginsF_OperatorDivide(void* self, double divisor) {
	auto&& __r = (*(QMarginsF*)self) / (divisor);
	return new QMarginsF(__r);
}
void* QMarginsF_OperatorBitwiseOr(void* self, void* m2) {
	auto&& __r = (*(QMarginsF*)self) | (*(QMarginsF*)m2);
	return new QMarginsF(__r);
}
void* QMarginsF_OperatorPlus3(void* self) {
	auto&& __r = +(*(QMarginsF*)self);
	return new QMarginsF(__r);
}
void* QMarginsF_OperatorMinus3(void* self) {
	auto&& __r = -(*(QMarginsF*)self);
	return new QMarginsF(__r);
}
void* QMarginsF_OperatorPlus4(void* self, void* rhs) {
	auto&& __r = (*(QMarginsF*)self) + (*(QRectF*)rhs);
	return new QRectF(__r);
}
void QMarginsF_delete(void* self) { delete (QMarginsF*)self; }

}

#include <QtCore/qplugin.h>
#include "gen_qpluginmetadata__magicheader.h"

extern "C" {

void* QPluginMetaData__MagicHeader_new() {
	return new QPluginMetaData::MagicHeader();
}
void* QPluginMetaData__MagicHeader_header_get(void* self) {
	auto __r = ((QPluginMetaData::MagicHeader*)self)->header;
	return new QPluginMetaData::Header(__r);
}
void QPluginMetaData__MagicHeader_header_set(void* self, void* header) {
	((QPluginMetaData::MagicHeader*)self)->header = *(QPluginMetaData::Header*)header;
}
void QPluginMetaData__MagicHeader_delete(void* self) { delete (QPluginMetaData::MagicHeader*)self; }

}

#include <QtCore/qsequentialanimationgroup.h>
#include <QtCore/qanimationgroup.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qpauseanimation.h>
#include <QtCore/qabstractanimation.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qsequentialanimationgroup.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QSequentialAnimationGroup__Go : public QSequentialAnimationGroup {
public:
	long long __goId;
	unsigned char __goOn[12] = {};
	template<typename... A> QSequentialAnimationGroup__Go(long long id, A&&... a) : QSequentialAnimationGroup(std::forward<A>(a)...), __goId(id) {}
	using QSequentialAnimationGroup::sender;
	using QSequentialAnimationGroup::senderSignalIndex;
	using QSequentialAnimationGroup::receivers;
	using QSequentialAnimationGroup::isSignalConnected;
	bool cppgoBase_event(QEvent* __a0) { return QSequentialAnimationGroup::event(__a0); }
	void cppgoBase_updateCurrentTime(int __a0) { QSequentialAnimationGroup::updateCurrentTime(__a0); }
	void cppgoBase_updateState(int __a0, int __a1) { QSequentialAnimationGroup::updateState((QAbstractAnimation::State)__a0, (QAbstractAnimation::State)__a1); }
	void cppgoBase_updateDirection(int __a0) { QSequentialAnimationGroup::updateDirection((QAbstractAnimation::Direction)__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QSequentialAnimationGroup::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QSequentialAnimationGroup::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QSequentialAnimationGroup::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QSequentialAnimationGroup::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QSequentialAnimationGroup::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QSequentialAnimationGroup::metaObject();
	}
	int duration() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8434675112194400115LL, &__h); if (__h) return __r; }
		return QSequentialAnimationGroup::duration();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QSequentialAnimationGroup::event(event);
	}
	void updateCurrentTime(int arg1) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -1342418961870150968LL, arg1, &__h); if (__h) return; }
		QSequentialAnimationGroup::updateCurrentTime(arg1);
	}
	void updateState(QAbstractAnimation::State newState, QAbstractAnimation::State oldState) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, 7692482089134926569LL, (int)(newState), (int)(oldState), &__h); if (__h) return; }
		QSequentialAnimationGroup::updateState(newState, oldState);
	}
	void updateDirection(QAbstractAnimation::Direction direction) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -3231888725407197643LL, (int)(direction), &__h); if (__h) return; }
		QSequentialAnimationGroup::updateDirection(direction);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QSequentialAnimationGroup::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QSequentialAnimationGroup::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QSequentialAnimationGroup::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QSequentialAnimationGroup::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QSequentialAnimationGroup::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QSequentialAnimationGroup::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QSequentialAnimationGroup;
static std::mutex __connsMu_QSequentialAnimationGroup;

extern "C" {

void* QSequentialAnimationGroup_base(void* self) { return (void*)static_cast<QAnimationGroup*>((QSequentialAnimationGroup*)self); }
void QSequentialAnimationGroup_chain(void* self, void** out) {
	QSequentialAnimationGroup* __p = (QSequentialAnimationGroup*)self;
	QAnimationGroup* __b0 = static_cast<QAnimationGroup*>(__p); out[0] = (void*)__b0;
	QAbstractAnimation* __b1 = static_cast<QAbstractAnimation*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QSequentialAnimationGroup__set_override(void* self, int idx) { ((QSequentialAnimationGroup__Go*)self)->__goOn[idx] = 1; }
void* QSequentialAnimationGroup_fromQObject(void* o) { return (void*)qobject_cast<QSequentialAnimationGroup*>((QObject*)o); }
void* QSequentialAnimationGroup_new(long long __goid) {
	return new QSequentialAnimationGroup__Go(__goid);
}
void* QSequentialAnimationGroup_new2(long long __goid, void* parent) {
	return new QSequentialAnimationGroup__Go(__goid, (QObject*)parent);
}
char* QSequentialAnimationGroup_tr(const char* s, long long* __out_len) {
	auto&& __r = QSequentialAnimationGroup::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QSequentialAnimationGroup_addPause(void* self, int msecs) {
	auto&& __r = ((QSequentialAnimationGroup*)self)->addPause(msecs);
	return (void*)(__r);
}
void* QSequentialAnimationGroup_insertPause(void* self, int index, int msecs) {
	auto&& __r = ((QSequentialAnimationGroup*)self)->insertPause(index, msecs);
	return (void*)(__r);
}
void* QSequentialAnimationGroup_currentAnimation(void* self) {
	auto&& __r = ((QSequentialAnimationGroup*)self)->currentAnimation();
	return (void*)(__r);
}
void QSequentialAnimationGroup_currentAnimationChanged(void* self, void* current) {
	((QSequentialAnimationGroup*)self)->currentAnimationChanged((QAbstractAnimation*)current);
}
void* QSequentialAnimationGroup_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QSequentialAnimationGroup::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QSequentialAnimationGroup_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QSequentialAnimationGroup::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QSequentialAnimationGroup_sender(void* self) {
	auto&& __r = ((QSequentialAnimationGroup__Go*)self)->sender();
	return (void*)(__r);
}
int QSequentialAnimationGroup_senderSignalIndex(void* self) {
	auto&& __r = ((QSequentialAnimationGroup__Go*)self)->senderSignalIndex();
	return __r;
}
int QSequentialAnimationGroup_receivers(void* self, const char* signal) {
	auto&& __r = ((QSequentialAnimationGroup__Go*)self)->receivers(signal);
	return __r;
}
int QSequentialAnimationGroup_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QSequentialAnimationGroup__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QSequentialAnimationGroup_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QSequentialAnimationGroup::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QSequentialAnimationGroup_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QSequentialAnimationGroup::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QSequentialAnimationGroup_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QSequentialAnimationGroup::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QSequentialAnimationGroup_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QSequentialAnimationGroup::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QSequentialAnimationGroup_metaObject_default(void* self) {
	auto&& __r = ((QSequentialAnimationGroup*)self)->QSequentialAnimationGroup::metaObject();
	return (void*)(__r);
}
int QSequentialAnimationGroup_duration_default(void* self) {
	auto&& __r = ((QSequentialAnimationGroup*)self)->QSequentialAnimationGroup::duration();
	return __r;
}
int QSequentialAnimationGroup_event_default(void* self, void* event) {
	auto&& __r = ((QSequentialAnimationGroup__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
void QSequentialAnimationGroup_updateCurrentTime_default(void* self, int arg1) {
	((QSequentialAnimationGroup__Go*)self)->cppgoBase_updateCurrentTime(arg1);
}
void QSequentialAnimationGroup_updateState_default(void* self, int newState, int oldState) {
	((QSequentialAnimationGroup__Go*)self)->cppgoBase_updateState(newState, oldState);
}
void QSequentialAnimationGroup_updateDirection_default(void* self, int direction) {
	((QSequentialAnimationGroup__Go*)self)->cppgoBase_updateDirection(direction);
}
int QSequentialAnimationGroup_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QSequentialAnimationGroup*)self)->QSequentialAnimationGroup::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QSequentialAnimationGroup_timerEvent_default(void* self, void* event) {
	((QSequentialAnimationGroup__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QSequentialAnimationGroup_childEvent_default(void* self, void* event) {
	((QSequentialAnimationGroup__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QSequentialAnimationGroup_customEvent_default(void* self, void* event) {
	((QSequentialAnimationGroup__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QSequentialAnimationGroup_connectNotify_default(void* self, void* signal) {
	((QSequentialAnimationGroup__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QSequentialAnimationGroup_disconnectNotify_default(void* self, void* signal) {
	((QSequentialAnimationGroup__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QSequentialAnimationGroup_connect_currentAnimationChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSequentialAnimationGroup*)self, &QSequentialAnimationGroup::currentAnimationChanged, (QSequentialAnimationGroup*)self, [__id](QAbstractAnimation* current){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)current);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSequentialAnimationGroup);
	__conns_QSequentialAnimationGroup[__id] = __c;
}
void QSequentialAnimationGroup_connect_finished(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSequentialAnimationGroup*)self, &QSequentialAnimationGroup::finished, (QSequentialAnimationGroup*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSequentialAnimationGroup);
	__conns_QSequentialAnimationGroup[__id] = __c;
}
void QSequentialAnimationGroup_connect_stateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSequentialAnimationGroup*)self, &QSequentialAnimationGroup::stateChanged, (QSequentialAnimationGroup*)self, [__id](QAbstractAnimation::State newState, QAbstractAnimation::State oldState){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, (int)(newState), (int)(oldState));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSequentialAnimationGroup);
	__conns_QSequentialAnimationGroup[__id] = __c;
}
void QSequentialAnimationGroup_connect_currentLoopChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSequentialAnimationGroup*)self, &QSequentialAnimationGroup::currentLoopChanged, (QSequentialAnimationGroup*)self, [__id](int currentLoop){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, currentLoop);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSequentialAnimationGroup);
	__conns_QSequentialAnimationGroup[__id] = __c;
}
void QSequentialAnimationGroup_connect_directionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSequentialAnimationGroup*)self, &QSequentialAnimationGroup::directionChanged, (QSequentialAnimationGroup*)self, [__id](QAbstractAnimation::Direction arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(arg1));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSequentialAnimationGroup);
	__conns_QSequentialAnimationGroup[__id] = __c;
}
void QSequentialAnimationGroup_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSequentialAnimationGroup*)self, &QSequentialAnimationGroup::destroyed, (QSequentialAnimationGroup*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSequentialAnimationGroup);
	__conns_QSequentialAnimationGroup[__id] = __c;
}
void QSequentialAnimationGroup_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSequentialAnimationGroup*)self, &QSequentialAnimationGroup::destroyed, (QSequentialAnimationGroup*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSequentialAnimationGroup);
	__conns_QSequentialAnimationGroup[__id] = __c;
}
void QSequentialAnimationGroup_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QSequentialAnimationGroup*)self, &QSequentialAnimationGroup::objectNameChanged, (QSequentialAnimationGroup*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QSequentialAnimationGroup);
	__conns_QSequentialAnimationGroup[__id] = __c;
}
void QSequentialAnimationGroup__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QSequentialAnimationGroup);
		auto __it = __conns_QSequentialAnimationGroup.find(id);
		if (__it == __conns_QSequentialAnimationGroup.end()) return;
		__c = __it->second;
		__conns_QSequentialAnimationGroup.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QSequentialAnimationGroup__Go_new(long long id, void* parent) { return new QSequentialAnimationGroup__Go(id, (QObject*)parent); }
void QSequentialAnimationGroup_delete(void* self) { delete (QSequentialAnimationGroup*)self; }

}

#include <QtCore/qwineventnotifier.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qnamespace.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qwineventnotifier.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QWinEventNotifier__Go : public QWinEventNotifier {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QWinEventNotifier__Go(long long id, A&&... a) : QWinEventNotifier(std::forward<A>(a)...), __goId(id) {}
	using QWinEventNotifier::sender;
	using QWinEventNotifier::senderSignalIndex;
	using QWinEventNotifier::receivers;
	using QWinEventNotifier::isSignalConnected;
	bool cppgoBase_event(QEvent* __a0) { return QWinEventNotifier::event(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QWinEventNotifier::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QWinEventNotifier::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QWinEventNotifier::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QWinEventNotifier::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QWinEventNotifier::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QWinEventNotifier::metaObject();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QWinEventNotifier::event(e);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QWinEventNotifier::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QWinEventNotifier::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QWinEventNotifier::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QWinEventNotifier::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QWinEventNotifier::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QWinEventNotifier::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QWinEventNotifier;
static std::mutex __connsMu_QWinEventNotifier;

extern "C" {

void* QWinEventNotifier_base(void* self) { return (void*)static_cast<QObject*>((QWinEventNotifier*)self); }
void QWinEventNotifier__set_override(void* self, int idx) { ((QWinEventNotifier__Go*)self)->__goOn[idx] = 1; }
void* QWinEventNotifier_fromQObject(void* o) { return (void*)qobject_cast<QWinEventNotifier*>((QObject*)o); }
void* QWinEventNotifier_new(long long __goid) {
	return new QWinEventNotifier__Go(__goid);
}
void* QWinEventNotifier_new2(long long __goid, void* parent) {
	return new QWinEventNotifier__Go(__goid, (QObject*)parent);
}
char* QWinEventNotifier_tr(const char* s, long long* __out_len) {
	auto&& __r = QWinEventNotifier::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QWinEventNotifier_isEnabled(void* self) {
	auto&& __r = ((QWinEventNotifier*)self)->isEnabled();
	return __r;
}
void QWinEventNotifier_setEnabled(void* self, int enable) {
	((QWinEventNotifier*)self)->setEnabled(enable);
}
void* QWinEventNotifier_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QWinEventNotifier::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QWinEventNotifier_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QWinEventNotifier::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QWinEventNotifier_sender(void* self) {
	auto&& __r = ((QWinEventNotifier__Go*)self)->sender();
	return (void*)(__r);
}
int QWinEventNotifier_senderSignalIndex(void* self) {
	auto&& __r = ((QWinEventNotifier__Go*)self)->senderSignalIndex();
	return __r;
}
int QWinEventNotifier_receivers(void* self, const char* signal) {
	auto&& __r = ((QWinEventNotifier__Go*)self)->receivers(signal);
	return __r;
}
int QWinEventNotifier_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QWinEventNotifier__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QWinEventNotifier_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QWinEventNotifier::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QWinEventNotifier_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QWinEventNotifier::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QWinEventNotifier_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QWinEventNotifier::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QWinEventNotifier_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QWinEventNotifier::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QWinEventNotifier_metaObject_default(void* self) {
	auto&& __r = ((QWinEventNotifier*)self)->QWinEventNotifier::metaObject();
	return (void*)(__r);
}
int QWinEventNotifier_event_default(void* self, void* e) {
	auto&& __r = ((QWinEventNotifier__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
int QWinEventNotifier_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QWinEventNotifier*)self)->QWinEventNotifier::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QWinEventNotifier_timerEvent_default(void* self, void* event) {
	((QWinEventNotifier__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QWinEventNotifier_childEvent_default(void* self, void* event) {
	((QWinEventNotifier__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QWinEventNotifier_customEvent_default(void* self, void* event) {
	((QWinEventNotifier__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QWinEventNotifier_connectNotify_default(void* self, void* signal) {
	((QWinEventNotifier__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QWinEventNotifier_disconnectNotify_default(void* self, void* signal) {
	((QWinEventNotifier__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QWinEventNotifier_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWinEventNotifier*)self, &QWinEventNotifier::destroyed, (QWinEventNotifier*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWinEventNotifier);
	__conns_QWinEventNotifier[__id] = __c;
}
void QWinEventNotifier_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWinEventNotifier*)self, &QWinEventNotifier::destroyed, (QWinEventNotifier*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWinEventNotifier);
	__conns_QWinEventNotifier[__id] = __c;
}
void QWinEventNotifier_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QWinEventNotifier*)self, &QWinEventNotifier::objectNameChanged, (QWinEventNotifier*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QWinEventNotifier);
	__conns_QWinEventNotifier[__id] = __c;
}
void QWinEventNotifier__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QWinEventNotifier);
		auto __it = __conns_QWinEventNotifier.find(id);
		if (__it == __conns_QWinEventNotifier.end()) return;
		__c = __it->second;
		__conns_QWinEventNotifier.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QWinEventNotifier__Go_new(long long id, void* parent) { return new QWinEventNotifier__Go(id, (QObject*)parent); }
void QWinEventNotifier_delete(void* self) { delete (QWinEventNotifier*)self; }

}

#include <QtGui/qscreen.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qtransform.h>
#include <QtGui/qpixmap.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qscreen.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QScreen;
static std::mutex __connsMu_QScreen;

extern "C" {

void* QScreen_base(void* self) { return (void*)static_cast<QObject*>((QScreen*)self); }
void* QScreen_fromQObject(void* o) { return (void*)qobject_cast<QScreen*>((QObject*)o); }
char* QScreen_tr(const char* s, long long* __out_len) {
	auto&& __r = QScreen::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QScreen_name(void* self, long long* __out_len) {
	auto&& __r = ((QScreen*)self)->name();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QScreen_manufacturer(void* self, long long* __out_len) {
	auto&& __r = ((QScreen*)self)->manufacturer();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QScreen_model(void* self, long long* __out_len) {
	auto&& __r = ((QScreen*)self)->model();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QScreen_serialNumber(void* self, long long* __out_len) {
	auto&& __r = ((QScreen*)self)->serialNumber();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QScreen_depth(void* self) {
	auto&& __r = ((QScreen*)self)->depth();
	return __r;
}
void* QScreen_size(void* self) {
	auto&& __r = ((QScreen*)self)->size();
	return new QSize(__r);
}
void* QScreen_geometry(void* self) {
	auto&& __r = ((QScreen*)self)->geometry();
	return new QRect(__r);
}
void* QScreen_physicalSize(void* self) {
	auto&& __r = ((QScreen*)self)->physicalSize();
	return new QSizeF(__r);
}
double QScreen_physicalDotsPerInchX(void* self) {
	auto&& __r = ((QScreen*)self)->physicalDotsPerInchX();
	return __r;
}
double QScreen_physicalDotsPerInchY(void* self) {
	auto&& __r = ((QScreen*)self)->physicalDotsPerInchY();
	return __r;
}
double QScreen_physicalDotsPerInch(void* self) {
	auto&& __r = ((QScreen*)self)->physicalDotsPerInch();
	return __r;
}
double QScreen_logicalDotsPerInchX(void* self) {
	auto&& __r = ((QScreen*)self)->logicalDotsPerInchX();
	return __r;
}
double QScreen_logicalDotsPerInchY(void* self) {
	auto&& __r = ((QScreen*)self)->logicalDotsPerInchY();
	return __r;
}
double QScreen_logicalDotsPerInch(void* self) {
	auto&& __r = ((QScreen*)self)->logicalDotsPerInch();
	return __r;
}
double QScreen_devicePixelRatio(void* self) {
	auto&& __r = ((QScreen*)self)->devicePixelRatio();
	return __r;
}
void* QScreen_availableSize(void* self) {
	auto&& __r = ((QScreen*)self)->availableSize();
	return new QSize(__r);
}
void* QScreen_availableGeometry(void* self) {
	auto&& __r = ((QScreen*)self)->availableGeometry();
	return new QRect(__r);
}
void** QScreen_virtualSiblings(void* self, long long* __out_len) {
	auto&& __r = ((QScreen*)self)->virtualSiblings();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = (void*)(__e);
	return __arr;
}
void* QScreen_virtualSiblingAt(void* self, void* point) {
	auto&& __r = ((QScreen*)self)->virtualSiblingAt(*(QPoint*)point);
	return (void*)(__r);
}
void* QScreen_virtualSize(void* self) {
	auto&& __r = ((QScreen*)self)->virtualSize();
	return new QSize(__r);
}
void* QScreen_virtualGeometry(void* self) {
	auto&& __r = ((QScreen*)self)->virtualGeometry();
	return new QRect(__r);
}
void* QScreen_availableVirtualSize(void* self) {
	auto&& __r = ((QScreen*)self)->availableVirtualSize();
	return new QSize(__r);
}
void* QScreen_availableVirtualGeometry(void* self) {
	auto&& __r = ((QScreen*)self)->availableVirtualGeometry();
	return new QRect(__r);
}
int QScreen_primaryOrientation(void* self) {
	auto&& __r = ((QScreen*)self)->primaryOrientation();
	return (int)(__r);
}
int QScreen_orientation(void* self) {
	auto&& __r = ((QScreen*)self)->orientation();
	return (int)(__r);
}
int QScreen_nativeOrientation(void* self) {
	auto&& __r = ((QScreen*)self)->nativeOrientation();
	return (int)(__r);
}
int QScreen_angleBetween(void* self, int a, int b) {
	auto&& __r = ((QScreen*)self)->angleBetween((Qt::ScreenOrientation)(a), (Qt::ScreenOrientation)(b));
	return __r;
}
void* QScreen_transformBetween(void* self, int a, int b, void* target) {
	auto&& __r = ((QScreen*)self)->transformBetween((Qt::ScreenOrientation)(a), (Qt::ScreenOrientation)(b), *(QRect*)target);
	return new QTransform(__r);
}
void* QScreen_mapBetween(void* self, int a, int b, void* rect) {
	auto&& __r = ((QScreen*)self)->mapBetween((Qt::ScreenOrientation)(a), (Qt::ScreenOrientation)(b), *(QRect*)rect);
	return new QRect(__r);
}
int QScreen_isPortrait(void* self, int orientation) {
	auto&& __r = ((QScreen*)self)->isPortrait((Qt::ScreenOrientation)(orientation));
	return __r;
}
int QScreen_isLandscape(void* self, int orientation) {
	auto&& __r = ((QScreen*)self)->isLandscape((Qt::ScreenOrientation)(orientation));
	return __r;
}
void* QScreen_grabWindow(void* self) {
	auto&& __r = ((QScreen*)self)->grabWindow();
	return new QPixmap(__r);
}
double QScreen_refreshRate(void* self) {
	auto&& __r = ((QScreen*)self)->refreshRate();
	return __r;
}
void QScreen_geometryChanged(void* self, void* geometry) {
	((QScreen*)self)->geometryChanged(*(QRect*)geometry);
}
void QScreen_availableGeometryChanged(void* self, void* geometry) {
	((QScreen*)self)->availableGeometryChanged(*(QRect*)geometry);
}
void QScreen_physicalSizeChanged(void* self, void* size) {
	((QScreen*)self)->physicalSizeChanged(*(QSizeF*)size);
}
void QScreen_physicalDotsPerInchChanged(void* self, double dpi) {
	((QScreen*)self)->physicalDotsPerInchChanged(dpi);
}
void QScreen_logicalDotsPerInchChanged(void* self, double dpi) {
	((QScreen*)self)->logicalDotsPerInchChanged(dpi);
}
void QScreen_virtualGeometryChanged(void* self, void* rect) {
	((QScreen*)self)->virtualGeometryChanged(*(QRect*)rect);
}
void QScreen_primaryOrientationChanged(void* self, int orientation) {
	((QScreen*)self)->primaryOrientationChanged((Qt::ScreenOrientation)(orientation));
}
void QScreen_orientationChanged(void* self, int orientation) {
	((QScreen*)self)->orientationChanged((Qt::ScreenOrientation)(orientation));
}
void QScreen_refreshRateChanged(void* self, double refreshRate) {
	((QScreen*)self)->refreshRateChanged(refreshRate);
}
void* QScreen_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QScreen::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QScreen_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QScreen::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
char* QScreen_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QScreen::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QScreen_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QScreen::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QScreen_grabWindow2(void* self, unsigned long long window) {
	auto&& __r = ((QScreen*)self)->grabWindow(window);
	return new QPixmap(__r);
}
void* QScreen_grabWindow3(void* self, unsigned long long window, int x) {
	auto&& __r = ((QScreen*)self)->grabWindow(window, x);
	return new QPixmap(__r);
}
void* QScreen_grabWindow4(void* self, unsigned long long window, int x, int y) {
	auto&& __r = ((QScreen*)self)->grabWindow(window, x, y);
	return new QPixmap(__r);
}
void* QScreen_grabWindow5(void* self, unsigned long long window, int x, int y, int w) {
	auto&& __r = ((QScreen*)self)->grabWindow(window, x, y, w);
	return new QPixmap(__r);
}
void* QScreen_grabWindow6(void* self, unsigned long long window, int x, int y, int w, int h) {
	auto&& __r = ((QScreen*)self)->grabWindow(window, x, y, w, h);
	return new QPixmap(__r);
}
void* QScreen_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QScreen::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QScreen_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QScreen::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QScreen_connect_geometryChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScreen*)self, &QScreen::geometryChanged, (QScreen*)self, [__id](const QRect& geometry){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&geometry);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScreen);
	__conns_QScreen[__id] = __c;
}
void QScreen_connect_availableGeometryChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScreen*)self, &QScreen::availableGeometryChanged, (QScreen*)self, [__id](const QRect& geometry){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&geometry);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScreen);
	__conns_QScreen[__id] = __c;
}
void QScreen_connect_physicalSizeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScreen*)self, &QScreen::physicalSizeChanged, (QScreen*)self, [__id](const QSizeF& size){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&size);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScreen);
	__conns_QScreen[__id] = __c;
}
void QScreen_connect_physicalDotsPerInchChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScreen*)self, &QScreen::physicalDotsPerInchChanged, (QScreen*)self, [__id](qreal dpi){
		if (cppgo_sig_void_unsignedlonglong) ((void(*)(long long, unsigned long long))cppgo_sig_void_unsignedlonglong)(__id, __builtin_bit_cast(unsigned long long, (double)(dpi)));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScreen);
	__conns_QScreen[__id] = __c;
}
void QScreen_connect_logicalDotsPerInchChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScreen*)self, &QScreen::logicalDotsPerInchChanged, (QScreen*)self, [__id](qreal dpi){
		if (cppgo_sig_void_unsignedlonglong) ((void(*)(long long, unsigned long long))cppgo_sig_void_unsignedlonglong)(__id, __builtin_bit_cast(unsigned long long, (double)(dpi)));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScreen);
	__conns_QScreen[__id] = __c;
}
void QScreen_connect_virtualGeometryChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScreen*)self, &QScreen::virtualGeometryChanged, (QScreen*)self, [__id](const QRect& rect){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&rect);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScreen);
	__conns_QScreen[__id] = __c;
}
void QScreen_connect_primaryOrientationChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScreen*)self, &QScreen::primaryOrientationChanged, (QScreen*)self, [__id](Qt::ScreenOrientation orientation){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(orientation));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScreen);
	__conns_QScreen[__id] = __c;
}
void QScreen_connect_orientationChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScreen*)self, &QScreen::orientationChanged, (QScreen*)self, [__id](Qt::ScreenOrientation orientation){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(orientation));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScreen);
	__conns_QScreen[__id] = __c;
}
void QScreen_connect_refreshRateChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScreen*)self, &QScreen::refreshRateChanged, (QScreen*)self, [__id](qreal refreshRate){
		if (cppgo_sig_void_unsignedlonglong) ((void(*)(long long, unsigned long long))cppgo_sig_void_unsignedlonglong)(__id, __builtin_bit_cast(unsigned long long, (double)(refreshRate)));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScreen);
	__conns_QScreen[__id] = __c;
}
void QScreen_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScreen*)self, &QScreen::destroyed, (QScreen*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScreen);
	__conns_QScreen[__id] = __c;
}
void QScreen_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScreen*)self, &QScreen::destroyed, (QScreen*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScreen);
	__conns_QScreen[__id] = __c;
}
void QScreen_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QScreen*)self, &QScreen::objectNameChanged, (QScreen*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QScreen);
	__conns_QScreen[__id] = __c;
}
void QScreen__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QScreen);
		auto __it = __conns_QScreen.find(id);
		if (__it == __conns_QScreen.end()) return;
		__c = __it->second;
		__conns_QScreen.erase(__it);
	}
	QObject::disconnect(__c);
}
void QScreen_delete(void* self) { delete (QScreen*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include "gen_qcloseevent.h"
#include "gen_purego_cb.h"

class QCloseEvent__Go : public QCloseEvent {
public:
	long long __goId;
	unsigned char __goOn[2] = {};
	template<typename... A> QCloseEvent__Go(long long id, A&&... a) : QCloseEvent(std::forward<A>(a)...), __goId(id) {}
	QCloseEvent* clone() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, -7032813281048616478LL, &__h); if (__h) return (QCloseEvent*)__r; }
		return QCloseEvent::clone();
	}
	void setAccepted(bool accepted) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6367896784077321574LL, accepted, &__h); if (__h) return; }
		QCloseEvent::setAccepted(accepted);
	}
};

extern "C" {

void* QCloseEvent_base(void* self) { return (void*)static_cast<QEvent*>((QCloseEvent*)self); }
void QCloseEvent__set_override(void* self, int idx) { ((QCloseEvent__Go*)self)->__goOn[idx] = 1; }
void* QCloseEvent_new(long long __goid) {
	return new QCloseEvent__Go(__goid);
}
int QCloseEvent_registerEventType() {
	auto&& __r = QCloseEvent::registerEventType();
	return __r;
}
int QCloseEvent_registerEventType2(int hint) {
	auto&& __r = QCloseEvent::registerEventType(hint);
	return __r;
}
void* QCloseEvent_clone_default(void* self) {
	auto&& __r = ((QCloseEvent*)self)->QCloseEvent::clone();
	return (void*)(__r);
}
void QCloseEvent_setAccepted_default(void* self, int accepted) {
	((QCloseEvent*)self)->QCloseEvent::setAccepted(accepted);
}
void* QCloseEvent__Go_new(long long id) { return new QCloseEvent__Go(id); }
void QCloseEvent_delete(void* self) { delete (QCloseEvent*)self; }

}

#include <QtGui/qevent.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qnamespace.h>
#include "gen_qapplicationstatechangeevent.h"

extern "C" {

void* QApplicationStateChangeEvent_base(void* self) { return (void*)static_cast<QEvent*>((QApplicationStateChangeEvent*)self); }
void* QApplicationStateChangeEvent_new(int state) {
	return new QApplicationStateChangeEvent((Qt::ApplicationState)(state));
}
int QApplicationStateChangeEvent_applicationState(void* self) {
	auto&& __r = ((QApplicationStateChangeEvent*)self)->applicationState();
	return (int)(__r);
}
int QApplicationStateChangeEvent_registerEventType() {
	auto&& __r = QApplicationStateChangeEvent::registerEventType();
	return __r;
}
int QApplicationStateChangeEvent_registerEventType2(int hint) {
	auto&& __r = QApplicationStateChangeEvent::registerEventType(hint);
	return __r;
}
void QApplicationStateChangeEvent_delete(void* self) { delete (QApplicationStateChangeEvent*)self; }

}

#include <QtGui/qtextformat.h>
#include <QtCore/qstring.h>
#include <QtGui/qfont.h>
#include <QtCore/qvariant.h>
#include <QtGui/qcolor.h>
#include <QtGui/qpen.h>
#include <QtGui/qbrush.h>
#include <QtCore/qnamespace.h>
#include "gen_qtextimageformat.h"

extern "C" {

void* QTextImageFormat_base(void* self) { return (void*)static_cast<QTextCharFormat*>((QTextImageFormat*)self); }
void QTextImageFormat_chain(void* self, void** out) {
	QTextImageFormat* __p = (QTextImageFormat*)self;
	QTextCharFormat* __b0 = static_cast<QTextCharFormat*>(__p); out[0] = (void*)__b0;
	QTextFormat* __b1 = static_cast<QTextFormat*>(__b0); out[1] = (void*)__b1;
}
void* QTextImageFormat_new() {
	return new QTextImageFormat();
}
void* QTextImageFormat_new2(void* arg1) {
	return new QTextImageFormat(*(QTextImageFormat*)arg1);
}
int QTextImageFormat_isValid(void* self) {
	auto&& __r = ((QTextImageFormat*)self)->isValid();
	return __r;
}
void QTextImageFormat_setName(void* self, const char* name, long long name_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	((QTextImageFormat*)self)->setName(name_qstr);
}
char* QTextImageFormat_name(void* self, long long* __out_len) {
	auto&& __r = ((QTextImageFormat*)self)->name();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextImageFormat_setWidth(void* self, double width) {
	((QTextImageFormat*)self)->setWidth(width);
}
double QTextImageFormat_width(void* self) {
	auto&& __r = ((QTextImageFormat*)self)->width();
	return __r;
}
void QTextImageFormat_setMaximumWidth(void* self, void* maxWidth) {
	((QTextImageFormat*)self)->setMaximumWidth(*(QTextLength*)maxWidth);
}
void* QTextImageFormat_maximumWidth(void* self) {
	auto&& __r = ((QTextImageFormat*)self)->maximumWidth();
	return new QTextLength(__r);
}
void QTextImageFormat_setHeight(void* self, double height) {
	((QTextImageFormat*)self)->setHeight(height);
}
double QTextImageFormat_height(void* self) {
	auto&& __r = ((QTextImageFormat*)self)->height();
	return __r;
}
void QTextImageFormat_setQuality(void* self, int quality) {
	((QTextImageFormat*)self)->setQuality(quality);
}
void QTextImageFormat_setQuality2(void* self) {
	((QTextImageFormat*)self)->setQuality();
}
int QTextImageFormat_quality(void* self) {
	auto&& __r = ((QTextImageFormat*)self)->quality();
	return __r;
}
void QTextImageFormat_delete(void* self) { delete (QTextImageFormat*)self; }

}

#include <QtGui/qaccessible.h>
#include <QtCore/qstring.h>
#include "gen_qaccessibleactioninterface.h"
#include "gen_purego_cb.h"

class QAccessibleActionInterface__Go : public QAccessibleActionInterface {
public:
	long long __goId;
	unsigned char __goOn[5] = {};
	template<typename... A> QAccessibleActionInterface__Go(long long id, A&&... a) : QAccessibleActionInterface(std::forward<A>(a)...), __goId(id) {}
	QList<QString> actionNames() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virtp_voidp) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, long long*, int*))cppgo_virtp_voidp)(__goId, 5274920687135865149LL, &__len, &__h);
			if (__h) {
				QList<QString> __r; __r.reserve(__len);
				char** __aa = (char**)__a;
				for (long long __i = 0; __i < __len; __i++) { __r.append(QString::fromUtf8(__aa[__i])); free(__aa[__i]); }
				if (__a) free(__a);
				return __r;
			}
		}
		return QList<QString>{};
	}
	void doAction(const QString& actionName) override {
		int __h = 0;
		QByteArray actionName_ba = actionName.toUtf8();
		if (__goOn[1] && cppgo_virt_void_constcharp_longlong) { ((void(*)(long long, long long, const char*, long long, int*))cppgo_virt_void_constcharp_longlong)(__goId, -219369122549130410LL, (char*)actionName_ba.constData(), actionName_ba.size(), &__h); if (__h) return; }
		return;
	}
	QList<QString> keyBindingsForAction(const QString& actionName) const override {
		int __h = 0;
		QByteArray actionName_ba = actionName.toUtf8();
		if (__goOn[2] && cppgo_virtp_voidp_constcharp_longlong) {
			long long __len = 0;
			void* __a = ((void*(*)(long long, long long, const char*, long long, long long*, int*))cppgo_virtp_voidp_constcharp_longlong)(__goId, -9040881633578371181LL, (char*)actionName_ba.constData(), actionName_ba.size(), &__len, &__h);
			if (__h) {
				QList<QString> __r; __r.reserve(__len);
				char** __aa = (char**)__a;
				for (long long __i = 0; __i < __len; __i++) { __r.append(QString::fromUtf8(__aa[__i])); free(__aa[__i]); }
				if (__a) free(__a);
				return __r;
			}
		}
		return QList<QString>{};
	}
	QString localizedActionName(const QString& name) const override {
		int __h = 0;
		QByteArray name_ba = name.toUtf8();
		if (__goOn[3] && cppgo_virtp_charp_constcharp_longlong) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, const char*, long long, long long*, int*))cppgo_virtp_charp_constcharp_longlong)(__goId, 5320856891804608331LL, (char*)name_ba.constData(), name_ba.size(), &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QAccessibleActionInterface::localizedActionName(name);
	}
	QString localizedActionDescription(const QString& name) const override {
		int __h = 0;
		QByteArray name_ba = name.toUtf8();
		if (__goOn[4] && cppgo_virtp_charp_constcharp_longlong) {
			long long __len = 0;
			char* __s = ((char*(*)(long long, long long, const char*, long long, long long*, int*))cppgo_virtp_charp_constcharp_longlong)(__goId, -2998875658810589822LL, (char*)name_ba.constData(), name_ba.size(), &__len, &__h);
			if (__h) { QString __r = QString::fromUtf8(__s, __len); free(__s); return __r; }
		}
		return QAccessibleActionInterface::localizedActionDescription(name);
	}
};

extern "C" {

void QAccessibleActionInterface__set_override(void* self, int idx) { ((QAccessibleActionInterface__Go*)self)->__goOn[idx] = 1; }
char* QAccessibleActionInterface_tr(const char* sourceText, long long* __out_len) {
	auto&& __r = QAccessibleActionInterface::tr(sourceText);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char** QAccessibleActionInterface_actionNames(void* self, long long* __out_len) {
	auto&& __r = ((QAccessibleActionInterface*)self)->actionNames();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char* QAccessibleActionInterface_localizedActionName(void* self, const char* name, long long name_len, long long* __out_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = ((QAccessibleActionInterface*)self)->localizedActionName(name_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleActionInterface_localizedActionDescription(void* self, const char* name, long long name_len, long long* __out_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = ((QAccessibleActionInterface*)self)->localizedActionDescription(name_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QAccessibleActionInterface_doAction(void* self, const char* actionName, long long actionName_len) {
	QString actionName_qstr = QString::fromUtf8(actionName, actionName_len);
	((QAccessibleActionInterface*)self)->doAction(actionName_qstr);
}
char** QAccessibleActionInterface_keyBindingsForAction(void* self, const char* actionName, long long actionName_len, long long* __out_len) {
	QString actionName_qstr = QString::fromUtf8(actionName, actionName_len);
	auto&& __r = ((QAccessibleActionInterface*)self)->keyBindingsForAction(actionName_qstr);
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
char* QAccessibleActionInterface_pressAction(long long* __out_len) {
	auto&& __r = QAccessibleActionInterface::pressAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleActionInterface_increaseAction(long long* __out_len) {
	auto&& __r = QAccessibleActionInterface::increaseAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleActionInterface_decreaseAction(long long* __out_len) {
	auto&& __r = QAccessibleActionInterface::decreaseAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleActionInterface_showMenuAction(long long* __out_len) {
	auto&& __r = QAccessibleActionInterface::showMenuAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleActionInterface_setFocusAction(long long* __out_len) {
	auto&& __r = QAccessibleActionInterface::setFocusAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleActionInterface_toggleAction(long long* __out_len) {
	auto&& __r = QAccessibleActionInterface::toggleAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleActionInterface_scrollLeftAction(long long* __out_len) {
	auto&& __r = QAccessibleActionInterface::scrollLeftAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleActionInterface_scrollRightAction(long long* __out_len) {
	auto&& __r = QAccessibleActionInterface::scrollRightAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleActionInterface_scrollUpAction(long long* __out_len) {
	auto&& __r = QAccessibleActionInterface::scrollUpAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleActionInterface_scrollDownAction(long long* __out_len) {
	auto&& __r = QAccessibleActionInterface::scrollDownAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleActionInterface_nextPageAction(long long* __out_len) {
	auto&& __r = QAccessibleActionInterface::nextPageAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleActionInterface_previousPageAction(long long* __out_len) {
	auto&& __r = QAccessibleActionInterface::previousPageAction();
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAccessibleActionInterface_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QAccessibleActionInterface*)self)->operator=(*(QAccessibleActionInterface*)arg1);
	return (void*)&(__r);
}
char* QAccessibleActionInterface_tr2(const char* sourceText, const char* disambiguation, long long* __out_len) {
	auto&& __r = QAccessibleActionInterface::tr(sourceText, disambiguation);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleActionInterface_tr3(const char* sourceText, const char* disambiguation, int n, long long* __out_len) {
	auto&& __r = QAccessibleActionInterface::tr(sourceText, disambiguation, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleActionInterface_localizedActionName_default(void* self, const char* name, long long name_len, long long* __out_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = ((QAccessibleActionInterface*)self)->QAccessibleActionInterface::localizedActionName(name_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QAccessibleActionInterface_localizedActionDescription_default(void* self, const char* name, long long name_len, long long* __out_len) {
	QString name_qstr = QString::fromUtf8(name, name_len);
	auto&& __r = ((QAccessibleActionInterface*)self)->QAccessibleActionInterface::localizedActionDescription(name_qstr);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QAccessibleActionInterface__Go_new(long long id) { return new QAccessibleActionInterface__Go(id); }
void QAccessibleActionInterface_delete(void* self) { delete (QAccessibleActionInterface*)self; }

}

#include <QtGui/qbitmap.h>
#include <QtGui/qpixmap.h>
#include <QtCore/qsize.h>
#include <QtCore/qstring.h>
#include <QtCore/qvariant.h>
#include <QtGui/qimage.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qtransform.h>
#include <QtCore/qrect.h>
#include <QtGui/qcolor.h>
#include <QtGui/qimagereader.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qiodevice.h>
#include <QtGui/qregion.h>
#include <QtGui/qpaintengine.h>
#include <QtGui/qpaintdevice.h>
#include <QtGui/qpainter.h>
#include <QtCore/qpoint.h>
#include "gen_qbitmap.h"
#include "gen_purego_cb.h"

class QBitmap__Go : public QBitmap {
public:
	long long __goId;
	unsigned char __goOn[6] = {};
	template<typename... A> QBitmap__Go(long long id, A&&... a) : QBitmap(std::forward<A>(a)...), __goId(id) {}
	using QBitmap::getDecodedMetricF;
	int cppgoBase_metric(int __a0) { return QBitmap::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QBitmap::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QBitmap::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QBitmap::sharedPainter(); }
	int devType() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QBitmap::devType();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QBitmap::paintEngine();
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QBitmap::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QBitmap::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QBitmap::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QBitmap::sharedPainter();
	}
};

extern "C" {

void* QBitmap_base(void* self) { return (void*)static_cast<QPixmap*>((QBitmap*)self); }
void QBitmap_chain(void* self, void** out) {
	QBitmap* __p = (QBitmap*)self;
	QPixmap* __b0 = static_cast<QPixmap*>(__p); out[0] = (void*)__b0;
	QPaintDevice* __b1 = static_cast<QPaintDevice*>(__b0); out[1] = (void*)__b1;
}
void QBitmap__set_override(void* self, int idx) { ((QBitmap__Go*)self)->__goOn[idx] = 1; }
void* QBitmap_new(long long __goid) {
	return new QBitmap__Go(__goid);
}
void* QBitmap_new2(long long __goid, void* arg1) {
	return new QBitmap__Go(__goid, *(QPixmap*)arg1);
}
void* QBitmap_new3(long long __goid, int w, int h) {
	return new QBitmap__Go(__goid, w, h);
}
void* QBitmap_new4(long long __goid, void* arg1) {
	return new QBitmap__Go(__goid, *(QSize*)arg1);
}
void* QBitmap_new5(long long __goid, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QBitmap__Go(__goid, fileName_qstr);
}
void* QBitmap_new6(long long __goid, void* arg1) {
	return new QBitmap__Go(__goid, *(QBitmap*)arg1);
}
void* QBitmap_new7(long long __goid, const char* fileName, long long fileName_len, const char* format) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	return new QBitmap__Go(__goid, fileName_qstr, format);
}
void* QBitmap_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QBitmap*)self)->operator=(*(QPixmap*)arg1);
	return (void*)&(__r);
}
void QBitmap_swap(void* self, void* other) {
	((QBitmap*)self)->swap(*(QBitmap*)other);
}
void* QBitmap_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QBitmap*)self);
	return new QVariant(__r);
}
void QBitmap_clear(void* self) {
	((QBitmap*)self)->clear();
}
void* QBitmap_fromImage(void* image) {
	auto&& __r = QBitmap::fromImage(*(QImage*)image);
	return new QBitmap(__r);
}
void* QBitmap_fromData(void* size, const unsigned char* bits) {
	auto&& __r = QBitmap::fromData(*(QSize*)size, bits);
	return new QBitmap(__r);
}
void* QBitmap_fromPixmap(void* pixmap) {
	auto&& __r = QBitmap::fromPixmap(*(QPixmap*)pixmap);
	return new QBitmap(__r);
}
void* QBitmap_transformed(void* self, void* matrix) {
	auto&& __r = ((QBitmap*)self)->transformed(*(QTransform*)matrix);
	return new QBitmap(__r);
}
void* QBitmap_OperatorAssign2(void* self, void* arg1) {
	auto&& __r = ((QBitmap*)self)->operator=(*(QBitmap*)arg1);
	return (void*)&(__r);
}
int QBitmap_defaultDepth() {
	auto&& __r = QBitmap::defaultDepth();
	return __r;
}
void* QBitmap_trueMatrix(void* m, int w, int h) {
	auto&& __r = QBitmap::trueMatrix(*(QTransform*)m, w, h);
	return new QTransform(__r);
}
void* QBitmap_fromImageReader(void* imageReader) {
	auto&& __r = QBitmap::fromImageReader((QImageReader*)imageReader);
	return new QPixmap(__r);
}
double QBitmap_devicePixelRatioFScale() {
	auto&& __r = QBitmap::devicePixelRatioFScale();
	return __r;
}
int QBitmap_encodeMetricF(int metric, double value) {
	auto&& __r = QBitmap::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QBitmap_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QBitmap__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
void* QBitmap_fromImage2(void* image, int flags) {
	auto&& __r = QBitmap::fromImage(*(QImage*)image, QFlags<Qt::ImageConversionFlag>(flags));
	return new QBitmap(__r);
}
void* QBitmap_fromData2(void* size, const unsigned char* bits, int monoFormat) {
	auto&& __r = QBitmap::fromData(*(QSize*)size, bits, (QImage::Format)(monoFormat));
	return new QBitmap(__r);
}
void* QBitmap_fromImageReader2(void* imageReader, int flags) {
	auto&& __r = QBitmap::fromImageReader((QImageReader*)imageReader, QFlags<Qt::ImageConversionFlag>(flags));
	return new QPixmap(__r);
}
void* QBitmap_toQVariant(void* self) {
	auto&& __r = QVariant::fromValue(*(QBitmap*)self);
	return new QVariant(__r);
}
int QBitmap_devType_default(void* self) {
	auto&& __r = ((QBitmap*)self)->QBitmap::devType();
	return __r;
}
void* QBitmap_paintEngine_default(void* self) {
	auto&& __r = ((QBitmap*)self)->QBitmap::paintEngine();
	return (void*)(__r);
}
int QBitmap_metric_default(void* self, int arg1) {
	auto&& __r = ((QBitmap__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QBitmap_initPainter_default(void* self, void* painter) {
	((QBitmap__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QBitmap_redirected_default(void* self, void* offset) {
	auto&& __r = ((QBitmap__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QBitmap_sharedPainter_default(void* self) {
	auto&& __r = ((QBitmap__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void* QBitmap__Go_new(long long id) { return new QBitmap__Go(id); }
void QBitmap_delete(void* self) { delete (QBitmap*)self; }

}

#include <QtGui/qmatrix4x4.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qtransform.h>
#include <QtGui/qvectornd.h>
#include <QtGui/qquaternion.h>
#include <QtCore/qrect.h>
#include <QtCore/qpoint.h>
#include <QtCore/qvariant.h>
#include "gen_qmatrix4x4.h"

extern "C" {

void* QMatrix4x4_new() {
	return new QMatrix4x4();
}
void* QMatrix4x4_new2(int arg1) {
	return new QMatrix4x4((Qt::Initialization)(arg1));
}
void* QMatrix4x4_new3(float m11, float m12, float m13, float m14, float m21, float m22, float m23, float m24, float m31, float m32, float m33, float m34, float m41, float m42, float m43, float m44) {
	return new QMatrix4x4(m11, m12, m13, m14, m21, m22, m23, m24, m31, m32, m33, m34, m41, m42, m43, m44);
}
void* QMatrix4x4_new4(void* transform) {
	return new QMatrix4x4(*(QTransform*)transform);
}
void* QMatrix4x4_new5(void* arg1) {
	return new QMatrix4x4(*(QMatrix4x4*)arg1);
}
void* QMatrix4x4_column(void* self, int index) {
	auto&& __r = ((QMatrix4x4*)self)->column(index);
	return new QVector4D(__r);
}
void QMatrix4x4_setColumn(void* self, int index, void* value) {
	((QMatrix4x4*)self)->setColumn(index, *(QVector4D*)value);
}
void* QMatrix4x4_row(void* self, int index) {
	auto&& __r = ((QMatrix4x4*)self)->row(index);
	return new QVector4D(__r);
}
void QMatrix4x4_setRow(void* self, int index, void* value) {
	((QMatrix4x4*)self)->setRow(index, *(QVector4D*)value);
}
int QMatrix4x4_isAffine(void* self) {
	auto&& __r = ((QMatrix4x4*)self)->isAffine();
	return __r;
}
int QMatrix4x4_isIdentity(void* self) {
	auto&& __r = ((QMatrix4x4*)self)->isIdentity();
	return __r;
}
void QMatrix4x4_setToIdentity(void* self) {
	((QMatrix4x4*)self)->setToIdentity();
}
void QMatrix4x4_fill(void* self, float value) {
	((QMatrix4x4*)self)->fill(value);
}
double QMatrix4x4_determinant(void* self) {
	auto&& __r = ((QMatrix4x4*)self)->determinant();
	return __r;
}
void* QMatrix4x4_inverted(void* self) {
	auto&& __r = ((QMatrix4x4*)self)->inverted();
	return new QMatrix4x4(__r);
}
void* QMatrix4x4_transposed(void* self) {
	auto&& __r = ((QMatrix4x4*)self)->transposed();
	return new QMatrix4x4(__r);
}
void* QMatrix4x4_OperatorPlusAssign(void* self, void* other) {
	auto&& __r = ((QMatrix4x4*)self)->operator+=(*(QMatrix4x4*)other);
	return (void*)&(__r);
}
void* QMatrix4x4_OperatorMinusAssign(void* self, void* other) {
	auto&& __r = ((QMatrix4x4*)self)->operator-=(*(QMatrix4x4*)other);
	return (void*)&(__r);
}
void* QMatrix4x4_OperatorMultiplyAssign(void* self, void* other) {
	auto&& __r = ((QMatrix4x4*)self)->operator*=(*(QMatrix4x4*)other);
	return (void*)&(__r);
}
void* QMatrix4x4_OperatorMultiplyAssign2(void* self, float factor) {
	auto&& __r = ((QMatrix4x4*)self)->operator*=(factor);
	return (void*)&(__r);
}
void* QMatrix4x4_OperatorDivideAssign(void* self, float divisor) {
	auto&& __r = ((QMatrix4x4*)self)->operator/=(divisor);
	return (void*)&(__r);
}
int QMatrix4x4_OperatorEqual(void* self, void* other) {
	auto&& __r = ((QMatrix4x4*)self)->operator==(*(QMatrix4x4*)other);
	return __r;
}
int QMatrix4x4_OperatorNotEqual(void* self, void* other) {
	auto&& __r = ((QMatrix4x4*)self)->operator!=(*(QMatrix4x4*)other);
	return __r;
}
void QMatrix4x4_scale(void* self, void* vector) {
	((QMatrix4x4*)self)->scale(*(QVector3D*)vector);
}
void QMatrix4x4_translate(void* self, void* vector) {
	((QMatrix4x4*)self)->translate(*(QVector3D*)vector);
}
void QMatrix4x4_rotate(void* self, float angle, void* vector) {
	((QMatrix4x4*)self)->rotate(angle, *(QVector3D*)vector);
}
void QMatrix4x4_scale2(void* self, float x, float y) {
	((QMatrix4x4*)self)->scale(x, y);
}
void QMatrix4x4_scale3(void* self, float x, float y, float z) {
	((QMatrix4x4*)self)->scale(x, y, z);
}
void QMatrix4x4_scale4(void* self, float factor) {
	((QMatrix4x4*)self)->scale(factor);
}
void QMatrix4x4_translate2(void* self, float x, float y) {
	((QMatrix4x4*)self)->translate(x, y);
}
void QMatrix4x4_translate3(void* self, float x, float y, float z) {
	((QMatrix4x4*)self)->translate(x, y, z);
}
void QMatrix4x4_rotate2(void* self, float angle, float x, float y) {
	((QMatrix4x4*)self)->rotate(angle, x, y);
}
void QMatrix4x4_rotate3(void* self, void* quaternion) {
	((QMatrix4x4*)self)->rotate(*(QQuaternion*)quaternion);
}
void QMatrix4x4_ortho(void* self, void* rect) {
	((QMatrix4x4*)self)->ortho(*(QRect*)rect);
}
void QMatrix4x4_ortho2(void* self, void* rect) {
	((QMatrix4x4*)self)->ortho(*(QRectF*)rect);
}
void QMatrix4x4_ortho3(void* self, float left, float right, float bottom, float top, float nearPlane, float farPlane) {
	((QMatrix4x4*)self)->ortho(left, right, bottom, top, nearPlane, farPlane);
}
void QMatrix4x4_frustum(void* self, float left, float right, float bottom, float top, float nearPlane, float farPlane) {
	((QMatrix4x4*)self)->frustum(left, right, bottom, top, nearPlane, farPlane);
}
void QMatrix4x4_perspective(void* self, float verticalAngle, float aspectRatio, float nearPlane, float farPlane) {
	((QMatrix4x4*)self)->perspective(verticalAngle, aspectRatio, nearPlane, farPlane);
}
void QMatrix4x4_lookAt(void* self, void* eye, void* center, void* up) {
	((QMatrix4x4*)self)->lookAt(*(QVector3D*)eye, *(QVector3D*)center, *(QVector3D*)up);
}
void QMatrix4x4_viewport(void* self, void* rect) {
	((QMatrix4x4*)self)->viewport(*(QRectF*)rect);
}
void QMatrix4x4_viewport2(void* self, float left, float bottom, float width, float height) {
	((QMatrix4x4*)self)->viewport(left, bottom, width, height);
}
void QMatrix4x4_flipCoordinates(void* self) {
	((QMatrix4x4*)self)->flipCoordinates();
}
void QMatrix4x4_copyDataTo(void* self, float* values) {
	float values_local{};
	((QMatrix4x4*)self)->copyDataTo(&values_local);
	*values = values_local;
}
void* QMatrix4x4_toTransform(void* self) {
	auto&& __r = ((QMatrix4x4*)self)->toTransform();
	return new QTransform(__r);
}
void* QMatrix4x4_toTransform2(void* self, float distanceToPlane) {
	auto&& __r = ((QMatrix4x4*)self)->toTransform(distanceToPlane);
	return new QTransform(__r);
}
void* QMatrix4x4_map(void* self, void* point) {
	auto&& __r = ((QMatrix4x4*)self)->map(*(QPoint*)point);
	return new QPoint(__r);
}
void* QMatrix4x4_map2(void* self, void* point) {
	auto&& __r = ((QMatrix4x4*)self)->map(*(QPointF*)point);
	return new QPointF(__r);
}
void* QMatrix4x4_map3(void* self, void* point) {
	auto&& __r = ((QMatrix4x4*)self)->map(*(QVector3D*)point);
	return new QVector3D(__r);
}
void* QMatrix4x4_mapVector(void* self, void* vector) {
	auto&& __r = ((QMatrix4x4*)self)->mapVector(*(QVector3D*)vector);
	return new QVector3D(__r);
}
void* QMatrix4x4_map4(void* self, void* point) {
	auto&& __r = ((QMatrix4x4*)self)->map(*(QVector4D*)point);
	return new QVector4D(__r);
}
void* QMatrix4x4_mapRect(void* self, void* rect) {
	auto&& __r = ((QMatrix4x4*)self)->mapRect(*(QRect*)rect);
	return new QRect(__r);
}
void* QMatrix4x4_mapRect2(void* self, void* rect) {
	auto&& __r = ((QMatrix4x4*)self)->mapRect(*(QRectF*)rect);
	return new QRectF(__r);
}
void QMatrix4x4_optimize(void* self) {
	((QMatrix4x4*)self)->optimize();
}
void* QMatrix4x4_ToQVariant(void* self) {
	auto&& __r = static_cast<QVariant>(*(QMatrix4x4*)self);
	return new QVariant(__r);
}
void QMatrix4x4_projectedRotate(void* self, float angle, float x, float y, float z, float distanceToPlane) {
	((QMatrix4x4*)self)->projectedRotate(angle, x, y, z, distanceToPlane);
}
void QMatrix4x4_projectedRotate2(void* self, float angle, float x, float y, float z) {
	((QMatrix4x4*)self)->projectedRotate(angle, x, y, z);
}
int QMatrix4x4_flags(void* self) {
	auto&& __r = ((QMatrix4x4*)self)->flags();
	return (int)(__r);
}
void* QMatrix4x4_inverted2(void* self, int* invertible) {
	bool invertible_local{};
	auto&& __r = ((QMatrix4x4*)self)->inverted(&invertible_local);
	*invertible = invertible_local;
	return new QMatrix4x4(__r);
}
void QMatrix4x4_rotate4(void* self, float angle, float x, float y, float z) {
	((QMatrix4x4*)self)->rotate(angle, x, y, z);
}
void QMatrix4x4_viewport3(void* self, float left, float bottom, float width, float height, float nearPlane) {
	((QMatrix4x4*)self)->viewport(left, bottom, width, height, nearPlane);
}
void QMatrix4x4_viewport4(void* self, float left, float bottom, float width, float height, float nearPlane, float farPlane) {
	((QMatrix4x4*)self)->viewport(left, bottom, width, height, nearPlane, farPlane);
}
void* QMatrix4x4_OperatorMultiply(void* self, void* vector) {
	auto&& __r = (*(QMatrix4x4*)self) * (*(QVector3D*)vector);
	return new QVector3D(__r);
}
void* QMatrix4x4_OperatorMultiply2(void* self, void* vector) {
	auto&& __r = (*(QMatrix4x4*)self) * (*(QVector4D*)vector);
	return new QVector4D(__r);
}
void* QMatrix4x4_OperatorPlus(void* self, void* m2) {
	auto&& __r = (*(QMatrix4x4*)self) + (*(QMatrix4x4*)m2);
	return new QMatrix4x4(__r);
}
void* QMatrix4x4_OperatorMinus(void* self, void* m2) {
	auto&& __r = (*(QMatrix4x4*)self) - (*(QMatrix4x4*)m2);
	return new QMatrix4x4(__r);
}
void* QMatrix4x4_OperatorMultiply3(void* self, void* m2) {
	auto&& __r = (*(QMatrix4x4*)self) * (*(QMatrix4x4*)m2);
	return new QMatrix4x4(__r);
}
void* QMatrix4x4_OperatorMinus2(void* self) {
	auto&& __r = -(*(QMatrix4x4*)self);
	return new QMatrix4x4(__r);
}
void* QMatrix4x4_OperatorMultiply4(void* self, void* point) {
	auto&& __r = (*(QMatrix4x4*)self) * (*(QPoint*)point);
	return new QPoint(__r);
}
void* QMatrix4x4_OperatorMultiply5(void* self, void* point) {
	auto&& __r = (*(QMatrix4x4*)self) * (*(QPointF*)point);
	return new QPointF(__r);
}
void* QMatrix4x4_OperatorMultiply6(void* self, float factor) {
	auto&& __r = (*(QMatrix4x4*)self) * (factor);
	return new QMatrix4x4(__r);
}
void* QMatrix4x4_OperatorDivide(void* self, float divisor) {
	auto&& __r = (*(QMatrix4x4*)self) / (divisor);
	return new QMatrix4x4(__r);
}
void QMatrix4x4_delete(void* self) { delete (QMatrix4x4*)self; }

}

#include <QtGui/qpicture.h>
#include <QtGui/qpaintdevice.h>
#include <QtGui/qpainter.h>
#include <QtCore/qiodevice.h>
#include <QtCore/qstring.h>
#include <QtCore/qrect.h>
#include <QtGui/qpaintengine.h>
#include <QtCore/qpoint.h>
#include "gen_qpicture.h"
#include "gen_purego_cb.h"

class QPicture__Go : public QPicture {
public:
	long long __goId;
	unsigned char __goOn[6] = {};
	template<typename... A> QPicture__Go(long long id, A&&... a) : QPicture(std::forward<A>(a)...), __goId(id) {}
	using QPicture::getDecodedMetricF;
	int cppgoBase_metric(int __a0) { return QPicture::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QPicture::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QPicture::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QPicture::sharedPainter(); }
	int devType() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QPicture::devType();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QPicture::paintEngine();
	}
	int metric(QPaintDevice::PaintDeviceMetric m) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(m), &__h); if (__h) return __r; }
		return QPicture::metric(m);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QPicture::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QPicture::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QPicture::sharedPainter();
	}
};

extern "C" {

void* QPicture_base(void* self) { return (void*)static_cast<QPaintDevice*>((QPicture*)self); }
void QPicture__set_override(void* self, int idx) { ((QPicture__Go*)self)->__goOn[idx] = 1; }
void* QPicture_new(long long __goid) {
	return new QPicture__Go(__goid);
}
void* QPicture_new2(long long __goid, void* arg1) {
	return new QPicture__Go(__goid, *(QPicture*)arg1);
}
void* QPicture_new3(long long __goid, int formatVersion) {
	return new QPicture__Go(__goid, formatVersion);
}
int QPicture_isNull(void* self) {
	auto&& __r = ((QPicture*)self)->isNull();
	return __r;
}
unsigned int QPicture_size(void* self) {
	auto&& __r = ((QPicture*)self)->size();
	return __r;
}
const char* QPicture_data(void* self) {
	auto&& __r = ((QPicture*)self)->data();
	return __r;
}
void QPicture_setData(void* self, const char* data, unsigned int size) {
	((QPicture*)self)->setData(data, size);
}
int QPicture_play(void* self, void* p) {
	auto&& __r = ((QPicture*)self)->play((QPainter*)p);
	return __r;
}
int QPicture_load(void* self, void* dev) {
	auto&& __r = ((QPicture*)self)->load((QIODevice*)dev);
	return __r;
}
int QPicture_load2(void* self, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QPicture*)self)->load(fileName_qstr);
	return __r;
}
int QPicture_save(void* self, void* dev) {
	auto&& __r = ((QPicture*)self)->save((QIODevice*)dev);
	return __r;
}
int QPicture_save2(void* self, const char* fileName, long long fileName_len) {
	QString fileName_qstr = QString::fromUtf8(fileName, fileName_len);
	auto&& __r = ((QPicture*)self)->save(fileName_qstr);
	return __r;
}
void* QPicture_boundingRect(void* self) {
	auto&& __r = ((QPicture*)self)->boundingRect();
	return new QRect(__r);
}
void QPicture_setBoundingRect(void* self, void* r) {
	((QPicture*)self)->setBoundingRect(*(QRect*)r);
}
void* QPicture_OperatorAssign(void* self, void* p) {
	auto&& __r = ((QPicture*)self)->operator=(*(QPicture*)p);
	return (void*)&(__r);
}
void QPicture_swap(void* self, void* other) {
	((QPicture*)self)->swap(*(QPicture*)other);
}
void QPicture_detach(void* self) {
	((QPicture*)self)->detach();
}
int QPicture_isDetached(void* self) {
	auto&& __r = ((QPicture*)self)->isDetached();
	return __r;
}
double QPicture_devicePixelRatioFScale() {
	auto&& __r = QPicture::devicePixelRatioFScale();
	return __r;
}
int QPicture_encodeMetricF(int metric, double value) {
	auto&& __r = QPicture::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QPicture_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QPicture__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
int QPicture_devType_default(void* self) {
	auto&& __r = ((QPicture*)self)->QPicture::devType();
	return __r;
}
void* QPicture_paintEngine_default(void* self) {
	auto&& __r = ((QPicture*)self)->QPicture::paintEngine();
	return (void*)(__r);
}
int QPicture_metric_default(void* self, int m) {
	auto&& __r = ((QPicture__Go*)self)->cppgoBase_metric(m);
	return __r;
}
void QPicture_initPainter_default(void* self, void* painter) {
	((QPicture__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QPicture_redirected_default(void* self, void* offset) {
	auto&& __r = ((QPicture__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QPicture_sharedPainter_default(void* self) {
	auto&& __r = ((QPicture__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void* QPicture__Go_new(long long id) { return new QPicture__Go(id); }
void QPicture_delete(void* self) { delete (QPicture*)self; }

}

#include <QtGui/qtexttable.h>
#include <QtGui/qtextformat.h>
#include <QtGui/qtextcursor.h>
#include "gen_qtexttablecell.h"

extern "C" {

void* QTextTableCell_new() {
	return new QTextTableCell();
}
void* QTextTableCell_new2(void* o) {
	return new QTextTableCell(*(QTextTableCell*)o);
}
void* QTextTableCell_OperatorAssign(void* self, void* o) {
	auto&& __r = ((QTextTableCell*)self)->operator=(*(QTextTableCell*)o);
	return (void*)&(__r);
}
void QTextTableCell_setFormat(void* self, void* format) {
	((QTextTableCell*)self)->setFormat(*(QTextCharFormat*)format);
}
void* QTextTableCell_format(void* self) {
	auto&& __r = ((QTextTableCell*)self)->format();
	return new QTextCharFormat(__r);
}
int QTextTableCell_row(void* self) {
	auto&& __r = ((QTextTableCell*)self)->row();
	return __r;
}
int QTextTableCell_column(void* self) {
	auto&& __r = ((QTextTableCell*)self)->column();
	return __r;
}
int QTextTableCell_rowSpan(void* self) {
	auto&& __r = ((QTextTableCell*)self)->rowSpan();
	return __r;
}
int QTextTableCell_columnSpan(void* self) {
	auto&& __r = ((QTextTableCell*)self)->columnSpan();
	return __r;
}
int QTextTableCell_isValid(void* self) {
	auto&& __r = ((QTextTableCell*)self)->isValid();
	return __r;
}
void* QTextTableCell_firstCursorPosition(void* self) {
	auto&& __r = ((QTextTableCell*)self)->firstCursorPosition();
	return new QTextCursor(__r);
}
void* QTextTableCell_lastCursorPosition(void* self) {
	auto&& __r = ((QTextTableCell*)self)->lastCursorPosition();
	return new QTextCursor(__r);
}
int QTextTableCell_firstPosition(void* self) {
	auto&& __r = ((QTextTableCell*)self)->firstPosition();
	return __r;
}
int QTextTableCell_lastPosition(void* self) {
	auto&& __r = ((QTextTableCell*)self)->lastPosition();
	return __r;
}
int QTextTableCell_OperatorEqual(void* self, void* other) {
	auto&& __r = ((QTextTableCell*)self)->operator==(*(QTextTableCell*)other);
	return __r;
}
int QTextTableCell_OperatorNotEqual(void* self, void* other) {
	auto&& __r = ((QTextTableCell*)self)->operator!=(*(QTextTableCell*)other);
	return __r;
}
int QTextTableCell_tableCellFormatIndex(void* self) {
	auto&& __r = ((QTextTableCell*)self)->tableCellFormatIndex();
	return __r;
}
void QTextTableCell_delete(void* self) { delete (QTextTableCell*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qwidget.h>
#include <QtWidgets/qframe.h>
#include "gen_qstyleoptionframe.h"

extern "C" {

void* QStyleOptionFrame_base(void* self) { return (void*)static_cast<QStyleOption*>((QStyleOptionFrame*)self); }
void* QStyleOptionFrame_new() {
	return new QStyleOptionFrame();
}
void* QStyleOptionFrame_new2(void* other) {
	return new QStyleOptionFrame(*(QStyleOptionFrame*)other);
}
void* QStyleOptionFrame_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleOptionFrame*)self)->operator=(*(QStyleOptionFrame*)arg1);
	return (void*)&(__r);
}
int QStyleOptionFrame_lineWidth_get(void* self) {
	auto __r = ((QStyleOptionFrame*)self)->lineWidth;
	return __r;
}
void QStyleOptionFrame_lineWidth_set(void* self, int lineWidth) {
	((QStyleOptionFrame*)self)->lineWidth = lineWidth;
}
int QStyleOptionFrame_midLineWidth_get(void* self) {
	auto __r = ((QStyleOptionFrame*)self)->midLineWidth;
	return __r;
}
void QStyleOptionFrame_midLineWidth_set(void* self, int midLineWidth) {
	((QStyleOptionFrame*)self)->midLineWidth = midLineWidth;
}
int QStyleOptionFrame_features_get(void* self) {
	auto __r = ((QStyleOptionFrame*)self)->features;
	return (int)(__r);
}
void QStyleOptionFrame_features_set(void* self, int features) {
	((QStyleOptionFrame*)self)->features = QFlags<QStyleOptionFrame::FrameFeature>(features);
}
int QStyleOptionFrame_frameShape_get(void* self) {
	auto __r = ((QStyleOptionFrame*)self)->frameShape;
	return (int)(__r);
}
void QStyleOptionFrame_frameShape_set(void* self, int frameShape) {
	((QStyleOptionFrame*)self)->frameShape = (QFrame::Shape)(frameShape);
}
void QStyleOptionFrame_delete(void* self) { delete (QStyleOptionFrame*)self; }

}

#include <QtWidgets/qstyleoption.h>
#include "gen_qstylehintreturn.h"

extern "C" {

void* QStyleHintReturn_new() {
	return new QStyleHintReturn();
}
void* QStyleHintReturn_new2(void* arg1) {
	return new QStyleHintReturn(*(QStyleHintReturn*)arg1);
}
void* QStyleHintReturn_new3(int version) {
	return new QStyleHintReturn(version);
}
void* QStyleHintReturn_new4(int version, int type_) {
	return new QStyleHintReturn(version, type_);
}
void* QStyleHintReturn_OperatorAssign(void* self, void* arg1) {
	auto&& __r = ((QStyleHintReturn*)self)->operator=(*(QStyleHintReturn*)arg1);
	return (void*)&(__r);
}
int QStyleHintReturn_version_get(void* self) {
	auto __r = ((QStyleHintReturn*)self)->version;
	return __r;
}
void QStyleHintReturn_version_set(void* self, int version) {
	((QStyleHintReturn*)self)->version = version;
}
int QStyleHintReturn_type_get(void* self) {
	auto __r = ((QStyleHintReturn*)self)->type;
	return __r;
}
void QStyleHintReturn_type_set(void* self, int type_) {
	((QStyleHintReturn*)self)->type = type_;
}
void QStyleHintReturn_delete(void* self) { delete (QStyleHintReturn*)self; }

}

#include <QtWidgets/qcolumnview.h>
#include <QtWidgets/qabstractitemview.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qpoint.h>
#include <QtCore/qsize.h>
#include <QtCore/qrect.h>
#include <QtCore/qitemselectionmodel.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qevent.h>
#include <QtGui/qregion.h>
#include <QtWidgets/qabstractitemdelegate.h>
#include <QtCore/qvariant.h>
#include <QtCore/qobject.h>
#include <QtCore/qcoreevent.h>
#include <QtWidgets/qstyleoption.h>
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
#include "gen_qcolumnview.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QColumnView__Go : public QColumnView {
public:
	long long __goId;
	unsigned char __goOn[92] = {};
	template<typename... A> QColumnView__Go(long long id, A&&... a) : QColumnView(std::forward<A>(a)...), __goId(id) {}
	using QColumnView::initializeColumn;
	using QColumnView::state;
	using QColumnView::scheduleDelayedItemsLayout;
	using QColumnView::executeDelayedItemsLayout;
	using QColumnView::setDirtyRegion;
	using QColumnView::scrollDirtyRegion;
	using QColumnView::dirtyRegionOffset;
	using QColumnView::startAutoScroll;
	using QColumnView::stopAutoScroll;
	using QColumnView::doAutoScroll;
	using QColumnView::dropIndicatorPosition;
	using QColumnView::setViewportMargins;
	using QColumnView::viewportMargins;
	using QColumnView::drawFrame;
	using QColumnView::updateMicroFocus;
	using QColumnView::create;
	using QColumnView::destroy;
	using QColumnView::focusNextChild;
	using QColumnView::focusPreviousChild;
	using QColumnView::sender;
	using QColumnView::senderSignalIndex;
	using QColumnView::receivers;
	using QColumnView::isSignalConnected;
	using QColumnView::getDecodedMetricF;
	bool cppgoBase_isIndexHidden(const QModelIndex& __a0) { return QColumnView::isIndexHidden(__a0); }
	QModelIndex cppgoBase_moveCursor(int __a0, QFlags<Qt::KeyboardModifier> __a1) { return QColumnView::moveCursor((QAbstractItemView::CursorAction)__a0, __a1); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QColumnView::resizeEvent(__a0); }
	void cppgoBase_setSelection(const QRect& __a0, QFlags<QItemSelectionModel::SelectionFlag> __a1) { QColumnView::setSelection(__a0, __a1); }
	QRegion cppgoBase_visualRegionForSelection(const QItemSelection& __a0) { return QColumnView::visualRegionForSelection(__a0); }
	int cppgoBase_horizontalOffset() { return QColumnView::horizontalOffset(); }
	int cppgoBase_verticalOffset() { return QColumnView::verticalOffset(); }
	void cppgoBase_rowsInserted(const QModelIndex& __a0, int __a1, int __a2) { QColumnView::rowsInserted(__a0, __a1, __a2); }
	void cppgoBase_currentChanged(const QModelIndex& __a0, const QModelIndex& __a1) { QColumnView::currentChanged(__a0, __a1); }
	void cppgoBase_scrollContentsBy(int __a0, int __a1) { QColumnView::scrollContentsBy(__a0, __a1); }
	QAbstractItemView* cppgoBase_createColumn(const QModelIndex& __a0) { return QColumnView::createColumn(__a0); }
	void cppgoBase_rowsAboutToBeRemoved(const QModelIndex& __a0, int __a1, int __a2) { QColumnView::rowsAboutToBeRemoved(__a0, __a1, __a2); }
	void cppgoBase_selectionChanged(const QItemSelection& __a0, const QItemSelection& __a1) { QColumnView::selectionChanged(__a0, __a1); }
	void cppgoBase_updateEditorData() { QColumnView::updateEditorData(); }
	void cppgoBase_updateEditorGeometries() { QColumnView::updateEditorGeometries(); }
	void cppgoBase_updateGeometries() { QColumnView::updateGeometries(); }
	void cppgoBase_verticalScrollbarAction(int __a0) { QColumnView::verticalScrollbarAction(__a0); }
	void cppgoBase_horizontalScrollbarAction(int __a0) { QColumnView::horizontalScrollbarAction(__a0); }
	void cppgoBase_verticalScrollbarValueChanged(int __a0) { QColumnView::verticalScrollbarValueChanged(__a0); }
	void cppgoBase_horizontalScrollbarValueChanged(int __a0) { QColumnView::horizontalScrollbarValueChanged(__a0); }
	void cppgoBase_closeEditor(QWidget* __a0, int __a1) { QColumnView::closeEditor(__a0, (QAbstractItemDelegate::EndEditHint)__a1); }
	void cppgoBase_commitData(QWidget* __a0) { QColumnView::commitData(__a0); }
	void cppgoBase_editorDestroyed(QObject* __a0) { QColumnView::editorDestroyed(__a0); }
	QList<QModelIndex> cppgoBase_selectedIndexes() { return QColumnView::selectedIndexes(); }
	bool cppgoBase_edit(const QModelIndex& __a0, int __a1, QEvent* __a2) { return QColumnView::edit(__a0, (QAbstractItemView::EditTrigger)__a1, __a2); }
	QFlags<QItemSelectionModel::SelectionFlag> cppgoBase_selectionCommand(const QModelIndex& __a0, const QEvent* __a1) { return QColumnView::selectionCommand(__a0, __a1); }
	void cppgoBase_startDrag(QFlags<Qt::DropAction> __a0) { QColumnView::startDrag(__a0); }
	void cppgoBase_initViewItemOption(QStyleOptionViewItem* __a0) { QColumnView::initViewItemOption(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QColumnView::focusNextPrevChild(__a0); }
	bool cppgoBase_event(QEvent* __a0) { return QColumnView::event(__a0); }
	bool cppgoBase_viewportEvent(QEvent* __a0) { return QColumnView::viewportEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QColumnView::mousePressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QColumnView::mouseMoveEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QColumnView::mouseReleaseEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QColumnView::mouseDoubleClickEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QColumnView::dragEnterEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QColumnView::dragMoveEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QColumnView::dragLeaveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QColumnView::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QColumnView::focusInEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QColumnView::focusOutEvent(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QColumnView::keyPressEvent(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QColumnView::timerEvent(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QColumnView::inputMethodEvent(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QColumnView::eventFilter(__a0, __a1); }
	QSize cppgoBase_viewportSizeHint() { return QColumnView::viewportSizeHint(); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QColumnView::paintEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QColumnView::wheelEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QColumnView::contextMenuEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QColumnView::changeEvent(__a0); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QColumnView::initStyleOption(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QColumnView::keyReleaseEvent(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QColumnView::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QColumnView::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QColumnView::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QColumnView::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QColumnView::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QColumnView::actionEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QColumnView::showEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QColumnView::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QColumnView::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QColumnView::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QColumnView::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QColumnView::sharedPainter(); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QColumnView::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QColumnView::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QColumnView::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QColumnView::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QColumnView::metaObject();
	}
	QModelIndex indexAt(const QPoint& point) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -7252823847285716802LL, (void*)&point, &__h); if (__h) return *(QModelIndex*)__r; }
		return QColumnView::indexAt(point);
	}
	void scrollTo(const QModelIndex& index, QAbstractItemView::ScrollHint hint) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -4832340528568269947LL, (void*)&index, (int)(hint), &__h); if (__h) return; }
		QColumnView::scrollTo(index, hint);
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QColumnView::sizeHint();
	}
	QRect visualRect(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -1267696851865372997LL, (void*)&index, &__h); if (__h) return *(QRect*)__r; }
		return QColumnView::visualRect(index);
	}
	void setModel(QAbstractItemModel* model) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4283922507359547360LL, (void*)model, &__h); if (__h) return; }
		QColumnView::setModel(model);
	}
	void setSelectionModel(QItemSelectionModel* selectionModel) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4689322347236366780LL, (void*)selectionModel, &__h); if (__h) return; }
		QColumnView::setSelectionModel(selectionModel);
	}
	void setRootIndex(const QModelIndex& index) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1574077111192992763LL, (void*)&index, &__h); if (__h) return; }
		QColumnView::setRootIndex(index);
	}
	void selectAll() override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 8758113809152574300LL, &__h); if (__h) return; }
		QColumnView::selectAll();
	}
	bool isIndexHidden(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 6057989667089378859LL, (void*)&index, &__h); if (__h) return __r; }
		return QColumnView::isIndexHidden(index);
	}
	QModelIndex moveCursor(QAbstractItemView::CursorAction cursorAction, QFlags<Qt::KeyboardModifier> modifiers) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_voidp_int_int) { void* __r = ((void*(*)(long long, long long, int, int, int*))cppgo_virt_voidp_int_int)(__goId, -158523882194220890LL, (int)(cursorAction), (int)(modifiers), &__h); if (__h) return *(QModelIndex*)__r; }
		return QColumnView::moveCursor(cursorAction, modifiers);
	}
	void resizeEvent(QResizeEvent* event) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)event, &__h); if (__h) return; }
		QColumnView::resizeEvent(event);
	}
	void setSelection(const QRect& rect, QFlags<QItemSelectionModel::SelectionFlag> command) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, -1539266838375037305LL, (void*)&rect, (int)(command), &__h); if (__h) return; }
		QColumnView::setSelection(rect, command);
	}
	QRegion visualRegionForSelection(const QItemSelection& selection) const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 7063795720810552074LL, (void*)&selection, &__h); if (__h) return *(QRegion*)__r; }
		return QColumnView::visualRegionForSelection(selection);
	}
	int horizontalOffset() const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -1500310024801917802LL, &__h); if (__h) return __r; }
		return QColumnView::horizontalOffset();
	}
	int verticalOffset() const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, 5218707810677921016LL, &__h); if (__h) return __r; }
		return QColumnView::verticalOffset();
	}
	void rowsInserted(const QModelIndex& parent, int start, int end) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, 4472434241583495620LL, (void*)&parent, start, end, &__h); if (__h) return; }
		QColumnView::rowsInserted(parent, start, end);
	}
	void currentChanged(const QModelIndex& current, const QModelIndex& previous) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -6978583972112316034LL, (void*)&current, (void*)&previous, &__h); if (__h) return; }
		QColumnView::currentChanged(current, previous);
	}
	void scrollContentsBy(int dx, int dy) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, -6118001150126564745LL, dx, dy, &__h); if (__h) return; }
		QColumnView::scrollContentsBy(dx, dy);
	}
	QAbstractItemView* createColumn(const QModelIndex& rootIndex) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -5176237958185448161LL, (void*)&rootIndex, &__h); if (__h) return (QAbstractItemView*)__r; }
		return QColumnView::createColumn(rootIndex);
	}
	void keyboardSearch(const QString& search) override {
		int __h = 0;
		QByteArray search_ba = search.toUtf8();
		if (__goOn[20] && cppgo_virt_void_constcharp_longlong) { ((void(*)(long long, long long, const char*, long long, int*))cppgo_virt_void_constcharp_longlong)(__goId, -5041452885022114200LL, (char*)search_ba.constData(), search_ba.size(), &__h); if (__h) return; }
		QColumnView::keyboardSearch(search);
	}
	int sizeHintForRow(int row) const override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -4084029023708556876LL, row, &__h); if (__h) return __r; }
		return QColumnView::sizeHintForRow(row);
	}
	int sizeHintForColumn(int column) const override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, -3666792239577137638LL, column, &__h); if (__h) return __r; }
		return QColumnView::sizeHintForColumn(column);
	}
	QAbstractItemDelegate* itemDelegateForIndex(const QModelIndex& index) const override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, -3811438228839363824LL, (void*)&index, &__h); if (__h) return (QAbstractItemDelegate*)__r; }
		return QColumnView::itemDelegateForIndex(index);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery query) const override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(query), &__h); if (__h) return *(QVariant*)__r; }
		return QColumnView::inputMethodQuery(query);
	}
	void reset() override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1086335023529244512LL, &__h); if (__h) return; }
		QColumnView::reset();
	}
	void doItemsLayout() override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1888629738259754754LL, &__h); if (__h) return; }
		QColumnView::doItemsLayout();
	}
	void rowsAboutToBeRemoved(const QModelIndex& parent, int start, int end) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp_int_int) { ((void(*)(long long, long long, void*, int, int, int*))cppgo_virt_void_voidp_int_int)(__goId, -8151049665126058157LL, (void*)&parent, start, end, &__h); if (__h) return; }
		QColumnView::rowsAboutToBeRemoved(parent, start, end);
	}
	void selectionChanged(const QItemSelection& selected, const QItemSelection& deselected) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp_voidp) { ((void(*)(long long, long long, void*, void*, int*))cppgo_virt_void_voidp_voidp)(__goId, -5892205462145366901LL, (void*)&selected, (void*)&deselected, &__h); if (__h) return; }
		QColumnView::selectionChanged(selected, deselected);
	}
	void updateEditorData() override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 9029941737538200621LL, &__h); if (__h) return; }
		QColumnView::updateEditorData();
	}
	void updateEditorGeometries() override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -8726300139121449691LL, &__h); if (__h) return; }
		QColumnView::updateEditorGeometries();
	}
	void updateGeometries() override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 1022771470898460576LL, &__h); if (__h) return; }
		QColumnView::updateGeometries();
	}
	void verticalScrollbarAction(int action) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -1155555122391250623LL, action, &__h); if (__h) return; }
		QColumnView::verticalScrollbarAction(action);
	}
	void horizontalScrollbarAction(int action) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -785551405282638793LL, action, &__h); if (__h) return; }
		QColumnView::horizontalScrollbarAction(action);
	}
	void verticalScrollbarValueChanged(int value) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 6366953527806415524LL, value, &__h); if (__h) return; }
		QColumnView::verticalScrollbarValueChanged(value);
	}
	void horizontalScrollbarValueChanged(int value) override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 1786430177580376614LL, value, &__h); if (__h) return; }
		QColumnView::horizontalScrollbarValueChanged(value);
	}
	void closeEditor(QWidget* editor, QAbstractItemDelegate::EndEditHint hint) override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, 8648484264628816806LL, (void*)editor, (int)(hint), &__h); if (__h) return; }
		QColumnView::closeEditor(editor, hint);
	}
	void commitData(QWidget* editor) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6021383853323633948LL, (void*)editor, &__h); if (__h) return; }
		QColumnView::commitData(editor);
	}
	void editorDestroyed(QObject* editor) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6979418566323745651LL, (void*)editor, &__h); if (__h) return; }
		QColumnView::editorDestroyed(editor);
	}
	QList<QModelIndex> selectedIndexes() const override {
		int __h = 0;
		if (__goOn[39] && cppgo_virtp_voidp) {
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
		return QColumnView::selectedIndexes();
	}
	bool edit(const QModelIndex& index, QAbstractItemView::EditTrigger trigger, QEvent* event) override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_int_voidp_int_voidp) { int __r = ((int(*)(long long, long long, void*, int, void*, int*))cppgo_virt_int_voidp_int_voidp)(__goId, 4206689019978932209LL, (void*)&index, (int)(trigger), (void*)event, &__h); if (__h) return __r; }
		return QColumnView::edit(index, trigger, event);
	}
	QFlags<QItemSelectionModel::SelectionFlag> selectionCommand(const QModelIndex& index, const QEvent* event) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, -2447490284072398302LL, (void*)&index, (void*)event, &__h); if (__h) return QFlags<QItemSelectionModel::SelectionFlag>(__r); }
		return QColumnView::selectionCommand(index, event);
	}
	void startDrag(QFlags<Qt::DropAction> supportedActions) override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -8112366604639385491LL, (int)(supportedActions), &__h); if (__h) return; }
		QColumnView::startDrag(supportedActions);
	}
	void initViewItemOption(QStyleOptionViewItem* option) const override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -7450787277202094556LL, (void*)option, &__h); if (__h) return; }
		QColumnView::initViewItemOption(option);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QColumnView::focusNextPrevChild(next);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QColumnView::event(event);
	}
	bool viewportEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -1971465653061861519LL, (void*)event, &__h); if (__h) return __r; }
		return QColumnView::viewportEvent(event);
	}
	void mousePressEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)event, &__h); if (__h) return; }
		QColumnView::mousePressEvent(event);
	}
	void mouseMoveEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)event, &__h); if (__h) return; }
		QColumnView::mouseMoveEvent(event);
	}
	void mouseReleaseEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)event, &__h); if (__h) return; }
		QColumnView::mouseReleaseEvent(event);
	}
	void mouseDoubleClickEvent(QMouseEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)event, &__h); if (__h) return; }
		QColumnView::mouseDoubleClickEvent(event);
	}
	void dragEnterEvent(QDragEnterEvent* event) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)event, &__h); if (__h) return; }
		QColumnView::dragEnterEvent(event);
	}
	void dragMoveEvent(QDragMoveEvent* event) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)event, &__h); if (__h) return; }
		QColumnView::dragMoveEvent(event);
	}
	void dragLeaveEvent(QDragLeaveEvent* event) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)event, &__h); if (__h) return; }
		QColumnView::dragLeaveEvent(event);
	}
	void dropEvent(QDropEvent* event) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)event, &__h); if (__h) return; }
		QColumnView::dropEvent(event);
	}
	void focusInEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)event, &__h); if (__h) return; }
		QColumnView::focusInEvent(event);
	}
	void focusOutEvent(QFocusEvent* event) override {
		int __h = 0;
		if (__goOn[56] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)event, &__h); if (__h) return; }
		QColumnView::focusOutEvent(event);
	}
	void keyPressEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[57] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)event, &__h); if (__h) return; }
		QColumnView::keyPressEvent(event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[58] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QColumnView::timerEvent(event);
	}
	void inputMethodEvent(QInputMethodEvent* event) override {
		int __h = 0;
		if (__goOn[59] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)event, &__h); if (__h) return; }
		QColumnView::inputMethodEvent(event);
	}
	bool eventFilter(QObject* object, QEvent* event) override {
		int __h = 0;
		if (__goOn[60] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)object, (void*)event, &__h); if (__h) return __r; }
		return QColumnView::eventFilter(object, event);
	}
	QSize viewportSizeHint() const override {
		int __h = 0;
		if (__goOn[61] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 503579067252048239LL, &__h); if (__h) return *(QSize*)__r; }
		return QColumnView::viewportSizeHint();
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[62] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QColumnView::minimumSizeHint();
	}
	void setupViewport(QWidget* viewport) override {
		int __h = 0;
		if (__goOn[63] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4531996381736730688LL, (void*)viewport, &__h); if (__h) return; }
		QColumnView::setupViewport(viewport);
	}
	void paintEvent(QPaintEvent* arg1) override {
		int __h = 0;
		if (__goOn[64] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)arg1, &__h); if (__h) return; }
		QColumnView::paintEvent(arg1);
	}
	void wheelEvent(QWheelEvent* arg1) override {
		int __h = 0;
		if (__goOn[65] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)arg1, &__h); if (__h) return; }
		QColumnView::wheelEvent(arg1);
	}
	void contextMenuEvent(QContextMenuEvent* arg1) override {
		int __h = 0;
		if (__goOn[66] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)arg1, &__h); if (__h) return; }
		QColumnView::contextMenuEvent(arg1);
	}
	void changeEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[67] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)arg1, &__h); if (__h) return; }
		QColumnView::changeEvent(arg1);
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[68] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QColumnView::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[69] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QColumnView::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[70] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QColumnView::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[71] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QColumnView::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[72] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QColumnView::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[73] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QColumnView::paintEngine();
	}
	void keyReleaseEvent(QKeyEvent* event) override {
		int __h = 0;
		if (__goOn[74] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)event, &__h); if (__h) return; }
		QColumnView::keyReleaseEvent(event);
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[75] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QColumnView::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[76] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QColumnView::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[77] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QColumnView::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[78] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QColumnView::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[79] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QColumnView::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[80] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QColumnView::actionEvent(event);
	}
	void showEvent(QShowEvent* event) override {
		int __h = 0;
		if (__goOn[81] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)event, &__h); if (__h) return; }
		QColumnView::showEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[82] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QColumnView::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[83] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QColumnView::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[84] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QColumnView::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[85] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QColumnView::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[86] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QColumnView::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[87] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QColumnView::sharedPainter();
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[88] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QColumnView::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[89] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QColumnView::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[90] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QColumnView::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[91] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QColumnView::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QColumnView;
static std::mutex __connsMu_QColumnView;

extern "C" {

void* QColumnView_base(void* self) { return (void*)static_cast<QAbstractItemView*>((QColumnView*)self); }
void QColumnView_chain(void* self, void** out) {
	QColumnView* __p = (QColumnView*)self;
	QAbstractItemView* __b0 = static_cast<QAbstractItemView*>(__p); out[0] = (void*)__b0;
	QAbstractScrollArea* __b1 = static_cast<QAbstractScrollArea*>(__b0); out[1] = (void*)__b1;
	QFrame* __b2 = static_cast<QFrame*>(__b1); out[2] = (void*)__b2;
	QWidget* __b3 = static_cast<QWidget*>(__b2); out[3] = (void*)__b3;
	QObject* __b4 = static_cast<QObject*>(__b3); out[4] = (void*)__b4;
}
void QColumnView__set_override(void* self, int idx) { ((QColumnView__Go*)self)->__goOn[idx] = 1; }
void* QColumnView_fromQObject(void* o) { return (void*)qobject_cast<QColumnView*>((QObject*)o); }
void* QColumnView_new(long long __goid, void* parent) {
	return new QColumnView__Go(__goid, (QWidget*)parent);
}
void* QColumnView_new2(long long __goid) {
	return new QColumnView__Go(__goid);
}
char* QColumnView_tr(const char* s, long long* __out_len) {
	auto&& __r = QColumnView::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QColumnView_updatePreviewWidget(void* self, void* index) {
	((QColumnView*)self)->updatePreviewWidget(*(QModelIndex*)index);
}
void QColumnView_setResizeGripsVisible(void* self, int visible) {
	((QColumnView*)self)->setResizeGripsVisible(visible);
}
int QColumnView_resizeGripsVisible(void* self) {
	auto&& __r = ((QColumnView*)self)->resizeGripsVisible();
	return __r;
}
void QColumnView_setPreviewColumnVisible(void* self, int visible) {
	((QColumnView*)self)->setPreviewColumnVisible(visible);
}
int QColumnView_isPreviewColumnVisible(void* self) {
	auto&& __r = ((QColumnView*)self)->isPreviewColumnVisible();
	return __r;
}
void* QColumnView_previewWidget(void* self) {
	auto&& __r = ((QColumnView*)self)->previewWidget();
	return (void*)(__r);
}
void QColumnView_setPreviewWidget(void* self, void* widget) {
	((QColumnView*)self)->setPreviewWidget((QWidget*)widget);
}
void QColumnView_setColumnWidths(void* self, int* list, long long list_len) {
	QList<int> list_list; list_list.reserve(list_len);
	for (long long __i = 0; __i < list_len; __i++) list_list.append(list[__i]);
	((QColumnView*)self)->setColumnWidths(list_list);
}
int* QColumnView_columnWidths(void* self, long long* __out_len) {
	auto&& __r = ((QColumnView*)self)->columnWidths();
	*__out_len = __r.size();
	int* __arr = (int*)malloc(sizeof(int) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = __e;
	return __arr;
}
void QColumnView_initializeColumn(void* self, void* column) {
	((QColumnView__Go*)self)->initializeColumn((QAbstractItemView*)column);
}
int QColumnView_state(void* self) {
	auto&& __r = ((QColumnView__Go*)self)->state();
	return (int)(__r);
}
void QColumnView_scheduleDelayedItemsLayout(void* self) {
	((QColumnView__Go*)self)->scheduleDelayedItemsLayout();
}
void QColumnView_executeDelayedItemsLayout(void* self) {
	((QColumnView__Go*)self)->executeDelayedItemsLayout();
}
void QColumnView_setDirtyRegion(void* self, void* region) {
	((QColumnView__Go*)self)->setDirtyRegion(*(QRegion*)region);
}
void QColumnView_scrollDirtyRegion(void* self, int dx, int dy) {
	((QColumnView__Go*)self)->scrollDirtyRegion(dx, dy);
}
void* QColumnView_dirtyRegionOffset(void* self) {
	auto&& __r = ((QColumnView__Go*)self)->dirtyRegionOffset();
	return new QPoint(__r);
}
void QColumnView_startAutoScroll(void* self) {
	((QColumnView__Go*)self)->startAutoScroll();
}
void QColumnView_stopAutoScroll(void* self) {
	((QColumnView__Go*)self)->stopAutoScroll();
}
void QColumnView_doAutoScroll(void* self) {
	((QColumnView__Go*)self)->doAutoScroll();
}
int QColumnView_dropIndicatorPosition(void* self) {
	auto&& __r = ((QColumnView__Go*)self)->dropIndicatorPosition();
	return (int)(__r);
}
void QColumnView_setViewportMargins(void* self, int left, int top, int right, int bottom) {
	((QColumnView__Go*)self)->setViewportMargins(left, top, right, bottom);
}
void QColumnView_setViewportMargins2(void* self, void* margins) {
	((QColumnView__Go*)self)->setViewportMargins(*(QMargins*)margins);
}
void* QColumnView_viewportMargins(void* self) {
	auto&& __r = ((QColumnView__Go*)self)->viewportMargins();
	return new QMargins(__r);
}
void QColumnView_drawFrame(void* self, void* arg1) {
	((QColumnView__Go*)self)->drawFrame((QPainter*)arg1);
}
void QColumnView_setTabOrder(void* arg1, void* arg2) {
	QColumnView::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QColumnView_mouseGrabber() {
	auto&& __r = QColumnView::mouseGrabber();
	return (void*)(__r);
}
void* QColumnView_keyboardGrabber() {
	auto&& __r = QColumnView::keyboardGrabber();
	return (void*)(__r);
}
void* QColumnView_find(unsigned long long arg1) {
	auto&& __r = QColumnView::find(arg1);
	return (void*)(__r);
}
void* QColumnView_createWindowContainer(void* window) {
	auto&& __r = QColumnView::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QColumnView_updateMicroFocus(void* self) {
	((QColumnView__Go*)self)->updateMicroFocus();
}
void QColumnView_create(void* self) {
	((QColumnView__Go*)self)->create();
}
void QColumnView_destroy(void* self) {
	((QColumnView__Go*)self)->destroy();
}
int QColumnView_focusNextChild(void* self) {
	auto&& __r = ((QColumnView__Go*)self)->focusNextChild();
	return __r;
}
int QColumnView_focusPreviousChild(void* self) {
	auto&& __r = ((QColumnView__Go*)self)->focusPreviousChild();
	return __r;
}
void* QColumnView_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QColumnView::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QColumnView_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QColumnView::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QColumnView_sender(void* self) {
	auto&& __r = ((QColumnView__Go*)self)->sender();
	return (void*)(__r);
}
int QColumnView_senderSignalIndex(void* self) {
	auto&& __r = ((QColumnView__Go*)self)->senderSignalIndex();
	return __r;
}
int QColumnView_receivers(void* self, const char* signal) {
	auto&& __r = ((QColumnView__Go*)self)->receivers(signal);
	return __r;
}
int QColumnView_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QColumnView__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QColumnView_devicePixelRatioFScale() {
	auto&& __r = QColumnView::devicePixelRatioFScale();
	return __r;
}
int QColumnView_encodeMetricF(int metric, double value) {
	auto&& __r = QColumnView::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QColumnView_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QColumnView__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QColumnView_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QColumnView::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QColumnView_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QColumnView::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QColumnView_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QColumnView::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QColumnView_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QColumnView::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QColumnView_updateMicroFocus2(void* self, int query) {
	((QColumnView__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QColumnView_create2(void* self, unsigned long long arg1) {
	((QColumnView__Go*)self)->create(arg1);
}
void QColumnView_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QColumnView__Go*)self)->create(arg1, initializeWindow);
}
void QColumnView_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QColumnView__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QColumnView_destroy2(void* self, int destroyWindow) {
	((QColumnView__Go*)self)->destroy(destroyWindow);
}
void QColumnView_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QColumnView__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QColumnView_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QColumnView::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QColumnView_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QColumnView::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QColumnView_metaObject_default(void* self) {
	auto&& __r = ((QColumnView*)self)->QColumnView::metaObject();
	return (void*)(__r);
}
void* QColumnView_indexAt_default(void* self, void* point) {
	auto&& __r = ((QColumnView*)self)->QColumnView::indexAt(*(QPoint*)point);
	return new QModelIndex(__r);
}
void QColumnView_scrollTo_default(void* self, void* index, int hint) {
	((QColumnView*)self)->QColumnView::scrollTo(*(QModelIndex*)index, (QAbstractItemView::ScrollHint)(hint));
}
void* QColumnView_sizeHint_default(void* self) {
	auto&& __r = ((QColumnView*)self)->QColumnView::sizeHint();
	return new QSize(__r);
}
void* QColumnView_visualRect_default(void* self, void* index) {
	auto&& __r = ((QColumnView*)self)->QColumnView::visualRect(*(QModelIndex*)index);
	return new QRect(__r);
}
void QColumnView_setModel_default(void* self, void* model) {
	((QColumnView*)self)->QColumnView::setModel((QAbstractItemModel*)model);
}
void QColumnView_setSelectionModel_default(void* self, void* selectionModel) {
	((QColumnView*)self)->QColumnView::setSelectionModel((QItemSelectionModel*)selectionModel);
}
void QColumnView_setRootIndex_default(void* self, void* index) {
	((QColumnView*)self)->QColumnView::setRootIndex(*(QModelIndex*)index);
}
void QColumnView_selectAll_default(void* self) {
	((QColumnView*)self)->QColumnView::selectAll();
}
int QColumnView_isIndexHidden_default(void* self, void* index) {
	auto&& __r = ((QColumnView__Go*)self)->cppgoBase_isIndexHidden(*(QModelIndex*)index);
	return __r;
}
void* QColumnView_moveCursor_default(void* self, int cursorAction, int modifiers) {
	auto&& __r = ((QColumnView__Go*)self)->cppgoBase_moveCursor(cursorAction, QFlags<Qt::KeyboardModifier>(modifiers));
	return new QModelIndex(__r);
}
void QColumnView_resizeEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)event);
}
void QColumnView_setSelection_default(void* self, void* rect, int command) {
	((QColumnView__Go*)self)->cppgoBase_setSelection(*(QRect*)rect, QFlags<QItemSelectionModel::SelectionFlag>(command));
}
void* QColumnView_visualRegionForSelection_default(void* self, void* selection) {
	auto&& __r = ((QColumnView__Go*)self)->cppgoBase_visualRegionForSelection(*(QItemSelection*)selection);
	return new QRegion(__r);
}
int QColumnView_horizontalOffset_default(void* self) {
	auto&& __r = ((QColumnView__Go*)self)->cppgoBase_horizontalOffset();
	return __r;
}
int QColumnView_verticalOffset_default(void* self) {
	auto&& __r = ((QColumnView__Go*)self)->cppgoBase_verticalOffset();
	return __r;
}
void QColumnView_rowsInserted_default(void* self, void* parent, int start, int end) {
	((QColumnView__Go*)self)->cppgoBase_rowsInserted(*(QModelIndex*)parent, start, end);
}
void QColumnView_currentChanged_default(void* self, void* current, void* previous) {
	((QColumnView__Go*)self)->cppgoBase_currentChanged(*(QModelIndex*)current, *(QModelIndex*)previous);
}
void QColumnView_scrollContentsBy_default(void* self, int dx, int dy) {
	((QColumnView__Go*)self)->cppgoBase_scrollContentsBy(dx, dy);
}
void* QColumnView_createColumn_default(void* self, void* rootIndex) {
	auto&& __r = ((QColumnView__Go*)self)->cppgoBase_createColumn(*(QModelIndex*)rootIndex);
	return (void*)(__r);
}
void QColumnView_keyboardSearch_default(void* self, const char* search, long long search_len) {
	QString search_qstr = QString::fromUtf8(search, search_len);
	((QColumnView*)self)->QColumnView::keyboardSearch(search_qstr);
}
int QColumnView_sizeHintForRow_default(void* self, int row) {
	auto&& __r = ((QColumnView*)self)->QColumnView::sizeHintForRow(row);
	return __r;
}
int QColumnView_sizeHintForColumn_default(void* self, int column) {
	auto&& __r = ((QColumnView*)self)->QColumnView::sizeHintForColumn(column);
	return __r;
}
void* QColumnView_itemDelegateForIndex_default(void* self, void* index) {
	auto&& __r = ((QColumnView*)self)->QColumnView::itemDelegateForIndex(*(QModelIndex*)index);
	return (void*)(__r);
}
void* QColumnView_inputMethodQuery_default(void* self, int query) {
	auto&& __r = ((QColumnView*)self)->QColumnView::inputMethodQuery((Qt::InputMethodQuery)(query));
	return new QVariant(__r);
}
void QColumnView_reset_default(void* self) {
	((QColumnView*)self)->QColumnView::reset();
}
void QColumnView_doItemsLayout_default(void* self) {
	((QColumnView*)self)->QColumnView::doItemsLayout();
}
void QColumnView_rowsAboutToBeRemoved_default(void* self, void* parent, int start, int end) {
	((QColumnView__Go*)self)->cppgoBase_rowsAboutToBeRemoved(*(QModelIndex*)parent, start, end);
}
void QColumnView_selectionChanged_default(void* self, void* selected, void* deselected) {
	((QColumnView__Go*)self)->cppgoBase_selectionChanged(*(QItemSelection*)selected, *(QItemSelection*)deselected);
}
void QColumnView_updateEditorData_default(void* self) {
	((QColumnView__Go*)self)->cppgoBase_updateEditorData();
}
void QColumnView_updateEditorGeometries_default(void* self) {
	((QColumnView__Go*)self)->cppgoBase_updateEditorGeometries();
}
void QColumnView_updateGeometries_default(void* self) {
	((QColumnView__Go*)self)->cppgoBase_updateGeometries();
}
void QColumnView_verticalScrollbarAction_default(void* self, int action) {
	((QColumnView__Go*)self)->cppgoBase_verticalScrollbarAction(action);
}
void QColumnView_horizontalScrollbarAction_default(void* self, int action) {
	((QColumnView__Go*)self)->cppgoBase_horizontalScrollbarAction(action);
}
void QColumnView_verticalScrollbarValueChanged_default(void* self, int value) {
	((QColumnView__Go*)self)->cppgoBase_verticalScrollbarValueChanged(value);
}
void QColumnView_horizontalScrollbarValueChanged_default(void* self, int value) {
	((QColumnView__Go*)self)->cppgoBase_horizontalScrollbarValueChanged(value);
}
void QColumnView_closeEditor_default(void* self, void* editor, int hint) {
	((QColumnView__Go*)self)->cppgoBase_closeEditor((QWidget*)editor, hint);
}
void QColumnView_commitData_default(void* self, void* editor) {
	((QColumnView__Go*)self)->cppgoBase_commitData((QWidget*)editor);
}
void QColumnView_editorDestroyed_default(void* self, void* editor) {
	((QColumnView__Go*)self)->cppgoBase_editorDestroyed((QObject*)editor);
}
void** QColumnView_selectedIndexes_default(void* self, long long* __out_len) {
	auto&& __r = ((QColumnView__Go*)self)->cppgoBase_selectedIndexes();
	*__out_len = __r.size();
	void** __arr = (void**)malloc(sizeof(void*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) __arr[__i++] = new QModelIndex(__e);
	return __arr;
}
int QColumnView_edit_default(void* self, void* index, int trigger, void* event) {
	auto&& __r = ((QColumnView__Go*)self)->cppgoBase_edit(*(QModelIndex*)index, trigger, (QEvent*)event);
	return __r;
}
int QColumnView_selectionCommand_default(void* self, void* index, void* event) {
	auto&& __r = ((QColumnView__Go*)self)->cppgoBase_selectionCommand(*(QModelIndex*)index, (QEvent*)event);
	return (int)(__r);
}
void QColumnView_startDrag_default(void* self, int supportedActions) {
	((QColumnView__Go*)self)->cppgoBase_startDrag(QFlags<Qt::DropAction>(supportedActions));
}
void QColumnView_initViewItemOption_default(void* self, void* option) {
	((QColumnView__Go*)self)->cppgoBase_initViewItemOption((QStyleOptionViewItem*)option);
}
int QColumnView_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QColumnView__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
int QColumnView_event_default(void* self, void* event) {
	auto&& __r = ((QColumnView__Go*)self)->cppgoBase_event((QEvent*)event);
	return __r;
}
int QColumnView_viewportEvent_default(void* self, void* event) {
	auto&& __r = ((QColumnView__Go*)self)->cppgoBase_viewportEvent((QEvent*)event);
	return __r;
}
void QColumnView_mousePressEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)event);
}
void QColumnView_mouseMoveEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)event);
}
void QColumnView_mouseReleaseEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)event);
}
void QColumnView_mouseDoubleClickEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)event);
}
void QColumnView_dragEnterEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)event);
}
void QColumnView_dragMoveEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)event);
}
void QColumnView_dragLeaveEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)event);
}
void QColumnView_dropEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_dropEvent((QDropEvent*)event);
}
void QColumnView_focusInEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)event);
}
void QColumnView_focusOutEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)event);
}
void QColumnView_keyPressEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)event);
}
void QColumnView_timerEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QColumnView_inputMethodEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)event);
}
int QColumnView_eventFilter_default(void* self, void* object, void* event) {
	auto&& __r = ((QColumnView__Go*)self)->cppgoBase_eventFilter((QObject*)object, (QEvent*)event);
	return __r;
}
void* QColumnView_viewportSizeHint_default(void* self) {
	auto&& __r = ((QColumnView__Go*)self)->cppgoBase_viewportSizeHint();
	return new QSize(__r);
}
void* QColumnView_minimumSizeHint_default(void* self) {
	auto&& __r = ((QColumnView*)self)->QColumnView::minimumSizeHint();
	return new QSize(__r);
}
void QColumnView_setupViewport_default(void* self, void* viewport) {
	((QColumnView*)self)->QColumnView::setupViewport((QWidget*)viewport);
}
void QColumnView_paintEvent_default(void* self, void* arg1) {
	((QColumnView__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)arg1);
}
void QColumnView_wheelEvent_default(void* self, void* arg1) {
	((QColumnView__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)arg1);
}
void QColumnView_contextMenuEvent_default(void* self, void* arg1) {
	((QColumnView__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)arg1);
}
void QColumnView_changeEvent_default(void* self, void* arg1) {
	((QColumnView__Go*)self)->cppgoBase_changeEvent((QEvent*)arg1);
}
void QColumnView_initStyleOption_default(void* self, void* option) {
	((QColumnView__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QColumnView_devType_default(void* self) {
	auto&& __r = ((QColumnView*)self)->QColumnView::devType();
	return __r;
}
void QColumnView_setVisible_default(void* self, int visible) {
	((QColumnView*)self)->QColumnView::setVisible(visible);
}
int QColumnView_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QColumnView*)self)->QColumnView::heightForWidth(arg1);
	return __r;
}
int QColumnView_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QColumnView*)self)->QColumnView::hasHeightForWidth();
	return __r;
}
void* QColumnView_paintEngine_default(void* self) {
	auto&& __r = ((QColumnView*)self)->QColumnView::paintEngine();
	return (void*)(__r);
}
void QColumnView_keyReleaseEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)event);
}
void QColumnView_enterEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QColumnView_leaveEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QColumnView_moveEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QColumnView_closeEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QColumnView_tabletEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QColumnView_actionEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QColumnView_showEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_showEvent((QShowEvent*)event);
}
void QColumnView_hideEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QColumnView_metric_default(void* self, int arg1) {
	auto&& __r = ((QColumnView__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QColumnView_initPainter_default(void* self, void* painter) {
	((QColumnView__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QColumnView_redirected_default(void* self, void* offset) {
	auto&& __r = ((QColumnView__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QColumnView_sharedPainter_default(void* self) {
	auto&& __r = ((QColumnView__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QColumnView_childEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QColumnView_customEvent_default(void* self, void* event) {
	((QColumnView__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QColumnView_connectNotify_default(void* self, void* signal) {
	((QColumnView__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QColumnView_disconnectNotify_default(void* self, void* signal) {
	((QColumnView__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QColumnView_connect_updatePreviewWidget(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColumnView*)self, &QColumnView::updatePreviewWidget, (QColumnView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColumnView);
	__conns_QColumnView[__id] = __c;
}
void QColumnView_connect_pressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColumnView*)self, &QColumnView::pressed, (QColumnView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColumnView);
	__conns_QColumnView[__id] = __c;
}
void QColumnView_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColumnView*)self, &QColumnView::clicked, (QColumnView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColumnView);
	__conns_QColumnView[__id] = __c;
}
void QColumnView_connect_doubleClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColumnView*)self, &QColumnView::doubleClicked, (QColumnView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColumnView);
	__conns_QColumnView[__id] = __c;
}
void QColumnView_connect_activated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColumnView*)self, &QColumnView::activated, (QColumnView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColumnView);
	__conns_QColumnView[__id] = __c;
}
void QColumnView_connect_entered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColumnView*)self, &QColumnView::entered, (QColumnView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColumnView);
	__conns_QColumnView[__id] = __c;
}
void QColumnView_connect_viewportEntered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColumnView*)self, &QColumnView::viewportEntered, (QColumnView*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColumnView);
	__conns_QColumnView[__id] = __c;
}
void QColumnView_connect_iconSizeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColumnView*)self, &QColumnView::iconSizeChanged, (QColumnView*)self, [__id](const QSize& size){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&size);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColumnView);
	__conns_QColumnView[__id] = __c;
}
void QColumnView_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColumnView*)self, &QColumnView::windowTitleChanged, (QColumnView*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColumnView);
	__conns_QColumnView[__id] = __c;
}
void QColumnView_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColumnView*)self, &QColumnView::windowIconChanged, (QColumnView*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColumnView);
	__conns_QColumnView[__id] = __c;
}
void QColumnView_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColumnView*)self, &QColumnView::windowIconTextChanged, (QColumnView*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColumnView);
	__conns_QColumnView[__id] = __c;
}
void QColumnView_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColumnView*)self, &QColumnView::customContextMenuRequested, (QColumnView*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColumnView);
	__conns_QColumnView[__id] = __c;
}
void QColumnView_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColumnView*)self, &QColumnView::destroyed, (QColumnView*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColumnView);
	__conns_QColumnView[__id] = __c;
}
void QColumnView_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColumnView*)self, &QColumnView::destroyed, (QColumnView*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColumnView);
	__conns_QColumnView[__id] = __c;
}
void QColumnView_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QColumnView*)self, &QColumnView::objectNameChanged, (QColumnView*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QColumnView);
	__conns_QColumnView[__id] = __c;
}
void QColumnView__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QColumnView);
		auto __it = __conns_QColumnView.find(id);
		if (__it == __conns_QColumnView.end()) return;
		__c = __it->second;
		__conns_QColumnView.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QColumnView__Go_new(long long id) { return new QColumnView__Go(id); }
void QColumnView_delete(void* self) { delete (QColumnView*)self; }

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
#include "gen_qpinchgesture.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QPinchGesture__Go : public QPinchGesture {
public:
	long long __goId;
	unsigned char __goOn[8] = {};
	template<typename... A> QPinchGesture__Go(long long id, A&&... a) : QPinchGesture(std::forward<A>(a)...), __goId(id) {}
	using QPinchGesture::sender;
	using QPinchGesture::senderSignalIndex;
	using QPinchGesture::receivers;
	using QPinchGesture::isSignalConnected;
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QPinchGesture::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QPinchGesture::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QPinchGesture::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QPinchGesture::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QPinchGesture::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QPinchGesture::metaObject();
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QPinchGesture::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QPinchGesture::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QPinchGesture::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QPinchGesture::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QPinchGesture::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QPinchGesture::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QPinchGesture::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QPinchGesture;
static std::mutex __connsMu_QPinchGesture;

extern "C" {

void* QPinchGesture_base(void* self) { return (void*)static_cast<QGesture*>((QPinchGesture*)self); }
void QPinchGesture_chain(void* self, void** out) {
	QPinchGesture* __p = (QPinchGesture*)self;
	QGesture* __b0 = static_cast<QGesture*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QPinchGesture__set_override(void* self, int idx) { ((QPinchGesture__Go*)self)->__goOn[idx] = 1; }
void* QPinchGesture_fromQObject(void* o) { return (void*)qobject_cast<QPinchGesture*>((QObject*)o); }
void* QPinchGesture_new(long long __goid) {
	return new QPinchGesture__Go(__goid);
}
void* QPinchGesture_new2(long long __goid, void* parent) {
	return new QPinchGesture__Go(__goid, (QObject*)parent);
}
char* QPinchGesture_tr(const char* s, long long* __out_len) {
	auto&& __r = QPinchGesture::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QPinchGesture_totalChangeFlags(void* self) {
	auto&& __r = ((QPinchGesture*)self)->totalChangeFlags();
	return (int)(__r);
}
void QPinchGesture_setTotalChangeFlags(void* self, int value) {
	((QPinchGesture*)self)->setTotalChangeFlags(QFlags<QPinchGesture::ChangeFlag>(value));
}
int QPinchGesture_changeFlags(void* self) {
	auto&& __r = ((QPinchGesture*)self)->changeFlags();
	return (int)(__r);
}
void QPinchGesture_setChangeFlags(void* self, int value) {
	((QPinchGesture*)self)->setChangeFlags(QFlags<QPinchGesture::ChangeFlag>(value));
}
void* QPinchGesture_startCenterPoint(void* self) {
	auto&& __r = ((QPinchGesture*)self)->startCenterPoint();
	return new QPointF(__r);
}
void* QPinchGesture_lastCenterPoint(void* self) {
	auto&& __r = ((QPinchGesture*)self)->lastCenterPoint();
	return new QPointF(__r);
}
void* QPinchGesture_centerPoint(void* self) {
	auto&& __r = ((QPinchGesture*)self)->centerPoint();
	return new QPointF(__r);
}
void QPinchGesture_setStartCenterPoint(void* self, void* value) {
	((QPinchGesture*)self)->setStartCenterPoint(*(QPointF*)value);
}
void QPinchGesture_setLastCenterPoint(void* self, void* value) {
	((QPinchGesture*)self)->setLastCenterPoint(*(QPointF*)value);
}
void QPinchGesture_setCenterPoint(void* self, void* value) {
	((QPinchGesture*)self)->setCenterPoint(*(QPointF*)value);
}
double QPinchGesture_totalScaleFactor(void* self) {
	auto&& __r = ((QPinchGesture*)self)->totalScaleFactor();
	return __r;
}
double QPinchGesture_lastScaleFactor(void* self) {
	auto&& __r = ((QPinchGesture*)self)->lastScaleFactor();
	return __r;
}
double QPinchGesture_scaleFactor(void* self) {
	auto&& __r = ((QPinchGesture*)self)->scaleFactor();
	return __r;
}
void QPinchGesture_setTotalScaleFactor(void* self, double value) {
	((QPinchGesture*)self)->setTotalScaleFactor(value);
}
void QPinchGesture_setLastScaleFactor(void* self, double value) {
	((QPinchGesture*)self)->setLastScaleFactor(value);
}
void QPinchGesture_setScaleFactor(void* self, double value) {
	((QPinchGesture*)self)->setScaleFactor(value);
}
double QPinchGesture_totalRotationAngle(void* self) {
	auto&& __r = ((QPinchGesture*)self)->totalRotationAngle();
	return __r;
}
double QPinchGesture_lastRotationAngle(void* self) {
	auto&& __r = ((QPinchGesture*)self)->lastRotationAngle();
	return __r;
}
double QPinchGesture_rotationAngle(void* self) {
	auto&& __r = ((QPinchGesture*)self)->rotationAngle();
	return __r;
}
void QPinchGesture_setTotalRotationAngle(void* self, double value) {
	((QPinchGesture*)self)->setTotalRotationAngle(value);
}
void QPinchGesture_setLastRotationAngle(void* self, double value) {
	((QPinchGesture*)self)->setLastRotationAngle(value);
}
void QPinchGesture_setRotationAngle(void* self, double value) {
	((QPinchGesture*)self)->setRotationAngle(value);
}
void* QPinchGesture_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QPinchGesture::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QPinchGesture_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QPinchGesture::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QPinchGesture_sender(void* self) {
	auto&& __r = ((QPinchGesture__Go*)self)->sender();
	return (void*)(__r);
}
int QPinchGesture_senderSignalIndex(void* self) {
	auto&& __r = ((QPinchGesture__Go*)self)->senderSignalIndex();
	return __r;
}
int QPinchGesture_receivers(void* self, const char* signal) {
	auto&& __r = ((QPinchGesture__Go*)self)->receivers(signal);
	return __r;
}
int QPinchGesture_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QPinchGesture__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QPinchGesture_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QPinchGesture::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QPinchGesture_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QPinchGesture::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QPinchGesture_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QPinchGesture::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QPinchGesture_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QPinchGesture::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QPinchGesture_metaObject_default(void* self) {
	auto&& __r = ((QPinchGesture*)self)->QPinchGesture::metaObject();
	return (void*)(__r);
}
int QPinchGesture_event_default(void* self, void* event) {
	auto&& __r = ((QPinchGesture*)self)->QPinchGesture::event((QEvent*)event);
	return __r;
}
int QPinchGesture_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QPinchGesture*)self)->QPinchGesture::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QPinchGesture_timerEvent_default(void* self, void* event) {
	((QPinchGesture__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QPinchGesture_childEvent_default(void* self, void* event) {
	((QPinchGesture__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QPinchGesture_customEvent_default(void* self, void* event) {
	((QPinchGesture__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QPinchGesture_connectNotify_default(void* self, void* signal) {
	((QPinchGesture__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QPinchGesture_disconnectNotify_default(void* self, void* signal) {
	((QPinchGesture__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QPinchGesture_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPinchGesture*)self, &QPinchGesture::destroyed, (QPinchGesture*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPinchGesture);
	__conns_QPinchGesture[__id] = __c;
}
void QPinchGesture_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPinchGesture*)self, &QPinchGesture::destroyed, (QPinchGesture*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPinchGesture);
	__conns_QPinchGesture[__id] = __c;
}
void QPinchGesture_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QPinchGesture*)self, &QPinchGesture::objectNameChanged, (QPinchGesture*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QPinchGesture);
	__conns_QPinchGesture[__id] = __c;
}
void QPinchGesture__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QPinchGesture);
		auto __it = __conns_QPinchGesture.find(id);
		if (__it == __conns_QPinchGesture.end()) return;
		__c = __it->second;
		__conns_QPinchGesture.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QPinchGesture__Go_new(long long id, void* parent) { return new QPinchGesture__Go(id, (QObject*)parent); }
void QPinchGesture_delete(void* self) { delete (QPinchGesture*)self; }

}

#include <QtWidgets/qgraphicseffect.h>
#include <QtCore/qobject.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qrect.h>
#include <QtGui/qpainter.h>
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
#include "gen_qgraphicsblureffect.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QGraphicsBlurEffect__Go : public QGraphicsBlurEffect {
public:
	long long __goId;
	unsigned char __goOn[11] = {};
	template<typename... A> QGraphicsBlurEffect__Go(long long id, A&&... a) : QGraphicsBlurEffect(std::forward<A>(a)...), __goId(id) {}
	using QGraphicsBlurEffect::updateBoundingRect;
	using QGraphicsBlurEffect::sourceIsPixmap;
	using QGraphicsBlurEffect::sourceBoundingRect;
	using QGraphicsBlurEffect::drawSource;
	using QGraphicsBlurEffect::sourcePixmap;
	using QGraphicsBlurEffect::sender;
	using QGraphicsBlurEffect::senderSignalIndex;
	using QGraphicsBlurEffect::receivers;
	using QGraphicsBlurEffect::isSignalConnected;
	void cppgoBase_draw(QPainter* __a0) { QGraphicsBlurEffect::draw(__a0); }
	void cppgoBase_sourceChanged(QFlags<QGraphicsEffect::ChangeFlag> __a0) { QGraphicsBlurEffect::sourceChanged(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QGraphicsBlurEffect::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QGraphicsBlurEffect::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QGraphicsBlurEffect::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QGraphicsBlurEffect::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QGraphicsBlurEffect::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QGraphicsBlurEffect::metaObject();
	}
	QRectF boundingRectFor(const QRectF& rect) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 5831744872130794138LL, (void*)&rect, &__h); if (__h) return *(QRectF*)__r; }
		return QGraphicsBlurEffect::boundingRectFor(rect);
	}
	void draw(QPainter* painter) override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1044652152830004333LL, (void*)painter, &__h); if (__h) return; }
		QGraphicsBlurEffect::draw(painter);
	}
	void sourceChanged(QFlags<QGraphicsEffect::ChangeFlag> flags) override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, -4520368087989481956LL, (int)(flags), &__h); if (__h) return; }
		QGraphicsBlurEffect::sourceChanged(flags);
	}
	bool event(QEvent* event) override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsBlurEffect::event(event);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QGraphicsBlurEffect::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QGraphicsBlurEffect::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QGraphicsBlurEffect::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QGraphicsBlurEffect::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsBlurEffect::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QGraphicsBlurEffect::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QGraphicsBlurEffect;
static std::mutex __connsMu_QGraphicsBlurEffect;

extern "C" {

void* QGraphicsBlurEffect_base(void* self) { return (void*)static_cast<QGraphicsEffect*>((QGraphicsBlurEffect*)self); }
void QGraphicsBlurEffect_chain(void* self, void** out) {
	QGraphicsBlurEffect* __p = (QGraphicsBlurEffect*)self;
	QGraphicsEffect* __b0 = static_cast<QGraphicsEffect*>(__p); out[0] = (void*)__b0;
	QObject* __b1 = static_cast<QObject*>(__b0); out[1] = (void*)__b1;
}
void QGraphicsBlurEffect__set_override(void* self, int idx) { ((QGraphicsBlurEffect__Go*)self)->__goOn[idx] = 1; }
void* QGraphicsBlurEffect_fromQObject(void* o) { return (void*)qobject_cast<QGraphicsBlurEffect*>((QObject*)o); }
void* QGraphicsBlurEffect_new(long long __goid) {
	return new QGraphicsBlurEffect__Go(__goid);
}
void* QGraphicsBlurEffect_new2(long long __goid, void* parent) {
	return new QGraphicsBlurEffect__Go(__goid, (QObject*)parent);
}
char* QGraphicsBlurEffect_tr(const char* s, long long* __out_len) {
	auto&& __r = QGraphicsBlurEffect::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
double QGraphicsBlurEffect_blurRadius(void* self) {
	auto&& __r = ((QGraphicsBlurEffect*)self)->blurRadius();
	return __r;
}
int QGraphicsBlurEffect_blurHints(void* self) {
	auto&& __r = ((QGraphicsBlurEffect*)self)->blurHints();
	return (int)(__r);
}
void QGraphicsBlurEffect_setBlurRadius(void* self, double blurRadius) {
	((QGraphicsBlurEffect*)self)->setBlurRadius(blurRadius);
}
void QGraphicsBlurEffect_setBlurHints(void* self, int hints) {
	((QGraphicsBlurEffect*)self)->setBlurHints(QFlags<QGraphicsBlurEffect::BlurHint>(hints));
}
void QGraphicsBlurEffect_blurRadiusChanged(void* self, double blurRadius) {
	((QGraphicsBlurEffect*)self)->blurRadiusChanged(blurRadius);
}
void QGraphicsBlurEffect_blurHintsChanged(void* self, int hints) {
	((QGraphicsBlurEffect*)self)->blurHintsChanged(QFlags<QGraphicsBlurEffect::BlurHint>(hints));
}
void QGraphicsBlurEffect_updateBoundingRect(void* self) {
	((QGraphicsBlurEffect__Go*)self)->updateBoundingRect();
}
int QGraphicsBlurEffect_sourceIsPixmap(void* self) {
	auto&& __r = ((QGraphicsBlurEffect__Go*)self)->sourceIsPixmap();
	return __r;
}
void* QGraphicsBlurEffect_sourceBoundingRect(void* self) {
	auto&& __r = ((QGraphicsBlurEffect__Go*)self)->sourceBoundingRect();
	return new QRectF(__r);
}
void QGraphicsBlurEffect_drawSource(void* self, void* painter) {
	((QGraphicsBlurEffect__Go*)self)->drawSource((QPainter*)painter);
}
void* QGraphicsBlurEffect_sourcePixmap(void* self) {
	auto&& __r = ((QGraphicsBlurEffect__Go*)self)->sourcePixmap();
	return new QPixmap(__r);
}
void* QGraphicsBlurEffect_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QGraphicsBlurEffect::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsBlurEffect_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QGraphicsBlurEffect::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QGraphicsBlurEffect_sender(void* self) {
	auto&& __r = ((QGraphicsBlurEffect__Go*)self)->sender();
	return (void*)(__r);
}
int QGraphicsBlurEffect_senderSignalIndex(void* self) {
	auto&& __r = ((QGraphicsBlurEffect__Go*)self)->senderSignalIndex();
	return __r;
}
int QGraphicsBlurEffect_receivers(void* self, const char* signal) {
	auto&& __r = ((QGraphicsBlurEffect__Go*)self)->receivers(signal);
	return __r;
}
int QGraphicsBlurEffect_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QGraphicsBlurEffect__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QGraphicsBlurEffect_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QGraphicsBlurEffect::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QGraphicsBlurEffect_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QGraphicsBlurEffect::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QGraphicsBlurEffect_sourceBoundingRect2(void* self, int system) {
	auto&& __r = ((QGraphicsBlurEffect__Go*)self)->sourceBoundingRect((Qt::CoordinateSystem)(system));
	return new QRectF(__r);
}
void* QGraphicsBlurEffect_sourcePixmap2(void* self, int system) {
	auto&& __r = ((QGraphicsBlurEffect__Go*)self)->sourcePixmap((Qt::CoordinateSystem)(system));
	return new QPixmap(__r);
}
void* QGraphicsBlurEffect_sourcePixmap3(void* self, int system, void* offset) {
	auto&& __r = ((QGraphicsBlurEffect__Go*)self)->sourcePixmap((Qt::CoordinateSystem)(system), (QPoint*)offset);
	return new QPixmap(__r);
}
void* QGraphicsBlurEffect_sourcePixmap4(void* self, int system, void* offset, int mode) {
	auto&& __r = ((QGraphicsBlurEffect__Go*)self)->sourcePixmap((Qt::CoordinateSystem)(system), (QPoint*)offset, (QGraphicsEffect::PixmapPadMode)(mode));
	return new QPixmap(__r);
}
void* QGraphicsBlurEffect_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QGraphicsBlurEffect::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsBlurEffect_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QGraphicsBlurEffect::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QGraphicsBlurEffect_metaObject_default(void* self) {
	auto&& __r = ((QGraphicsBlurEffect*)self)->QGraphicsBlurEffect::metaObject();
	return (void*)(__r);
}
void* QGraphicsBlurEffect_boundingRectFor_default(void* self, void* rect) {
	auto&& __r = ((QGraphicsBlurEffect*)self)->QGraphicsBlurEffect::boundingRectFor(*(QRectF*)rect);
	return new QRectF(__r);
}
void QGraphicsBlurEffect_draw_default(void* self, void* painter) {
	((QGraphicsBlurEffect__Go*)self)->cppgoBase_draw((QPainter*)painter);
}
void QGraphicsBlurEffect_sourceChanged_default(void* self, int flags) {
	((QGraphicsBlurEffect__Go*)self)->cppgoBase_sourceChanged(QFlags<QGraphicsEffect::ChangeFlag>(flags));
}
int QGraphicsBlurEffect_event_default(void* self, void* event) {
	auto&& __r = ((QGraphicsBlurEffect*)self)->QGraphicsBlurEffect::event((QEvent*)event);
	return __r;
}
int QGraphicsBlurEffect_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QGraphicsBlurEffect*)self)->QGraphicsBlurEffect::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QGraphicsBlurEffect_timerEvent_default(void* self, void* event) {
	((QGraphicsBlurEffect__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QGraphicsBlurEffect_childEvent_default(void* self, void* event) {
	((QGraphicsBlurEffect__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QGraphicsBlurEffect_customEvent_default(void* self, void* event) {
	((QGraphicsBlurEffect__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QGraphicsBlurEffect_connectNotify_default(void* self, void* signal) {
	((QGraphicsBlurEffect__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QGraphicsBlurEffect_disconnectNotify_default(void* self, void* signal) {
	((QGraphicsBlurEffect__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QGraphicsBlurEffect_connect_blurRadiusChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsBlurEffect*)self, &QGraphicsBlurEffect::blurRadiusChanged, (QGraphicsBlurEffect*)self, [__id](qreal blurRadius){
		if (cppgo_sig_void_unsignedlonglong) ((void(*)(long long, unsigned long long))cppgo_sig_void_unsignedlonglong)(__id, __builtin_bit_cast(unsigned long long, (double)(blurRadius)));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsBlurEffect);
	__conns_QGraphicsBlurEffect[__id] = __c;
}
void QGraphicsBlurEffect_connect_blurHintsChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsBlurEffect*)self, &QGraphicsBlurEffect::blurHintsChanged, (QGraphicsBlurEffect*)self, [__id](QFlags<QGraphicsBlurEffect::BlurHint> hints){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, (int)(hints));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsBlurEffect);
	__conns_QGraphicsBlurEffect[__id] = __c;
}
void QGraphicsBlurEffect_connect_enabledChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsBlurEffect*)self, &QGraphicsBlurEffect::enabledChanged, (QGraphicsBlurEffect*)self, [__id](bool enabled){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, enabled);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsBlurEffect);
	__conns_QGraphicsBlurEffect[__id] = __c;
}
void QGraphicsBlurEffect_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsBlurEffect*)self, &QGraphicsBlurEffect::destroyed, (QGraphicsBlurEffect*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsBlurEffect);
	__conns_QGraphicsBlurEffect[__id] = __c;
}
void QGraphicsBlurEffect_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsBlurEffect*)self, &QGraphicsBlurEffect::destroyed, (QGraphicsBlurEffect*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsBlurEffect);
	__conns_QGraphicsBlurEffect[__id] = __c;
}
void QGraphicsBlurEffect_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QGraphicsBlurEffect*)self, &QGraphicsBlurEffect::objectNameChanged, (QGraphicsBlurEffect*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsBlurEffect);
	__conns_QGraphicsBlurEffect[__id] = __c;
}
void QGraphicsBlurEffect__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QGraphicsBlurEffect);
		auto __it = __conns_QGraphicsBlurEffect.find(id);
		if (__it == __conns_QGraphicsBlurEffect.end()) return;
		__c = __it->second;
		__conns_QGraphicsBlurEffect.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QGraphicsBlurEffect__Go_new(long long id, void* parent) { return new QGraphicsBlurEffect__Go(id, (QObject*)parent); }
void QGraphicsBlurEffect_delete(void* self) { delete (QGraphicsBlurEffect*)self; }

}

#include <QtWidgets/qheaderview.h>
#include <QtWidgets/qabstractitemview.h>
#include <QtCore/qnamespace.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qabstractitemmodel.h>
#include <QtCore/qsize.h>
#include <QtCore/qpoint.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtGui/qpainter.h>
#include <QtCore/qrect.h>
#include <QtCore/qitemselectionmodel.h>
#include <QtGui/qregion.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qabstractitemdelegate.h>
#include <QtCore/qvariant.h>
#include <QtCore/qobject.h>
#include <QtWidgets/qscrollbar.h>
#include <QtWidgets/qabstractscrollarea.h>
#include <QtCore/qmargins.h>
#include <QtWidgets/qframe.h>
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
#include "gen_qheaderview.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

static std::unordered_map<long long, QMetaObject::Connection> __conns_QHeaderView;
static std::mutex __connsMu_QHeaderView;

extern "C" {

void* QHeaderView_base(void* self) { return (void*)static_cast<QAbstractItemView*>((QHeaderView*)self); }
void QHeaderView_chain(void* self, void** out) {
	QHeaderView* __p = (QHeaderView*)self;
	QAbstractItemView* __b0 = static_cast<QAbstractItemView*>(__p); out[0] = (void*)__b0;
	QAbstractScrollArea* __b1 = static_cast<QAbstractScrollArea*>(__b0); out[1] = (void*)__b1;
	QFrame* __b2 = static_cast<QFrame*>(__b1); out[2] = (void*)__b2;
	QWidget* __b3 = static_cast<QWidget*>(__b2); out[3] = (void*)__b3;
	QObject* __b4 = static_cast<QObject*>(__b3); out[4] = (void*)__b4;
}
void* QHeaderView_fromQObject(void* o) { return (void*)qobject_cast<QHeaderView*>((QObject*)o); }
void* QHeaderView_new(int orientation) {
	return new QHeaderView((Qt::Orientation)(orientation));
}
void* QHeaderView_new2(int orientation, void* parent) {
	return new QHeaderView((Qt::Orientation)(orientation), (QWidget*)parent);
}
char* QHeaderView_tr(const char* s, long long* __out_len) {
	auto&& __r = QHeaderView::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QHeaderView_orientation(void* self) {
	auto&& __r = ((QHeaderView*)self)->orientation();
	return (int)(__r);
}
int QHeaderView_offset(void* self) {
	auto&& __r = ((QHeaderView*)self)->offset();
	return __r;
}
int QHeaderView_length(void* self) {
	auto&& __r = ((QHeaderView*)self)->length();
	return __r;
}
int QHeaderView_sectionSizeHint(void* self, int logicalIndex) {
	auto&& __r = ((QHeaderView*)self)->sectionSizeHint(logicalIndex);
	return __r;
}
int QHeaderView_visualIndexAt(void* self, int position) {
	auto&& __r = ((QHeaderView*)self)->visualIndexAt(position);
	return __r;
}
int QHeaderView_logicalIndexAt(void* self, int position) {
	auto&& __r = ((QHeaderView*)self)->logicalIndexAt(position);
	return __r;
}
int QHeaderView_logicalIndexAt2(void* self, int x, int y) {
	auto&& __r = ((QHeaderView*)self)->logicalIndexAt(x, y);
	return __r;
}
int QHeaderView_logicalIndexAt3(void* self, void* pos) {
	auto&& __r = ((QHeaderView*)self)->logicalIndexAt(*(QPoint*)pos);
	return __r;
}
int QHeaderView_sectionSize(void* self, int logicalIndex) {
	auto&& __r = ((QHeaderView*)self)->sectionSize(logicalIndex);
	return __r;
}
int QHeaderView_sectionPosition(void* self, int logicalIndex) {
	auto&& __r = ((QHeaderView*)self)->sectionPosition(logicalIndex);
	return __r;
}
int QHeaderView_sectionViewportPosition(void* self, int logicalIndex) {
	auto&& __r = ((QHeaderView*)self)->sectionViewportPosition(logicalIndex);
	return __r;
}
void QHeaderView_moveSection(void* self, int from, int to) {
	((QHeaderView*)self)->moveSection(from, to);
}
void QHeaderView_swapSections(void* self, int first, int second) {
	((QHeaderView*)self)->swapSections(first, second);
}
void QHeaderView_resizeSection(void* self, int logicalIndex, int size) {
	((QHeaderView*)self)->resizeSection(logicalIndex, size);
}
void QHeaderView_resizeSections(void* self, int mode) {
	((QHeaderView*)self)->resizeSections((QHeaderView::ResizeMode)(mode));
}
int QHeaderView_isSectionHidden(void* self, int logicalIndex) {
	auto&& __r = ((QHeaderView*)self)->isSectionHidden(logicalIndex);
	return __r;
}
void QHeaderView_setSectionHidden(void* self, int logicalIndex, int hide) {
	((QHeaderView*)self)->setSectionHidden(logicalIndex, hide);
}
int QHeaderView_hiddenSectionCount(void* self) {
	auto&& __r = ((QHeaderView*)self)->hiddenSectionCount();
	return __r;
}
void QHeaderView_hideSection(void* self, int logicalIndex) {
	((QHeaderView*)self)->hideSection(logicalIndex);
}
void QHeaderView_showSection(void* self, int logicalIndex) {
	((QHeaderView*)self)->showSection(logicalIndex);
}
int QHeaderView_count(void* self) {
	auto&& __r = ((QHeaderView*)self)->count();
	return __r;
}
int QHeaderView_visualIndex(void* self, int logicalIndex) {
	auto&& __r = ((QHeaderView*)self)->visualIndex(logicalIndex);
	return __r;
}
int QHeaderView_logicalIndex(void* self, int visualIndex) {
	auto&& __r = ((QHeaderView*)self)->logicalIndex(visualIndex);
	return __r;
}
void QHeaderView_setSectionsMovable(void* self, int movable) {
	((QHeaderView*)self)->setSectionsMovable(movable);
}
int QHeaderView_sectionsMovable(void* self) {
	auto&& __r = ((QHeaderView*)self)->sectionsMovable();
	return __r;
}
void QHeaderView_setFirstSectionMovable(void* self, int movable) {
	((QHeaderView*)self)->setFirstSectionMovable(movable);
}
int QHeaderView_isFirstSectionMovable(void* self) {
	auto&& __r = ((QHeaderView*)self)->isFirstSectionMovable();
	return __r;
}
void QHeaderView_setSectionsClickable(void* self, int clickable) {
	((QHeaderView*)self)->setSectionsClickable(clickable);
}
int QHeaderView_sectionsClickable(void* self) {
	auto&& __r = ((QHeaderView*)self)->sectionsClickable();
	return __r;
}
void QHeaderView_setHighlightSections(void* self, int highlight) {
	((QHeaderView*)self)->setHighlightSections(highlight);
}
int QHeaderView_highlightSections(void* self) {
	auto&& __r = ((QHeaderView*)self)->highlightSections();
	return __r;
}
int QHeaderView_sectionResizeMode(void* self, int logicalIndex) {
	auto&& __r = ((QHeaderView*)self)->sectionResizeMode(logicalIndex);
	return (int)(__r);
}
void QHeaderView_setSectionResizeMode(void* self, int mode) {
	((QHeaderView*)self)->setSectionResizeMode((QHeaderView::ResizeMode)(mode));
}
void QHeaderView_setSectionResizeMode2(void* self, int logicalIndex, int mode) {
	((QHeaderView*)self)->setSectionResizeMode(logicalIndex, (QHeaderView::ResizeMode)(mode));
}
void QHeaderView_setResizeContentsPrecision(void* self, int precision) {
	((QHeaderView*)self)->setResizeContentsPrecision(precision);
}
int QHeaderView_resizeContentsPrecision(void* self) {
	auto&& __r = ((QHeaderView*)self)->resizeContentsPrecision();
	return __r;
}
int QHeaderView_stretchSectionCount(void* self) {
	auto&& __r = ((QHeaderView*)self)->stretchSectionCount();
	return __r;
}
void QHeaderView_setSortIndicatorShown(void* self, int show) {
	((QHeaderView*)self)->setSortIndicatorShown(show);
}
int QHeaderView_isSortIndicatorShown(void* self) {
	auto&& __r = ((QHeaderView*)self)->isSortIndicatorShown();
	return __r;
}
void QHeaderView_setSortIndicator(void* self, int logicalIndex, int order) {
	((QHeaderView*)self)->setSortIndicator(logicalIndex, (Qt::SortOrder)(order));
}
int QHeaderView_sortIndicatorSection(void* self) {
	auto&& __r = ((QHeaderView*)self)->sortIndicatorSection();
	return __r;
}
int QHeaderView_sortIndicatorOrder(void* self) {
	auto&& __r = ((QHeaderView*)self)->sortIndicatorOrder();
	return (int)(__r);
}
void QHeaderView_setSortIndicatorClearable(void* self, int clearable) {
	((QHeaderView*)self)->setSortIndicatorClearable(clearable);
}
int QHeaderView_isSortIndicatorClearable(void* self) {
	auto&& __r = ((QHeaderView*)self)->isSortIndicatorClearable();
	return __r;
}
int QHeaderView_stretchLastSection(void* self) {
	auto&& __r = ((QHeaderView*)self)->stretchLastSection();
	return __r;
}
void QHeaderView_setStretchLastSection(void* self, int stretch) {
	((QHeaderView*)self)->setStretchLastSection(stretch);
}
int QHeaderView_cascadingSectionResizes(void* self) {
	auto&& __r = ((QHeaderView*)self)->cascadingSectionResizes();
	return __r;
}
void QHeaderView_setCascadingSectionResizes(void* self, int enable) {
	((QHeaderView*)self)->setCascadingSectionResizes(enable);
}
int QHeaderView_defaultSectionSize(void* self) {
	auto&& __r = ((QHeaderView*)self)->defaultSectionSize();
	return __r;
}
void QHeaderView_setDefaultSectionSize(void* self, int size) {
	((QHeaderView*)self)->setDefaultSectionSize(size);
}
void QHeaderView_resetDefaultSectionSize(void* self) {
	((QHeaderView*)self)->resetDefaultSectionSize();
}
int QHeaderView_minimumSectionSize(void* self) {
	auto&& __r = ((QHeaderView*)self)->minimumSectionSize();
	return __r;
}
void QHeaderView_setMinimumSectionSize(void* self, int size) {
	((QHeaderView*)self)->setMinimumSectionSize(size);
}
int QHeaderView_maximumSectionSize(void* self) {
	auto&& __r = ((QHeaderView*)self)->maximumSectionSize();
	return __r;
}
void QHeaderView_setMaximumSectionSize(void* self, int size) {
	((QHeaderView*)self)->setMaximumSectionSize(size);
}
int QHeaderView_defaultAlignment(void* self) {
	auto&& __r = ((QHeaderView*)self)->defaultAlignment();
	return (int)(__r);
}
void QHeaderView_setDefaultAlignment(void* self, int alignment) {
	((QHeaderView*)self)->setDefaultAlignment(QFlags<Qt::AlignmentFlag>(alignment));
}
int QHeaderView_sectionsMoved(void* self) {
	auto&& __r = ((QHeaderView*)self)->sectionsMoved();
	return __r;
}
int QHeaderView_sectionsHidden(void* self) {
	auto&& __r = ((QHeaderView*)self)->sectionsHidden();
	return __r;
}
void* QHeaderView_saveState(void* self) {
	auto&& __r = ((QHeaderView*)self)->saveState();
	return new QByteArray(__r);
}
int QHeaderView_restoreState(void* self, void* state) {
	auto&& __r = ((QHeaderView*)self)->restoreState(*(QByteArray*)state);
	return __r;
}
void QHeaderView_setOffset(void* self, int offset) {
	((QHeaderView*)self)->setOffset(offset);
}
void QHeaderView_setOffsetToSectionPosition(void* self, int visualIndex) {
	((QHeaderView*)self)->setOffsetToSectionPosition(visualIndex);
}
void QHeaderView_setOffsetToLastSection(void* self) {
	((QHeaderView*)self)->setOffsetToLastSection();
}
void QHeaderView_headerDataChanged(void* self, int orientation, int logicalFirst, int logicalLast) {
	((QHeaderView*)self)->headerDataChanged((Qt::Orientation)(orientation), logicalFirst, logicalLast);
}
void QHeaderView_sectionMoved(void* self, int logicalIndex, int oldVisualIndex, int newVisualIndex) {
	((QHeaderView*)self)->sectionMoved(logicalIndex, oldVisualIndex, newVisualIndex);
}
void QHeaderView_sectionResized(void* self, int logicalIndex, int oldSize, int newSize) {
	((QHeaderView*)self)->sectionResized(logicalIndex, oldSize, newSize);
}
void QHeaderView_sectionPressed(void* self, int logicalIndex) {
	((QHeaderView*)self)->sectionPressed(logicalIndex);
}
void QHeaderView_sectionClicked(void* self, int logicalIndex) {
	((QHeaderView*)self)->sectionClicked(logicalIndex);
}
void QHeaderView_sectionEntered(void* self, int logicalIndex) {
	((QHeaderView*)self)->sectionEntered(logicalIndex);
}
void QHeaderView_sectionDoubleClicked(void* self, int logicalIndex) {
	((QHeaderView*)self)->sectionDoubleClicked(logicalIndex);
}
void QHeaderView_sectionCountChanged(void* self, int oldCount, int newCount) {
	((QHeaderView*)self)->sectionCountChanged(oldCount, newCount);
}
void QHeaderView_sectionHandleDoubleClicked(void* self, int logicalIndex) {
	((QHeaderView*)self)->sectionHandleDoubleClicked(logicalIndex);
}
void QHeaderView_geometriesChanged(void* self) {
	((QHeaderView*)self)->geometriesChanged();
}
void QHeaderView_sortIndicatorChanged(void* self, int logicalIndex, int order) {
	((QHeaderView*)self)->sortIndicatorChanged(logicalIndex, (Qt::SortOrder)(order));
}
void QHeaderView_sortIndicatorClearableChanged(void* self, int clearable) {
	((QHeaderView*)self)->sortIndicatorClearableChanged(clearable);
}
void QHeaderView_setTabOrder(void* arg1, void* arg2) {
	QHeaderView::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QHeaderView_mouseGrabber() {
	auto&& __r = QHeaderView::mouseGrabber();
	return (void*)(__r);
}
void* QHeaderView_keyboardGrabber() {
	auto&& __r = QHeaderView::keyboardGrabber();
	return (void*)(__r);
}
void* QHeaderView_find(unsigned long long arg1) {
	auto&& __r = QHeaderView::find(arg1);
	return (void*)(__r);
}
void* QHeaderView_createWindowContainer(void* window) {
	auto&& __r = QHeaderView::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void* QHeaderView_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QHeaderView::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QHeaderView_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QHeaderView::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
double QHeaderView_devicePixelRatioFScale() {
	auto&& __r = QHeaderView::devicePixelRatioFScale();
	return __r;
}
int QHeaderView_encodeMetricF(int metric, double value) {
	auto&& __r = QHeaderView::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
char* QHeaderView_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QHeaderView::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QHeaderView_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QHeaderView::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QHeaderView_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QHeaderView::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QHeaderView_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QHeaderView::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void* QHeaderView_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QHeaderView::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QHeaderView_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QHeaderView::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QHeaderView_connect_sectionMoved(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::sectionMoved, (QHeaderView*)self, [__id](int logicalIndex, int oldVisualIndex, int newVisualIndex){
		if (cppgo_sig_void_int_int_int) ((void(*)(long long, int, int, int))cppgo_sig_void_int_int_int)(__id, logicalIndex, oldVisualIndex, newVisualIndex);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_sectionResized(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::sectionResized, (QHeaderView*)self, [__id](int logicalIndex, int oldSize, int newSize){
		if (cppgo_sig_void_int_int_int) ((void(*)(long long, int, int, int))cppgo_sig_void_int_int_int)(__id, logicalIndex, oldSize, newSize);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_sectionPressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::sectionPressed, (QHeaderView*)self, [__id](int logicalIndex){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, logicalIndex);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_sectionClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::sectionClicked, (QHeaderView*)self, [__id](int logicalIndex){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, logicalIndex);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_sectionEntered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::sectionEntered, (QHeaderView*)self, [__id](int logicalIndex){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, logicalIndex);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_sectionDoubleClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::sectionDoubleClicked, (QHeaderView*)self, [__id](int logicalIndex){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, logicalIndex);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_sectionCountChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::sectionCountChanged, (QHeaderView*)self, [__id](int oldCount, int newCount){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, oldCount, newCount);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_sectionHandleDoubleClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::sectionHandleDoubleClicked, (QHeaderView*)self, [__id](int logicalIndex){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, logicalIndex);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_geometriesChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::geometriesChanged, (QHeaderView*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_sortIndicatorChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::sortIndicatorChanged, (QHeaderView*)self, [__id](int logicalIndex, Qt::SortOrder order){
		if (cppgo_sig_void_int_int) ((void(*)(long long, int, int))cppgo_sig_void_int_int)(__id, logicalIndex, (int)(order));
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_sortIndicatorClearableChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::sortIndicatorClearableChanged, (QHeaderView*)self, [__id](bool clearable){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, clearable);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_pressed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::pressed, (QHeaderView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_clicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::clicked, (QHeaderView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_doubleClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::doubleClicked, (QHeaderView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_activated(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::activated, (QHeaderView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_entered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::entered, (QHeaderView*)self, [__id](const QModelIndex& index){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&index);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_viewportEntered(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::viewportEntered, (QHeaderView*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_iconSizeChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::iconSizeChanged, (QHeaderView*)self, [__id](const QSize& size){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&size);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::windowTitleChanged, (QHeaderView*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::windowIconChanged, (QHeaderView*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::windowIconTextChanged, (QHeaderView*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::customContextMenuRequested, (QHeaderView*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::destroyed, (QHeaderView*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::destroyed, (QHeaderView*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QHeaderView*)self, &QHeaderView::objectNameChanged, (QHeaderView*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
	__conns_QHeaderView[__id] = __c;
}
void QHeaderView__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QHeaderView);
		auto __it = __conns_QHeaderView.find(id);
		if (__it == __conns_QHeaderView.end()) return;
		__c = __it->second;
		__conns_QHeaderView.erase(__it);
	}
	QObject::disconnect(__c);
}
void QHeaderView_delete(void* self) { delete (QHeaderView*)self; }

}

#include <QtWidgets/qproxystyle.h>
#include <QtWidgets/qcommonstyle.h>
#include <QtWidgets/qstyle.h>
#include <QtCore/qstring.h>
#include <QtCore/qobjectdefs.h>
#include <QtWidgets/qstyleoption.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qrect.h>
#include <QtGui/qpalette.h>
#include <QtGui/qpixmap.h>
#include <QtCore/qsize.h>
#include <QtGui/qfontmetrics.h>
#include <QtCore/qpoint.h>
#include <QtWidgets/qsizepolicy.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qicon.h>
#include <QtWidgets/qapplication.h>
#include <QtCore/qcoreevent.h>
#include <QtCore/qobject.h>
#include <QtCore/qanystringview.h>
#include <QtCore/qthread.h>
#include <QtCore/qmetaobject.h>
#include <QtCore/qvariant.h>
#include <QtCore/qbytearray.h>
#include <QtCore/qbindingstorage.h>
#include "gen_qproxystyle.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QProxyStyle__Go : public QProxyStyle {
public:
	long long __goId;
	unsigned char __goOn[31] = {};
	template<typename... A> QProxyStyle__Go(long long id, A&&... a) : QProxyStyle(std::forward<A>(a)...), __goId(id) {}
	using QProxyStyle::sender;
	using QProxyStyle::senderSignalIndex;
	using QProxyStyle::receivers;
	using QProxyStyle::isSignalConnected;
	bool cppgoBase_event(QEvent* __a0) { return QProxyStyle::event(__a0); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QProxyStyle::timerEvent(__a0); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QProxyStyle::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QProxyStyle::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QProxyStyle::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QProxyStyle::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QProxyStyle::metaObject();
	}
	void drawPrimitive(QStyle::PrimitiveElement element, const QStyleOption* option, QPainter* painter, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_void_int_voidp_voidp_voidp) { ((void(*)(long long, long long, int, void*, void*, void*, int*))cppgo_virt_void_int_voidp_voidp_voidp)(__goId, 6098275085281943360LL, (int)(element), (void*)option, (void*)painter, (void*)widget, &__h); if (__h) return; }
		QProxyStyle::drawPrimitive(element, option, painter, widget);
	}
	void drawControl(QStyle::ControlElement element, const QStyleOption* option, QPainter* painter, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void_int_voidp_voidp_voidp) { ((void(*)(long long, long long, int, void*, void*, void*, int*))cppgo_virt_void_int_voidp_voidp_voidp)(__goId, -8697913287309583788LL, (int)(element), (void*)option, (void*)painter, (void*)widget, &__h); if (__h) return; }
		QProxyStyle::drawControl(element, option, painter, widget);
	}
	void drawComplexControl(QStyle::ComplexControl control, const QStyleOptionComplex* option, QPainter* painter, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void_int_voidp_voidp_voidp) { ((void(*)(long long, long long, int, void*, void*, void*, int*))cppgo_virt_void_int_voidp_voidp_voidp)(__goId, -778167327566211030LL, (int)(control), (void*)option, (void*)painter, (void*)widget, &__h); if (__h) return; }
		QProxyStyle::drawComplexControl(control, option, painter, widget);
	}
	void drawItemText(QPainter* painter, const QRect& rect, int flags, const QPalette& pal, bool enabled, const QString& text, QPalette::ColorRole textRole) const override {
		int __h = 0;
		QByteArray text_ba = text.toUtf8();
		if (__goOn[4] && cppgo_virt_void_voidp_voidp_int_voidp_int_constcharp_longlong_int) { ((void(*)(long long, long long, void*, void*, int, void*, int, const char*, long long, int, int*))cppgo_virt_void_voidp_voidp_int_voidp_int_constcharp_longlong_int)(__goId, 2315857538243927747LL, (void*)painter, (void*)&rect, flags, (void*)&pal, enabled, (char*)text_ba.constData(), text_ba.size(), (int)(textRole), &__h); if (__h) return; }
		QProxyStyle::drawItemText(painter, rect, flags, pal, enabled, text, textRole);
	}
	void drawItemPixmap(QPainter* painter, const QRect& rect, int alignment, const QPixmap& pixmap) const override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void_voidp_voidp_int_voidp) { ((void(*)(long long, long long, void*, void*, int, void*, int*))cppgo_virt_void_voidp_voidp_int_voidp)(__goId, -3740415202673381007LL, (void*)painter, (void*)&rect, alignment, (void*)&pixmap, &__h); if (__h) return; }
		QProxyStyle::drawItemPixmap(painter, rect, alignment, pixmap);
	}
	QSize sizeFromContents(QStyle::ContentsType type_, const QStyleOption* option, const QSize& size, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_voidp_int_voidp_voidp_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, void*, void*, int*))cppgo_virt_voidp_int_voidp_voidp_voidp)(__goId, -3316981301817172730LL, (int)(type_), (void*)option, (void*)&size, (void*)widget, &__h); if (__h) return *(QSize*)__r; }
		return QProxyStyle::sizeFromContents(type_, option, size, widget);
	}
	QRect subElementRect(QStyle::SubElement element, const QStyleOption* option, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_voidp_int_voidp_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, void*, int*))cppgo_virt_voidp_int_voidp_voidp)(__goId, -1580195784239416471LL, (int)(element), (void*)option, (void*)widget, &__h); if (__h) return *(QRect*)__r; }
		return QProxyStyle::subElementRect(element, option, widget);
	}
	QRect subControlRect(QStyle::ComplexControl cc, const QStyleOptionComplex* opt, QStyle::SubControl sc, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_voidp_int_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int, void*, int*))cppgo_virt_voidp_int_voidp_int_voidp)(__goId, -8244607819117461060LL, (int)(cc), (void*)opt, (int)(sc), (void*)widget, &__h); if (__h) return *(QRect*)__r; }
		return QProxyStyle::subControlRect(cc, opt, sc, widget);
	}
	QRect itemTextRect(const QFontMetrics& fm, const QRect& r, int flags, bool enabled, const QString& text) const override {
		int __h = 0;
		QByteArray text_ba = text.toUtf8();
		if (__goOn[9] && cppgo_virt_voidp_voidp_voidp_int_int_constcharp_longlong) { void* __r = ((void*(*)(long long, long long, void*, void*, int, int, const char*, long long, int*))cppgo_virt_voidp_voidp_voidp_int_int_constcharp_longlong)(__goId, 8974830244642632823LL, (void*)&fm, (void*)&r, flags, enabled, (char*)text_ba.constData(), text_ba.size(), &__h); if (__h) return *(QRect*)__r; }
		return QProxyStyle::itemTextRect(fm, r, flags, enabled, text);
	}
	QRect itemPixmapRect(const QRect& r, int flags, const QPixmap& pixmap) const override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_voidp_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, void*, int, void*, int*))cppgo_virt_voidp_voidp_int_voidp)(__goId, 3847657773274941773LL, (void*)&r, flags, (void*)&pixmap, &__h); if (__h) return *(QRect*)__r; }
		return QProxyStyle::itemPixmapRect(r, flags, pixmap);
	}
	QStyle::SubControl hitTestComplexControl(QStyle::ComplexControl control, const QStyleOptionComplex* option, const QPoint& pos, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_int_int_voidp_voidp_voidp) { int __r = ((int(*)(long long, long long, int, void*, void*, void*, int*))cppgo_virt_int_int_voidp_voidp_voidp)(__goId, 7058885155526130727LL, (int)(control), (void*)option, (void*)&pos, (void*)widget, &__h); if (__h) return (QStyle::SubControl)(__r); }
		return QProxyStyle::hitTestComplexControl(control, option, pos, widget);
	}
	int styleHint(QStyle::StyleHint hint, const QStyleOption* option, const QWidget* widget, QStyleHintReturn* returnData) const override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_int_voidp_voidp_voidp) { int __r = ((int(*)(long long, long long, int, void*, void*, void*, int*))cppgo_virt_int_int_voidp_voidp_voidp)(__goId, 2017029407462380579LL, (int)(hint), (void*)option, (void*)widget, (void*)returnData, &__h); if (__h) return __r; }
		return QProxyStyle::styleHint(hint, option, widget, returnData);
	}
	int pixelMetric(QStyle::PixelMetric metric, const QStyleOption* option, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_int_int_voidp_voidp) { int __r = ((int(*)(long long, long long, int, void*, void*, int*))cppgo_virt_int_int_voidp_voidp)(__goId, -7621686327863150821LL, (int)(metric), (void*)option, (void*)widget, &__h); if (__h) return __r; }
		return QProxyStyle::pixelMetric(metric, option, widget);
	}
	int layoutSpacing(QSizePolicy::ControlType control1, QSizePolicy::ControlType control2, Qt::Orientation orientation, const QStyleOption* option, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_int_int_int_int_voidp_voidp) { int __r = ((int(*)(long long, long long, int, int, int, void*, void*, int*))cppgo_virt_int_int_int_int_voidp_voidp)(__goId, -4874522738816254324LL, (int)(control1), (int)(control2), (int)(orientation), (void*)option, (void*)widget, &__h); if (__h) return __r; }
		return QProxyStyle::layoutSpacing(control1, control2, orientation, option, widget);
	}
	QIcon standardIcon(QStyle::StandardPixmap standardIcon, const QStyleOption* option, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_voidp_int_voidp_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, void*, int*))cppgo_virt_voidp_int_voidp_voidp)(__goId, -1870022100809264807LL, (int)(standardIcon), (void*)option, (void*)widget, &__h); if (__h) return *(QIcon*)__r; }
		return QProxyStyle::standardIcon(standardIcon, option, widget);
	}
	QPixmap standardPixmap(QStyle::StandardPixmap standardPixmap, const QStyleOption* opt, const QWidget* widget) const override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_voidp_int_voidp_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, void*, int*))cppgo_virt_voidp_int_voidp_voidp)(__goId, -5288619860977043091LL, (int)(standardPixmap), (void*)opt, (void*)widget, &__h); if (__h) return *(QPixmap*)__r; }
		return QProxyStyle::standardPixmap(standardPixmap, opt, widget);
	}
	QPixmap generatedIconPixmap(QIcon::Mode iconMode, const QPixmap& pixmap, const QStyleOption* opt) const override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_voidp_int_voidp_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, void*, int*))cppgo_virt_voidp_int_voidp_voidp)(__goId, -4387834948099211678LL, (int)(iconMode), (void*)&pixmap, (void*)opt, &__h); if (__h) return *(QPixmap*)__r; }
		return QProxyStyle::generatedIconPixmap(iconMode, pixmap, opt);
	}
	QPalette standardPalette() const override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7278964765692882119LL, &__h); if (__h) return *(QPalette*)__r; }
		return QProxyStyle::standardPalette();
	}
	void polish(QWidget* widget) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4633126903937721620LL, (void*)widget, &__h); if (__h) return; }
		QProxyStyle::polish(widget);
	}
	void polish(QPalette& pal) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4263633720130086714LL, (void*)&pal, &__h); if (__h) return; }
		QProxyStyle::polish(pal);
	}
	void polish(QApplication* app) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4263634819641714925LL, (void*)app, &__h); if (__h) return; }
		QProxyStyle::polish(app);
	}
	void unpolish(QWidget* widget) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5547720835653541585LL, (void*)widget, &__h); if (__h) return; }
		QProxyStyle::unpolish(widget);
	}
	void unpolish(QApplication* app) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 265826155645830215LL, (void*)app, &__h); if (__h) return; }
		QProxyStyle::unpolish(app);
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QProxyStyle::event(e);
	}
	bool eventFilter(QObject* watched, QEvent* event) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)watched, (void*)event, &__h); if (__h) return __r; }
		return QProxyStyle::eventFilter(watched, event);
	}
	void timerEvent(QTimerEvent* event) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)event, &__h); if (__h) return; }
		QProxyStyle::timerEvent(event);
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QProxyStyle::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QProxyStyle::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QProxyStyle::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QProxyStyle::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QProxyStyle;
static std::mutex __connsMu_QProxyStyle;

extern "C" {

void* QProxyStyle_base(void* self) { return (void*)static_cast<QCommonStyle*>((QProxyStyle*)self); }
void QProxyStyle_chain(void* self, void** out) {
	QProxyStyle* __p = (QProxyStyle*)self;
	QCommonStyle* __b0 = static_cast<QCommonStyle*>(__p); out[0] = (void*)__b0;
	QStyle* __b1 = static_cast<QStyle*>(__b0); out[1] = (void*)__b1;
	QObject* __b2 = static_cast<QObject*>(__b1); out[2] = (void*)__b2;
}
void QProxyStyle__set_override(void* self, int idx) { ((QProxyStyle__Go*)self)->__goOn[idx] = 1; }
void* QProxyStyle_fromQObject(void* o) { return (void*)qobject_cast<QProxyStyle*>((QObject*)o); }
void* QProxyStyle_new(long long __goid) {
	return new QProxyStyle__Go(__goid);
}
void* QProxyStyle_new2(long long __goid, const char* key, long long key_len) {
	QString key_qstr = QString::fromUtf8(key, key_len);
	return new QProxyStyle__Go(__goid, key_qstr);
}
void* QProxyStyle_new3(long long __goid, void* style) {
	return new QProxyStyle__Go(__goid, (QStyle*)style);
}
char* QProxyStyle_tr(const char* s, long long* __out_len) {
	auto&& __r = QProxyStyle::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QProxyStyle_baseStyle(void* self) {
	auto&& __r = ((QProxyStyle*)self)->baseStyle();
	return (void*)(__r);
}
void QProxyStyle_setBaseStyle(void* self, void* style) {
	((QProxyStyle*)self)->setBaseStyle((QStyle*)style);
}
void* QProxyStyle_visualRect(int direction, void* boundingRect, void* logicalRect) {
	auto&& __r = QProxyStyle::visualRect((Qt::LayoutDirection)(direction), *(QRect*)boundingRect, *(QRect*)logicalRect);
	return new QRect(__r);
}
void* QProxyStyle_visualPos(int direction, void* boundingRect, void* logicalPos) {
	auto&& __r = QProxyStyle::visualPos((Qt::LayoutDirection)(direction), *(QRect*)boundingRect, *(QPoint*)logicalPos);
	return new QPoint(__r);
}
int QProxyStyle_sliderPositionFromValue(int min, int max, int val, int space) {
	auto&& __r = QProxyStyle::sliderPositionFromValue(min, max, val, space);
	return __r;
}
int QProxyStyle_sliderValueFromPosition(int min, int max, int pos, int space) {
	auto&& __r = QProxyStyle::sliderValueFromPosition(min, max, pos, space);
	return __r;
}
int QProxyStyle_visualAlignment(int direction, int alignment) {
	auto&& __r = QProxyStyle::visualAlignment((Qt::LayoutDirection)(direction), QFlags<Qt::AlignmentFlag>(alignment));
	return (int)(__r);
}
void* QProxyStyle_alignedRect(int direction, int alignment, void* size, void* rectangle) {
	auto&& __r = QProxyStyle::alignedRect((Qt::LayoutDirection)(direction), QFlags<Qt::AlignmentFlag>(alignment), *(QSize*)size, *(QRect*)rectangle);
	return new QRect(__r);
}
void* QProxyStyle_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QProxyStyle::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QProxyStyle_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QProxyStyle::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QProxyStyle_sender(void* self) {
	auto&& __r = ((QProxyStyle__Go*)self)->sender();
	return (void*)(__r);
}
int QProxyStyle_senderSignalIndex(void* self) {
	auto&& __r = ((QProxyStyle__Go*)self)->senderSignalIndex();
	return __r;
}
int QProxyStyle_receivers(void* self, const char* signal) {
	auto&& __r = ((QProxyStyle__Go*)self)->receivers(signal);
	return __r;
}
int QProxyStyle_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QProxyStyle__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
char* QProxyStyle_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QProxyStyle::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QProxyStyle_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QProxyStyle::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
int QProxyStyle_sliderPositionFromValue2(int min, int max, int val, int space, int upsideDown) {
	auto&& __r = QProxyStyle::sliderPositionFromValue(min, max, val, space, upsideDown);
	return __r;
}
int QProxyStyle_sliderValueFromPosition2(int min, int max, int pos, int space, int upsideDown) {
	auto&& __r = QProxyStyle::sliderValueFromPosition(min, max, pos, space, upsideDown);
	return __r;
}
void* QProxyStyle_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QProxyStyle::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QProxyStyle_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QProxyStyle::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void* QProxyStyle_metaObject_default(void* self) {
	auto&& __r = ((QProxyStyle*)self)->QProxyStyle::metaObject();
	return (void*)(__r);
}
void QProxyStyle_drawPrimitive_default(void* self, int element, void* option, void* painter, void* widget) {
	((QProxyStyle*)self)->QProxyStyle::drawPrimitive((QStyle::PrimitiveElement)(element), (QStyleOption*)option, (QPainter*)painter, (QWidget*)widget);
}
void QProxyStyle_drawControl_default(void* self, int element, void* option, void* painter, void* widget) {
	((QProxyStyle*)self)->QProxyStyle::drawControl((QStyle::ControlElement)(element), (QStyleOption*)option, (QPainter*)painter, (QWidget*)widget);
}
void QProxyStyle_drawComplexControl_default(void* self, int control, void* option, void* painter, void* widget) {
	((QProxyStyle*)self)->QProxyStyle::drawComplexControl((QStyle::ComplexControl)(control), (QStyleOptionComplex*)option, (QPainter*)painter, (QWidget*)widget);
}
void QProxyStyle_drawItemText_default(void* self, void* painter, void* rect, int flags, void* pal, int enabled, const char* text, long long text_len, int textRole) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	((QProxyStyle*)self)->QProxyStyle::drawItemText((QPainter*)painter, *(QRect*)rect, flags, *(QPalette*)pal, enabled, text_qstr, (QPalette::ColorRole)(textRole));
}
void QProxyStyle_drawItemPixmap_default(void* self, void* painter, void* rect, int alignment, void* pixmap) {
	((QProxyStyle*)self)->QProxyStyle::drawItemPixmap((QPainter*)painter, *(QRect*)rect, alignment, *(QPixmap*)pixmap);
}
void* QProxyStyle_sizeFromContents_default(void* self, int type_, void* option, void* size, void* widget) {
	auto&& __r = ((QProxyStyle*)self)->QProxyStyle::sizeFromContents((QStyle::ContentsType)(type_), (QStyleOption*)option, *(QSize*)size, (QWidget*)widget);
	return new QSize(__r);
}
void* QProxyStyle_subElementRect_default(void* self, int element, void* option, void* widget) {
	auto&& __r = ((QProxyStyle*)self)->QProxyStyle::subElementRect((QStyle::SubElement)(element), (QStyleOption*)option, (QWidget*)widget);
	return new QRect(__r);
}
void* QProxyStyle_subControlRect_default(void* self, int cc, void* opt, int sc, void* widget) {
	auto&& __r = ((QProxyStyle*)self)->QProxyStyle::subControlRect((QStyle::ComplexControl)(cc), (QStyleOptionComplex*)opt, (QStyle::SubControl)(sc), (QWidget*)widget);
	return new QRect(__r);
}
void* QProxyStyle_itemTextRect_default(void* self, void* fm, void* r, int flags, int enabled, const char* text, long long text_len) {
	QString text_qstr = QString::fromUtf8(text, text_len);
	auto&& __r = ((QProxyStyle*)self)->QProxyStyle::itemTextRect(*(QFontMetrics*)fm, *(QRect*)r, flags, enabled, text_qstr);
	return new QRect(__r);
}
void* QProxyStyle_itemPixmapRect_default(void* self, void* r, int flags, void* pixmap) {
	auto&& __r = ((QProxyStyle*)self)->QProxyStyle::itemPixmapRect(*(QRect*)r, flags, *(QPixmap*)pixmap);
	return new QRect(__r);
}
int QProxyStyle_hitTestComplexControl_default(void* self, int control, void* option, void* pos, void* widget) {
	auto&& __r = ((QProxyStyle*)self)->QProxyStyle::hitTestComplexControl((QStyle::ComplexControl)(control), (QStyleOptionComplex*)option, *(QPoint*)pos, (QWidget*)widget);
	return (int)(__r);
}
int QProxyStyle_styleHint_default(void* self, int hint, void* option, void* widget, void* returnData) {
	auto&& __r = ((QProxyStyle*)self)->QProxyStyle::styleHint((QStyle::StyleHint)(hint), (QStyleOption*)option, (QWidget*)widget, (QStyleHintReturn*)returnData);
	return __r;
}
int QProxyStyle_pixelMetric_default(void* self, int metric, void* option, void* widget) {
	auto&& __r = ((QProxyStyle*)self)->QProxyStyle::pixelMetric((QStyle::PixelMetric)(metric), (QStyleOption*)option, (QWidget*)widget);
	return __r;
}
int QProxyStyle_layoutSpacing_default(void* self, int control1, int control2, int orientation, void* option, void* widget) {
	auto&& __r = ((QProxyStyle*)self)->QProxyStyle::layoutSpacing((QSizePolicy::ControlType)(control1), (QSizePolicy::ControlType)(control2), (Qt::Orientation)(orientation), (QStyleOption*)option, (QWidget*)widget);
	return __r;
}
void* QProxyStyle_standardIcon_default(void* self, int standardIcon, void* option, void* widget) {
	auto&& __r = ((QProxyStyle*)self)->QProxyStyle::standardIcon((QStyle::StandardPixmap)(standardIcon), (QStyleOption*)option, (QWidget*)widget);
	return new QIcon(__r);
}
void* QProxyStyle_standardPixmap_default(void* self, int standardPixmap, void* opt, void* widget) {
	auto&& __r = ((QProxyStyle*)self)->QProxyStyle::standardPixmap((QStyle::StandardPixmap)(standardPixmap), (QStyleOption*)opt, (QWidget*)widget);
	return new QPixmap(__r);
}
void* QProxyStyle_generatedIconPixmap_default(void* self, int iconMode, void* pixmap, void* opt) {
	auto&& __r = ((QProxyStyle*)self)->QProxyStyle::generatedIconPixmap((QIcon::Mode)(iconMode), *(QPixmap*)pixmap, (QStyleOption*)opt);
	return new QPixmap(__r);
}
void* QProxyStyle_standardPalette_default(void* self) {
	auto&& __r = ((QProxyStyle*)self)->QProxyStyle::standardPalette();
	return new QPalette(__r);
}
void QProxyStyle_polish_default(void* self, void* widget) {
	((QProxyStyle*)self)->QProxyStyle::polish((QWidget*)widget);
}
void QProxyStyle_unpolish_default(void* self, void* widget) {
	((QProxyStyle*)self)->QProxyStyle::unpolish((QWidget*)widget);
}
int QProxyStyle_event_default(void* self, void* e) {
	auto&& __r = ((QProxyStyle__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
int QProxyStyle_eventFilter_default(void* self, void* watched, void* event) {
	auto&& __r = ((QProxyStyle*)self)->QProxyStyle::eventFilter((QObject*)watched, (QEvent*)event);
	return __r;
}
void QProxyStyle_timerEvent_default(void* self, void* event) {
	((QProxyStyle__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)event);
}
void QProxyStyle_childEvent_default(void* self, void* event) {
	((QProxyStyle__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QProxyStyle_customEvent_default(void* self, void* event) {
	((QProxyStyle__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QProxyStyle_connectNotify_default(void* self, void* signal) {
	((QProxyStyle__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QProxyStyle_disconnectNotify_default(void* self, void* signal) {
	((QProxyStyle__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QProxyStyle_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProxyStyle*)self, &QProxyStyle::destroyed, (QProxyStyle*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProxyStyle);
	__conns_QProxyStyle[__id] = __c;
}
void QProxyStyle_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProxyStyle*)self, &QProxyStyle::destroyed, (QProxyStyle*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProxyStyle);
	__conns_QProxyStyle[__id] = __c;
}
void QProxyStyle_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QProxyStyle*)self, &QProxyStyle::objectNameChanged, (QProxyStyle*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QProxyStyle);
	__conns_QProxyStyle[__id] = __c;
}
void QProxyStyle__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QProxyStyle);
		auto __it = __conns_QProxyStyle.find(id);
		if (__it == __conns_QProxyStyle.end()) return;
		__c = __it->second;
		__conns_QProxyStyle.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QProxyStyle__Go_new(long long id) { return new QProxyStyle__Go(id); }
void QProxyStyle_delete(void* self) { delete (QProxyStyle*)self; }

}

#include <QtWidgets/qtextbrowser.h>
#include <QtWidgets/qtextedit.h>
#include <QtWidgets/qwidget.h>
#include <QtCore/qobjectdefs.h>
#include <QtCore/qstring.h>
#include <QtCore/qurl.h>
#include <QtGui/qtextdocument.h>
#include <QtCore/qvariant.h>
#include <QtCore/qcoreevent.h>
#include <QtGui/qevent.h>
#include <QtGui/qtextcursor.h>
#include <QtCore/qnamespace.h>
#include <QtGui/qcolor.h>
#include <QtGui/qfont.h>
#include <QtGui/qtextformat.h>
#include <QtGui/qtextoption.h>
#include <QtCore/qregularexpression.h>
#include <QtWidgets/qmenu.h>
#include <QtCore/qpoint.h>
#include <QtCore/qrect.h>
#include <QtGui/qpagedpaintdevice.h>
#include <QtCore/qmimedata.h>
#include <QtWidgets/qscrollbar.h>
#include <QtCore/qsize.h>
#include <QtWidgets/qabstractscrollarea.h>
#include <QtCore/qmargins.h>
#include <QtCore/qobject.h>
#include <QtWidgets/qframe.h>
#include <QtGui/qpainter.h>
#include <QtWidgets/qstyleoption.h>
#include <QtWidgets/qstyle.h>
#include <QtGui/qregion.h>
#include <QtGui/qpalette.h>
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
#include "gen_qtextbrowser.h"
#include <unordered_map>
#include <mutex>
#include "gen_purego_cb.h"

class QTextBrowser__Go : public QTextBrowser {
public:
	long long __goId;
	unsigned char __goOn[63] = {};
	template<typename... A> QTextBrowser__Go(long long id, A&&... a) : QTextBrowser(std::forward<A>(a)...), __goId(id) {}
	using QTextBrowser::zoomInF;
	using QTextBrowser::setViewportMargins;
	using QTextBrowser::viewportMargins;
	using QTextBrowser::drawFrame;
	using QTextBrowser::updateMicroFocus;
	using QTextBrowser::create;
	using QTextBrowser::destroy;
	using QTextBrowser::focusNextChild;
	using QTextBrowser::focusPreviousChild;
	using QTextBrowser::sender;
	using QTextBrowser::senderSignalIndex;
	using QTextBrowser::receivers;
	using QTextBrowser::isSignalConnected;
	using QTextBrowser::getDecodedMetricF;
	bool cppgoBase_event(QEvent* __a0) { return QTextBrowser::event(__a0); }
	void cppgoBase_keyPressEvent(QKeyEvent* __a0) { QTextBrowser::keyPressEvent(__a0); }
	void cppgoBase_mouseMoveEvent(QMouseEvent* __a0) { QTextBrowser::mouseMoveEvent(__a0); }
	void cppgoBase_mousePressEvent(QMouseEvent* __a0) { QTextBrowser::mousePressEvent(__a0); }
	void cppgoBase_mouseReleaseEvent(QMouseEvent* __a0) { QTextBrowser::mouseReleaseEvent(__a0); }
	void cppgoBase_focusOutEvent(QFocusEvent* __a0) { QTextBrowser::focusOutEvent(__a0); }
	bool cppgoBase_focusNextPrevChild(bool __a0) { return QTextBrowser::focusNextPrevChild(__a0); }
	void cppgoBase_paintEvent(QPaintEvent* __a0) { QTextBrowser::paintEvent(__a0); }
	void cppgoBase_doSetSource(const QUrl& __a0, int __a1) { QTextBrowser::doSetSource(__a0, (QTextDocument::ResourceType)__a1); }
	void cppgoBase_timerEvent(QTimerEvent* __a0) { QTextBrowser::timerEvent(__a0); }
	void cppgoBase_keyReleaseEvent(QKeyEvent* __a0) { QTextBrowser::keyReleaseEvent(__a0); }
	void cppgoBase_resizeEvent(QResizeEvent* __a0) { QTextBrowser::resizeEvent(__a0); }
	void cppgoBase_mouseDoubleClickEvent(QMouseEvent* __a0) { QTextBrowser::mouseDoubleClickEvent(__a0); }
	void cppgoBase_contextMenuEvent(QContextMenuEvent* __a0) { QTextBrowser::contextMenuEvent(__a0); }
	void cppgoBase_dragEnterEvent(QDragEnterEvent* __a0) { QTextBrowser::dragEnterEvent(__a0); }
	void cppgoBase_dragLeaveEvent(QDragLeaveEvent* __a0) { QTextBrowser::dragLeaveEvent(__a0); }
	void cppgoBase_dragMoveEvent(QDragMoveEvent* __a0) { QTextBrowser::dragMoveEvent(__a0); }
	void cppgoBase_dropEvent(QDropEvent* __a0) { QTextBrowser::dropEvent(__a0); }
	void cppgoBase_focusInEvent(QFocusEvent* __a0) { QTextBrowser::focusInEvent(__a0); }
	void cppgoBase_showEvent(QShowEvent* __a0) { QTextBrowser::showEvent(__a0); }
	void cppgoBase_changeEvent(QEvent* __a0) { QTextBrowser::changeEvent(__a0); }
	void cppgoBase_wheelEvent(QWheelEvent* __a0) { QTextBrowser::wheelEvent(__a0); }
	QMimeData* cppgoBase_createMimeDataFromSelection() { return QTextBrowser::createMimeDataFromSelection(); }
	bool cppgoBase_canInsertFromMimeData(const QMimeData* __a0) { return QTextBrowser::canInsertFromMimeData(__a0); }
	void cppgoBase_insertFromMimeData(const QMimeData* __a0) { QTextBrowser::insertFromMimeData(__a0); }
	void cppgoBase_inputMethodEvent(QInputMethodEvent* __a0) { QTextBrowser::inputMethodEvent(__a0); }
	void cppgoBase_scrollContentsBy(int __a0, int __a1) { QTextBrowser::scrollContentsBy(__a0, __a1); }
	void cppgoBase_doSetTextCursor(const QTextCursor& __a0) { QTextBrowser::doSetTextCursor(__a0); }
	bool cppgoBase_eventFilter(QObject* __a0, QEvent* __a1) { return QTextBrowser::eventFilter(__a0, __a1); }
	bool cppgoBase_viewportEvent(QEvent* __a0) { return QTextBrowser::viewportEvent(__a0); }
	QSize cppgoBase_viewportSizeHint() { return QTextBrowser::viewportSizeHint(); }
	void cppgoBase_initStyleOption(QStyleOptionFrame* __a0) { QTextBrowser::initStyleOption(__a0); }
	void cppgoBase_enterEvent(QEnterEvent* __a0) { QTextBrowser::enterEvent(__a0); }
	void cppgoBase_leaveEvent(QEvent* __a0) { QTextBrowser::leaveEvent(__a0); }
	void cppgoBase_moveEvent(QMoveEvent* __a0) { QTextBrowser::moveEvent(__a0); }
	void cppgoBase_closeEvent(QCloseEvent* __a0) { QTextBrowser::closeEvent(__a0); }
	void cppgoBase_tabletEvent(QTabletEvent* __a0) { QTextBrowser::tabletEvent(__a0); }
	void cppgoBase_actionEvent(QActionEvent* __a0) { QTextBrowser::actionEvent(__a0); }
	void cppgoBase_hideEvent(QHideEvent* __a0) { QTextBrowser::hideEvent(__a0); }
	int cppgoBase_metric(int __a0) { return QTextBrowser::metric((QPaintDevice::PaintDeviceMetric)__a0); }
	void cppgoBase_initPainter(QPainter* __a0) { QTextBrowser::initPainter(__a0); }
	QPaintDevice* cppgoBase_redirected(QPoint* __a0) { return QTextBrowser::redirected(__a0); }
	QPainter* cppgoBase_sharedPainter() { return QTextBrowser::sharedPainter(); }
	void cppgoBase_childEvent(QChildEvent* __a0) { QTextBrowser::childEvent(__a0); }
	void cppgoBase_customEvent(QEvent* __a0) { QTextBrowser::customEvent(__a0); }
	void cppgoBase_connectNotify(const QMetaMethod& __a0) { QTextBrowser::connectNotify(__a0); }
	void cppgoBase_disconnectNotify(const QMetaMethod& __a0) { QTextBrowser::disconnectNotify(__a0); }
	const QMetaObject* metaObject() const override {
		int __h = 0;
		if (__goOn[0] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 1607084412591819503LL, &__h); if (__h) return (const QMetaObject*)__r; }
		return QTextBrowser::metaObject();
	}
	QVariant loadResource(int type_, const QUrl& name) override {
		int __h = 0;
		if (__goOn[1] && cppgo_virt_voidp_int_voidp) { void* __r = ((void*(*)(long long, long long, int, void*, int*))cppgo_virt_voidp_int_voidp)(__goId, -519616681208050567LL, type_, (void*)&name, &__h); if (__h) return *(QVariant*)__r; }
		return QTextBrowser::loadResource(type_, name);
	}
	void backward() override {
		int __h = 0;
		if (__goOn[2] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 2479296857741040516LL, &__h); if (__h) return; }
		QTextBrowser::backward();
	}
	void forward() override {
		int __h = 0;
		if (__goOn[3] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 2190319719396888442LL, &__h); if (__h) return; }
		QTextBrowser::forward();
	}
	void home() override {
		int __h = 0;
		if (__goOn[4] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, 4624382957487889774LL, &__h); if (__h) return; }
		QTextBrowser::home();
	}
	void reload() override {
		int __h = 0;
		if (__goOn[5] && cppgo_virt_void) { ((void(*)(long long, long long, int*))cppgo_virt_void)(__goId, -1731273083434011732LL, &__h); if (__h) return; }
		QTextBrowser::reload();
	}
	bool event(QEvent* e) override {
		int __h = 0;
		if (__goOn[6] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, 1237236456312824991LL, (void*)e, &__h); if (__h) return __r; }
		return QTextBrowser::event(e);
	}
	void keyPressEvent(QKeyEvent* ev) override {
		int __h = 0;
		if (__goOn[7] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1023791977679793185LL, (void*)ev, &__h); if (__h) return; }
		QTextBrowser::keyPressEvent(ev);
	}
	void mouseMoveEvent(QMouseEvent* ev) override {
		int __h = 0;
		if (__goOn[8] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4255494004099806581LL, (void*)ev, &__h); if (__h) return; }
		QTextBrowser::mouseMoveEvent(ev);
	}
	void mousePressEvent(QMouseEvent* ev) override {
		int __h = 0;
		if (__goOn[9] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -332090898078543319LL, (void*)ev, &__h); if (__h) return; }
		QTextBrowser::mousePressEvent(ev);
	}
	void mouseReleaseEvent(QMouseEvent* ev) override {
		int __h = 0;
		if (__goOn[10] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4415920049297795213LL, (void*)ev, &__h); if (__h) return; }
		QTextBrowser::mouseReleaseEvent(ev);
	}
	void focusOutEvent(QFocusEvent* ev) override {
		int __h = 0;
		if (__goOn[11] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -385596010576048993LL, (void*)ev, &__h); if (__h) return; }
		QTextBrowser::focusOutEvent(ev);
	}
	bool focusNextPrevChild(bool next) override {
		int __h = 0;
		if (__goOn[12] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 5898328517505784191LL, next, &__h); if (__h) return __r; }
		return QTextBrowser::focusNextPrevChild(next);
	}
	void paintEvent(QPaintEvent* e) override {
		int __h = 0;
		if (__goOn[13] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5290442163167458465LL, (void*)e, &__h); if (__h) return; }
		QTextBrowser::paintEvent(e);
	}
	void doSetSource(const QUrl& name, QTextDocument::ResourceType type_) override {
		int __h = 0;
		if (__goOn[14] && cppgo_virt_void_voidp_int) { ((void(*)(long long, long long, void*, int, int*))cppgo_virt_void_voidp_int)(__goId, 3615073669235633897LL, (void*)&name, (int)(type_), &__h); if (__h) return; }
		QTextBrowser::doSetSource(name, type_);
	}
	QVariant inputMethodQuery(Qt::InputMethodQuery property) const override {
		int __h = 0;
		if (__goOn[15] && cppgo_virt_voidp_int) { void* __r = ((void*(*)(long long, long long, int, int*))cppgo_virt_voidp_int)(__goId, 7345853097340709638LL, (int)(property), &__h); if (__h) return *(QVariant*)__r; }
		return QTextBrowser::inputMethodQuery(property);
	}
	void timerEvent(QTimerEvent* e) override {
		int __h = 0;
		if (__goOn[16] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -5409927581812374498LL, (void*)e, &__h); if (__h) return; }
		QTextBrowser::timerEvent(e);
	}
	void keyReleaseEvent(QKeyEvent* e) override {
		int __h = 0;
		if (__goOn[17] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 772598475950050299LL, (void*)e, &__h); if (__h) return; }
		QTextBrowser::keyReleaseEvent(e);
	}
	void resizeEvent(QResizeEvent* e) override {
		int __h = 0;
		if (__goOn[18] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4302917285655014661LL, (void*)e, &__h); if (__h) return; }
		QTextBrowser::resizeEvent(e);
	}
	void mouseDoubleClickEvent(QMouseEvent* e) override {
		int __h = 0;
		if (__goOn[19] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -4029073905697023887LL, (void*)e, &__h); if (__h) return; }
		QTextBrowser::mouseDoubleClickEvent(e);
	}
	void contextMenuEvent(QContextMenuEvent* e) override {
		int __h = 0;
		if (__goOn[20] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6698043222227006015LL, (void*)e, &__h); if (__h) return; }
		QTextBrowser::contextMenuEvent(e);
	}
	void dragEnterEvent(QDragEnterEvent* e) override {
		int __h = 0;
		if (__goOn[21] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 206122715919917589LL, (void*)e, &__h); if (__h) return; }
		QTextBrowser::dragEnterEvent(e);
	}
	void dragLeaveEvent(QDragLeaveEvent* e) override {
		int __h = 0;
		if (__goOn[22] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -3139904158373725006LL, (void*)e, &__h); if (__h) return; }
		QTextBrowser::dragLeaveEvent(e);
	}
	void dragMoveEvent(QDragMoveEvent* e) override {
		int __h = 0;
		if (__goOn[23] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1781545137352745978LL, (void*)e, &__h); if (__h) return; }
		QTextBrowser::dragMoveEvent(e);
	}
	void dropEvent(QDropEvent* e) override {
		int __h = 0;
		if (__goOn[24] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 9107058682754591716LL, (void*)e, &__h); if (__h) return; }
		QTextBrowser::dropEvent(e);
	}
	void focusInEvent(QFocusEvent* e) override {
		int __h = 0;
		if (__goOn[25] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 975043767033552552LL, (void*)e, &__h); if (__h) return; }
		QTextBrowser::focusInEvent(e);
	}
	void showEvent(QShowEvent* arg1) override {
		int __h = 0;
		if (__goOn[26] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -6825369174751916LL, (void*)arg1, &__h); if (__h) return; }
		QTextBrowser::showEvent(arg1);
	}
	void changeEvent(QEvent* e) override {
		int __h = 0;
		if (__goOn[27] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4959376063826165943LL, (void*)e, &__h); if (__h) return; }
		QTextBrowser::changeEvent(e);
	}
	void wheelEvent(QWheelEvent* e) override {
		int __h = 0;
		if (__goOn[28] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1919988340934511316LL, (void*)e, &__h); if (__h) return; }
		QTextBrowser::wheelEvent(e);
	}
	QMimeData* createMimeDataFromSelection() const override {
		int __h = 0;
		if (__goOn[29] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 3493558106800165317LL, &__h); if (__h) return (QMimeData*)__r; }
		return QTextBrowser::createMimeDataFromSelection();
	}
	bool canInsertFromMimeData(const QMimeData* source) const override {
		int __h = 0;
		if (__goOn[30] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -2605205276815256896LL, (void*)source, &__h); if (__h) return __r; }
		return QTextBrowser::canInsertFromMimeData(source);
	}
	void insertFromMimeData(const QMimeData* source) override {
		int __h = 0;
		if (__goOn[31] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -466553387102125056LL, (void*)source, &__h); if (__h) return; }
		QTextBrowser::insertFromMimeData(source);
	}
	void inputMethodEvent(QInputMethodEvent* arg1) override {
		int __h = 0;
		if (__goOn[32] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 5438819661942843222LL, (void*)arg1, &__h); if (__h) return; }
		QTextBrowser::inputMethodEvent(arg1);
	}
	void scrollContentsBy(int dx, int dy) override {
		int __h = 0;
		if (__goOn[33] && cppgo_virt_void_int_int) { ((void(*)(long long, long long, int, int, int*))cppgo_virt_void_int_int)(__goId, -6118001150126564745LL, dx, dy, &__h); if (__h) return; }
		QTextBrowser::scrollContentsBy(dx, dy);
	}
	void doSetTextCursor(const QTextCursor& cursor) override {
		int __h = 0;
		if (__goOn[34] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4090661491944147209LL, (void*)&cursor, &__h); if (__h) return; }
		QTextBrowser::doSetTextCursor(cursor);
	}
	QSize minimumSizeHint() const override {
		int __h = 0;
		if (__goOn[35] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 7049264599837532491LL, &__h); if (__h) return *(QSize*)__r; }
		return QTextBrowser::minimumSizeHint();
	}
	QSize sizeHint() const override {
		int __h = 0;
		if (__goOn[36] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 729690261457684101LL, &__h); if (__h) return *(QSize*)__r; }
		return QTextBrowser::sizeHint();
	}
	void setupViewport(QWidget* viewport) override {
		int __h = 0;
		if (__goOn[37] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 4531996381736730688LL, (void*)viewport, &__h); if (__h) return; }
		QTextBrowser::setupViewport(viewport);
	}
	bool eventFilter(QObject* arg1, QEvent* arg2) override {
		int __h = 0;
		if (__goOn[38] && cppgo_virt_int_voidp_voidp) { int __r = ((int(*)(long long, long long, void*, void*, int*))cppgo_virt_int_voidp_voidp)(__goId, 5217623514630858297LL, (void*)arg1, (void*)arg2, &__h); if (__h) return __r; }
		return QTextBrowser::eventFilter(arg1, arg2);
	}
	bool viewportEvent(QEvent* arg1) override {
		int __h = 0;
		if (__goOn[39] && cppgo_virt_int_voidp) { int __r = ((int(*)(long long, long long, void*, int*))cppgo_virt_int_voidp)(__goId, -1971465653061861519LL, (void*)arg1, &__h); if (__h) return __r; }
		return QTextBrowser::viewportEvent(arg1);
	}
	QSize viewportSizeHint() const override {
		int __h = 0;
		if (__goOn[40] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 503579067252048239LL, &__h); if (__h) return *(QSize*)__r; }
		return QTextBrowser::viewportSizeHint();
	}
	void initStyleOption(QStyleOptionFrame* option) const override {
		int __h = 0;
		if (__goOn[41] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -488693769771308531LL, (void*)option, &__h); if (__h) return; }
		QTextBrowser::initStyleOption(option);
	}
	int devType() const override {
		int __h = 0;
		if (__goOn[42] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8659679893868524044LL, &__h); if (__h) return __r; }
		return QTextBrowser::devType();
	}
	void setVisible(bool visible) override {
		int __h = 0;
		if (__goOn[43] && cppgo_virt_void_int) { ((void(*)(long long, long long, int, int*))cppgo_virt_void_int)(__goId, 2110182719375245575LL, visible, &__h); if (__h) return; }
		QTextBrowser::setVisible(visible);
	}
	int heightForWidth(int arg1) const override {
		int __h = 0;
		if (__goOn[44] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 4311534060495737339LL, arg1, &__h); if (__h) return __r; }
		return QTextBrowser::heightForWidth(arg1);
	}
	bool hasHeightForWidth() const override {
		int __h = 0;
		if (__goOn[45] && cppgo_virt_int) { int __r = ((int(*)(long long, long long, int*))cppgo_virt_int)(__goId, -8266762397995423915LL, &__h); if (__h) return __r; }
		return QTextBrowser::hasHeightForWidth();
	}
	QPaintEngine* paintEngine() const override {
		int __h = 0;
		if (__goOn[46] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 6398007404756346381LL, &__h); if (__h) return (QPaintEngine*)__r; }
		return QTextBrowser::paintEngine();
	}
	void enterEvent(QEnterEvent* event) override {
		int __h = 0;
		if (__goOn[47] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6316270346668237895LL, (void*)event, &__h); if (__h) return; }
		QTextBrowser::enterEvent(event);
	}
	void leaveEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[48] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -551070009623711640LL, (void*)event, &__h); if (__h) return; }
		QTextBrowser::leaveEvent(event);
	}
	void moveEvent(QMoveEvent* event) override {
		int __h = 0;
		if (__goOn[49] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -9161949150489525396LL, (void*)event, &__h); if (__h) return; }
		QTextBrowser::moveEvent(event);
	}
	void closeEvent(QCloseEvent* event) override {
		int __h = 0;
		if (__goOn[50] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1989937050019464783LL, (void*)event, &__h); if (__h) return; }
		QTextBrowser::closeEvent(event);
	}
	void tabletEvent(QTabletEvent* event) override {
		int __h = 0;
		if (__goOn[51] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -212068866860701141LL, (void*)event, &__h); if (__h) return; }
		QTextBrowser::tabletEvent(event);
	}
	void actionEvent(QActionEvent* event) override {
		int __h = 0;
		if (__goOn[52] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 1697444140088555581LL, (void*)event, &__h); if (__h) return; }
		QTextBrowser::actionEvent(event);
	}
	void hideEvent(QHideEvent* event) override {
		int __h = 0;
		if (__goOn[53] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -199718707876719129LL, (void*)event, &__h); if (__h) return; }
		QTextBrowser::hideEvent(event);
	}
	bool nativeEvent(const QByteArray& eventType, void* message, qintptr* result) override {
		int __h = 0;
		if (__goOn[54] && cppgo_virt_int_voidp_voidp_longlongp) { int __r = ((int(*)(long long, long long, void*, void*, long long*, int*))cppgo_virt_int_voidp_voidp_longlongp)(__goId, -2039166586246443706LL, (void*)&eventType, (void*)message, (long long*)result, &__h); if (__h) return __r; }
		return QTextBrowser::nativeEvent(eventType, message, result);
	}
	int metric(QPaintDevice::PaintDeviceMetric arg1) const override {
		int __h = 0;
		if (__goOn[55] && cppgo_virt_int_int) { int __r = ((int(*)(long long, long long, int, int*))cppgo_virt_int_int)(__goId, 427629649267311777LL, (int)(arg1), &__h); if (__h) return __r; }
		return QTextBrowser::metric(arg1);
	}
	void initPainter(QPainter* painter) const override {
		int __h = 0;
		if (__goOn[56] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 6477500296980050210LL, (void*)painter, &__h); if (__h) return; }
		QTextBrowser::initPainter(painter);
	}
	QPaintDevice* redirected(QPoint* offset) const override {
		int __h = 0;
		if (__goOn[57] && cppgo_virt_voidp_voidp) { void* __r = ((void*(*)(long long, long long, void*, int*))cppgo_virt_voidp_voidp)(__goId, 4127240633310655048LL, (void*)offset, &__h); if (__h) return (QPaintDevice*)__r; }
		return QTextBrowser::redirected(offset);
	}
	QPainter* sharedPainter() const override {
		int __h = 0;
		if (__goOn[58] && cppgo_virt_voidp) { void* __r = ((void*(*)(long long, long long, int*))cppgo_virt_voidp)(__goId, 5792771328741774991LL, &__h); if (__h) return (QPainter*)__r; }
		return QTextBrowser::sharedPainter();
	}
	void childEvent(QChildEvent* event) override {
		int __h = 0;
		if (__goOn[59] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -1185315367526857833LL, (void*)event, &__h); if (__h) return; }
		QTextBrowser::childEvent(event);
	}
	void customEvent(QEvent* event) override {
		int __h = 0;
		if (__goOn[60] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 3769868229123868474LL, (void*)event, &__h); if (__h) return; }
		QTextBrowser::customEvent(event);
	}
	void connectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[61] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, 7379029982170503008LL, (void*)&signal, &__h); if (__h) return; }
		QTextBrowser::connectNotify(signal);
	}
	void disconnectNotify(const QMetaMethod& signal) override {
		int __h = 0;
		if (__goOn[62] && cppgo_virt_void_voidp) { ((void(*)(long long, long long, void*, int*))cppgo_virt_void_voidp)(__goId, -568369772982215962LL, (void*)&signal, &__h); if (__h) return; }
		QTextBrowser::disconnectNotify(signal);
	}
};

static std::unordered_map<long long, QMetaObject::Connection> __conns_QTextBrowser;
static std::mutex __connsMu_QTextBrowser;

extern "C" {

void* QTextBrowser_base(void* self) { return (void*)static_cast<QTextEdit*>((QTextBrowser*)self); }
void QTextBrowser_chain(void* self, void** out) {
	QTextBrowser* __p = (QTextBrowser*)self;
	QTextEdit* __b0 = static_cast<QTextEdit*>(__p); out[0] = (void*)__b0;
	QAbstractScrollArea* __b1 = static_cast<QAbstractScrollArea*>(__b0); out[1] = (void*)__b1;
	QFrame* __b2 = static_cast<QFrame*>(__b1); out[2] = (void*)__b2;
	QWidget* __b3 = static_cast<QWidget*>(__b2); out[3] = (void*)__b3;
	QObject* __b4 = static_cast<QObject*>(__b3); out[4] = (void*)__b4;
}
void QTextBrowser__set_override(void* self, int idx) { ((QTextBrowser__Go*)self)->__goOn[idx] = 1; }
void* QTextBrowser_fromQObject(void* o) { return (void*)qobject_cast<QTextBrowser*>((QObject*)o); }
void* QTextBrowser_new(long long __goid, void* parent) {
	return new QTextBrowser__Go(__goid, (QWidget*)parent);
}
void* QTextBrowser_new2(long long __goid) {
	return new QTextBrowser__Go(__goid);
}
char* QTextBrowser_tr(const char* s, long long* __out_len) {
	auto&& __r = QTextBrowser::tr(s);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTextBrowser_source(void* self) {
	auto&& __r = ((QTextBrowser*)self)->source();
	return new QUrl(__r);
}
int QTextBrowser_sourceType(void* self) {
	auto&& __r = ((QTextBrowser*)self)->sourceType();
	return (int)(__r);
}
char** QTextBrowser_searchPaths(void* self, long long* __out_len) {
	auto&& __r = ((QTextBrowser*)self)->searchPaths();
	*__out_len = __r.size();
	char** __arr = (char**)malloc(sizeof(char*) * (__r.size() > 0 ? __r.size() : 1));
	long long __i = 0;
	for (auto&& __e : __r) { QByteArray __eb = __e.toUtf8(); char* __s = (char*)malloc(__eb.size() + 1); memcpy(__s, __eb.constData(), __eb.size()); __s[__eb.size()] = 0; __arr[__i++] = __s; }
	return __arr;
}
void QTextBrowser_setSearchPaths(void* self, char** paths, long long paths_len) {
	QList<QString> paths_list; paths_list.reserve(paths_len);
	for (long long __i = 0; __i < paths_len; __i++) paths_list.append(QString::fromUtf8(paths[__i]));
	((QTextBrowser*)self)->setSearchPaths(paths_list);
}
int QTextBrowser_isBackwardAvailable(void* self) {
	auto&& __r = ((QTextBrowser*)self)->isBackwardAvailable();
	return __r;
}
int QTextBrowser_isForwardAvailable(void* self) {
	auto&& __r = ((QTextBrowser*)self)->isForwardAvailable();
	return __r;
}
void QTextBrowser_clearHistory(void* self) {
	((QTextBrowser*)self)->clearHistory();
}
char* QTextBrowser_historyTitle(void* self, int arg1, long long* __out_len) {
	auto&& __r = ((QTextBrowser*)self)->historyTitle(arg1);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void* QTextBrowser_historyUrl(void* self, int arg1) {
	auto&& __r = ((QTextBrowser*)self)->historyUrl(arg1);
	return new QUrl(__r);
}
int QTextBrowser_backwardHistoryCount(void* self) {
	auto&& __r = ((QTextBrowser*)self)->backwardHistoryCount();
	return __r;
}
int QTextBrowser_forwardHistoryCount(void* self) {
	auto&& __r = ((QTextBrowser*)self)->forwardHistoryCount();
	return __r;
}
int QTextBrowser_openExternalLinks(void* self) {
	auto&& __r = ((QTextBrowser*)self)->openExternalLinks();
	return __r;
}
void QTextBrowser_setOpenExternalLinks(void* self, int open) {
	((QTextBrowser*)self)->setOpenExternalLinks(open);
}
int QTextBrowser_openLinks(void* self) {
	auto&& __r = ((QTextBrowser*)self)->openLinks();
	return __r;
}
void QTextBrowser_setOpenLinks(void* self, int open) {
	((QTextBrowser*)self)->setOpenLinks(open);
}
void QTextBrowser_setSource(void* self, void* name) {
	((QTextBrowser*)self)->setSource(*(QUrl*)name);
}
void QTextBrowser_backward(void* self) {
	((QTextBrowser*)self)->backward();
}
void QTextBrowser_forward(void* self) {
	((QTextBrowser*)self)->forward();
}
void QTextBrowser_home(void* self) {
	((QTextBrowser*)self)->home();
}
void QTextBrowser_reload(void* self) {
	((QTextBrowser*)self)->reload();
}
void QTextBrowser_backwardAvailable(void* self, int arg1) {
	((QTextBrowser*)self)->backwardAvailable(arg1);
}
void QTextBrowser_forwardAvailable(void* self, int arg1) {
	((QTextBrowser*)self)->forwardAvailable(arg1);
}
void QTextBrowser_historyChanged(void* self) {
	((QTextBrowser*)self)->historyChanged();
}
void QTextBrowser_sourceChanged(void* self, void* arg1) {
	((QTextBrowser*)self)->sourceChanged(*(QUrl*)arg1);
}
void QTextBrowser_highlighted(void* self, void* arg1) {
	((QTextBrowser*)self)->highlighted(*(QUrl*)arg1);
}
void QTextBrowser_anchorClicked(void* self, void* arg1) {
	((QTextBrowser*)self)->anchorClicked(*(QUrl*)arg1);
}
void QTextBrowser_zoomInF(void* self, float range_) {
	((QTextBrowser__Go*)self)->zoomInF(range_);
}
void QTextBrowser_setViewportMargins(void* self, int left, int top, int right, int bottom) {
	((QTextBrowser__Go*)self)->setViewportMargins(left, top, right, bottom);
}
void QTextBrowser_setViewportMargins2(void* self, void* margins) {
	((QTextBrowser__Go*)self)->setViewportMargins(*(QMargins*)margins);
}
void* QTextBrowser_viewportMargins(void* self) {
	auto&& __r = ((QTextBrowser__Go*)self)->viewportMargins();
	return new QMargins(__r);
}
void QTextBrowser_drawFrame(void* self, void* arg1) {
	((QTextBrowser__Go*)self)->drawFrame((QPainter*)arg1);
}
void QTextBrowser_setTabOrder(void* arg1, void* arg2) {
	QTextBrowser::setTabOrder((QWidget*)arg1, (QWidget*)arg2);
}
void* QTextBrowser_mouseGrabber() {
	auto&& __r = QTextBrowser::mouseGrabber();
	return (void*)(__r);
}
void* QTextBrowser_keyboardGrabber() {
	auto&& __r = QTextBrowser::keyboardGrabber();
	return (void*)(__r);
}
void* QTextBrowser_createWindowContainer(void* window) {
	auto&& __r = QTextBrowser::createWindowContainer((QWindow*)window);
	return (void*)(__r);
}
void QTextBrowser_updateMicroFocus(void* self) {
	((QTextBrowser__Go*)self)->updateMicroFocus();
}
void QTextBrowser_create(void* self) {
	((QTextBrowser__Go*)self)->create();
}
void QTextBrowser_destroy(void* self) {
	((QTextBrowser__Go*)self)->destroy();
}
int QTextBrowser_focusNextChild(void* self) {
	auto&& __r = ((QTextBrowser__Go*)self)->focusNextChild();
	return __r;
}
int QTextBrowser_focusPreviousChild(void* self) {
	auto&& __r = ((QTextBrowser__Go*)self)->focusPreviousChild();
	return __r;
}
void* QTextBrowser_connect(void* sender, const char* signal, void* receiver, const char* member) {
	auto&& __r = QTextBrowser::connect((QObject*)sender, signal, (QObject*)receiver, member);
	return new QMetaObject::Connection(__r);
}
void* QTextBrowser_connect2(void* sender, void* signal, void* receiver, void* method) {
	auto&& __r = QTextBrowser::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method);
	return new QMetaObject::Connection(__r);
}
void* QTextBrowser_sender(void* self) {
	auto&& __r = ((QTextBrowser__Go*)self)->sender();
	return (void*)(__r);
}
int QTextBrowser_senderSignalIndex(void* self) {
	auto&& __r = ((QTextBrowser__Go*)self)->senderSignalIndex();
	return __r;
}
int QTextBrowser_receivers(void* self, const char* signal) {
	auto&& __r = ((QTextBrowser__Go*)self)->receivers(signal);
	return __r;
}
int QTextBrowser_isSignalConnected(void* self, void* signal) {
	auto&& __r = ((QTextBrowser__Go*)self)->isSignalConnected(*(QMetaMethod*)signal);
	return __r;
}
double QTextBrowser_devicePixelRatioFScale() {
	auto&& __r = QTextBrowser::devicePixelRatioFScale();
	return __r;
}
int QTextBrowser_encodeMetricF(int metric, double value) {
	auto&& __r = QTextBrowser::encodeMetricF((QPaintDevice::PaintDeviceMetric)(metric), value);
	return __r;
}
double QTextBrowser_getDecodedMetricF(void* self, int metricA, int metricB) {
	auto&& __r = ((QTextBrowser__Go*)self)->getDecodedMetricF((QPaintDevice::PaintDeviceMetric)(metricA), (QPaintDevice::PaintDeviceMetric)(metricB));
	return __r;
}
char* QTextBrowser_tr2(const char* s, const char* c, long long* __out_len) {
	auto&& __r = QTextBrowser::tr(s, c);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
char* QTextBrowser_tr3(const char* s, const char* c, int n, long long* __out_len) {
	auto&& __r = QTextBrowser::tr(s, c, n);
	QByteArray __ba = __r.toUtf8();
	*__out_len = __ba.size();
	char* __s = (char*)malloc(__ba.size() > 0 ? __ba.size() : 1);
	memcpy(__s, __ba.constData(), __ba.size());
	return __s;
}
void QTextBrowser_setSource2(void* self, void* name, int type_) {
	((QTextBrowser*)self)->setSource(*(QUrl*)name, (QTextDocument::ResourceType)(type_));
}
void* QTextBrowser_createWindowContainer2(void* window, void* parent) {
	auto&& __r = QTextBrowser::createWindowContainer((QWindow*)window, (QWidget*)parent);
	return (void*)(__r);
}
void* QTextBrowser_createWindowContainer3(void* window, void* parent, int flags) {
	auto&& __r = QTextBrowser::createWindowContainer((QWindow*)window, (QWidget*)parent, QFlags<Qt::WindowType>(flags));
	return (void*)(__r);
}
void QTextBrowser_updateMicroFocus2(void* self, int query) {
	((QTextBrowser__Go*)self)->updateMicroFocus((Qt::InputMethodQuery)(query));
}
void QTextBrowser_create2(void* self, unsigned long long arg1) {
	((QTextBrowser__Go*)self)->create(arg1);
}
void QTextBrowser_create3(void* self, unsigned long long arg1, int initializeWindow) {
	((QTextBrowser__Go*)self)->create(arg1, initializeWindow);
}
void QTextBrowser_create4(void* self, unsigned long long arg1, int initializeWindow, int destroyOldWindow) {
	((QTextBrowser__Go*)self)->create(arg1, initializeWindow, destroyOldWindow);
}
void QTextBrowser_destroy2(void* self, int destroyWindow) {
	((QTextBrowser__Go*)self)->destroy(destroyWindow);
}
void QTextBrowser_destroy3(void* self, int destroyWindow, int destroySubWindows) {
	((QTextBrowser__Go*)self)->destroy(destroyWindow, destroySubWindows);
}
void* QTextBrowser_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5) {
	auto&& __r = QTextBrowser::connect((QObject*)sender, signal, (QObject*)receiver, member, (Qt::ConnectionType)(arg5));
	return new QMetaObject::Connection(__r);
}
void* QTextBrowser_connect4(void* sender, void* signal, void* receiver, void* method, int type_) {
	auto&& __r = QTextBrowser::connect((QObject*)sender, *(QMetaMethod*)signal, (QObject*)receiver, *(QMetaMethod*)method, (Qt::ConnectionType)(type_));
	return new QMetaObject::Connection(__r);
}
void QTextBrowser_backward_default(void* self) {
	((QTextBrowser*)self)->QTextBrowser::backward();
}
void QTextBrowser_forward_default(void* self) {
	((QTextBrowser*)self)->QTextBrowser::forward();
}
void QTextBrowser_home_default(void* self) {
	((QTextBrowser*)self)->QTextBrowser::home();
}
void QTextBrowser_reload_default(void* self) {
	((QTextBrowser*)self)->QTextBrowser::reload();
}
void* QTextBrowser_metaObject_default(void* self) {
	auto&& __r = ((QTextBrowser*)self)->QTextBrowser::metaObject();
	return (void*)(__r);
}
void* QTextBrowser_loadResource_default(void* self, int type_, void* name) {
	auto&& __r = ((QTextBrowser*)self)->QTextBrowser::loadResource(type_, *(QUrl*)name);
	return new QVariant(__r);
}
int QTextBrowser_event_default(void* self, void* e) {
	auto&& __r = ((QTextBrowser__Go*)self)->cppgoBase_event((QEvent*)e);
	return __r;
}
void QTextBrowser_keyPressEvent_default(void* self, void* ev) {
	((QTextBrowser__Go*)self)->cppgoBase_keyPressEvent((QKeyEvent*)ev);
}
void QTextBrowser_mouseMoveEvent_default(void* self, void* ev) {
	((QTextBrowser__Go*)self)->cppgoBase_mouseMoveEvent((QMouseEvent*)ev);
}
void QTextBrowser_mousePressEvent_default(void* self, void* ev) {
	((QTextBrowser__Go*)self)->cppgoBase_mousePressEvent((QMouseEvent*)ev);
}
void QTextBrowser_mouseReleaseEvent_default(void* self, void* ev) {
	((QTextBrowser__Go*)self)->cppgoBase_mouseReleaseEvent((QMouseEvent*)ev);
}
void QTextBrowser_focusOutEvent_default(void* self, void* ev) {
	((QTextBrowser__Go*)self)->cppgoBase_focusOutEvent((QFocusEvent*)ev);
}
int QTextBrowser_focusNextPrevChild_default(void* self, int next) {
	auto&& __r = ((QTextBrowser__Go*)self)->cppgoBase_focusNextPrevChild(next);
	return __r;
}
void QTextBrowser_paintEvent_default(void* self, void* e) {
	((QTextBrowser__Go*)self)->cppgoBase_paintEvent((QPaintEvent*)e);
}
void QTextBrowser_doSetSource_default(void* self, void* name, int type_) {
	((QTextBrowser__Go*)self)->cppgoBase_doSetSource(*(QUrl*)name, type_);
}
void* QTextBrowser_inputMethodQuery_default(void* self, int property) {
	auto&& __r = ((QTextBrowser*)self)->QTextBrowser::inputMethodQuery((Qt::InputMethodQuery)(property));
	return new QVariant(__r);
}
void QTextBrowser_timerEvent_default(void* self, void* e) {
	((QTextBrowser__Go*)self)->cppgoBase_timerEvent((QTimerEvent*)e);
}
void QTextBrowser_keyReleaseEvent_default(void* self, void* e) {
	((QTextBrowser__Go*)self)->cppgoBase_keyReleaseEvent((QKeyEvent*)e);
}
void QTextBrowser_resizeEvent_default(void* self, void* e) {
	((QTextBrowser__Go*)self)->cppgoBase_resizeEvent((QResizeEvent*)e);
}
void QTextBrowser_mouseDoubleClickEvent_default(void* self, void* e) {
	((QTextBrowser__Go*)self)->cppgoBase_mouseDoubleClickEvent((QMouseEvent*)e);
}
void QTextBrowser_contextMenuEvent_default(void* self, void* e) {
	((QTextBrowser__Go*)self)->cppgoBase_contextMenuEvent((QContextMenuEvent*)e);
}
void QTextBrowser_dragEnterEvent_default(void* self, void* e) {
	((QTextBrowser__Go*)self)->cppgoBase_dragEnterEvent((QDragEnterEvent*)e);
}
void QTextBrowser_dragLeaveEvent_default(void* self, void* e) {
	((QTextBrowser__Go*)self)->cppgoBase_dragLeaveEvent((QDragLeaveEvent*)e);
}
void QTextBrowser_dragMoveEvent_default(void* self, void* e) {
	((QTextBrowser__Go*)self)->cppgoBase_dragMoveEvent((QDragMoveEvent*)e);
}
void QTextBrowser_dropEvent_default(void* self, void* e) {
	((QTextBrowser__Go*)self)->cppgoBase_dropEvent((QDropEvent*)e);
}
void QTextBrowser_focusInEvent_default(void* self, void* e) {
	((QTextBrowser__Go*)self)->cppgoBase_focusInEvent((QFocusEvent*)e);
}
void QTextBrowser_showEvent_default(void* self, void* arg1) {
	((QTextBrowser__Go*)self)->cppgoBase_showEvent((QShowEvent*)arg1);
}
void QTextBrowser_changeEvent_default(void* self, void* e) {
	((QTextBrowser__Go*)self)->cppgoBase_changeEvent((QEvent*)e);
}
void QTextBrowser_wheelEvent_default(void* self, void* e) {
	((QTextBrowser__Go*)self)->cppgoBase_wheelEvent((QWheelEvent*)e);
}
void* QTextBrowser_createMimeDataFromSelection_default(void* self) {
	auto&& __r = ((QTextBrowser__Go*)self)->cppgoBase_createMimeDataFromSelection();
	return (void*)(__r);
}
int QTextBrowser_canInsertFromMimeData_default(void* self, void* source) {
	auto&& __r = ((QTextBrowser__Go*)self)->cppgoBase_canInsertFromMimeData((QMimeData*)source);
	return __r;
}
void QTextBrowser_insertFromMimeData_default(void* self, void* source) {
	((QTextBrowser__Go*)self)->cppgoBase_insertFromMimeData((QMimeData*)source);
}
void QTextBrowser_inputMethodEvent_default(void* self, void* arg1) {
	((QTextBrowser__Go*)self)->cppgoBase_inputMethodEvent((QInputMethodEvent*)arg1);
}
void QTextBrowser_scrollContentsBy_default(void* self, int dx, int dy) {
	((QTextBrowser__Go*)self)->cppgoBase_scrollContentsBy(dx, dy);
}
void QTextBrowser_doSetTextCursor_default(void* self, void* cursor) {
	((QTextBrowser__Go*)self)->cppgoBase_doSetTextCursor(*(QTextCursor*)cursor);
}
void* QTextBrowser_minimumSizeHint_default(void* self) {
	auto&& __r = ((QTextBrowser*)self)->QTextBrowser::minimumSizeHint();
	return new QSize(__r);
}
void* QTextBrowser_sizeHint_default(void* self) {
	auto&& __r = ((QTextBrowser*)self)->QTextBrowser::sizeHint();
	return new QSize(__r);
}
void QTextBrowser_setupViewport_default(void* self, void* viewport) {
	((QTextBrowser*)self)->QTextBrowser::setupViewport((QWidget*)viewport);
}
int QTextBrowser_eventFilter_default(void* self, void* arg1, void* arg2) {
	auto&& __r = ((QTextBrowser__Go*)self)->cppgoBase_eventFilter((QObject*)arg1, (QEvent*)arg2);
	return __r;
}
int QTextBrowser_viewportEvent_default(void* self, void* arg1) {
	auto&& __r = ((QTextBrowser__Go*)self)->cppgoBase_viewportEvent((QEvent*)arg1);
	return __r;
}
void* QTextBrowser_viewportSizeHint_default(void* self) {
	auto&& __r = ((QTextBrowser__Go*)self)->cppgoBase_viewportSizeHint();
	return new QSize(__r);
}
void QTextBrowser_initStyleOption_default(void* self, void* option) {
	((QTextBrowser__Go*)self)->cppgoBase_initStyleOption((QStyleOptionFrame*)option);
}
int QTextBrowser_devType_default(void* self) {
	auto&& __r = ((QTextBrowser*)self)->QTextBrowser::devType();
	return __r;
}
void QTextBrowser_setVisible_default(void* self, int visible) {
	((QTextBrowser*)self)->QTextBrowser::setVisible(visible);
}
int QTextBrowser_heightForWidth_default(void* self, int arg1) {
	auto&& __r = ((QTextBrowser*)self)->QTextBrowser::heightForWidth(arg1);
	return __r;
}
int QTextBrowser_hasHeightForWidth_default(void* self) {
	auto&& __r = ((QTextBrowser*)self)->QTextBrowser::hasHeightForWidth();
	return __r;
}
void* QTextBrowser_paintEngine_default(void* self) {
	auto&& __r = ((QTextBrowser*)self)->QTextBrowser::paintEngine();
	return (void*)(__r);
}
void QTextBrowser_enterEvent_default(void* self, void* event) {
	((QTextBrowser__Go*)self)->cppgoBase_enterEvent((QEnterEvent*)event);
}
void QTextBrowser_leaveEvent_default(void* self, void* event) {
	((QTextBrowser__Go*)self)->cppgoBase_leaveEvent((QEvent*)event);
}
void QTextBrowser_moveEvent_default(void* self, void* event) {
	((QTextBrowser__Go*)self)->cppgoBase_moveEvent((QMoveEvent*)event);
}
void QTextBrowser_closeEvent_default(void* self, void* event) {
	((QTextBrowser__Go*)self)->cppgoBase_closeEvent((QCloseEvent*)event);
}
void QTextBrowser_tabletEvent_default(void* self, void* event) {
	((QTextBrowser__Go*)self)->cppgoBase_tabletEvent((QTabletEvent*)event);
}
void QTextBrowser_actionEvent_default(void* self, void* event) {
	((QTextBrowser__Go*)self)->cppgoBase_actionEvent((QActionEvent*)event);
}
void QTextBrowser_hideEvent_default(void* self, void* event) {
	((QTextBrowser__Go*)self)->cppgoBase_hideEvent((QHideEvent*)event);
}
int QTextBrowser_metric_default(void* self, int arg1) {
	auto&& __r = ((QTextBrowser__Go*)self)->cppgoBase_metric(arg1);
	return __r;
}
void QTextBrowser_initPainter_default(void* self, void* painter) {
	((QTextBrowser__Go*)self)->cppgoBase_initPainter((QPainter*)painter);
}
void* QTextBrowser_redirected_default(void* self, void* offset) {
	auto&& __r = ((QTextBrowser__Go*)self)->cppgoBase_redirected((QPoint*)offset);
	return (void*)(__r);
}
void* QTextBrowser_sharedPainter_default(void* self) {
	auto&& __r = ((QTextBrowser__Go*)self)->cppgoBase_sharedPainter();
	return (void*)(__r);
}
void QTextBrowser_childEvent_default(void* self, void* event) {
	((QTextBrowser__Go*)self)->cppgoBase_childEvent((QChildEvent*)event);
}
void QTextBrowser_customEvent_default(void* self, void* event) {
	((QTextBrowser__Go*)self)->cppgoBase_customEvent((QEvent*)event);
}
void QTextBrowser_connectNotify_default(void* self, void* signal) {
	((QTextBrowser__Go*)self)->cppgoBase_connectNotify(*(QMetaMethod*)signal);
}
void QTextBrowser_disconnectNotify_default(void* self, void* signal) {
	((QTextBrowser__Go*)self)->cppgoBase_disconnectNotify(*(QMetaMethod*)signal);
}
void QTextBrowser_connect_backwardAvailable(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::backwardAvailable, (QTextBrowser*)self, [__id](bool arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_forwardAvailable(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::forwardAvailable, (QTextBrowser*)self, [__id](bool arg1){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_historyChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::historyChanged, (QTextBrowser*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_sourceChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::sourceChanged, (QTextBrowser*)self, [__id](const QUrl& arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_highlighted(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::highlighted, (QTextBrowser*)self, [__id](const QUrl& arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_anchorClicked(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::anchorClicked, (QTextBrowser*)self, [__id](const QUrl& arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_textChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::textChanged, (QTextBrowser*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_undoAvailable(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::undoAvailable, (QTextBrowser*)self, [__id](bool b){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, b);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_redoAvailable(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::redoAvailable, (QTextBrowser*)self, [__id](bool b){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, b);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_currentCharFormatChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::currentCharFormatChanged, (QTextBrowser*)self, [__id](const QTextCharFormat& format){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&format);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_copyAvailable(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::copyAvailable, (QTextBrowser*)self, [__id](bool b){
		if (cppgo_sig_void_int) ((void(*)(long long, int))cppgo_sig_void_int)(__id, b);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_selectionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::selectionChanged, (QTextBrowser*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_cursorPositionChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::cursorPositionChanged, (QTextBrowser*)self, [__id](){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_windowTitleChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::windowTitleChanged, (QTextBrowser*)self, [__id](const QString& title){
		QByteArray title_ba = title.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)title_ba.constData(), title_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_windowIconChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::windowIconChanged, (QTextBrowser*)self, [__id](const QIcon& icon){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&icon);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_windowIconTextChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::windowIconTextChanged, (QTextBrowser*)self, [__id](const QString& iconText){
		QByteArray iconText_ba = iconText.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)iconText_ba.constData(), iconText_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_customContextMenuRequested(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::customContextMenuRequested, (QTextBrowser*)self, [__id](const QPoint& pos){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)&pos);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_destroyed(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::destroyed, (QTextBrowser*)self, [__id](QObject* arg1){
		if (cppgo_sig_void) ((void(*)(long long))cppgo_sig_void)(__id);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_destroyed2(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::destroyed, (QTextBrowser*)self, [__id](QObject* arg1){
		if (cppgo_sig_void_voidp) ((void(*)(long long, void*))cppgo_sig_void_voidp)(__id, (void*)arg1);
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser_connect_objectNameChanged(void* self, long long __id) {
	QMetaObject::Connection __c = QObject::connect((QTextBrowser*)self, &QTextBrowser::objectNameChanged, (QTextBrowser*)self, [__id](const QString& objectName){
		QByteArray objectName_ba = objectName.toUtf8();
		if (cppgo_sig_void_constcharp_longlong) ((void(*)(long long, const char*, long long))cppgo_sig_void_constcharp_longlong)(__id, (char*)objectName_ba.constData(), objectName_ba.size());
	});
	std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
	__conns_QTextBrowser[__id] = __c;
}
void QTextBrowser__disconnect(long long id) {
	QMetaObject::Connection __c;
	{
		std::lock_guard<std::mutex> __lk(__connsMu_QTextBrowser);
		auto __it = __conns_QTextBrowser.find(id);
		if (__it == __conns_QTextBrowser.end()) return;
		__c = __it->second;
		__conns_QTextBrowser.erase(__it);
	}
	QObject::disconnect(__c);
}
void* QTextBrowser__Go_new(long long id) { return new QTextBrowser__Go(id); }
void QTextBrowser_delete(void* self) { delete (QTextBrowser*)self; }

}

