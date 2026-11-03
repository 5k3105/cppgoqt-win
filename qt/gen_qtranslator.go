package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QTranslator_base func(uintptr) uintptr
	addr_QTranslator_new uintptr
	addr_QTranslator_new2 uintptr
	addr_QTranslator_tr uintptr
	addr_QTranslator_translate uintptr
	addr_QTranslator_isEmpty uintptr
	addr_QTranslator_language uintptr
	addr_QTranslator_filePath uintptr
	addr_QTranslator_load uintptr
	addr_QTranslator_load2 uintptr
	addr_QTranslator_load3 uintptr
	addr_QTranslator_connect uintptr
	addr_QTranslator_connect2 uintptr
	addr_QTranslator_sender uintptr
	addr_QTranslator_senderSignalIndex uintptr
	addr_QTranslator_receivers uintptr
	addr_QTranslator_isSignalConnected uintptr
	addr_QTranslator_tr2 uintptr
	addr_QTranslator_tr3 uintptr
	addr_QTranslator_load4 uintptr
	addr_QTranslator_load5 uintptr
	addr_QTranslator_load6 uintptr
	addr_QTranslator_load7 uintptr
	addr_QTranslator_load8 uintptr
	addr_QTranslator_load9 uintptr
	addr_QTranslator_load10 uintptr
	addr_QTranslator_connect3 uintptr
	addr_QTranslator_connect4 uintptr
	addr_QTranslator_isEmpty_default uintptr
	addr_QTranslator_metaObject_default uintptr
	addr_QTranslator_event_default uintptr
	addr_QTranslator_eventFilter_default uintptr
	addr_QTranslator_timerEvent_default uintptr
	addr_QTranslator_childEvent_default uintptr
	addr_QTranslator_customEvent_default uintptr
	addr_QTranslator_connectNotify_default uintptr
	addr_QTranslator_disconnectNotify_default uintptr
	pg_QTranslator_connect_destroyed func(uintptr, int64)
	pg_QTranslator_connect_destroyed2 func(uintptr, int64)
	pg_QTranslator_connect_objectNameChanged func(uintptr, int64)
	pg_QTranslator__disconnect func(int64)
	pg_QTranslator__set_override func(uintptr, int32)
	pg_QTranslator_delete func(uintptr)
)

func init() {
	reg(&pg_QTranslator_base, "QTranslator_base")
	symaddr(&addr_QTranslator_new, "QTranslator_new")
	symaddr(&addr_QTranslator_new2, "QTranslator_new2")
	symaddr(&addr_QTranslator_tr, "QTranslator_tr")
	symaddr(&addr_QTranslator_translate, "QTranslator_translate")
	symaddr(&addr_QTranslator_isEmpty, "QTranslator_isEmpty")
	symaddr(&addr_QTranslator_language, "QTranslator_language")
	symaddr(&addr_QTranslator_filePath, "QTranslator_filePath")
	symaddr(&addr_QTranslator_load, "QTranslator_load")
	symaddr(&addr_QTranslator_load2, "QTranslator_load2")
	symaddr(&addr_QTranslator_load3, "QTranslator_load3")
	symaddr(&addr_QTranslator_connect, "QTranslator_connect")
	symaddr(&addr_QTranslator_connect2, "QTranslator_connect2")
	symaddr(&addr_QTranslator_sender, "QTranslator_sender")
	symaddr(&addr_QTranslator_senderSignalIndex, "QTranslator_senderSignalIndex")
	symaddr(&addr_QTranslator_receivers, "QTranslator_receivers")
	symaddr(&addr_QTranslator_isSignalConnected, "QTranslator_isSignalConnected")
	symaddr(&addr_QTranslator_tr2, "QTranslator_tr2")
	symaddr(&addr_QTranslator_tr3, "QTranslator_tr3")
	symaddr(&addr_QTranslator_load4, "QTranslator_load4")
	symaddr(&addr_QTranslator_load5, "QTranslator_load5")
	symaddr(&addr_QTranslator_load6, "QTranslator_load6")
	symaddr(&addr_QTranslator_load7, "QTranslator_load7")
	symaddr(&addr_QTranslator_load8, "QTranslator_load8")
	symaddr(&addr_QTranslator_load9, "QTranslator_load9")
	symaddr(&addr_QTranslator_load10, "QTranslator_load10")
	symaddr(&addr_QTranslator_connect3, "QTranslator_connect3")
	symaddr(&addr_QTranslator_connect4, "QTranslator_connect4")
	symaddr(&addr_QTranslator_isEmpty_default, "QTranslator_isEmpty_default")
	symaddr(&addr_QTranslator_metaObject_default, "QTranslator_metaObject_default")
	symaddr(&addr_QTranslator_event_default, "QTranslator_event_default")
	symaddr(&addr_QTranslator_eventFilter_default, "QTranslator_eventFilter_default")
	symaddr(&addr_QTranslator_timerEvent_default, "QTranslator_timerEvent_default")
	symaddr(&addr_QTranslator_childEvent_default, "QTranslator_childEvent_default")
	symaddr(&addr_QTranslator_customEvent_default, "QTranslator_customEvent_default")
	symaddr(&addr_QTranslator_connectNotify_default, "QTranslator_connectNotify_default")
	symaddr(&addr_QTranslator_disconnectNotify_default, "QTranslator_disconnectNotify_default")
	reg(&pg_QTranslator_connect_destroyed, "QTranslator_connect_destroyed")
	reg(&pg_QTranslator_connect_destroyed2, "QTranslator_connect_destroyed2")
	reg(&pg_QTranslator_connect_objectNameChanged, "QTranslator_connect_objectNameChanged")
	reg(&pg_QTranslator__disconnect, "QTranslator__disconnect")
	reg(&pg_QTranslator__set_override, "QTranslator__set_override")
	reg(&pg_QTranslator_delete, "QTranslator_delete")
}

type QTranslator struct {
	h uintptr
	*QObject
}

func bindQTranslator(h uintptr) *QTranslator {
	if h == 0 {
		return nil
	}
	return &QTranslator{h: h, QObject: bindQObject(pg_QTranslator_base(h))}
}

func newownedQTranslator(h uintptr) *QTranslator { return bindQTranslator(h) }

func (this *QTranslator) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTranslator) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTranslator(h unsafe.Pointer) *QTranslator { return bindQTranslator(uintptr(h)) }
func UnsafeNewQTranslatorOwned(h unsafe.Pointer) *QTranslator { return newownedQTranslator(uintptr(h)) }

func (this *QTranslator) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQTranslator() *QTranslator {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTranslator_new, uintptr(__goid))
	__o := newownedQTranslator(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTranslator_virtIdx, QTranslator_enable)
	return __o
}

func NewQTranslator2(parent *QObject) *QTranslator {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QTranslator_new2, uintptr(__goid), uintptr(parent.UnsafePointer()))
	__o := newownedQTranslator(__r)
	runtime.KeepAlive(parent)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QTranslator_virtIdx, QTranslator_enable)
	return __o
}

func QTranslator_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTranslator_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QTranslator) Translate(context string, sourceText string, disambiguation string, n int) string {
	context_b := append([]byte(context), 0)
	sourceText_b := append([]byte(sourceText), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTranslator_translate, this.h, uintptr(unsafe.Pointer(&context_b[0])), uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(context_b)
	runtime.KeepAlive(sourceText_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func (this *QTranslator) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QTranslator_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTranslator) Language() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTranslator_language, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTranslator) FilePath() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTranslator_filePath, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QTranslator) Load(filename string) bool {
	var filename_p uintptr
	if len(filename) > 0 {
		filename_p = uintptr(unsafe.Pointer(unsafe.StringData(filename)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTranslator_load, this.h, filename_p, uintptr(len(filename)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filename)
	return __r != 0
}

func (this *QTranslator) Load2(locale *QLocale, filename string) bool {
	var filename_p uintptr
	if len(filename) > 0 {
		filename_p = uintptr(unsafe.Pointer(unsafe.StringData(filename)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTranslator_load2, this.h, uintptr(locale.UnsafePointer()), filename_p, uintptr(len(filename)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
	runtime.KeepAlive(filename)
	return __r != 0
}

func (this *QTranslator) Load3(data []byte, len_ int) bool {
	var data_p uintptr
	if len(data) > 0 {
		data_p = uintptr(unsafe.Pointer(&data[0]))
	}
	__r0, _, _ := purego.SyscallN(addr_QTranslator_load3, this.h, data_p, uintptr(len_))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	return __r != 0
}

func QTranslator_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTranslator_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTranslator_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTranslator_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTranslator) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QTranslator_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QTranslator) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QTranslator_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QTranslator) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QTranslator_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QTranslator) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QTranslator_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QTranslator_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTranslator_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QTranslator_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTranslator_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QTranslator) Load4(filename string, directory string) bool {
	var filename_p uintptr
	if len(filename) > 0 {
		filename_p = uintptr(unsafe.Pointer(unsafe.StringData(filename)))
	}
	var directory_p uintptr
	if len(directory) > 0 {
		directory_p = uintptr(unsafe.Pointer(unsafe.StringData(directory)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTranslator_load4, this.h, filename_p, uintptr(len(filename)), directory_p, uintptr(len(directory)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(directory)
	return __r != 0
}

func (this *QTranslator) Load5(filename string, directory string, search_delimiters string) bool {
	var filename_p uintptr
	if len(filename) > 0 {
		filename_p = uintptr(unsafe.Pointer(unsafe.StringData(filename)))
	}
	var directory_p uintptr
	if len(directory) > 0 {
		directory_p = uintptr(unsafe.Pointer(unsafe.StringData(directory)))
	}
	var search_delimiters_p uintptr
	if len(search_delimiters) > 0 {
		search_delimiters_p = uintptr(unsafe.Pointer(unsafe.StringData(search_delimiters)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTranslator_load5, this.h, filename_p, uintptr(len(filename)), directory_p, uintptr(len(directory)), search_delimiters_p, uintptr(len(search_delimiters)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(directory)
	runtime.KeepAlive(search_delimiters)
	return __r != 0
}

func (this *QTranslator) Load6(filename string, directory string, search_delimiters string, suffix string) bool {
	var filename_p uintptr
	if len(filename) > 0 {
		filename_p = uintptr(unsafe.Pointer(unsafe.StringData(filename)))
	}
	var directory_p uintptr
	if len(directory) > 0 {
		directory_p = uintptr(unsafe.Pointer(unsafe.StringData(directory)))
	}
	var search_delimiters_p uintptr
	if len(search_delimiters) > 0 {
		search_delimiters_p = uintptr(unsafe.Pointer(unsafe.StringData(search_delimiters)))
	}
	var suffix_p uintptr
	if len(suffix) > 0 {
		suffix_p = uintptr(unsafe.Pointer(unsafe.StringData(suffix)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTranslator_load6, this.h, filename_p, uintptr(len(filename)), directory_p, uintptr(len(directory)), search_delimiters_p, uintptr(len(search_delimiters)), suffix_p, uintptr(len(suffix)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(directory)
	runtime.KeepAlive(search_delimiters)
	runtime.KeepAlive(suffix)
	return __r != 0
}

func (this *QTranslator) Load7(locale *QLocale, filename string, prefix string) bool {
	var filename_p uintptr
	if len(filename) > 0 {
		filename_p = uintptr(unsafe.Pointer(unsafe.StringData(filename)))
	}
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTranslator_load7, this.h, uintptr(locale.UnsafePointer()), filename_p, uintptr(len(filename)), prefix_p, uintptr(len(prefix)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(prefix)
	return __r != 0
}

func (this *QTranslator) Load8(locale *QLocale, filename string, prefix string, directory string) bool {
	var filename_p uintptr
	if len(filename) > 0 {
		filename_p = uintptr(unsafe.Pointer(unsafe.StringData(filename)))
	}
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	var directory_p uintptr
	if len(directory) > 0 {
		directory_p = uintptr(unsafe.Pointer(unsafe.StringData(directory)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTranslator_load8, this.h, uintptr(locale.UnsafePointer()), filename_p, uintptr(len(filename)), prefix_p, uintptr(len(prefix)), directory_p, uintptr(len(directory)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(prefix)
	runtime.KeepAlive(directory)
	return __r != 0
}

func (this *QTranslator) Load9(locale *QLocale, filename string, prefix string, directory string, suffix string) bool {
	var filename_p uintptr
	if len(filename) > 0 {
		filename_p = uintptr(unsafe.Pointer(unsafe.StringData(filename)))
	}
	var prefix_p uintptr
	if len(prefix) > 0 {
		prefix_p = uintptr(unsafe.Pointer(unsafe.StringData(prefix)))
	}
	var directory_p uintptr
	if len(directory) > 0 {
		directory_p = uintptr(unsafe.Pointer(unsafe.StringData(directory)))
	}
	var suffix_p uintptr
	if len(suffix) > 0 {
		suffix_p = uintptr(unsafe.Pointer(unsafe.StringData(suffix)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTranslator_load9, this.h, uintptr(locale.UnsafePointer()), filename_p, uintptr(len(filename)), prefix_p, uintptr(len(prefix)), directory_p, uintptr(len(directory)), suffix_p, uintptr(len(suffix)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(locale)
	runtime.KeepAlive(filename)
	runtime.KeepAlive(prefix)
	runtime.KeepAlive(directory)
	runtime.KeepAlive(suffix)
	return __r != 0
}

func (this *QTranslator) Load10(data []byte, len_ int, directory string) bool {
	var data_p uintptr
	if len(data) > 0 {
		data_p = uintptr(unsafe.Pointer(&data[0]))
	}
	var directory_p uintptr
	if len(directory) > 0 {
		directory_p = uintptr(unsafe.Pointer(unsafe.StringData(directory)))
	}
	__r0, _, _ := purego.SyscallN(addr_QTranslator_load10, this.h, data_p, uintptr(len_), directory_p, uintptr(len(directory)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
	runtime.KeepAlive(directory)
	return __r != 0
}

func QTranslator_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QTranslator_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QTranslator_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QTranslator_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QTranslator) IsEmptyDefault() bool {
	__r0, _, _ := purego.SyscallN(addr_QTranslator_isEmpty_default, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTranslator) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QTranslator_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QTranslator) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTranslator_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTranslator) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QTranslator_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QTranslator) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QTranslator_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTranslator) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QTranslator_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTranslator) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QTranslator_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QTranslator) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTranslator_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTranslator) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QTranslator_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QTranslator) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QTranslator_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTranslator) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QTranslator_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTranslator) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QTranslator_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QTranslator) Disconnect(handle int64) {
	pg_QTranslator__disconnect(handle)
	unregisterCallback(handle)
}

func QTranslator_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case -626822430251098498:
		return 1
	case 1237236456312824991:
		return 2
	case 5217623514630858297:
		return 3
	case -5409927581812374498:
		return 4
	case -1185315367526857833:
		return 5
	case 3769868229123868474:
		return 6
	case 7379029982170503008:
		return 7
	case -568369772982215962:
		return 8
	}
	return -1
}

func QTranslator_enable(h uintptr, idx int32) { pg_QTranslator__set_override(h, idx) }

func (this *QTranslator) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QTranslator) OnIsEmpty(fn func() bool) {
	registerVirt(this.UnsafeGoID(), -626822430251098498, func() int32 { return b2i(fn()) }) // slot "isEmpty"
}

func (this *QTranslator) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QTranslator) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QTranslator) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QTranslator) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QTranslator) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QTranslator) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QTranslator) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QTranslator) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTranslator_delete(this.h)
	}
	this.UnsafeRelease()
}
