package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QClipboard_base func(uintptr) uintptr
	addr_QClipboard_tr uintptr
	addr_QClipboard_clear uintptr
	addr_QClipboard_supportsSelection uintptr
	addr_QClipboard_supportsFindBuffer uintptr
	addr_QClipboard_ownsSelection uintptr
	addr_QClipboard_ownsClipboard uintptr
	addr_QClipboard_ownsFindBuffer uintptr
	addr_QClipboard_text uintptr
	addr_QClipboard_text2 uintptr
	addr_QClipboard_setText uintptr
	addr_QClipboard_mimeData uintptr
	addr_QClipboard_setMimeData uintptr
	addr_QClipboard_image uintptr
	addr_QClipboard_pixmap uintptr
	addr_QClipboard_setImage uintptr
	addr_QClipboard_setPixmap uintptr
	addr_QClipboard_changed uintptr
	addr_QClipboard_selectionChanged uintptr
	addr_QClipboard_findBufferChanged uintptr
	addr_QClipboard_dataChanged uintptr
	addr_QClipboard_connect uintptr
	addr_QClipboard_connect2 uintptr
	addr_QClipboard_tr2 uintptr
	addr_QClipboard_tr3 uintptr
	addr_QClipboard_clear2 uintptr
	addr_QClipboard_text3 uintptr
	addr_QClipboard_text4 uintptr
	addr_QClipboard_setText2 uintptr
	addr_QClipboard_mimeData2 uintptr
	addr_QClipboard_setMimeData2 uintptr
	addr_QClipboard_image2 uintptr
	addr_QClipboard_pixmap2 uintptr
	addr_QClipboard_setImage2 uintptr
	addr_QClipboard_setPixmap2 uintptr
	addr_QClipboard_connect3 uintptr
	addr_QClipboard_connect4 uintptr
	pg_QClipboard_connect_changed func(uintptr, int64)
	pg_QClipboard_connect_selectionChanged func(uintptr, int64)
	pg_QClipboard_connect_findBufferChanged func(uintptr, int64)
	pg_QClipboard_connect_dataChanged func(uintptr, int64)
	pg_QClipboard_connect_destroyed func(uintptr, int64)
	pg_QClipboard_connect_destroyed2 func(uintptr, int64)
	pg_QClipboard_connect_objectNameChanged func(uintptr, int64)
	pg_QClipboard__disconnect func(int64)
)

func init() {
	reg(&pg_QClipboard_base, "QClipboard_base")
	symaddr(&addr_QClipboard_tr, "QClipboard_tr")
	symaddr(&addr_QClipboard_clear, "QClipboard_clear")
	symaddr(&addr_QClipboard_supportsSelection, "QClipboard_supportsSelection")
	symaddr(&addr_QClipboard_supportsFindBuffer, "QClipboard_supportsFindBuffer")
	symaddr(&addr_QClipboard_ownsSelection, "QClipboard_ownsSelection")
	symaddr(&addr_QClipboard_ownsClipboard, "QClipboard_ownsClipboard")
	symaddr(&addr_QClipboard_ownsFindBuffer, "QClipboard_ownsFindBuffer")
	symaddr(&addr_QClipboard_text, "QClipboard_text")
	symaddr(&addr_QClipboard_text2, "QClipboard_text2")
	symaddr(&addr_QClipboard_setText, "QClipboard_setText")
	symaddr(&addr_QClipboard_mimeData, "QClipboard_mimeData")
	symaddr(&addr_QClipboard_setMimeData, "QClipboard_setMimeData")
	symaddr(&addr_QClipboard_image, "QClipboard_image")
	symaddr(&addr_QClipboard_pixmap, "QClipboard_pixmap")
	symaddr(&addr_QClipboard_setImage, "QClipboard_setImage")
	symaddr(&addr_QClipboard_setPixmap, "QClipboard_setPixmap")
	symaddr(&addr_QClipboard_changed, "QClipboard_changed")
	symaddr(&addr_QClipboard_selectionChanged, "QClipboard_selectionChanged")
	symaddr(&addr_QClipboard_findBufferChanged, "QClipboard_findBufferChanged")
	symaddr(&addr_QClipboard_dataChanged, "QClipboard_dataChanged")
	symaddr(&addr_QClipboard_connect, "QClipboard_connect")
	symaddr(&addr_QClipboard_connect2, "QClipboard_connect2")
	symaddr(&addr_QClipboard_tr2, "QClipboard_tr2")
	symaddr(&addr_QClipboard_tr3, "QClipboard_tr3")
	symaddr(&addr_QClipboard_clear2, "QClipboard_clear2")
	symaddr(&addr_QClipboard_text3, "QClipboard_text3")
	symaddr(&addr_QClipboard_text4, "QClipboard_text4")
	symaddr(&addr_QClipboard_setText2, "QClipboard_setText2")
	symaddr(&addr_QClipboard_mimeData2, "QClipboard_mimeData2")
	symaddr(&addr_QClipboard_setMimeData2, "QClipboard_setMimeData2")
	symaddr(&addr_QClipboard_image2, "QClipboard_image2")
	symaddr(&addr_QClipboard_pixmap2, "QClipboard_pixmap2")
	symaddr(&addr_QClipboard_setImage2, "QClipboard_setImage2")
	symaddr(&addr_QClipboard_setPixmap2, "QClipboard_setPixmap2")
	symaddr(&addr_QClipboard_connect3, "QClipboard_connect3")
	symaddr(&addr_QClipboard_connect4, "QClipboard_connect4")
	reg(&pg_QClipboard_connect_changed, "QClipboard_connect_changed")
	reg(&pg_QClipboard_connect_selectionChanged, "QClipboard_connect_selectionChanged")
	reg(&pg_QClipboard_connect_findBufferChanged, "QClipboard_connect_findBufferChanged")
	reg(&pg_QClipboard_connect_dataChanged, "QClipboard_connect_dataChanged")
	reg(&pg_QClipboard_connect_destroyed, "QClipboard_connect_destroyed")
	reg(&pg_QClipboard_connect_destroyed2, "QClipboard_connect_destroyed2")
	reg(&pg_QClipboard_connect_objectNameChanged, "QClipboard_connect_objectNameChanged")
	reg(&pg_QClipboard__disconnect, "QClipboard__disconnect")
}

type QClipboard struct {
	h uintptr
	*QObject
}

func bindQClipboard(h uintptr) *QClipboard {
	if h == 0 {
		return nil
	}
	return &QClipboard{h: h, QObject: bindQObject(pg_QClipboard_base(h))}
}

func newownedQClipboard(h uintptr) *QClipboard { return bindQClipboard(h) }

func (this *QClipboard) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QClipboard) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQClipboard(h unsafe.Pointer) *QClipboard { return bindQClipboard(uintptr(h)) }
func UnsafeNewQClipboardOwned(h unsafe.Pointer) *QClipboard { return newownedQClipboard(uintptr(h)) }

func (this *QClipboard) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QClipboard_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QClipboard_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QClipboard) Clear() {
	purego.SyscallN(addr_QClipboard_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QClipboard) SupportsSelection() bool {
	__r0, _, _ := purego.SyscallN(addr_QClipboard_supportsSelection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QClipboard) SupportsFindBuffer() bool {
	__r0, _, _ := purego.SyscallN(addr_QClipboard_supportsFindBuffer, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QClipboard) OwnsSelection() bool {
	__r0, _, _ := purego.SyscallN(addr_QClipboard_ownsSelection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QClipboard) OwnsClipboard() bool {
	__r0, _, _ := purego.SyscallN(addr_QClipboard_ownsClipboard, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QClipboard) OwnsFindBuffer() bool {
	__r0, _, _ := purego.SyscallN(addr_QClipboard_ownsFindBuffer, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QClipboard) Text() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QClipboard_text, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QClipboard) TextWithSubtype(subtype *QString) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QClipboard_text2, this.h, uintptr(subtype.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subtype)
	return pgStr(__r, __n)
}

func (this *QClipboard) SetText(arg1 string) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QClipboard_setText, this.h, arg1_p, uintptr(len(arg1)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QClipboard) MimeData() *QMimeData {
	__r, _, _ := purego.SyscallN(addr_QClipboard_mimeData, this.h)
	runtime.KeepAlive(this)
	return bindQMimeData(__r)
}

func (this *QClipboard) SetMimeData(data *QMimeData) {
	purego.SyscallN(addr_QClipboard_setMimeData, this.h, uintptr(data.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
}

func (this *QClipboard) Image() *QImage {
	__r, _, _ := purego.SyscallN(addr_QClipboard_image, this.h)
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QClipboard) Pixmap() *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QClipboard_pixmap, this.h)
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QClipboard) SetImage(arg1 *QImage) {
	purego.SyscallN(addr_QClipboard_setImage, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QClipboard) SetPixmap(arg1 *QPixmap) {
	purego.SyscallN(addr_QClipboard_setPixmap, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QClipboard) Changed(mode QClipboard__Mode) {
	purego.SyscallN(addr_QClipboard_changed, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QClipboard) SelectionChanged() {
	purego.SyscallN(addr_QClipboard_selectionChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QClipboard) FindBufferChanged() {
	purego.SyscallN(addr_QClipboard_findBufferChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QClipboard) DataChanged() {
	purego.SyscallN(addr_QClipboard_dataChanged, this.h)
	runtime.KeepAlive(this)
}

func QClipboard_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QClipboard_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QClipboard_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QClipboard_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QClipboard_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QClipboard_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QClipboard_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QClipboard_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QClipboard) ClearWithMode(mode QClipboard__Mode) {
	purego.SyscallN(addr_QClipboard_clear2, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QClipboard) TextWithMode(mode QClipboard__Mode) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QClipboard_text3, this.h, uintptr(mode), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QClipboard) Text2(subtype *QString, mode QClipboard__Mode) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QClipboard_text4, this.h, uintptr(subtype.UnsafePointer()), uintptr(mode), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(subtype)
	return pgStr(__r, __n)
}

func (this *QClipboard) SetText2(arg1 string, mode QClipboard__Mode) {
	var arg1_p uintptr
	if len(arg1) > 0 {
		arg1_p = uintptr(unsafe.Pointer(unsafe.StringData(arg1)))
	}
	purego.SyscallN(addr_QClipboard_setText2, this.h, arg1_p, uintptr(len(arg1)), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QClipboard) MimeDataWithMode(mode QClipboard__Mode) *QMimeData {
	__r, _, _ := purego.SyscallN(addr_QClipboard_mimeData2, this.h, uintptr(mode))
	runtime.KeepAlive(this)
	return bindQMimeData(__r)
}

func (this *QClipboard) SetMimeData2(data *QMimeData, mode QClipboard__Mode) {
	purego.SyscallN(addr_QClipboard_setMimeData2, this.h, uintptr(data.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(data)
}

func (this *QClipboard) ImageWithMode(mode QClipboard__Mode) *QImage {
	__r, _, _ := purego.SyscallN(addr_QClipboard_image2, this.h, uintptr(mode))
	runtime.KeepAlive(this)
	return newownedQImage(__r)
}

func (this *QClipboard) PixmapWithMode(mode QClipboard__Mode) *QPixmap {
	__r, _, _ := purego.SyscallN(addr_QClipboard_pixmap2, this.h, uintptr(mode))
	runtime.KeepAlive(this)
	return newownedQPixmap(__r)
}

func (this *QClipboard) SetImage2(arg1 *QImage, mode QClipboard__Mode) {
	purego.SyscallN(addr_QClipboard_setImage2, this.h, uintptr(arg1.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QClipboard) SetPixmap2(arg1 *QPixmap, mode QClipboard__Mode) {
	purego.SyscallN(addr_QClipboard_setPixmap2, this.h, uintptr(arg1.UnsafePointer()), uintptr(mode))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QClipboard_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QClipboard_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QClipboard_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QClipboard_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QClipboard) OnChanged(fn func(QClipboard__Mode)) int64 {
	id := registerCallback(func(mode int32) { fn(QClipboard__Mode(mode)) })
	pg_QClipboard_connect_changed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QClipboard) OnSelectionChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QClipboard_connect_selectionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QClipboard) OnFindBufferChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QClipboard_connect_findBufferChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QClipboard) OnDataChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QClipboard_connect_dataChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QClipboard) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QClipboard_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QClipboard) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QClipboard_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QClipboard) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QClipboard_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QClipboard) Disconnect(handle int64) {
	pg_QClipboard__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QClipboard) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	this.UnsafeRelease()
}
