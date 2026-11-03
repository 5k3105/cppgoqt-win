package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QSettings_base func(uintptr) uintptr
	addr_QSettings_new uintptr
	addr_QSettings_new2 uintptr
	addr_QSettings_new3 uintptr
	addr_QSettings_new4 uintptr
	addr_QSettings_new5 uintptr
	addr_QSettings_new6 uintptr
	addr_QSettings_new7 uintptr
	addr_QSettings_new8 uintptr
	addr_QSettings_new9 uintptr
	addr_QSettings_new10 uintptr
	addr_QSettings_new11 uintptr
	addr_QSettings_new12 uintptr
	addr_QSettings_new13 uintptr
	addr_QSettings_new14 uintptr
	addr_QSettings_new15 uintptr
	addr_QSettings_tr uintptr
	addr_QSettings_clear uintptr
	addr_QSettings_sync uintptr
	addr_QSettings_status uintptr
	addr_QSettings_isAtomicSyncRequired uintptr
	addr_QSettings_setAtomicSyncRequired uintptr
	addr_QSettings_beginGroup uintptr
	addr_QSettings_endGroup uintptr
	addr_QSettings_group uintptr
	addr_QSettings_beginReadArray uintptr
	addr_QSettings_beginWriteArray uintptr
	addr_QSettings_endArray uintptr
	addr_QSettings_setArrayIndex uintptr
	addr_QSettings_allKeys uintptr
	addr_QSettings_childKeys uintptr
	addr_QSettings_childGroups uintptr
	addr_QSettings_isWritable uintptr
	addr_QSettings_setValue uintptr
	addr_QSettings_value uintptr
	addr_QSettings_value2 uintptr
	addr_QSettings_remove uintptr
	addr_QSettings_contains uintptr
	addr_QSettings_setFallbacksEnabled uintptr
	addr_QSettings_fallbacksEnabled uintptr
	addr_QSettings_fileName uintptr
	addr_QSettings_format uintptr
	addr_QSettings_scope uintptr
	addr_QSettings_organizationName uintptr
	addr_QSettings_applicationName uintptr
	addr_QSettings_setDefaultFormat uintptr
	addr_QSettings_defaultFormat uintptr
	addr_QSettings_setPath uintptr
	addr_QSettings_connect uintptr
	addr_QSettings_connect2 uintptr
	addr_QSettings_sender uintptr
	addr_QSettings_senderSignalIndex uintptr
	addr_QSettings_receivers uintptr
	addr_QSettings_isSignalConnected uintptr
	addr_QSettings_tr2 uintptr
	addr_QSettings_tr3 uintptr
	addr_QSettings_beginWriteArray2 uintptr
	addr_QSettings_connect3 uintptr
	addr_QSettings_connect4 uintptr
	addr_QSettings_metaObject_default uintptr
	addr_QSettings_event_default uintptr
	addr_QSettings_eventFilter_default uintptr
	addr_QSettings_timerEvent_default uintptr
	addr_QSettings_childEvent_default uintptr
	addr_QSettings_customEvent_default uintptr
	addr_QSettings_connectNotify_default uintptr
	addr_QSettings_disconnectNotify_default uintptr
	pg_QSettings_connect_destroyed func(uintptr, int64)
	pg_QSettings_connect_destroyed2 func(uintptr, int64)
	pg_QSettings_connect_objectNameChanged func(uintptr, int64)
	pg_QSettings__disconnect func(int64)
	pg_QSettings__set_override func(uintptr, int32)
	pg_QSettings_delete func(uintptr)
)

func init() {
	reg(&pg_QSettings_base, "QSettings_base")
	symaddr(&addr_QSettings_new, "QSettings_new")
	symaddr(&addr_QSettings_new2, "QSettings_new2")
	symaddr(&addr_QSettings_new3, "QSettings_new3")
	symaddr(&addr_QSettings_new4, "QSettings_new4")
	symaddr(&addr_QSettings_new5, "QSettings_new5")
	symaddr(&addr_QSettings_new6, "QSettings_new6")
	symaddr(&addr_QSettings_new7, "QSettings_new7")
	symaddr(&addr_QSettings_new8, "QSettings_new8")
	symaddr(&addr_QSettings_new9, "QSettings_new9")
	symaddr(&addr_QSettings_new10, "QSettings_new10")
	symaddr(&addr_QSettings_new11, "QSettings_new11")
	symaddr(&addr_QSettings_new12, "QSettings_new12")
	symaddr(&addr_QSettings_new13, "QSettings_new13")
	symaddr(&addr_QSettings_new14, "QSettings_new14")
	symaddr(&addr_QSettings_new15, "QSettings_new15")
	symaddr(&addr_QSettings_tr, "QSettings_tr")
	symaddr(&addr_QSettings_clear, "QSettings_clear")
	symaddr(&addr_QSettings_sync, "QSettings_sync")
	symaddr(&addr_QSettings_status, "QSettings_status")
	symaddr(&addr_QSettings_isAtomicSyncRequired, "QSettings_isAtomicSyncRequired")
	symaddr(&addr_QSettings_setAtomicSyncRequired, "QSettings_setAtomicSyncRequired")
	symaddr(&addr_QSettings_beginGroup, "QSettings_beginGroup")
	symaddr(&addr_QSettings_endGroup, "QSettings_endGroup")
	symaddr(&addr_QSettings_group, "QSettings_group")
	symaddr(&addr_QSettings_beginReadArray, "QSettings_beginReadArray")
	symaddr(&addr_QSettings_beginWriteArray, "QSettings_beginWriteArray")
	symaddr(&addr_QSettings_endArray, "QSettings_endArray")
	symaddr(&addr_QSettings_setArrayIndex, "QSettings_setArrayIndex")
	symaddr(&addr_QSettings_allKeys, "QSettings_allKeys")
	symaddr(&addr_QSettings_childKeys, "QSettings_childKeys")
	symaddr(&addr_QSettings_childGroups, "QSettings_childGroups")
	symaddr(&addr_QSettings_isWritable, "QSettings_isWritable")
	symaddr(&addr_QSettings_setValue, "QSettings_setValue")
	symaddr(&addr_QSettings_value, "QSettings_value")
	symaddr(&addr_QSettings_value2, "QSettings_value2")
	symaddr(&addr_QSettings_remove, "QSettings_remove")
	symaddr(&addr_QSettings_contains, "QSettings_contains")
	symaddr(&addr_QSettings_setFallbacksEnabled, "QSettings_setFallbacksEnabled")
	symaddr(&addr_QSettings_fallbacksEnabled, "QSettings_fallbacksEnabled")
	symaddr(&addr_QSettings_fileName, "QSettings_fileName")
	symaddr(&addr_QSettings_format, "QSettings_format")
	symaddr(&addr_QSettings_scope, "QSettings_scope")
	symaddr(&addr_QSettings_organizationName, "QSettings_organizationName")
	symaddr(&addr_QSettings_applicationName, "QSettings_applicationName")
	symaddr(&addr_QSettings_setDefaultFormat, "QSettings_setDefaultFormat")
	symaddr(&addr_QSettings_defaultFormat, "QSettings_defaultFormat")
	symaddr(&addr_QSettings_setPath, "QSettings_setPath")
	symaddr(&addr_QSettings_connect, "QSettings_connect")
	symaddr(&addr_QSettings_connect2, "QSettings_connect2")
	symaddr(&addr_QSettings_sender, "QSettings_sender")
	symaddr(&addr_QSettings_senderSignalIndex, "QSettings_senderSignalIndex")
	symaddr(&addr_QSettings_receivers, "QSettings_receivers")
	symaddr(&addr_QSettings_isSignalConnected, "QSettings_isSignalConnected")
	symaddr(&addr_QSettings_tr2, "QSettings_tr2")
	symaddr(&addr_QSettings_tr3, "QSettings_tr3")
	symaddr(&addr_QSettings_beginWriteArray2, "QSettings_beginWriteArray2")
	symaddr(&addr_QSettings_connect3, "QSettings_connect3")
	symaddr(&addr_QSettings_connect4, "QSettings_connect4")
	symaddr(&addr_QSettings_metaObject_default, "QSettings_metaObject_default")
	symaddr(&addr_QSettings_event_default, "QSettings_event_default")
	symaddr(&addr_QSettings_eventFilter_default, "QSettings_eventFilter_default")
	symaddr(&addr_QSettings_timerEvent_default, "QSettings_timerEvent_default")
	symaddr(&addr_QSettings_childEvent_default, "QSettings_childEvent_default")
	symaddr(&addr_QSettings_customEvent_default, "QSettings_customEvent_default")
	symaddr(&addr_QSettings_connectNotify_default, "QSettings_connectNotify_default")
	symaddr(&addr_QSettings_disconnectNotify_default, "QSettings_disconnectNotify_default")
	reg(&pg_QSettings_connect_destroyed, "QSettings_connect_destroyed")
	reg(&pg_QSettings_connect_destroyed2, "QSettings_connect_destroyed2")
	reg(&pg_QSettings_connect_objectNameChanged, "QSettings_connect_objectNameChanged")
	reg(&pg_QSettings__disconnect, "QSettings__disconnect")
	reg(&pg_QSettings__set_override, "QSettings__set_override")
	reg(&pg_QSettings_delete, "QSettings_delete")
}

type QSettings struct {
	h uintptr
	*QObject
}

func bindQSettings(h uintptr) *QSettings {
	if h == 0 {
		return nil
	}
	return &QSettings{h: h, QObject: bindQObject(pg_QSettings_base(h))}
}

func newownedQSettings(h uintptr) *QSettings { return bindQSettings(h) }

func (this *QSettings) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QSettings) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQSettings(h unsafe.Pointer) *QSettings { return bindQSettings(uintptr(h)) }
func UnsafeNewQSettingsOwned(h unsafe.Pointer) *QSettings { return newownedQSettings(uintptr(h)) }

func (this *QSettings) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQSettings(organization string) *QSettings {
	__goid := nextVirtID()
	var organization_p uintptr
	if len(organization) > 0 {
		organization_p = uintptr(unsafe.Pointer(unsafe.StringData(organization)))
	}
	__r, _, _ := purego.SyscallN(addr_QSettings_new, uintptr(__goid), organization_p, uintptr(len(organization)))
	__o := newownedQSettings(__r)
	runtime.KeepAlive(organization)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSettings_virtIdx, QSettings_enable)
	return __o
}

func NewQSettings2(scope QSettings__Scope, organization string) *QSettings {
	__goid := nextVirtID()
	var organization_p uintptr
	if len(organization) > 0 {
		organization_p = uintptr(unsafe.Pointer(unsafe.StringData(organization)))
	}
	__r, _, _ := purego.SyscallN(addr_QSettings_new2, uintptr(__goid), uintptr(scope), organization_p, uintptr(len(organization)))
	__o := newownedQSettings(__r)
	runtime.KeepAlive(organization)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSettings_virtIdx, QSettings_enable)
	return __o
}

func NewQSettings3(format QSettings__Format, scope QSettings__Scope, organization string) *QSettings {
	__goid := nextVirtID()
	var organization_p uintptr
	if len(organization) > 0 {
		organization_p = uintptr(unsafe.Pointer(unsafe.StringData(organization)))
	}
	__r, _, _ := purego.SyscallN(addr_QSettings_new3, uintptr(__goid), uintptr(format), uintptr(scope), organization_p, uintptr(len(organization)))
	__o := newownedQSettings(__r)
	runtime.KeepAlive(organization)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSettings_virtIdx, QSettings_enable)
	return __o
}

func NewQSettings4(fileName string, format QSettings__Format) *QSettings {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QSettings_new4, uintptr(__goid), fileName_p, uintptr(len(fileName)), uintptr(format))
	__o := newownedQSettings(__r)
	runtime.KeepAlive(fileName)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSettings_virtIdx, QSettings_enable)
	return __o
}

func NewQSettings5() *QSettings {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSettings_new5, uintptr(__goid))
	__o := newownedQSettings(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSettings_virtIdx, QSettings_enable)
	return __o
}

func NewQSettings6(scope QSettings__Scope) *QSettings {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSettings_new6, uintptr(__goid), uintptr(scope))
	__o := newownedQSettings(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSettings_virtIdx, QSettings_enable)
	return __o
}

func NewQSettings7(organization string, application string) *QSettings {
	__goid := nextVirtID()
	var organization_p uintptr
	if len(organization) > 0 {
		organization_p = uintptr(unsafe.Pointer(unsafe.StringData(organization)))
	}
	var application_p uintptr
	if len(application) > 0 {
		application_p = uintptr(unsafe.Pointer(unsafe.StringData(application)))
	}
	__r, _, _ := purego.SyscallN(addr_QSettings_new7, uintptr(__goid), organization_p, uintptr(len(organization)), application_p, uintptr(len(application)))
	__o := newownedQSettings(__r)
	runtime.KeepAlive(organization)
	runtime.KeepAlive(application)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSettings_virtIdx, QSettings_enable)
	return __o
}

func NewQSettings8(organization string, application string, parent *QObject) *QSettings {
	__goid := nextVirtID()
	var organization_p uintptr
	if len(organization) > 0 {
		organization_p = uintptr(unsafe.Pointer(unsafe.StringData(organization)))
	}
	var application_p uintptr
	if len(application) > 0 {
		application_p = uintptr(unsafe.Pointer(unsafe.StringData(application)))
	}
	__r, _, _ := purego.SyscallN(addr_QSettings_new8, uintptr(__goid), organization_p, uintptr(len(organization)), application_p, uintptr(len(application)), uintptr(parent.UnsafePointer()))
	__o := newownedQSettings(__r)
	runtime.KeepAlive(organization)
	runtime.KeepAlive(application)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSettings_virtIdx, QSettings_enable)
	return __o
}

func NewQSettings9(scope QSettings__Scope, organization string, application string) *QSettings {
	__goid := nextVirtID()
	var organization_p uintptr
	if len(organization) > 0 {
		organization_p = uintptr(unsafe.Pointer(unsafe.StringData(organization)))
	}
	var application_p uintptr
	if len(application) > 0 {
		application_p = uintptr(unsafe.Pointer(unsafe.StringData(application)))
	}
	__r, _, _ := purego.SyscallN(addr_QSettings_new9, uintptr(__goid), uintptr(scope), organization_p, uintptr(len(organization)), application_p, uintptr(len(application)))
	__o := newownedQSettings(__r)
	runtime.KeepAlive(organization)
	runtime.KeepAlive(application)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSettings_virtIdx, QSettings_enable)
	return __o
}

func NewQSettings10(scope QSettings__Scope, organization string, application string, parent *QObject) *QSettings {
	__goid := nextVirtID()
	var organization_p uintptr
	if len(organization) > 0 {
		organization_p = uintptr(unsafe.Pointer(unsafe.StringData(organization)))
	}
	var application_p uintptr
	if len(application) > 0 {
		application_p = uintptr(unsafe.Pointer(unsafe.StringData(application)))
	}
	__r, _, _ := purego.SyscallN(addr_QSettings_new10, uintptr(__goid), uintptr(scope), organization_p, uintptr(len(organization)), application_p, uintptr(len(application)), uintptr(parent.UnsafePointer()))
	__o := newownedQSettings(__r)
	runtime.KeepAlive(organization)
	runtime.KeepAlive(application)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSettings_virtIdx, QSettings_enable)
	return __o
}

func NewQSettings11(format QSettings__Format, scope QSettings__Scope, organization string, application string) *QSettings {
	__goid := nextVirtID()
	var organization_p uintptr
	if len(organization) > 0 {
		organization_p = uintptr(unsafe.Pointer(unsafe.StringData(organization)))
	}
	var application_p uintptr
	if len(application) > 0 {
		application_p = uintptr(unsafe.Pointer(unsafe.StringData(application)))
	}
	__r, _, _ := purego.SyscallN(addr_QSettings_new11, uintptr(__goid), uintptr(format), uintptr(scope), organization_p, uintptr(len(organization)), application_p, uintptr(len(application)))
	__o := newownedQSettings(__r)
	runtime.KeepAlive(organization)
	runtime.KeepAlive(application)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSettings_virtIdx, QSettings_enable)
	return __o
}

func NewQSettings12(format QSettings__Format, scope QSettings__Scope, organization string, application string, parent *QObject) *QSettings {
	__goid := nextVirtID()
	var organization_p uintptr
	if len(organization) > 0 {
		organization_p = uintptr(unsafe.Pointer(unsafe.StringData(organization)))
	}
	var application_p uintptr
	if len(application) > 0 {
		application_p = uintptr(unsafe.Pointer(unsafe.StringData(application)))
	}
	__r, _, _ := purego.SyscallN(addr_QSettings_new12, uintptr(__goid), uintptr(format), uintptr(scope), organization_p, uintptr(len(organization)), application_p, uintptr(len(application)), uintptr(parent.UnsafePointer()))
	__o := newownedQSettings(__r)
	runtime.KeepAlive(organization)
	runtime.KeepAlive(application)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSettings_virtIdx, QSettings_enable)
	return __o
}

func NewQSettings13(fileName string, format QSettings__Format, parent *QObject) *QSettings {
	__goid := nextVirtID()
	var fileName_p uintptr
	if len(fileName) > 0 {
		fileName_p = uintptr(unsafe.Pointer(unsafe.StringData(fileName)))
	}
	__r, _, _ := purego.SyscallN(addr_QSettings_new13, uintptr(__goid), fileName_p, uintptr(len(fileName)), uintptr(format), uintptr(parent.UnsafePointer()))
	__o := newownedQSettings(__r)
	runtime.KeepAlive(fileName)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSettings_virtIdx, QSettings_enable)
	return __o
}

func NewQSettings14(parent *QObject) *QSettings {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSettings_new14, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQSettings(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSettings_virtIdx, QSettings_enable)
	return __o
}

func NewQSettings15(scope QSettings__Scope, parent *QObject) *QSettings {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QSettings_new15, uintptr(__goid), uintptr(scope), uintptr(parent.UnsafePointer()))
	__o := newownedQSettings(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QSettings_virtIdx, QSettings_enable)
	return __o
}

func QSettings_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSettings_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QSettings) Clear() {
	purego.SyscallN(addr_QSettings_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QSettings) Sync() {
	purego.SyscallN(addr_QSettings_sync, this.h)
	runtime.KeepAlive(this)
}

func (this *QSettings) Status() QSettings__Status {
	__r0, _, _ := purego.SyscallN(addr_QSettings_status, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSettings__Status(__r)
}

func (this *QSettings) IsAtomicSyncRequired() bool {
	__r0, _, _ := purego.SyscallN(addr_QSettings_isAtomicSyncRequired, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSettings) SetAtomicSyncRequired(enable bool) {
	purego.SyscallN(addr_QSettings_setAtomicSyncRequired, this.h, uintptr(b2i(enable)))
	runtime.KeepAlive(this)
}

func (this *QSettings) BeginGroup(prefix string) {
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	purego.SyscallN(addr_QSettings_beginGroup, this.h, prefix_p, uintptr(len(prefix)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(prefix)
}

func (this *QSettings) EndGroup() {
	purego.SyscallN(addr_QSettings_endGroup, this.h)
	runtime.KeepAlive(this)
}

func (this *QSettings) Group() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSettings_group, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QSettings) BeginReadArray(prefix string) int {
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	__r0, _, _ := purego.SyscallN(addr_QSettings_beginReadArray, this.h, prefix_p, uintptr(len(prefix)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(prefix)
	return int(__r)
}

func (this *QSettings) BeginWriteArray(prefix string) {
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	purego.SyscallN(addr_QSettings_beginWriteArray, this.h, prefix_p, uintptr(len(prefix)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(prefix)
}

func (this *QSettings) EndArray() {
	purego.SyscallN(addr_QSettings_endArray, this.h)
	runtime.KeepAlive(this)
}

func (this *QSettings) SetArrayIndex(i int) {
	purego.SyscallN(addr_QSettings_setArrayIndex, this.h, uintptr(i))
	runtime.KeepAlive(this)
}

func (this *QSettings) AllKeys() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSettings_allKeys, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QSettings) ChildKeys() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSettings_childKeys, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QSettings) ChildGroups() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSettings_childGroups, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QSettings) IsWritable() bool {
	__r0, _, _ := purego.SyscallN(addr_QSettings_isWritable, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSettings) SetValue(key string, value *QVariant) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	purego.SyscallN(addr_QSettings_setValue, this.h, key_p, uintptr(len(key)), uintptr(value.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	runtime.KeepAlive(value)
}

func (this *QSettings) Value(key string, defaultValue *QVariant) *QVariant {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QSettings_value, this.h, key_p, uintptr(len(key)), uintptr(defaultValue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	runtime.KeepAlive(defaultValue)
	return newownedQVariant(__r)
}

func (this *QSettings) ValueWithKey(key string) *QVariant {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r, _, _ := purego.SyscallN(addr_QSettings_value2, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return newownedQVariant(__r)
}

func (this *QSettings) Remove(key string) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	purego.SyscallN(addr_QSettings_remove, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QSettings) Contains(key string) bool {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r0, _, _ := purego.SyscallN(addr_QSettings_contains, this.h, key_p, uintptr(len(key)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return __r != 0
}

func (this *QSettings) SetFallbacksEnabled(b bool) {
	purego.SyscallN(addr_QSettings_setFallbacksEnabled, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QSettings) FallbacksEnabled() bool {
	__r0, _, _ := purego.SyscallN(addr_QSettings_fallbacksEnabled, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QSettings) FileName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSettings_fileName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QSettings) Format() QSettings__Format {
	__r0, _, _ := purego.SyscallN(addr_QSettings_format, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSettings__Format(__r)
}

func (this *QSettings) Scope() QSettings__Scope {
	__r0, _, _ := purego.SyscallN(addr_QSettings_scope, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QSettings__Scope(__r)
}

func (this *QSettings) OrganizationName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSettings_organizationName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QSettings) ApplicationName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSettings_applicationName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QSettings_SetDefaultFormat(format QSettings__Format) {
	purego.SyscallN(addr_QSettings_setDefaultFormat, uintptr(format))
}

func QSettings_DefaultFormat() QSettings__Format {
	__r0, _, _ := purego.SyscallN(addr_QSettings_defaultFormat)
	__r := int32(uint32(__r0))
	return QSettings__Format(__r)
}

func QSettings_SetPath(format QSettings__Format, scope QSettings__Scope, path string) {
	var path_p uintptr
	if len(path) > 0 {
		path_p = uintptr(unsafe.Pointer(unsafe.StringData(path)))
	}
	purego.SyscallN(addr_QSettings_setPath, uintptr(format), uintptr(scope), path_p, uintptr(len(path)))
	runtime.KeepAlive(path)
}

func QSettings_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSettings_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSettings_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSettings_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSettings) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QSettings_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QSettings) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QSettings_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QSettings) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QSettings_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QSettings) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QSettings_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QSettings_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSettings_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QSettings_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QSettings_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QSettings) BeginWriteArray2(prefix string, size int) {
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	purego.SyscallN(addr_QSettings_beginWriteArray2, this.h, prefix_p, uintptr(len(prefix)), uintptr(size))
	runtime.KeepAlive(this)
	runtime.KeepAlive(prefix)
}

func QSettings_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QSettings_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QSettings_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QSettings_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QSettings) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QSettings_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QSettings) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSettings_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSettings) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QSettings_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QSettings) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QSettings_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSettings) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QSettings_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSettings) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QSettings_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QSettings) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSettings_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSettings) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QSettings_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QSettings) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QSettings_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSettings) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QSettings_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSettings) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QSettings_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QSettings) Disconnect(handle int64) {
	pg_QSettings__disconnect(handle)
	unregisterCallback(handle)
}

func QSettings_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 1237236456312824991:
		return 1
	case 5217623514630858297:
		return 2
	case -5409927581812374498:
		return 3
	case -1185315367526857833:
		return 4
	case 3769868229123868474:
		return 5
	case 7379029982170503008:
		return 6
	case -568369772982215962:
		return 7
	}
	return -1
}

func QSettings_enable(h uintptr, idx int32) { pg_QSettings__set_override(h, idx) }

func (this *QSettings) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QSettings) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QSettings) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QSettings) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QSettings) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QSettings) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QSettings) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QSettings) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QSettings) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QSettings_delete(this.h)
	}
	this.UnsafeRelease()
}
