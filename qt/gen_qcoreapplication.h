#pragma once
#include <stddef.h>
#ifdef __cplusplus
extern "C" {
#endif

void* QCoreApplication_base(void* self);
void* QCoreApplication_fromQObject(void* o);
char* QCoreApplication_tr(const char* s, long long* __out_len);
char** QCoreApplication_arguments(long long* __out_len);
void QCoreApplication_setAttribute(int attribute);
int QCoreApplication_testAttribute(int attribute);
void QCoreApplication_setOrganizationDomain(const char* orgDomain, long long orgDomain_len);
char* QCoreApplication_organizationDomain(long long* __out_len);
void QCoreApplication_setOrganizationName(const char* orgName, long long orgName_len);
char* QCoreApplication_organizationName(long long* __out_len);
void QCoreApplication_setApplicationName(const char* application, long long application_len);
char* QCoreApplication_applicationName(long long* __out_len);
void QCoreApplication_setApplicationVersion(const char* version, long long version_len);
char* QCoreApplication_applicationVersion(long long* __out_len);
void QCoreApplication_setSetuidAllowed(int allow);
int QCoreApplication_isSetuidAllowed();
void* QCoreApplication_instance();
int QCoreApplication_instanceExists();
int QCoreApplication_exec();
void QCoreApplication_processEvents();
void QCoreApplication_processEvents2(int flags, int maxtime);
void QCoreApplication_processEvents3(int flags, void* deadline);
int QCoreApplication_sendEvent(void* receiver, void* event);
void QCoreApplication_postEvent(void* receiver, void* event);
void QCoreApplication_sendPostedEvents();
void QCoreApplication_removePostedEvents(void* receiver);
void* QCoreApplication_eventDispatcher();
void QCoreApplication_setEventDispatcher(void* eventDispatcher);
int QCoreApplication_notify(void* self, void* arg1, void* arg2);
int QCoreApplication_startingUp();
int QCoreApplication_closingDown();
char* QCoreApplication_applicationDirPath(long long* __out_len);
char* QCoreApplication_applicationFilePath(long long* __out_len);
long long QCoreApplication_applicationPid();
int QCoreApplication_checkPermission(void* self, void* permission);
void QCoreApplication_setLibraryPaths(char** arg1, long long arg1_len);
char** QCoreApplication_libraryPaths(long long* __out_len);
void QCoreApplication_addLibraryPath(const char* arg1, long long arg1_len);
void QCoreApplication_removeLibraryPath(const char* arg1, long long arg1_len);
int QCoreApplication_installTranslator(void* messageFile);
int QCoreApplication_removeTranslator(void* messageFile);
char* QCoreApplication_translate(const char* context, const char* key, long long* __out_len);
void QCoreApplication_installNativeEventFilter(void* self, void* filterObj);
void QCoreApplication_removeNativeEventFilter(void* self, void* filterObj);
int QCoreApplication_isQuitLockEnabled();
void QCoreApplication_setQuitLockEnabled(int enabled);
void QCoreApplication_quit();
void QCoreApplication_exit();
void QCoreApplication_organizationNameChanged(void* self);
void QCoreApplication_organizationDomainChanged(void* self);
void QCoreApplication_applicationNameChanged(void* self);
void QCoreApplication_applicationVersionChanged(void* self);
void* QCoreApplication_connect(void* sender, const char* signal, void* receiver, const char* member);
void* QCoreApplication_connect2(void* sender, void* signal, void* receiver, void* method);
char* QCoreApplication_tr2(const char* s, const char* c, long long* __out_len);
char* QCoreApplication_tr3(const char* s, const char* c, int n, long long* __out_len);
void QCoreApplication_setAttribute2(int attribute, int on);
void QCoreApplication_processEvents4(int flags);
void QCoreApplication_postEvent2(void* receiver, void* event, int priority);
void QCoreApplication_sendPostedEvents2(void* receiver);
void QCoreApplication_sendPostedEvents3(void* receiver, int event_type);
void QCoreApplication_removePostedEvents2(void* receiver, int eventType);
char* QCoreApplication_translate2(const char* context, const char* key, const char* disambiguation, long long* __out_len);
char* QCoreApplication_translate3(const char* context, const char* key, const char* disambiguation, int n, long long* __out_len);
void QCoreApplication_exit2(int retcode);
void* QCoreApplication_connect3(void* sender, const char* signal, void* receiver, const char* member, int arg5);
void* QCoreApplication_connect4(void* sender, void* signal, void* receiver, void* method, int type_);
void QCoreApplication_connect_aboutToQuit(void* self, long long id);
void QCoreApplication_connect_organizationNameChanged(void* self, long long id);
void QCoreApplication_connect_organizationDomainChanged(void* self, long long id);
void QCoreApplication_connect_applicationNameChanged(void* self, long long id);
void QCoreApplication_connect_applicationVersionChanged(void* self, long long id);
void QCoreApplication_connect_destroyed(void* self, long long id);
void QCoreApplication_connect_destroyed2(void* self, long long id);
void QCoreApplication_connect_objectNameChanged(void* self, long long id);
void QCoreApplication__disconnect(long long id);
void* QCoreApplication_new_args(int* argc, char** argv);
void QCoreApplication_delete(void* self);

#ifdef __cplusplus
}
#endif
