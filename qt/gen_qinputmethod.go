package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QInputMethod_base func(uintptr) uintptr
	addr_QInputMethod_tr uintptr
	addr_QInputMethod_inputItemTransform uintptr
	addr_QInputMethod_setInputItemTransform uintptr
	addr_QInputMethod_inputItemRectangle uintptr
	addr_QInputMethod_setInputItemRectangle uintptr
	addr_QInputMethod_cursorRectangle uintptr
	addr_QInputMethod_anchorRectangle uintptr
	addr_QInputMethod_keyboardRectangle uintptr
	addr_QInputMethod_inputItemClipRectangle uintptr
	addr_QInputMethod_isVisible uintptr
	addr_QInputMethod_setVisible uintptr
	addr_QInputMethod_isAnimating uintptr
	addr_QInputMethod_locale uintptr
	addr_QInputMethod_inputDirection uintptr
	addr_QInputMethod_queryFocusObject uintptr
	addr_QInputMethod_show uintptr
	addr_QInputMethod_hide uintptr
	addr_QInputMethod_update uintptr
	addr_QInputMethod_reset uintptr
	addr_QInputMethod_commit uintptr
	addr_QInputMethod_invokeAction uintptr
	addr_QInputMethod_cursorRectangleChanged uintptr
	addr_QInputMethod_anchorRectangleChanged uintptr
	addr_QInputMethod_keyboardRectangleChanged uintptr
	addr_QInputMethod_inputItemClipRectangleChanged uintptr
	addr_QInputMethod_visibleChanged uintptr
	addr_QInputMethod_animatingChanged uintptr
	addr_QInputMethod_localeChanged uintptr
	addr_QInputMethod_inputDirectionChanged uintptr
	addr_QInputMethod_connect uintptr
	addr_QInputMethod_connect2 uintptr
	addr_QInputMethod_tr2 uintptr
	addr_QInputMethod_tr3 uintptr
	addr_QInputMethod_connect3 uintptr
	addr_QInputMethod_connect4 uintptr
	pg_QInputMethod_connect_cursorRectangleChanged func(uintptr, int64)
	pg_QInputMethod_connect_anchorRectangleChanged func(uintptr, int64)
	pg_QInputMethod_connect_keyboardRectangleChanged func(uintptr, int64)
	pg_QInputMethod_connect_inputItemClipRectangleChanged func(uintptr, int64)
	pg_QInputMethod_connect_visibleChanged func(uintptr, int64)
	pg_QInputMethod_connect_animatingChanged func(uintptr, int64)
	pg_QInputMethod_connect_localeChanged func(uintptr, int64)
	pg_QInputMethod_connect_inputDirectionChanged func(uintptr, int64)
	pg_QInputMethod_connect_destroyed func(uintptr, int64)
	pg_QInputMethod_connect_destroyed2 func(uintptr, int64)
	pg_QInputMethod_connect_objectNameChanged func(uintptr, int64)
	pg_QInputMethod__disconnect func(int64)
)

func init() {
	reg(&pg_QInputMethod_base, "QInputMethod_base")
	symaddr(&addr_QInputMethod_tr, "QInputMethod_tr")
	symaddr(&addr_QInputMethod_inputItemTransform, "QInputMethod_inputItemTransform")
	symaddr(&addr_QInputMethod_setInputItemTransform, "QInputMethod_setInputItemTransform")
	symaddr(&addr_QInputMethod_inputItemRectangle, "QInputMethod_inputItemRectangle")
	symaddr(&addr_QInputMethod_setInputItemRectangle, "QInputMethod_setInputItemRectangle")
	symaddr(&addr_QInputMethod_cursorRectangle, "QInputMethod_cursorRectangle")
	symaddr(&addr_QInputMethod_anchorRectangle, "QInputMethod_anchorRectangle")
	symaddr(&addr_QInputMethod_keyboardRectangle, "QInputMethod_keyboardRectangle")
	symaddr(&addr_QInputMethod_inputItemClipRectangle, "QInputMethod_inputItemClipRectangle")
	symaddr(&addr_QInputMethod_isVisible, "QInputMethod_isVisible")
	symaddr(&addr_QInputMethod_setVisible, "QInputMethod_setVisible")
	symaddr(&addr_QInputMethod_isAnimating, "QInputMethod_isAnimating")
	symaddr(&addr_QInputMethod_locale, "QInputMethod_locale")
	symaddr(&addr_QInputMethod_inputDirection, "QInputMethod_inputDirection")
	symaddr(&addr_QInputMethod_queryFocusObject, "QInputMethod_queryFocusObject")
	symaddr(&addr_QInputMethod_show, "QInputMethod_show")
	symaddr(&addr_QInputMethod_hide, "QInputMethod_hide")
	symaddr(&addr_QInputMethod_update, "QInputMethod_update")
	symaddr(&addr_QInputMethod_reset, "QInputMethod_reset")
	symaddr(&addr_QInputMethod_commit, "QInputMethod_commit")
	symaddr(&addr_QInputMethod_invokeAction, "QInputMethod_invokeAction")
	symaddr(&addr_QInputMethod_cursorRectangleChanged, "QInputMethod_cursorRectangleChanged")
	symaddr(&addr_QInputMethod_anchorRectangleChanged, "QInputMethod_anchorRectangleChanged")
	symaddr(&addr_QInputMethod_keyboardRectangleChanged, "QInputMethod_keyboardRectangleChanged")
	symaddr(&addr_QInputMethod_inputItemClipRectangleChanged, "QInputMethod_inputItemClipRectangleChanged")
	symaddr(&addr_QInputMethod_visibleChanged, "QInputMethod_visibleChanged")
	symaddr(&addr_QInputMethod_animatingChanged, "QInputMethod_animatingChanged")
	symaddr(&addr_QInputMethod_localeChanged, "QInputMethod_localeChanged")
	symaddr(&addr_QInputMethod_inputDirectionChanged, "QInputMethod_inputDirectionChanged")
	symaddr(&addr_QInputMethod_connect, "QInputMethod_connect")
	symaddr(&addr_QInputMethod_connect2, "QInputMethod_connect2")
	symaddr(&addr_QInputMethod_tr2, "QInputMethod_tr2")
	symaddr(&addr_QInputMethod_tr3, "QInputMethod_tr3")
	symaddr(&addr_QInputMethod_connect3, "QInputMethod_connect3")
	symaddr(&addr_QInputMethod_connect4, "QInputMethod_connect4")
	reg(&pg_QInputMethod_connect_cursorRectangleChanged, "QInputMethod_connect_cursorRectangleChanged")
	reg(&pg_QInputMethod_connect_anchorRectangleChanged, "QInputMethod_connect_anchorRectangleChanged")
	reg(&pg_QInputMethod_connect_keyboardRectangleChanged, "QInputMethod_connect_keyboardRectangleChanged")
	reg(&pg_QInputMethod_connect_inputItemClipRectangleChanged, "QInputMethod_connect_inputItemClipRectangleChanged")
	reg(&pg_QInputMethod_connect_visibleChanged, "QInputMethod_connect_visibleChanged")
	reg(&pg_QInputMethod_connect_animatingChanged, "QInputMethod_connect_animatingChanged")
	reg(&pg_QInputMethod_connect_localeChanged, "QInputMethod_connect_localeChanged")
	reg(&pg_QInputMethod_connect_inputDirectionChanged, "QInputMethod_connect_inputDirectionChanged")
	reg(&pg_QInputMethod_connect_destroyed, "QInputMethod_connect_destroyed")
	reg(&pg_QInputMethod_connect_destroyed2, "QInputMethod_connect_destroyed2")
	reg(&pg_QInputMethod_connect_objectNameChanged, "QInputMethod_connect_objectNameChanged")
	reg(&pg_QInputMethod__disconnect, "QInputMethod__disconnect")
}

type QInputMethod struct {
	h uintptr
	*QObject
}

func bindQInputMethod(h uintptr) *QInputMethod {
	if h == 0 {
		return nil
	}
	return &QInputMethod{h: h, QObject: bindQObject(pg_QInputMethod_base(h))}
}

func newownedQInputMethod(h uintptr) *QInputMethod { return bindQInputMethod(h) }

func (this *QInputMethod) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QInputMethod) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQInputMethod(h unsafe.Pointer) *QInputMethod { return bindQInputMethod(uintptr(h)) }
func UnsafeNewQInputMethodOwned(h unsafe.Pointer) *QInputMethod { return newownedQInputMethod(uintptr(h)) }

func (this *QInputMethod) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QInputMethod_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputMethod_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QInputMethod) InputItemTransform() *QTransform {
	__r, _, _ := purego.SyscallN(addr_QInputMethod_inputItemTransform, this.h)
	runtime.KeepAlive(this)
	return newownedQTransform(__r)
}

func (this *QInputMethod) SetInputItemTransform(transform *QTransform) {
	purego.SyscallN(addr_QInputMethod_setInputItemTransform, this.h, uintptr(transform.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transform)
}

func (this *QInputMethod) InputItemRectangle() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QInputMethod_inputItemRectangle, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QInputMethod) SetInputItemRectangle(rect *QRectF) {
	purego.SyscallN(addr_QInputMethod_setInputItemRectangle, this.h, uintptr(rect.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QInputMethod) CursorRectangle() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QInputMethod_cursorRectangle, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QInputMethod) AnchorRectangle() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QInputMethod_anchorRectangle, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QInputMethod) KeyboardRectangle() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QInputMethod_keyboardRectangle, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QInputMethod) InputItemClipRectangle() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QInputMethod_inputItemClipRectangle, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QInputMethod) IsVisible() bool {
	__r0, _, _ := purego.SyscallN(addr_QInputMethod_isVisible, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QInputMethod) SetVisible(visible bool) {
	purego.SyscallN(addr_QInputMethod_setVisible, this.h, uintptr(b2i(visible)))
	runtime.KeepAlive(this)
}

func (this *QInputMethod) IsAnimating() bool {
	__r0, _, _ := purego.SyscallN(addr_QInputMethod_isAnimating, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QInputMethod) Locale() *QLocale {
	__r, _, _ := purego.SyscallN(addr_QInputMethod_locale, this.h)
	runtime.KeepAlive(this)
	return newownedQLocale(__r)
}

func (this *QInputMethod) InputDirection() LayoutDirection {
	__r0, _, _ := purego.SyscallN(addr_QInputMethod_inputDirection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return LayoutDirection(__r)
}

func QInputMethod_QueryFocusObject(query InputMethodQuery, argument *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QInputMethod_queryFocusObject, uintptr(query), uintptr(argument.UnsafePointer()))
	runtime.KeepAlive(argument)
	return newownedQVariant(__r)
}

func (this *QInputMethod) Show() {
	purego.SyscallN(addr_QInputMethod_show, this.h)
	runtime.KeepAlive(this)
}

func (this *QInputMethod) Hide() {
	purego.SyscallN(addr_QInputMethod_hide, this.h)
	runtime.KeepAlive(this)
}

func (this *QInputMethod) Update(queries InputMethodQuery) {
	purego.SyscallN(addr_QInputMethod_update, this.h, uintptr(queries))
	runtime.KeepAlive(this)
}

func (this *QInputMethod) Reset() {
	purego.SyscallN(addr_QInputMethod_reset, this.h)
	runtime.KeepAlive(this)
}

func (this *QInputMethod) Commit() {
	purego.SyscallN(addr_QInputMethod_commit, this.h)
	runtime.KeepAlive(this)
}

func (this *QInputMethod) InvokeAction(a QInputMethod__Action, cursorPosition int) {
	purego.SyscallN(addr_QInputMethod_invokeAction, this.h, uintptr(a), uintptr(cursorPosition))
	runtime.KeepAlive(this)
}

func (this *QInputMethod) CursorRectangleChanged() {
	purego.SyscallN(addr_QInputMethod_cursorRectangleChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QInputMethod) AnchorRectangleChanged() {
	purego.SyscallN(addr_QInputMethod_anchorRectangleChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QInputMethod) KeyboardRectangleChanged() {
	purego.SyscallN(addr_QInputMethod_keyboardRectangleChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QInputMethod) InputItemClipRectangleChanged() {
	purego.SyscallN(addr_QInputMethod_inputItemClipRectangleChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QInputMethod) VisibleChanged() {
	purego.SyscallN(addr_QInputMethod_visibleChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QInputMethod) AnimatingChanged() {
	purego.SyscallN(addr_QInputMethod_animatingChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QInputMethod) LocaleChanged() {
	purego.SyscallN(addr_QInputMethod_localeChanged, this.h)
	runtime.KeepAlive(this)
}

func (this *QInputMethod) InputDirectionChanged(newDirection LayoutDirection) {
	purego.SyscallN(addr_QInputMethod_inputDirectionChanged, this.h, uintptr(newDirection))
	runtime.KeepAlive(this)
}

func QInputMethod_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QInputMethod_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QInputMethod_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QInputMethod_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QInputMethod_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputMethod_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QInputMethod_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QInputMethod_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QInputMethod_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QInputMethod_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QInputMethod_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QInputMethod_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QInputMethod) OnCursorRectangleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QInputMethod_connect_cursorRectangleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputMethod) OnAnchorRectangleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QInputMethod_connect_anchorRectangleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputMethod) OnKeyboardRectangleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QInputMethod_connect_keyboardRectangleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputMethod) OnInputItemClipRectangleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QInputMethod_connect_inputItemClipRectangleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputMethod) OnVisibleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QInputMethod_connect_visibleChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputMethod) OnAnimatingChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QInputMethod_connect_animatingChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputMethod) OnLocaleChanged(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QInputMethod_connect_localeChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputMethod) OnInputDirectionChanged(fn func(LayoutDirection)) int64 {
	id := registerCallback(func(newDirection int32) { fn(LayoutDirection(newDirection)) })
	pg_QInputMethod_connect_inputDirectionChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputMethod) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QInputMethod_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputMethod) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QInputMethod_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputMethod) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QInputMethod_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QInputMethod) Disconnect(handle int64) {
	pg_QInputMethod__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QInputMethod) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	this.UnsafeRelease()
}
