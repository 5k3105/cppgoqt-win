package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QAbstractTextDocumentLayout_base func(uintptr) uintptr
	addr_QAbstractTextDocumentLayout_tr uintptr
	addr_QAbstractTextDocumentLayout_draw uintptr
	addr_QAbstractTextDocumentLayout_hitTest uintptr
	addr_QAbstractTextDocumentLayout_anchorAt uintptr
	addr_QAbstractTextDocumentLayout_imageAt uintptr
	addr_QAbstractTextDocumentLayout_formatAt uintptr
	addr_QAbstractTextDocumentLayout_blockWithMarkerAt uintptr
	addr_QAbstractTextDocumentLayout_pageCount uintptr
	addr_QAbstractTextDocumentLayout_documentSize uintptr
	addr_QAbstractTextDocumentLayout_frameBoundingRect uintptr
	addr_QAbstractTextDocumentLayout_blockBoundingRect uintptr
	addr_QAbstractTextDocumentLayout_setPaintDevice uintptr
	addr_QAbstractTextDocumentLayout_paintDevice uintptr
	addr_QAbstractTextDocumentLayout_document uintptr
	addr_QAbstractTextDocumentLayout_registerHandler uintptr
	addr_QAbstractTextDocumentLayout_unregisterHandler uintptr
	addr_QAbstractTextDocumentLayout_handlerForObject uintptr
	addr_QAbstractTextDocumentLayout_update uintptr
	addr_QAbstractTextDocumentLayout_updateBlock uintptr
	addr_QAbstractTextDocumentLayout_documentSizeChanged uintptr
	addr_QAbstractTextDocumentLayout_pageCountChanged uintptr
	addr_QAbstractTextDocumentLayout_connect uintptr
	addr_QAbstractTextDocumentLayout_connect2 uintptr
	addr_QAbstractTextDocumentLayout_tr2 uintptr
	addr_QAbstractTextDocumentLayout_tr3 uintptr
	addr_QAbstractTextDocumentLayout_unregisterHandler2 uintptr
	addr_QAbstractTextDocumentLayout_update2 uintptr
	addr_QAbstractTextDocumentLayout_connect3 uintptr
	addr_QAbstractTextDocumentLayout_connect4 uintptr
	pg_QAbstractTextDocumentLayout_connect_update func(uintptr, int64)
	pg_QAbstractTextDocumentLayout_connect_update2 func(uintptr, int64)
	pg_QAbstractTextDocumentLayout_connect_updateBlock func(uintptr, int64)
	pg_QAbstractTextDocumentLayout_connect_documentSizeChanged func(uintptr, int64)
	pg_QAbstractTextDocumentLayout_connect_pageCountChanged func(uintptr, int64)
	pg_QAbstractTextDocumentLayout_connect_destroyed func(uintptr, int64)
	pg_QAbstractTextDocumentLayout_connect_destroyed2 func(uintptr, int64)
	pg_QAbstractTextDocumentLayout_connect_objectNameChanged func(uintptr, int64)
	pg_QAbstractTextDocumentLayout__disconnect func(int64)
	pg_QAbstractTextDocumentLayout_delete func(uintptr)
)

func init() {
	reg(&pg_QAbstractTextDocumentLayout_base, "QAbstractTextDocumentLayout_base")
	symaddr(&addr_QAbstractTextDocumentLayout_tr, "QAbstractTextDocumentLayout_tr")
	symaddr(&addr_QAbstractTextDocumentLayout_draw, "QAbstractTextDocumentLayout_draw")
	symaddr(&addr_QAbstractTextDocumentLayout_hitTest, "QAbstractTextDocumentLayout_hitTest")
	symaddr(&addr_QAbstractTextDocumentLayout_anchorAt, "QAbstractTextDocumentLayout_anchorAt")
	symaddr(&addr_QAbstractTextDocumentLayout_imageAt, "QAbstractTextDocumentLayout_imageAt")
	symaddr(&addr_QAbstractTextDocumentLayout_formatAt, "QAbstractTextDocumentLayout_formatAt")
	symaddr(&addr_QAbstractTextDocumentLayout_blockWithMarkerAt, "QAbstractTextDocumentLayout_blockWithMarkerAt")
	symaddr(&addr_QAbstractTextDocumentLayout_pageCount, "QAbstractTextDocumentLayout_pageCount")
	symaddr(&addr_QAbstractTextDocumentLayout_documentSize, "QAbstractTextDocumentLayout_documentSize")
	symaddr(&addr_QAbstractTextDocumentLayout_frameBoundingRect, "QAbstractTextDocumentLayout_frameBoundingRect")
	symaddr(&addr_QAbstractTextDocumentLayout_blockBoundingRect, "QAbstractTextDocumentLayout_blockBoundingRect")
	symaddr(&addr_QAbstractTextDocumentLayout_setPaintDevice, "QAbstractTextDocumentLayout_setPaintDevice")
	symaddr(&addr_QAbstractTextDocumentLayout_paintDevice, "QAbstractTextDocumentLayout_paintDevice")
	symaddr(&addr_QAbstractTextDocumentLayout_document, "QAbstractTextDocumentLayout_document")
	symaddr(&addr_QAbstractTextDocumentLayout_registerHandler, "QAbstractTextDocumentLayout_registerHandler")
	symaddr(&addr_QAbstractTextDocumentLayout_unregisterHandler, "QAbstractTextDocumentLayout_unregisterHandler")
	symaddr(&addr_QAbstractTextDocumentLayout_handlerForObject, "QAbstractTextDocumentLayout_handlerForObject")
	symaddr(&addr_QAbstractTextDocumentLayout_update, "QAbstractTextDocumentLayout_update")
	symaddr(&addr_QAbstractTextDocumentLayout_updateBlock, "QAbstractTextDocumentLayout_updateBlock")
	symaddr(&addr_QAbstractTextDocumentLayout_documentSizeChanged, "QAbstractTextDocumentLayout_documentSizeChanged")
	symaddr(&addr_QAbstractTextDocumentLayout_pageCountChanged, "QAbstractTextDocumentLayout_pageCountChanged")
	symaddr(&addr_QAbstractTextDocumentLayout_connect, "QAbstractTextDocumentLayout_connect")
	symaddr(&addr_QAbstractTextDocumentLayout_connect2, "QAbstractTextDocumentLayout_connect2")
	symaddr(&addr_QAbstractTextDocumentLayout_tr2, "QAbstractTextDocumentLayout_tr2")
	symaddr(&addr_QAbstractTextDocumentLayout_tr3, "QAbstractTextDocumentLayout_tr3")
	symaddr(&addr_QAbstractTextDocumentLayout_unregisterHandler2, "QAbstractTextDocumentLayout_unregisterHandler2")
	symaddr(&addr_QAbstractTextDocumentLayout_update2, "QAbstractTextDocumentLayout_update2")
	symaddr(&addr_QAbstractTextDocumentLayout_connect3, "QAbstractTextDocumentLayout_connect3")
	symaddr(&addr_QAbstractTextDocumentLayout_connect4, "QAbstractTextDocumentLayout_connect4")
	reg(&pg_QAbstractTextDocumentLayout_connect_update, "QAbstractTextDocumentLayout_connect_update")
	reg(&pg_QAbstractTextDocumentLayout_connect_update2, "QAbstractTextDocumentLayout_connect_update2")
	reg(&pg_QAbstractTextDocumentLayout_connect_updateBlock, "QAbstractTextDocumentLayout_connect_updateBlock")
	reg(&pg_QAbstractTextDocumentLayout_connect_documentSizeChanged, "QAbstractTextDocumentLayout_connect_documentSizeChanged")
	reg(&pg_QAbstractTextDocumentLayout_connect_pageCountChanged, "QAbstractTextDocumentLayout_connect_pageCountChanged")
	reg(&pg_QAbstractTextDocumentLayout_connect_destroyed, "QAbstractTextDocumentLayout_connect_destroyed")
	reg(&pg_QAbstractTextDocumentLayout_connect_destroyed2, "QAbstractTextDocumentLayout_connect_destroyed2")
	reg(&pg_QAbstractTextDocumentLayout_connect_objectNameChanged, "QAbstractTextDocumentLayout_connect_objectNameChanged")
	reg(&pg_QAbstractTextDocumentLayout__disconnect, "QAbstractTextDocumentLayout__disconnect")
	reg(&pg_QAbstractTextDocumentLayout_delete, "QAbstractTextDocumentLayout_delete")
}

type QAbstractTextDocumentLayout struct {
	h uintptr
	*QObject
}

func bindQAbstractTextDocumentLayout(h uintptr) *QAbstractTextDocumentLayout {
	if h == 0 {
		return nil
	}
	return &QAbstractTextDocumentLayout{h: h, QObject: bindQObject(pg_QAbstractTextDocumentLayout_base(h))}
}

func newownedQAbstractTextDocumentLayout(h uintptr) *QAbstractTextDocumentLayout { return bindQAbstractTextDocumentLayout(h) }

func (this *QAbstractTextDocumentLayout) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QAbstractTextDocumentLayout) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQAbstractTextDocumentLayout(h unsafe.Pointer) *QAbstractTextDocumentLayout { return bindQAbstractTextDocumentLayout(uintptr(h)) }
func UnsafeNewQAbstractTextDocumentLayoutOwned(h unsafe.Pointer) *QAbstractTextDocumentLayout { return newownedQAbstractTextDocumentLayout(uintptr(h)) }

func (this *QAbstractTextDocumentLayout) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QAbstractTextDocumentLayout_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QAbstractTextDocumentLayout) Draw(painter *QPainter, context *QAbstractTextDocumentLayout__PaintContext) {
	purego.SyscallN(addr_QAbstractTextDocumentLayout_draw, this.h, uintptr(painter.UnsafePointer()), uintptr(context.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(painter)
	runtime.KeepAlive(context)
}

func (this *QAbstractTextDocumentLayout) HitTest(point *QPointF, accuracy HitTestAccuracy) int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_hitTest, this.h, uintptr(point.UnsafePointer()), uintptr(accuracy))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(point)
	return int(__r)
}

func (this *QAbstractTextDocumentLayout) AnchorAt(pos *QPointF) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_anchorAt, this.h, uintptr(pos.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return pgStr(__r, __n)
}

func (this *QAbstractTextDocumentLayout) ImageAt(pos *QPointF) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_imageAt, this.h, uintptr(pos.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return pgStr(__r, __n)
}

func (this *QAbstractTextDocumentLayout) FormatAt(pos *QPointF) *QTextFormat {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_formatAt, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return newownedQTextFormat(__r)
}

func (this *QAbstractTextDocumentLayout) BlockWithMarkerAt(pos *QPointF) *QTextBlock {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_blockWithMarkerAt, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
	return newownedQTextBlock(__r)
}

func (this *QAbstractTextDocumentLayout) PageCount() int {
	__r0, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_pageCount, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QAbstractTextDocumentLayout) DocumentSize() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_documentSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QAbstractTextDocumentLayout) FrameBoundingRect(frame *QTextFrame) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_frameBoundingRect, this.h, uintptr(frame.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(frame)
	return newownedQRectF(__r)
}

func (this *QAbstractTextDocumentLayout) BlockBoundingRect(block *QTextBlock) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_blockBoundingRect, this.h, uintptr(block.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(block)
	return newownedQRectF(__r)
}

func (this *QAbstractTextDocumentLayout) SetPaintDevice(device *QPaintDevice) {
	purego.SyscallN(addr_QAbstractTextDocumentLayout_setPaintDevice, this.h, uintptr(device.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(device)
}

func (this *QAbstractTextDocumentLayout) PaintDevice() *QPaintDevice {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_paintDevice, this.h)
	runtime.KeepAlive(this)
	return bindQPaintDevice(__r)
}

func (this *QAbstractTextDocumentLayout) Document() *QTextDocument {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_document, this.h)
	runtime.KeepAlive(this)
	return bindQTextDocument(__r)
}

func (this *QAbstractTextDocumentLayout) RegisterHandler(objectType int, component *QObject) {
	purego.SyscallN(addr_QAbstractTextDocumentLayout_registerHandler, this.h, uintptr(objectType), uintptr(component.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(component)
}

func (this *QAbstractTextDocumentLayout) UnregisterHandler(objectType int) {
	purego.SyscallN(addr_QAbstractTextDocumentLayout_unregisterHandler, this.h, uintptr(objectType))
	runtime.KeepAlive(this)
}

func (this *QAbstractTextDocumentLayout) HandlerForObject(objectType int) *QTextObjectInterface {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_handlerForObject, this.h, uintptr(objectType))
	runtime.KeepAlive(this)
	return bindQTextObjectInterface(__r)
}

func (this *QAbstractTextDocumentLayout) Update() {
	purego.SyscallN(addr_QAbstractTextDocumentLayout_update, this.h)
	runtime.KeepAlive(this)
}

func (this *QAbstractTextDocumentLayout) UpdateBlock(block *QTextBlock) {
	purego.SyscallN(addr_QAbstractTextDocumentLayout_updateBlock, this.h, uintptr(block.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(block)
}

func (this *QAbstractTextDocumentLayout) DocumentSizeChanged(newSize *QSizeF) {
	purego.SyscallN(addr_QAbstractTextDocumentLayout_documentSizeChanged, this.h, uintptr(newSize.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(newSize)
}

func (this *QAbstractTextDocumentLayout) PageCountChanged(newPages int) {
	purego.SyscallN(addr_QAbstractTextDocumentLayout_pageCountChanged, this.h, uintptr(newPages))
	runtime.KeepAlive(this)
}

func QAbstractTextDocumentLayout_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractTextDocumentLayout_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractTextDocumentLayout_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QAbstractTextDocumentLayout_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QAbstractTextDocumentLayout) UnregisterHandler2(objectType int, component *QObject) {
	purego.SyscallN(addr_QAbstractTextDocumentLayout_unregisterHandler2, this.h, uintptr(objectType), uintptr(component.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(component)
}

func (this *QAbstractTextDocumentLayout) UpdateWithQRectF(arg1 *QRectF) {
	purego.SyscallN(addr_QAbstractTextDocumentLayout_update2, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QAbstractTextDocumentLayout_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QAbstractTextDocumentLayout_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QAbstractTextDocumentLayout_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QAbstractTextDocumentLayout) OnUpdate(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractTextDocumentLayout_connect_update(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTextDocumentLayout) OnUpdateWithQRectF(fn func(*QRectF)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQRectF(unsafe.Pointer(arg1))) })
	pg_QAbstractTextDocumentLayout_connect_update2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTextDocumentLayout) OnUpdateBlock(fn func(*QTextBlock)) int64 {
	id := registerCallback(func(block uintptr) { fn(UnsafeNewQTextBlock(unsafe.Pointer(block))) })
	pg_QAbstractTextDocumentLayout_connect_updateBlock(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTextDocumentLayout) OnDocumentSizeChanged(fn func(*QSizeF)) int64 {
	id := registerCallback(func(newSize uintptr) { fn(UnsafeNewQSizeF(unsafe.Pointer(newSize))) })
	pg_QAbstractTextDocumentLayout_connect_documentSizeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTextDocumentLayout) OnPageCountChanged(fn func(int)) int64 {
	id := registerCallback(func(newPages int32) { fn(int(newPages)) })
	pg_QAbstractTextDocumentLayout_connect_pageCountChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTextDocumentLayout) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QAbstractTextDocumentLayout_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTextDocumentLayout) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QAbstractTextDocumentLayout_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTextDocumentLayout) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QAbstractTextDocumentLayout_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QAbstractTextDocumentLayout) Disconnect(handle int64) {
	pg_QAbstractTextDocumentLayout__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QAbstractTextDocumentLayout) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QAbstractTextDocumentLayout_delete(this.h)
	}
	this.UnsafeRelease()
}
