package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QCoreApplication_base func(uintptr) uintptr
	addr_QCoreApplication_tr uintptr
	addr_QCoreApplication_arguments uintptr
	addr_QCoreApplication_setAttribute uintptr
	addr_QCoreApplication_testAttribute uintptr
	addr_QCoreApplication_setOrganizationDomain uintptr
	addr_QCoreApplication_organizationDomain uintptr
	addr_QCoreApplication_setOrganizationName uintptr
	addr_QCoreApplication_organizationName uintptr
	addr_QCoreApplication_setApplicationName uintptr
	addr_QCoreApplication_applicationName uintptr
	addr_QCoreApplication_setApplicationVersion uintptr
	addr_QCoreApplication_applicationVersion uintptr
	addr_QCoreApplication_setSetuidAllowed uintptr
	addr_QCoreApplication_isSetuidAllowed uintptr
	addr_QCoreApplication_instance uintptr
	addr_QCoreApplication_instanceExists uintptr
	addr_QCoreApplication_exec uintptr
	addr_QCoreApplication_processEvents uintptr
	addr_QCoreApplication_processEvents2 uintptr
	addr_QCoreApplication_processEvents3 uintptr
	addr_QCoreApplication_sendEvent uintptr
	addr_QCoreApplication_postEvent uintptr
	addr_QCoreApplication_sendPostedEvents uintptr
	addr_QCoreApplication_removePostedEvents uintptr
	addr_QCoreApplication_eventDispatcher uintptr
	addr_QCoreApplication_setEventDispatcher uintptr
	addr_QCoreApplication_notify uintptr
	addr_QCoreApplication_startingUp uintptr
	addr_QCoreApplication_closingDown uintptr
	addr_QCoreApplication_applicationDirPath uintptr
	addr_QCoreApplication_applicationFilePath uintptr
	addr_QCoreApplication_applicationPid uintptr
	addr_QCoreApplication_checkPermission uintptr
	addr_QCoreApplication_setLibraryPaths uintptr
	addr_QCoreApplication_libraryPaths uintptr
	addr_QCoreApplication_addLibraryPath uintptr
	addr_QCoreApplication_removeLibraryPath uintptr
	addr_QCoreApplication_installTranslator uintptr
	addr_QCoreApplication_removeTranslator uintptr
	addr_QCoreApplication_translate uintptr
	addr_QCoreApplication_installNativeEventFilter uintptr
	addr_QCoreApplication_removeNativeEventFilter uintptr
	addr_QCoreApplication_isQuitLockEnabled uintptr
	addr_QCoreApplication_setQuitLockEnabled uintptr
	addr_QCoreApplication_quit uintptr
	addr_QCoreApplication_exit uintptr
	addr_QCoreApplication_organizationNameChanged uintptr
	addr_QCoreApplication_organizationDomainChanged uintptr
	addr_QCoreApplication_applicationNameChanged uintptr
	addr_QCoreApplication_applicationVersionChanged uintptr
	addr_QCoreApplication_connect uintptr
	addr_QCoreApplication_connect2 uintptr
	addr_QCoreApplication_tr2 uintptr
	addr_QCoreApplication_tr3 uintptr
	addr_QCoreApplication_setAttribute2 uintptr
	addr_QCoreApplication_processEvents4 uintptr
	addr_QCoreApplication_postEvent2 uintptr
	addr_QCoreApplication_sendPostedEvents2 uintptr
	addr_QCoreApplication_sendPostedEvents3 uintptr
	addr_QCoreApplication_removePostedEvents2 uintptr
	addr_QCoreApplication_translate2 uintptr
	addr_QCoreApplication_translate3 uintptr
	addr_QCoreApplication_exit2 uintptr
	addr_QCoreApplication_connect3 uintptr
	addr_QCoreApplication_connect4 uintptr
	pg_QCoreApplication_connect_aboutToQuit func(uintptr, int64)
	pg_QCoreApplication_connect_organizationNameChanged func(uintptr, int64)
	pg_QCoreApplication_connect_organizationDomainChanged func(uintptr, int64)
	pg_QCoreApplication_connect_applicationNameChanged func(uintptr, int64)
	pg_QCoreApplication_connect_applicationVersionChanged func(uintptr, int64)
	pg_QCoreApplication_connect_destroyed func(uintptr, int64)
	pg_QCoreApplication_connect_destroyed2 func(uintptr, int64)
	pg_QCoreApplication_connect_objectNameChanged func(uintptr, int64)
	pg_QCoreApplication__disconnect func(int64)
	pg_QCoreApplication_new_args func(uintptr, uintptr) uintptr
	pg_QCoreApplication_delete func(uintptr)
)

func init() {
	reg(&pg_QCoreApplication_base, "QCoreApplication_base")
	symaddr(&addr_QCoreApplication_tr, "QCoreApplication_tr")
	symaddr(&addr_QCoreApplication_arguments, "QCoreApplication_arguments")
	symaddr(&addr_QCoreApplication_setAttribute, "QCoreApplication_setAttribute")
	symaddr(&addr_QCoreApplication_testAttribute, "QCoreApplication_testAttribute")
	symaddr(&addr_QCoreApplication_setOrganizationDomain, "QCoreApplication_setOrganizationDomain")
	symaddr(&addr_QCoreApplication_organizationDomain, "QCoreApplication_organizationDomain")
	symaddr(&addr_QCoreApplication_setOrganizationName, "QCoreApplication_setOrganizationName")
	symaddr(&addr_QCoreApplication_organizationName, "QCoreApplication_organizationName")
	symaddr(&addr_QCoreApplication_setApplicationName, "QCoreApplication_setApplicationName")
	symaddr(&addr_QCoreApplication_applicationName, "QCoreApplication_applicationName")
	symaddr(&addr_QCoreApplication_setApplicationVersion, "QCoreApplication_setApplicationVersion")
	symaddr(&addr_QCoreApplication_applicationVersion, "QCoreApplication_applicationVersion")
	symaddr(&addr_QCoreApplication_setSetuidAllowed, "QCoreApplication_setSetuidAllowed")
	symaddr(&addr_QCoreApplication_isSetuidAllowed, "QCoreApplication_isSetuidAllowed")
	symaddr(&addr_QCoreApplication_instance, "QCoreApplication_instance")
	symaddr(&addr_QCoreApplication_instanceExists, "QCoreApplication_instanceExists")
	symaddr(&addr_QCoreApplication_exec, "QCoreApplication_exec")
	symaddr(&addr_QCoreApplication_processEvents, "QCoreApplication_processEvents")
	symaddr(&addr_QCoreApplication_processEvents2, "QCoreApplication_processEvents2")
	symaddr(&addr_QCoreApplication_processEvents3, "QCoreApplication_processEvents3")
	symaddr(&addr_QCoreApplication_sendEvent, "QCoreApplication_sendEvent")
	symaddr(&addr_QCoreApplication_postEvent, "QCoreApplication_postEvent")
	symaddr(&addr_QCoreApplication_sendPostedEvents, "QCoreApplication_sendPostedEvents")
	symaddr(&addr_QCoreApplication_removePostedEvents, "QCoreApplication_removePostedEvents")
	symaddr(&addr_QCoreApplication_eventDispatcher, "QCoreApplication_eventDispatcher")
	symaddr(&addr_QCoreApplication_setEventDispatcher, "QCoreApplication_setEventDispatcher")
	symaddr(&addr_QCoreApplication_notify, "QCoreApplication_notify")
	symaddr(&addr_QCoreApplication_startingUp, "QCoreApplication_startingUp")
	symaddr(&addr_QCoreApplication_closingDown, "QCoreApplication_closingDown")
	symaddr(&addr_QCoreApplication_applicationDirPath, "QCoreApplication_applicationDirPath")
	symaddr(&addr_QCoreApplication_applicationFilePath, "QCoreApplication_applicationFilePath")
	symaddr(&addr_QCoreApplication_applicationPid, "QCoreApplication_applicationPid")
	symaddr(&addr_QCoreApplication_checkPermission, "QCoreApplication_checkPermission")
	symaddr(&addr_QCoreApplication_setLibraryPaths, "QCoreApplication_setLibraryPaths")
	symaddr(&addr_QCoreApplication_libraryPaths, "QCoreApplication_libraryPaths")
	symaddr(&addr_QCoreApplication_addLibraryPath, "QCoreApplication_addLibraryPath")
	symaddr(&addr_QCoreApplication_removeLibraryPath, "QCoreApplication_removeLibraryPath")
	symaddr(&addr_QCoreApplication_installTranslator, "QCoreApplication_installTranslator")
	symaddr(&addr_QCoreApplication_removeTranslator, "QCoreApplication_removeTranslator")
	symaddr(&addr_QCoreApplication_translate, "QCoreApplication_translate")
	symaddr(&addr_QCoreApplication_installNativeEventFilter, "QCoreApplication_installNativeEventFilter")
	symaddr(&addr_QCoreApplication_removeNativeEventFilter, "QCoreApplication_removeNativeEventFilter")
	symaddr(&addr_QCoreApplication_isQuitLockEnabled, "QCoreApplication_isQuitLockEnabled")
	symaddr(&addr_QCoreApplication_setQuitLockEnabled, "QCoreApplication_setQuitLockEnabled")
	symaddr(&addr_QCoreApplication_quit, "QCoreApplication_quit")
	symaddr(&addr_QCoreApplication_exit, "QCoreApplication_exit")
	symaddr(&addr_QCoreApplication_organizationNameChanged, "QCoreApplication_organizationNameChanged")
	symaddr(&addr_QCoreApplication_organizationDomainChanged, "QCoreApplication_organizationDomainChanged")
	symaddr(&addr_QCoreApplication_applicationNameChanged, "QCoreApplication_applicationNameChanged")
	symaddr(&addr_QCoreApplication_applicationVersionChanged, "QCoreApplication_applicationVersionChanged")
	symaddr(&addr_QCoreApplication_connect, "QCoreApplication_connect")
	symaddr(&addr_QCoreApplication_connect2, "QCoreApplication_connect2")
	symaddr(&addr_QCoreApplication_tr2, "QCoreApplication_tr2")
	symaddr(&addr_QCoreApplication_tr3, "QCoreApplication_tr3")
	symaddr(&addr_QCoreApplication_setAttribute2, "QCoreApplication_setAttribute2")
	symaddr(&addr_QCoreApplication_processEvents4, "QCoreApplication_processEvents4")
	symaddr(&addr_QCoreApplication_postEvent2, "QCoreApplication_postEvent2")
	symaddr(&addr_QCoreApplication_sendPostedEvents2, "QCoreApplication_sendPostedEvents2")
	symaddr(&addr_QCoreApplication_sendPostedEvents3, "QCoreApplication_sendPostedEvents3")
	symaddr(&addr_QCoreApplication_removePostedEvents2, "QCoreApplication_removePostedEvents2")
	symaddr(&addr_QCoreApplication_translate2, "QCoreApplication_translate2")
	symaddr(&addr_QCoreApplication_translate3, "QCoreApplication_translate3")
	symaddr(&addr_QCoreApplication_exit2, "QCoreApplication_exit2")
	symaddr(&addr_QCoreApplication_connect3, "QCoreApplication_connect3")
	symaddr(&addr_QCoreApplication_connect4, "QCoreApplication_connect4")
	reg(&pg_QCoreApplication_connect_aboutToQuit, "QCoreApplication_connect_aboutToQuit")
	reg(&pg_QCoreApplication_connect_organizationNameChanged, "QCoreApplication_connect_organizationNameChanged")
	reg(&pg_QCoreApplication_connect_organizationDomainChanged, "QCoreApplication_connect_organizationDomainChanged")
	reg(&pg_QCoreApplication_connect_applicationNameChanged, "QCoreApplication_connect_applicationNameChanged")
	reg(&pg_QCoreApplication_connect_applicationVersionChanged, "QCoreApplication_connect_applicationVersionChanged")
	reg(&pg_QCoreApplication_connect_destroyed, "QCoreApplication_connect_destroyed")
	reg(&pg_QCoreApplication_connect_destroyed2, "QCoreApplication_connect_destroyed2")
	reg(&pg_QCoreApplication_connect_objectNameChanged, "QCoreApplication_connect_objectNameChanged")
	reg(&pg_QCoreApplication__disconnect, "QCoreApplication__disconnect")
	reg(&pg_QCoreApplication_new_args, "QCoreApplication_new_args")
	reg(&pg_QCoreApplication_delete, "QCoreApplication_delete")
}

type QCoreApplication struct {
	h uintptr
	*QObject
}

func bindQCoreApplication(h uintptr) *QCoreApplication {
	if h == 0 {
		return nil
	}
	return &QCoreApplication{h: h, QObject: bindQObject(pg_QCoreApplication_base(h))}
}

func newownedQCoreApplication(h uintptr) *QCoreApplication { return bindQCoreApplication(h) }

func (this *QCoreApplication) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCoreApplication) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCoreApplication(h unsafe.Pointer) *QCoreApplication { return bindQCoreApplication(uintptr(h)) }
func UnsafeNewQCoreApplicationOwned(h unsafe.Pointer) *QCoreApplication { return newownedQCoreApplication(uintptr(h)) }

func (this *QCoreApplication) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QCoreApplication_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func QCoreApplication_Arguments() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_arguments, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QCoreApplication_SetAttribute(attribute ApplicationAttribute) {
	purego.SyscallN(addr_QCoreApplication_setAttribute, uintptr(attribute))
}

func QCoreApplication_TestAttribute(attribute ApplicationAttribute) bool {
	__r0, _, _ := purego.SyscallN(addr_QCoreApplication_testAttribute, uintptr(attribute))
	__r := int32(uint32(__r0))
	return __r != 0
}

func QCoreApplication_SetOrganizationDomain(orgDomain string) {
	var orgDomain_p uintptr
	if len(orgDomain) > 0 {
		orgDomain_p = uintptr(unsafe.Pointer(unsafe.StringData(orgDomain)))
	}
	purego.SyscallN(addr_QCoreApplication_setOrganizationDomain, orgDomain_p, uintptr(len(orgDomain)))
	runtime.KeepAlive(orgDomain)
}

func QCoreApplication_OrganizationDomain() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_organizationDomain, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QCoreApplication_SetOrganizationName(orgName string) {
	var orgName_p uintptr
	if len(orgName) > 0 {
		orgName_p = uintptr(unsafe.Pointer(unsafe.StringData(orgName)))
	}
	purego.SyscallN(addr_QCoreApplication_setOrganizationName, orgName_p, uintptr(len(orgName)))
	runtime.KeepAlive(orgName)
}

func QCoreApplication_OrganizationName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_organizationName, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QCoreApplication_SetApplicationName(application string) {
	var application_p uintptr
	if len(application) > 0 {
		application_p = uintptr(unsafe.Pointer(unsafe.StringData(application)))
	}
	purego.SyscallN(addr_QCoreApplication_setApplicationName, application_p, uintptr(len(application)))
	runtime.KeepAlive(application)
}

func QCoreApplication_ApplicationName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_applicationName, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QCoreApplication_SetApplicationVersion(version string) {
	var version_p uintptr
	if len(version) > 0 {
		version_p = uintptr(unsafe.Pointer(unsafe.StringData(version)))
	}
	purego.SyscallN(addr_QCoreApplication_setApplicationVersion, version_p, uintptr(len(version)))
	runtime.KeepAlive(version)
}

func QCoreApplication_ApplicationVersion() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_applicationVersion, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QCoreApplication_SetSetuidAllowed(allow bool) {
	purego.SyscallN(addr_QCoreApplication_setSetuidAllowed, uintptr(b2i(allow)))
}

func QCoreApplication_IsSetuidAllowed() bool {
	__r0, _, _ := purego.SyscallN(addr_QCoreApplication_isSetuidAllowed)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QCoreApplication_Instance() *QCoreApplication {
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_instance)
	return bindQCoreApplication(__r)
}

func QCoreApplication_InstanceExists() bool {
	__r0, _, _ := purego.SyscallN(addr_QCoreApplication_instanceExists)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QCoreApplication_Exec() int {
	__r0, _, _ := purego.SyscallN(addr_QCoreApplication_exec)
	__r := int32(uint32(__r0))
	return int(__r)
}

func QCoreApplication_ProcessEvents() {
	purego.SyscallN(addr_QCoreApplication_processEvents)
}

func QCoreApplication_ProcessEvents2(flags QEventLoop__ProcessEventsFlag, maxtime int) {
	purego.SyscallN(addr_QCoreApplication_processEvents2, uintptr(flags), uintptr(maxtime))
}

func QCoreApplication_ProcessEvents3(flags QEventLoop__ProcessEventsFlag, deadline *QDeadlineTimer) {
	purego.SyscallN(addr_QCoreApplication_processEvents3, uintptr(flags), uintptr(deadline.UnsafePointer()))
	runtime.KeepAlive(deadline)
}

func QCoreApplication_SendEvent(receiver *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QCoreApplication_sendEvent, uintptr(receiver.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(event)
	return __r != 0
}

func QCoreApplication_PostEvent(receiver *QObject, event *QEvent) {
	purego.SyscallN(addr_QCoreApplication_postEvent, uintptr(receiver.UnsafePointer()), uintptr(event.UnsafePointer()))
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(event)
}

func QCoreApplication_SendPostedEvents() {
	purego.SyscallN(addr_QCoreApplication_sendPostedEvents)
}

func QCoreApplication_RemovePostedEvents(receiver *QObject) {
	purego.SyscallN(addr_QCoreApplication_removePostedEvents, uintptr(receiver.UnsafePointer()))
	runtime.KeepAlive(receiver)
}

func QCoreApplication_EventDispatcher() *QAbstractEventDispatcher {
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_eventDispatcher)
	return bindQAbstractEventDispatcher(__r)
}

func QCoreApplication_SetEventDispatcher(eventDispatcher *QAbstractEventDispatcher) {
	purego.SyscallN(addr_QCoreApplication_setEventDispatcher, uintptr(eventDispatcher.UnsafePointer()))
	runtime.KeepAlive(eventDispatcher)
}

func (this *QCoreApplication) Notify(arg1 *QObject, arg2 *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QCoreApplication_notify, this.h, uintptr(arg1.UnsafePointer()), uintptr(arg2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
	runtime.KeepAlive(arg2)
	return __r != 0
}

func QCoreApplication_StartingUp() bool {
	__r0, _, _ := purego.SyscallN(addr_QCoreApplication_startingUp)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QCoreApplication_ClosingDown() bool {
	__r0, _, _ := purego.SyscallN(addr_QCoreApplication_closingDown)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QCoreApplication_ApplicationDirPath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_applicationDirPath, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QCoreApplication_ApplicationFilePath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_applicationFilePath, uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QCoreApplication_ApplicationPid() int64 {
	__r0, _, _ := purego.SyscallN(addr_QCoreApplication_applicationPid)
	__r := int64(__r0)
	return __r
}

func (this *QCoreApplication) CheckPermission(permission *QPermission) PermissionStatus {
	__r0, _, _ := purego.SyscallN(addr_QCoreApplication_checkPermission, this.h, uintptr(permission.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(permission)
	return PermissionStatus(__r)
}

func QCoreApplication_SetLibraryPaths(arg1 []string) {
	arg1_arr := make([]uintptr, len(arg1))
	for __i, __e := range arg1 {
		arg1_arr[__i] = pgCString(__e)
	}
	var arg1_p uintptr
	if len(arg1_arr) > 0 {
		arg1_p = uintptr(unsafe.Pointer(&arg1_arr[0]))
	}
	purego.SyscallN(addr_QCoreApplication_setLibraryPaths, arg1_p, uintptr(len(arg1)))
	for __i := range arg1_arr {
		pgFree(arg1_arr[__i])
	}
	runtime.KeepAlive(arg1_arr)
}

func QCoreApplication_LibraryPaths() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_libraryPaths, uintptr(unsafe.Pointer(&__n)))
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func QCoreApplication_AddLibraryPath(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QCoreApplication_addLibraryPath, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(arg1)
}

func QCoreApplication_RemoveLibraryPath(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QCoreApplication_removeLibraryPath, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(arg1)
}

func QCoreApplication_InstallTranslator(messageFile *QTranslator) bool {
	__r0, _, _ := purego.SyscallN(addr_QCoreApplication_installTranslator, uintptr(messageFile.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(messageFile)
	return __r != 0
}

func QCoreApplication_RemoveTranslator(messageFile *QTranslator) bool {
	__r0, _, _ := purego.SyscallN(addr_QCoreApplication_removeTranslator, uintptr(messageFile.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(messageFile)
	return __r != 0
}

func QCoreApplication_Translate(context string, key string) string {
	context_b := append([]byte(context), 0)
	key_b := append([]byte(key), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_translate, uintptr(unsafe.Pointer(&context_b[0])), uintptr(unsafe.Pointer(&key_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(context_b)
	runtime.KeepAlive(key_b)
	return pgStr(__r, __n)
}

func (this *QCoreApplication) InstallNativeEventFilter(filterObj *QAbstractNativeEventFilter) {
	purego.SyscallN(addr_QCoreApplication_installNativeEventFilter, this.h, uintptr(filterObj.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filterObj)
}

func (this *QCoreApplication) RemoveNativeEventFilter(filterObj *QAbstractNativeEventFilter) {
	purego.SyscallN(addr_QCoreApplication_removeNativeEventFilter, this.h, uintptr(filterObj.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filterObj)
}

func QCoreApplication_IsQuitLockEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QCoreApplication_isQuitLockEnabled)
	__r := int32(uint32(__r0))
	return __r != 0
}

func QCoreApplication_SetQuitLockEnabled(enabled bool) {
	purego.SyscallN(addr_QCoreApplication_setQuitLockEnabled, uintptr(b2i(enabled)))
}

func QCoreApplication_Quit() {
	purego.SyscallN(addr_QCoreApplication_quit)
}

func QCoreApplication_Exit() {
	purego.SyscallN(addr_QCoreApplication_exit)
}

func (this *QCoreApplication) OrganizationNameChanged() {
	purego.SyscallN(addr_QCoreApplication_organizationNameChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QCoreApplication) OrganizationDomainChanged() {
	purego.SyscallN(addr_QCoreApplication_organizationDomainChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QCoreApplication) ApplicationNameChanged() {
	purego.SyscallN(addr_QCoreApplication_applicationNameChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QCoreApplication) ApplicationVersionChanged() {
	purego.SyscallN(addr_QCoreApplication_applicationVersionChanged, this.h)
	runtime.KeepAlive(this)
}

func QCoreApplication_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QCoreApplication_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QCoreApplication_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QCoreApplication_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QCoreApplication_SetAttribute2(attribute ApplicationAttribute, on bool) {
	purego.SyscallN(addr_QCoreApplication_setAttribute2, uintptr(attribute), uintptr(b2i(on)))
}

func QCoreApplication_ProcessEventsWithFlags(flags QEventLoop__ProcessEventsFlag) {
	purego.SyscallN(addr_QCoreApplication_processEvents4, uintptr(flags))
}

func QCoreApplication_PostEvent2(receiver *QObject, event *QEvent, priority int) {
	purego.SyscallN(addr_QCoreApplication_postEvent2, uintptr(receiver.UnsafePointer()), uintptr(event.UnsafePointer()), uintptr(priority))
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(event)
}

func QCoreApplication_SendPostedEventsWithReceiver(receiver *QObject) {
	purego.SyscallN(addr_QCoreApplication_sendPostedEvents2, uintptr(receiver.UnsafePointer()))
	runtime.KeepAlive(receiver)
}

func QCoreApplication_SendPostedEvents2(receiver *QObject, event_type int) {
	purego.SyscallN(addr_QCoreApplication_sendPostedEvents3, uintptr(receiver.UnsafePointer()), uintptr(event_type))
	runtime.KeepAlive(receiver)
}

func QCoreApplication_RemovePostedEvents2(receiver *QObject, eventType int) {
	purego.SyscallN(addr_QCoreApplication_removePostedEvents2, uintptr(receiver.UnsafePointer()), uintptr(eventType))
	runtime.KeepAlive(receiver)
}

func QCoreApplication_Translate2(context string, key string, disambiguation string) string {
	context_b := append([]byte(context), 0)
	key_b := append([]byte(key), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_translate2, uintptr(unsafe.Pointer(&context_b[0])), uintptr(unsafe.Pointer(&key_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(context_b)
	runtime.KeepAlive(key_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func QCoreApplication_Translate3(context string, key string, disambiguation string, n int) string {
	context_b := append([]byte(context), 0)
	key_b := append([]byte(key), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_translate3, uintptr(unsafe.Pointer(&context_b[0])), uintptr(unsafe.Pointer(&key_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(context_b)
	runtime.KeepAlive(key_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func QCoreApplication_ExitWithRetcode(retcode int) {
	purego.SyscallN(addr_QCoreApplication_exit2, uintptr(retcode))
}

func QCoreApplication_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QCoreApplication_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QCoreApplication_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QCoreApplication) OnAboutToQuit(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QCoreApplication_connect_aboutToQuit(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCoreApplication) OnOrganizationNameChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QCoreApplication_connect_organizationNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCoreApplication) OnOrganizationDomainChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QCoreApplication_connect_organizationDomainChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCoreApplication) OnApplicationNameChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QCoreApplication_connect_applicationNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCoreApplication) OnApplicationVersionChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QCoreApplication_connect_applicationVersionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCoreApplication) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QCoreApplication_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCoreApplication) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QCoreApplication_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCoreApplication) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QCoreApplication_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QCoreApplication) Disconnect(handle int64) {
	pg_QCoreApplication__disconnect(handle)
	unregisterCallback(handle)
}

func NewQCoreApplication(args []string) *QCoreApplication {
	argc := pgMalloc(unsafe.Sizeof(int32(0)))
	*(*int32)(unsafe.Pointer(argc)) = int32(len(args))
	argv := pgMalloc(uintptr(len(args)+1) * unsafe.Sizeof(uintptr(0)))
	__av := unsafe.Slice((*uintptr)(unsafe.Pointer(argv)), len(args)+1)
	for i := range args {
		__av[i] = pgCString(args[i])
	}
	__av[len(args)] = 0
	runtime.LockOSThread()
	return newownedQCoreApplication(pg_QCoreApplication_new_args(argc, argv))
}

func (this *QCoreApplication) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCoreApplication_delete(this.h)
	}
	this.UnsafeRelease()
}
