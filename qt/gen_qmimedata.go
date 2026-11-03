package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QMimeData_base func(uintptr) uintptr
	addr_QMimeData_new uintptr
	addr_QMimeData_tr uintptr
	addr_QMimeData_urls uintptr
	addr_QMimeData_setUrls uintptr
	addr_QMimeData_hasUrls uintptr
	addr_QMimeData_text uintptr
	addr_QMimeData_setText uintptr
	addr_QMimeData_hasText uintptr
	addr_QMimeData_html uintptr
	addr_QMimeData_setHtml uintptr
	addr_QMimeData_hasHtml uintptr
	addr_QMimeData_imageData uintptr
	addr_QMimeData_setImageData uintptr
	addr_QMimeData_hasImage uintptr
	addr_QMimeData_colorData uintptr
	addr_QMimeData_setColorData uintptr
	addr_QMimeData_hasColor uintptr
	addr_QMimeData_data uintptr
	addr_QMimeData_setData uintptr
	addr_QMimeData_removeFormat uintptr
	addr_QMimeData_hasFormat uintptr
	addr_QMimeData_formats uintptr
	addr_QMimeData_clear uintptr
	addr_QMimeData_connect uintptr
	addr_QMimeData_connect2 uintptr
	addr_QMimeData_sender uintptr
	addr_QMimeData_senderSignalIndex uintptr
	addr_QMimeData_receivers uintptr
	addr_QMimeData_isSignalConnected uintptr
	addr_QMimeData_tr2 uintptr
	addr_QMimeData_tr3 uintptr
	addr_QMimeData_connect3 uintptr
	addr_QMimeData_connect4 uintptr
	addr_QMimeData_hasFormat_default uintptr
	addr_QMimeData_formats_default uintptr
	addr_QMimeData_metaObject_default uintptr
	addr_QMimeData_retrieveData_default uintptr
	addr_QMimeData_event_default uintptr
	addr_QMimeData_eventFilter_default uintptr
	addr_QMimeData_timerEvent_default uintptr
	addr_QMimeData_childEvent_default uintptr
	addr_QMimeData_customEvent_default uintptr
	addr_QMimeData_connectNotify_default uintptr
	addr_QMimeData_disconnectNotify_default uintptr
	pg_QMimeData_connect_destroyed func(uintptr, int64)
	pg_QMimeData_connect_destroyed2 func(uintptr, int64)
	pg_QMimeData_connect_objectNameChanged func(uintptr, int64)
	pg_QMimeData__disconnect func(int64)
	pg_QMimeData__set_override func(uintptr, int32)
	pg_QMimeData_delete func(uintptr)
)

func init() {
	reg(&pg_QMimeData_base, "QMimeData_base")
	symaddr(&addr_QMimeData_new, "QMimeData_new")
	symaddr(&addr_QMimeData_tr, "QMimeData_tr")
	symaddr(&addr_QMimeData_urls, "QMimeData_urls")
	symaddr(&addr_QMimeData_setUrls, "QMimeData_setUrls")
	symaddr(&addr_QMimeData_hasUrls, "QMimeData_hasUrls")
	symaddr(&addr_QMimeData_text, "QMimeData_text")
	symaddr(&addr_QMimeData_setText, "QMimeData_setText")
	symaddr(&addr_QMimeData_hasText, "QMimeData_hasText")
	symaddr(&addr_QMimeData_html, "QMimeData_html")
	symaddr(&addr_QMimeData_setHtml, "QMimeData_setHtml")
	symaddr(&addr_QMimeData_hasHtml, "QMimeData_hasHtml")
	symaddr(&addr_QMimeData_imageData, "QMimeData_imageData")
	symaddr(&addr_QMimeData_setImageData, "QMimeData_setImageData")
	symaddr(&addr_QMimeData_hasImage, "QMimeData_hasImage")
	symaddr(&addr_QMimeData_colorData, "QMimeData_colorData")
	symaddr(&addr_QMimeData_setColorData, "QMimeData_setColorData")
	symaddr(&addr_QMimeData_hasColor, "QMimeData_hasColor")
	symaddr(&addr_QMimeData_data, "QMimeData_data")
	symaddr(&addr_QMimeData_setData, "QMimeData_setData")
	symaddr(&addr_QMimeData_removeFormat, "QMimeData_removeFormat")
	symaddr(&addr_QMimeData_hasFormat, "QMimeData_hasFormat")
	symaddr(&addr_QMimeData_formats, "QMimeData_formats")
	symaddr(&addr_QMimeData_clear, "QMimeData_clear")
	symaddr(&addr_QMimeData_connect, "QMimeData_connect")
	symaddr(&addr_QMimeData_connect2, "QMimeData_connect2")
	symaddr(&addr_QMimeData_sender, "QMimeData_sender")
	symaddr(&addr_QMimeData_senderSignalIndex, "QMimeData_senderSignalIndex")
	symaddr(&addr_QMimeData_receivers, "QMimeData_receivers")
	symaddr(&addr_QMimeData_isSignalConnected, "QMimeData_isSignalConnected")
	symaddr(&addr_QMimeData_tr2, "QMimeData_tr2")
	symaddr(&addr_QMimeData_tr3, "QMimeData_tr3")
	symaddr(&addr_QMimeData_connect3, "QMimeData_connect3")
	symaddr(&addr_QMimeData_connect4, "QMimeData_connect4")
	symaddr(&addr_QMimeData_hasFormat_default, "QMimeData_hasFormat_default")
	symaddr(&addr_QMimeData_formats_default, "QMimeData_formats_default")
	symaddr(&addr_QMimeData_metaObject_default, "QMimeData_metaObject_default")
	symaddr(&addr_QMimeData_retrieveData_default, "QMimeData_retrieveData_default")
	symaddr(&addr_QMimeData_event_default, "QMimeData_event_default")
	symaddr(&addr_QMimeData_eventFilter_default, "QMimeData_eventFilter_default")
	symaddr(&addr_QMimeData_timerEvent_default, "QMimeData_timerEvent_default")
	symaddr(&addr_QMimeData_childEvent_default, "QMimeData_childEvent_default")
	symaddr(&addr_QMimeData_customEvent_default, "QMimeData_customEvent_default")
	symaddr(&addr_QMimeData_connectNotify_default, "QMimeData_connectNotify_default")
	symaddr(&addr_QMimeData_disconnectNotify_default, "QMimeData_disconnectNotify_default")
	reg(&pg_QMimeData_connect_destroyed, "QMimeData_connect_destroyed")
	reg(&pg_QMimeData_connect_destroyed2, "QMimeData_connect_destroyed2")
	reg(&pg_QMimeData_connect_objectNameChanged, "QMimeData_connect_objectNameChanged")
	reg(&pg_QMimeData__disconnect, "QMimeData__disconnect")
	reg(&pg_QMimeData__set_override, "QMimeData__set_override")
	reg(&pg_QMimeData_delete, "QMimeData_delete")
}

type QMimeData struct {
	h uintptr
	*QObject
}

func bindQMimeData(h uintptr) *QMimeData {
	if h == 0 {
		return nil
	}
	return &QMimeData{h: h, QObject: bindQObject(pg_QMimeData_base(h))}
}

func newownedQMimeData(h uintptr) *QMimeData { return bindQMimeData(h) }

func (this *QMimeData) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QMimeData) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQMimeData(h unsafe.Pointer) *QMimeData { return bindQMimeData(uintptr(h)) }
func UnsafeNewQMimeDataOwned(h unsafe.Pointer) *QMimeData { return newownedQMimeData(uintptr(h)) }

func (this *QMimeData) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func NewQMimeData() *QMimeData {
	__goid := nextVirtID()
	__r, _, _ := purego.SyscallN(addr_QMimeData_new, uintptr(__goid))
	__o := newownedQMimeData(__r)
	__o.UnsafeSetGoID(__goid)
	registerInstance(__goid, __o.h, QMimeData_virtIdx, QMimeData_enable)
	return __o
}

func QMimeData_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeData_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QMimeData) Urls() []*QUrl {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeData_urls, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QUrl, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQUrl(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QMimeData) SetUrls(urls []*QUrl) {
	urls_arr := make([]uintptr, len(urls))
	for __i, __e := range urls {
		urls_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var urls_p uintptr
	if len(urls_arr) > 0 {
		urls_p = uintptr(unsafe.Pointer(&urls_arr[0]))
	}
	purego.SyscallN(addr_QMimeData_setUrls, this.h, urls_p, uintptr(len(urls)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(urls)
	runtime.KeepAlive(urls_arr)
}

func (this *QMimeData) HasUrls() bool {
	__r0, _, _ := purego.SyscallN(addr_QMimeData_hasUrls, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMimeData) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeData_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMimeData) SetText(text string) {
	var text_p uintptr
	if len(text) > 0 {
		text_p = uintptr(unsafe.Pointer(unsafe.StringData(text)))
	}
	purego.SyscallN(addr_QMimeData_setText, this.h, text_p, uintptr(len(text)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(text)
}

func (this *QMimeData) HasText() bool {
	__r0, _, _ := purego.SyscallN(addr_QMimeData_hasText, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMimeData) Html() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeData_html, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QMimeData) SetHtml(html string) {
	var html_p uintptr
	if len(html) > 0 {
		html_p = uintptr(unsafe.Pointer(unsafe.StringData(html)))
	}
	purego.SyscallN(addr_QMimeData_setHtml, this.h, html_p, uintptr(len(html)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(html)
}

func (this *QMimeData) HasHtml() bool {
	__r0, _, _ := purego.SyscallN(addr_QMimeData_hasHtml, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMimeData) ImageData() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QMimeData_imageData, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QMimeData) SetImageData(image *QVariant) {
	purego.SyscallN(addr_QMimeData_setImageData, this.h, uintptr(image.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(image)
}

func (this *QMimeData) HasImage() bool {
	__r0, _, _ := purego.SyscallN(addr_QMimeData_hasImage, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMimeData) ColorData() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QMimeData_colorData, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QMimeData) SetColorData(color *QVariant) {
	purego.SyscallN(addr_QMimeData_setColorData, this.h, uintptr(color.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(color)
}

func (this *QMimeData) HasColor() bool {
	__r0, _, _ := purego.SyscallN(addr_QMimeData_hasColor, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QMimeData) Data(mimetype string) *QByteArray {
	var mimetype_p uintptr
	if len(mimetype) > 0 {
		mimetype_p = uintptr(unsafe.Pointer(unsafe.StringData(mimetype)))
	}
	__r, _, _ := purego.SyscallN(addr_QMimeData_data, this.h, mimetype_p, uintptr(len(mimetype)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(mimetype)
	return newownedQByteArray(__r)
}

func (this *QMimeData) SetData(mimetype string, data *QByteArray) {
	var mimetype_p uintptr
	if len(mimetype) > 0 {
		mimetype_p = uintptr(unsafe.Pointer(unsafe.StringData(mimetype)))
	}
	purego.SyscallN(addr_QMimeData_setData, this.h, mimetype_p, uintptr(len(mimetype)), uintptr(data.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(mimetype)
	runtime.KeepAlive(data)
}

func (this *QMimeData) RemoveFormat(mimetype string) {
	var mimetype_p uintptr
	if len(mimetype) > 0 {
		mimetype_p = uintptr(unsafe.Pointer(unsafe.StringData(mimetype)))
	}
	purego.SyscallN(addr_QMimeData_removeFormat, this.h, mimetype_p, uintptr(len(mimetype)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(mimetype)
}

func (this *QMimeData) HasFormat(mimetype string) bool {
	var mimetype_p uintptr
	if len(mimetype) > 0 {
		mimetype_p = uintptr(unsafe.Pointer(unsafe.StringData(mimetype)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMimeData_hasFormat, this.h, mimetype_p, uintptr(len(mimetype)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(mimetype)
	return __r != 0
}

func (this *QMimeData) Formats() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeData_formats, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QMimeData) Clear() {
	purego.SyscallN(addr_QMimeData_clear, this.h)
	runtime.KeepAlive(this)
}

func QMimeData_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QMimeData_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QMimeData_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMimeData_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QMimeData) Sender() *QObject {
	__r, _, _ := purego.SyscallN(addr_QMimeData_sender, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QMimeData) SenderSignalIndex() int {
	__r0, _, _ := purego.SyscallN(addr_QMimeData_senderSignalIndex, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QMimeData) Receivers(signal string) int {
	signal_b := append([]byte(signal), 0)
	__r0, _, _ := purego.SyscallN(addr_QMimeData_receivers, this.h, uintptr(unsafe.Pointer(&signal_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal_b)
	return int(__r)
}

func (this *QMimeData) IsSignalConnected(signal *QMetaMethod) bool {
	__r0, _, _ := purego.SyscallN(addr_QMimeData_isSignalConnected, this.h, uintptr(signal.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
	return __r != 0
}

func QMimeData_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeData_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QMimeData_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeData_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QMimeData_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QMimeData_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QMimeData_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QMimeData_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QMimeData) HasFormatDefault(mimetype string) bool {
	var mimetype_p uintptr
	if len(mimetype) > 0 {
		mimetype_p = uintptr(unsafe.Pointer(unsafe.StringData(mimetype)))
	}
	__r0, _, _ := purego.SyscallN(addr_QMimeData_hasFormat_default, this.h, mimetype_p, uintptr(len(mimetype)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(mimetype)
	return __r != 0
}

func (this *QMimeData) FormatsDefault() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QMimeData_formats_default, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QMimeData) MetaObjectDefault() *QMetaObject {
	__r, _, _ := purego.SyscallN(addr_QMimeData_metaObject_default, this.h)
	runtime.KeepAlive(this)
	return bindQMetaObject(__r)
}

func (this *QMimeData) RetrieveDataDefault(mimetype string, preferredType *QMetaType) *QVariant {
	var mimetype_p uintptr
	if len(mimetype) > 0 {
		mimetype_p = uintptr(unsafe.Pointer(unsafe.StringData(mimetype)))
	}
	__r, _, _ := purego.SyscallN(addr_QMimeData_retrieveData_default, this.h, mimetype_p, uintptr(len(mimetype)), uintptr(preferredType.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(mimetype)
	runtime.KeepAlive(preferredType)
	return newownedQVariant(__r)
}

func (this *QMimeData) EventDefault(event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QMimeData_event_default, this.h, uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QMimeData) EventFilterDefault(watched *QObject, event *QEvent) bool {
	__r0, _, _ := purego.SyscallN(addr_QMimeData_eventFilter_default, this.h, uintptr(watched.UnsafePointer()), uintptr(event.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(watched)
	runtime.KeepAlive(event)
	return __r != 0
}

func (this *QMimeData) TimerEventDefault(event *QTimerEvent) {
	purego.SyscallN(addr_QMimeData_timerEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMimeData) ChildEventDefault(event *QChildEvent) {
	purego.SyscallN(addr_QMimeData_childEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMimeData) CustomEventDefault(event *QEvent) {
	purego.SyscallN(addr_QMimeData_customEvent_default, this.h, uintptr(event.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(event)
}

func (this *QMimeData) ConnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QMimeData_connectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QMimeData) DisconnectNotifyDefault(signal *QMetaMethod) {
	purego.SyscallN(addr_QMimeData_disconnectNotify_default, this.h, uintptr(signal.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(signal)
}

func (this *QMimeData) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QMimeData_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMimeData) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QMimeData_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMimeData) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QMimeData_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QMimeData) Disconnect(handle int64) {
	pg_QMimeData__disconnect(handle)
	unregisterCallback(handle)
}

func QMimeData_virtIdx(slot int64) int32 {
	switch slot {
	case 1607084412591819503:
		return 0
	case 7523410874251807628:
		return 1
	case 6926839562432501171:
		return 2
	case -2679356254762874311:
		return 3
	case 1237236456312824991:
		return 4
	case 5217623514630858297:
		return 5
	case -5409927581812374498:
		return 6
	case -1185315367526857833:
		return 7
	case 3769868229123868474:
		return 8
	case 7379029982170503008:
		return 9
	case -568369772982215962:
		return 10
	}
	return -1
}

func QMimeData_enable(h uintptr, idx int32) { pg_QMimeData__set_override(h, idx) }

func (this *QMimeData) OnMetaObject(fn func() *QMetaObject) {
	registerVirt(this.UnsafeGoID(), 1607084412591819503, func() uintptr { return uintptr(fn().UnsafePointer()) }) // slot "metaObject"
}

func (this *QMimeData) OnHasFormat(fn func(string) bool) {
	registerVirt(this.UnsafeGoID(), 7523410874251807628, func(mimetype uintptr, mimetype_len int64) int32 { return b2i(fn(pgStrNoFree(mimetype, mimetype_len))) }) // slot "hasFormat"
}

func (this *QMimeData) OnFormats(fn func() []string) {
	registerVirt(this.UnsafeGoID(), 6926839562432501171, func() (uintptr, int64) {
		__rs := fn()
		if len(__rs) == 0 {
			return 0, 0
		}
		__a := pgMalloc(uintptr(len(__rs)) * unsafe.Sizeof(uintptr(0)))
		__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__a)), len(__rs))
		for __i, __e := range __rs {
			__sl[__i] = pgCString(__e)
		}
		return __a, int64(len(__rs))
	}) // slot "formats"
}

func (this *QMimeData) OnRetrieveData(fn func(string, *QMetaType) *QVariant) {
	registerVirt(this.UnsafeGoID(), -2679356254762874311, func(mimetype uintptr, mimetype_len int64, preferredType uintptr) uintptr { return uintptr(fn(pgStrNoFree(mimetype, mimetype_len), UnsafeNewQMetaType(unsafe.Pointer(preferredType))).UnsafePointer()) }) // slot "retrieveData"
}

func (this *QMimeData) OnEvent(fn func(*QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 1237236456312824991, func(event uintptr) int32 { return b2i(fn(UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "event"
}

func (this *QMimeData) OnEventFilter(fn func(*QObject, *QEvent) bool) {
	registerVirt(this.UnsafeGoID(), 5217623514630858297, func(watched uintptr, event uintptr) int32 { return b2i(fn(UnsafeNewQObject(unsafe.Pointer(watched)), UnsafeNewQEvent(unsafe.Pointer(event)))) }) // slot "eventFilter"
}

func (this *QMimeData) OnTimerEvent(fn func(*QTimerEvent)) {
	registerVirt(this.UnsafeGoID(), -5409927581812374498, func(event uintptr) { fn(UnsafeNewQTimerEvent(unsafe.Pointer(event))) }) // slot "timerEvent"
}

func (this *QMimeData) OnChildEvent(fn func(*QChildEvent)) {
	registerVirt(this.UnsafeGoID(), -1185315367526857833, func(event uintptr) { fn(UnsafeNewQChildEvent(unsafe.Pointer(event))) }) // slot "childEvent"
}

func (this *QMimeData) OnCustomEvent(fn func(*QEvent)) {
	registerVirt(this.UnsafeGoID(), 3769868229123868474, func(event uintptr) { fn(UnsafeNewQEvent(unsafe.Pointer(event))) }) // slot "customEvent"
}

func (this *QMimeData) OnConnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), 7379029982170503008, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "connectNotify"
}

func (this *QMimeData) OnDisconnectNotify(fn func(*QMetaMethod)) {
	registerVirt(this.UnsafeGoID(), -568369772982215962, func(signal uintptr) { fn(UnsafeNewQMetaMethod(unsafe.Pointer(signal))) }) // slot "disconnectNotify"
}

func (this *QMimeData) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QMimeData_delete(this.h)
	}
	this.UnsafeRelease()
}
