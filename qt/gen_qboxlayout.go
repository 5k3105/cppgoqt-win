package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QBoxLayout_chain uintptr
	addr_QBoxLayout_new uintptr
	addr_QBoxLayout_new2 uintptr
	addr_QBoxLayout_tr uintptr
	addr_QBoxLayout_direction uintptr
	addr_QBoxLayout_setDirection uintptr
	addr_QBoxLayout_addSpacing uintptr
	addr_QBoxLayout_addStretch uintptr
	addr_QBoxLayout_addSpacerItem uintptr
	addr_QBoxLayout_addWidget uintptr
	addr_QBoxLayout_addLayout uintptr
	addr_QBoxLayout_addStrut uintptr
	addr_QBoxLayout_insertSpacing uintptr
	addr_QBoxLayout_insertStretch uintptr
	addr_QBoxLayout_insertSpacerItem uintptr
	addr_QBoxLayout_insertWidget uintptr
	addr_QBoxLayout_insertLayout uintptr
	addr_QBoxLayout_insertItem uintptr
	addr_QBoxLayout_setStretchFactor uintptr
	addr_QBoxLayout_setStretchFactor2 uintptr
	addr_QBoxLayout_setStretch uintptr
	addr_QBoxLayout_stretch uintptr
	addr_QBoxLayout_closestAcceptableSize uintptr
	addr_QBoxLayout_connect uintptr
	addr_QBoxLayout_connect2 uintptr
	addr_QBoxLayout_tr2 uintptr
	addr_QBoxLayout_tr3 uintptr
	addr_QBoxLayout_addStretch2 uintptr
	addr_QBoxLayout_addWidget2 uintptr
	addr_QBoxLayout_addWidget3 uintptr
	addr_QBoxLayout_addLayout2 uintptr
	addr_QBoxLayout_insertStretch2 uintptr
	addr_QBoxLayout_insertWidget2 uintptr
	addr_QBoxLayout_insertWidget3 uintptr
	addr_QBoxLayout_insertLayout2 uintptr
	addr_QBoxLayout_connect3 uintptr
	addr_QBoxLayout_connect4 uintptr
	pg_QBoxLayout_connect_destroyed func(uintptr, int64)
	pg_QBoxLayout_connect_destroyed2 func(uintptr, int64)
	pg_QBoxLayout_connect_objectNameChanged func(uintptr, int64)
	pg_QBoxLayout__disconnect func(int64)
	pg_QBoxLayout_delete func(uintptr)
)

func init() {
	symaddr(&addr_QBoxLayout_chain, "QBoxLayout_chain")
	symaddr(&addr_QBoxLayout_new, "QBoxLayout_new")
	symaddr(&addr_QBoxLayout_new2, "QBoxLayout_new2")
	symaddr(&addr_QBoxLayout_tr, "QBoxLayout_tr")
	symaddr(&addr_QBoxLayout_direction, "QBoxLayout_direction")
	symaddr(&addr_QBoxLayout_setDirection, "QBoxLayout_setDirection")
	symaddr(&addr_QBoxLayout_addSpacing, "QBoxLayout_addSpacing")
	symaddr(&addr_QBoxLayout_addStretch, "QBoxLayout_addStretch")
	symaddr(&addr_QBoxLayout_addSpacerItem, "QBoxLayout_addSpacerItem")
	symaddr(&addr_QBoxLayout_addWidget, "QBoxLayout_addWidget")
	symaddr(&addr_QBoxLayout_addLayout, "QBoxLayout_addLayout")
	symaddr(&addr_QBoxLayout_addStrut, "QBoxLayout_addStrut")
	symaddr(&addr_QBoxLayout_insertSpacing, "QBoxLayout_insertSpacing")
	symaddr(&addr_QBoxLayout_insertStretch, "QBoxLayout_insertStretch")
	symaddr(&addr_QBoxLayout_insertSpacerItem, "QBoxLayout_insertSpacerItem")
	symaddr(&addr_QBoxLayout_insertWidget, "QBoxLayout_insertWidget")
	symaddr(&addr_QBoxLayout_insertLayout, "QBoxLayout_insertLayout")
	symaddr(&addr_QBoxLayout_insertItem, "QBoxLayout_insertItem")
	symaddr(&addr_QBoxLayout_setStretchFactor, "QBoxLayout_setStretchFactor")
	symaddr(&addr_QBoxLayout_setStretchFactor2, "QBoxLayout_setStretchFactor2")
	symaddr(&addr_QBoxLayout_setStretch, "QBoxLayout_setStretch")
	symaddr(&addr_QBoxLayout_stretch, "QBoxLayout_stretch")
	symaddr(&addr_QBoxLayout_closestAcceptableSize, "QBoxLayout_closestAcceptableSize")
	symaddr(&addr_QBoxLayout_connect, "QBoxLayout_connect")
	symaddr(&addr_QBoxLayout_connect2, "QBoxLayout_connect2")
	symaddr(&addr_QBoxLayout_tr2, "QBoxLayout_tr2")
	symaddr(&addr_QBoxLayout_tr3, "QBoxLayout_tr3")
	symaddr(&addr_QBoxLayout_addStretch2, "QBoxLayout_addStretch2")
	symaddr(&addr_QBoxLayout_addWidget2, "QBoxLayout_addWidget2")
	symaddr(&addr_QBoxLayout_addWidget3, "QBoxLayout_addWidget3")
	symaddr(&addr_QBoxLayout_addLayout2, "QBoxLayout_addLayout2")
	symaddr(&addr_QBoxLayout_insertStretch2, "QBoxLayout_insertStretch2")
	symaddr(&addr_QBoxLayout_insertWidget2, "QBoxLayout_insertWidget2")
	symaddr(&addr_QBoxLayout_insertWidget3, "QBoxLayout_insertWidget3")
	symaddr(&addr_QBoxLayout_insertLayout2, "QBoxLayout_insertLayout2")
	symaddr(&addr_QBoxLayout_connect3, "QBoxLayout_connect3")
	symaddr(&addr_QBoxLayout_connect4, "QBoxLayout_connect4")
	reg(&pg_QBoxLayout_connect_destroyed, "QBoxLayout_connect_destroyed")
	reg(&pg_QBoxLayout_connect_destroyed2, "QBoxLayout_connect_destroyed2")
	reg(&pg_QBoxLayout_connect_objectNameChanged, "QBoxLayout_connect_objectNameChanged")
	reg(&pg_QBoxLayout__disconnect, "QBoxLayout__disconnect")
	reg(&pg_QBoxLayout_delete, "QBoxLayout_delete")
}

type QBoxLayout struct {
	h uintptr
	*QLayout
}

func bindQBoxLayout(h uintptr) *QBoxLayout {
	if h == 0 {
		return nil
	}
	var __c [2]uintptr
	purego.SyscallN(addr_QBoxLayout_chain, h, uintptr(unsafe.Pointer(&__c[0])))
	return &QBoxLayout{h: h, QLayout: &QLayout{h: __c[0], QObject: &QObject{h: __c[1]}}}
}

func newownedQBoxLayout(h uintptr) *QBoxLayout { return bindQBoxLayout(h) }

func (this *QBoxLayout) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QBoxLayout) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQBoxLayout(h unsafe.Pointer) *QBoxLayout { return bindQBoxLayout(uintptr(h)) }
func UnsafeNewQBoxLayoutOwned(h unsafe.Pointer) *QBoxLayout { return newownedQBoxLayout(uintptr(h)) }

func (this *QBoxLayout) UnsafeRelease() {
	this.h = 0
	if this.QLayout != nil {
		this.QLayout.UnsafeRelease()
	}
}

func NewQBoxLayout(arg1 QBoxLayout__Direction) *QBoxLayout {
	__r, _, _ := purego.SyscallN(addr_QBoxLayout_new, uintptr(arg1))
	__o := newownedQBoxLayout(__r)
	return __o
}

func NewQBoxLayout2(arg1 QBoxLayout__Direction, parent *QWidget) *QBoxLayout {
	__r, _, _ := purego.SyscallN(addr_QBoxLayout_new2, uintptr(arg1), uintptr(parent.UnsafePointer()))
	__o := newownedQBoxLayout(__r)
	runtime.KeepAlive(parent)
	return __o
}

func QBoxLayout_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QBoxLayout_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func (this *QBoxLayout) Direction() QBoxLayout__Direction {
	__r0, _, _ := purego.SyscallN(addr_QBoxLayout_direction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QBoxLayout__Direction(__r)
}

func (this *QBoxLayout) SetDirection(arg1 QBoxLayout__Direction) {
	purego.SyscallN(addr_QBoxLayout_setDirection, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QBoxLayout) AddSpacing(size int) {
	purego.SyscallN(addr_QBoxLayout_addSpacing, this.h, uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QBoxLayout) AddStretch() {
	purego.SyscallN(addr_QBoxLayout_addStretch, this.h)
	runtime.KeepAlive(this)
}

func (this *QBoxLayout) AddSpacerItem(spacerItem *QSpacerItem) {
	purego.SyscallN(addr_QBoxLayout_addSpacerItem, this.h, uintptr(spacerItem.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(spacerItem)
}

func (this *QBoxLayout) AddWidget(arg1 *QWidget) {
	purego.SyscallN(addr_QBoxLayout_addWidget, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QBoxLayout) AddLayout(layout *QLayout) {
	purego.SyscallN(addr_QBoxLayout_addLayout, this.h, uintptr(layout.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
}

func (this *QBoxLayout) AddStrut(arg1 int) {
	purego.SyscallN(addr_QBoxLayout_addStrut, this.h, uintptr(arg1))
	runtime.KeepAlive(this)
}

func (this *QBoxLayout) InsertSpacing(index int, size int) {
	purego.SyscallN(addr_QBoxLayout_insertSpacing, this.h, uintptr(index), uintptr(size))
	runtime.KeepAlive(this)
}

func (this *QBoxLayout) InsertStretch(index int) {
	purego.SyscallN(addr_QBoxLayout_insertStretch, this.h, uintptr(index))
	runtime.KeepAlive(this)
}

func (this *QBoxLayout) InsertSpacerItem(index int, spacerItem *QSpacerItem) {
	purego.SyscallN(addr_QBoxLayout_insertSpacerItem, this.h, uintptr(index), uintptr(spacerItem.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(spacerItem)
}

func (this *QBoxLayout) InsertWidget(index int, widget *QWidget) {
	purego.SyscallN(addr_QBoxLayout_insertWidget, this.h, uintptr(index), uintptr(widget.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QBoxLayout) InsertLayout(index int, layout *QLayout) {
	purego.SyscallN(addr_QBoxLayout_insertLayout, this.h, uintptr(index), uintptr(layout.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
}

func (this *QBoxLayout) InsertItem(index int, arg2 *QLayoutItem) {
	purego.SyscallN(addr_QBoxLayout_insertItem, this.h, uintptr(index), uintptr(arg2.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg2)
}

func (this *QBoxLayout) SetStretchFactor(w *QWidget, stretch int) bool {
	__r0, _, _ := purego.SyscallN(addr_QBoxLayout_setStretchFactor, this.h, uintptr(w.UnsafePointer()), uintptr(stretch))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(w)
	return __r != 0
}

func (this *QBoxLayout) SetStretchFactor2(l *QLayout, stretch int) bool {
	__r0, _, _ := purego.SyscallN(addr_QBoxLayout_setStretchFactor2, this.h, uintptr(l.UnsafePointer()), uintptr(stretch))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(l)
	return __r != 0
}

func (this *QBoxLayout) SetStretch(index int, stretch int) {
	purego.SyscallN(addr_QBoxLayout_setStretch, this.h, uintptr(index), uintptr(stretch))
	runtime.KeepAlive(this)
}

func (this *QBoxLayout) Stretch(index int) int {
	__r0, _, _ := purego.SyscallN(addr_QBoxLayout_stretch, this.h, uintptr(index))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func QBoxLayout_ClosestAcceptableSize(w *QWidget, s *QSize) *QSize {
	__r, _, _ := purego.SyscallN(addr_QBoxLayout_closestAcceptableSize, uintptr(w.UnsafePointer()), uintptr(s.UnsafePointer()))
	runtime.KeepAlive(w)
	runtime.KeepAlive(s)
	return newownedQSize(__r)
}

func QBoxLayout_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QBoxLayout_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QBoxLayout_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QBoxLayout_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QBoxLayout_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QBoxLayout_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QBoxLayout_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QBoxLayout_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func (this *QBoxLayout) AddStretchWithStretch(stretch int) {
	purego.SyscallN(addr_QBoxLayout_addStretch2, this.h, uintptr(stretch))
	runtime.KeepAlive(this)
}

func (this *QBoxLayout) AddWidget2(arg1 *QWidget, stretch int) {
	purego.SyscallN(addr_QBoxLayout_addWidget2, this.h, uintptr(arg1.UnsafePointer()), uintptr(stretch))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QBoxLayout) AddWidget3(arg1 *QWidget, stretch int, alignment AlignmentFlag) {
	purego.SyscallN(addr_QBoxLayout_addWidget3, this.h, uintptr(arg1.UnsafePointer()), uintptr(stretch), uintptr(alignment))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func (this *QBoxLayout) AddLayout2(layout *QLayout, stretch int) {
	purego.SyscallN(addr_QBoxLayout_addLayout2, this.h, uintptr(layout.UnsafePointer()), uintptr(stretch))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
}

func (this *QBoxLayout) InsertStretch2(index int, stretch int) {
	purego.SyscallN(addr_QBoxLayout_insertStretch2, this.h, uintptr(index), uintptr(stretch))
	runtime.KeepAlive(this)
}

func (this *QBoxLayout) InsertWidget2(index int, widget *QWidget, stretch int) {
	purego.SyscallN(addr_QBoxLayout_insertWidget2, this.h, uintptr(index), uintptr(widget.UnsafePointer()), uintptr(stretch))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QBoxLayout) InsertWidget3(index int, widget *QWidget, stretch int, alignment AlignmentFlag) {
	purego.SyscallN(addr_QBoxLayout_insertWidget3, this.h, uintptr(index), uintptr(widget.UnsafePointer()), uintptr(stretch), uintptr(alignment))
	runtime.KeepAlive(this)
	runtime.KeepAlive(widget)
}

func (this *QBoxLayout) InsertLayout2(index int, layout *QLayout, stretch int) {
	purego.SyscallN(addr_QBoxLayout_insertLayout2, this.h, uintptr(index), uintptr(layout.UnsafePointer()), uintptr(stretch))
	runtime.KeepAlive(this)
	runtime.KeepAlive(layout)
}

func QBoxLayout_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QBoxLayout_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QBoxLayout_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QBoxLayout_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QBoxLayout) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QBoxLayout_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QBoxLayout) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QBoxLayout_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QBoxLayout) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QBoxLayout_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QBoxLayout) Disconnect(handle int64) {
	pg_QBoxLayout__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QBoxLayout) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QBoxLayout_delete(this.h)
	}
	this.UnsafeRelease()
}
