package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	pg_QScroller_base func(uintptr) uintptr
	addr_QScroller_tr uintptr
	addr_QScroller_hasScroller uintptr
	addr_QScroller_scroller uintptr
	addr_QScroller_scroller2 uintptr
	addr_QScroller_grabGesture uintptr
	addr_QScroller_grabbedGesture uintptr
	addr_QScroller_ungrabGesture uintptr
	addr_QScroller_activeScrollers uintptr
	addr_QScroller_target uintptr
	addr_QScroller_state uintptr
	addr_QScroller_handleInput uintptr
	addr_QScroller_stop uintptr
	addr_QScroller_velocity uintptr
	addr_QScroller_finalPosition uintptr
	addr_QScroller_pixelPerMeter uintptr
	addr_QScroller_scrollerProperties uintptr
	addr_QScroller_setSnapPositionsX uintptr
	pg_QScroller_setSnapPositionsX2 func(uintptr, float64, float64)
	addr_QScroller_setSnapPositionsY uintptr
	pg_QScroller_setSnapPositionsY2 func(uintptr, float64, float64)
	addr_QScroller_setScrollerProperties uintptr
	addr_QScroller_scrollTo uintptr
	addr_QScroller_scrollTo2 uintptr
	pg_QScroller_ensureVisible func(uintptr, uintptr, float64, float64)
	pg_QScroller_ensureVisible2 func(uintptr, uintptr, float64, float64, int32)
	addr_QScroller_resendPrepareEvent uintptr
	addr_QScroller_stateChanged uintptr
	addr_QScroller_scrollerPropertiesChanged uintptr
	addr_QScroller_connect uintptr
	addr_QScroller_connect2 uintptr
	addr_QScroller_tr2 uintptr
	addr_QScroller_tr3 uintptr
	addr_QScroller_grabGesture2 uintptr
	addr_QScroller_handleInput2 uintptr
	addr_QScroller_connect3 uintptr
	addr_QScroller_connect4 uintptr
	pg_QScroller_connect_stateChanged func(uintptr, int64)
	pg_QScroller_connect_scrollerPropertiesChanged func(uintptr, int64)
	pg_QScroller_connect_destroyed func(uintptr, int64)
	pg_QScroller_connect_destroyed2 func(uintptr, int64)
	pg_QScroller_connect_objectNameChanged func(uintptr, int64)
	pg_QScroller__disconnect func(int64)
)

func init() {
	reg(&pg_QScroller_base, "QScroller_base")
	symaddr(&addr_QScroller_tr, "QScroller_tr")
	symaddr(&addr_QScroller_hasScroller, "QScroller_hasScroller")
	symaddr(&addr_QScroller_scroller, "QScroller_scroller")
	symaddr(&addr_QScroller_scroller2, "QScroller_scroller2")
	symaddr(&addr_QScroller_grabGesture, "QScroller_grabGesture")
	symaddr(&addr_QScroller_grabbedGesture, "QScroller_grabbedGesture")
	symaddr(&addr_QScroller_ungrabGesture, "QScroller_ungrabGesture")
	symaddr(&addr_QScroller_activeScrollers, "QScroller_activeScrollers")
	symaddr(&addr_QScroller_target, "QScroller_target")
	symaddr(&addr_QScroller_state, "QScroller_state")
	symaddr(&addr_QScroller_handleInput, "QScroller_handleInput")
	symaddr(&addr_QScroller_stop, "QScroller_stop")
	symaddr(&addr_QScroller_velocity, "QScroller_velocity")
	symaddr(&addr_QScroller_finalPosition, "QScroller_finalPosition")
	symaddr(&addr_QScroller_pixelPerMeter, "QScroller_pixelPerMeter")
	symaddr(&addr_QScroller_scrollerProperties, "QScroller_scrollerProperties")
	symaddr(&addr_QScroller_setSnapPositionsX, "QScroller_setSnapPositionsX")
	reg(&pg_QScroller_setSnapPositionsX2, "QScroller_setSnapPositionsX2")
	symaddr(&addr_QScroller_setSnapPositionsY, "QScroller_setSnapPositionsY")
	reg(&pg_QScroller_setSnapPositionsY2, "QScroller_setSnapPositionsY2")
	symaddr(&addr_QScroller_setScrollerProperties, "QScroller_setScrollerProperties")
	symaddr(&addr_QScroller_scrollTo, "QScroller_scrollTo")
	symaddr(&addr_QScroller_scrollTo2, "QScroller_scrollTo2")
	reg(&pg_QScroller_ensureVisible, "QScroller_ensureVisible")
	reg(&pg_QScroller_ensureVisible2, "QScroller_ensureVisible2")
	symaddr(&addr_QScroller_resendPrepareEvent, "QScroller_resendPrepareEvent")
	symaddr(&addr_QScroller_stateChanged, "QScroller_stateChanged")
	symaddr(&addr_QScroller_scrollerPropertiesChanged, "QScroller_scrollerPropertiesChanged")
	symaddr(&addr_QScroller_connect, "QScroller_connect")
	symaddr(&addr_QScroller_connect2, "QScroller_connect2")
	symaddr(&addr_QScroller_tr2, "QScroller_tr2")
	symaddr(&addr_QScroller_tr3, "QScroller_tr3")
	symaddr(&addr_QScroller_grabGesture2, "QScroller_grabGesture2")
	symaddr(&addr_QScroller_handleInput2, "QScroller_handleInput2")
	symaddr(&addr_QScroller_connect3, "QScroller_connect3")
	symaddr(&addr_QScroller_connect4, "QScroller_connect4")
	reg(&pg_QScroller_connect_stateChanged, "QScroller_connect_stateChanged")
	reg(&pg_QScroller_connect_scrollerPropertiesChanged, "QScroller_connect_scrollerPropertiesChanged")
	reg(&pg_QScroller_connect_destroyed, "QScroller_connect_destroyed")
	reg(&pg_QScroller_connect_destroyed2, "QScroller_connect_destroyed2")
	reg(&pg_QScroller_connect_objectNameChanged, "QScroller_connect_objectNameChanged")
	reg(&pg_QScroller__disconnect, "QScroller__disconnect")
}

type QScroller struct {
	h uintptr
	*QObject
}

func bindQScroller(h uintptr) *QScroller {
	if h == 0 {
		return nil
	}
	return &QScroller{h: h, QObject: bindQObject(pg_QScroller_base(h))}
}

func newownedQScroller(h uintptr) *QScroller { return bindQScroller(h) }

func (this *QScroller) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QScroller) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQScroller(h unsafe.Pointer) *QScroller { return bindQScroller(uintptr(h)) }
func UnsafeNewQScrollerOwned(h unsafe.Pointer) *QScroller { return newownedQScroller(uintptr(h)) }

func (this *QScroller) UnsafeRelease() {
	this.h = 0
	if this.QObject != nil {
		this.QObject.UnsafeRelease()
	}
}

func QScroller_Tr(s string) string {
	s_b := append([]byte(s), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QScroller_tr, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	return pgStr(__r, __n)
}

func QScroller_HasScroller(target *QObject) bool {
	__r0, _, _ := purego.SyscallN(addr_QScroller_hasScroller, uintptr(target.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(target)
	return __r != 0
}

func QScroller_Scroller(target *QObject) *QScroller {
	__r, _, _ := purego.SyscallN(addr_QScroller_scroller, uintptr(target.UnsafePointer()))
	runtime.KeepAlive(target)
	return bindQScroller(__r)
}

func QScroller_ScrollerWithTarget(target *QObject) *QScroller {
	__r, _, _ := purego.SyscallN(addr_QScroller_scroller2, uintptr(target.UnsafePointer()))
	runtime.KeepAlive(target)
	return bindQScroller(__r)
}

func QScroller_GrabGesture(target *QObject) GestureType {
	__r0, _, _ := purego.SyscallN(addr_QScroller_grabGesture, uintptr(target.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(target)
	return GestureType(__r)
}

func QScroller_GrabbedGesture(target *QObject) GestureType {
	__r0, _, _ := purego.SyscallN(addr_QScroller_grabbedGesture, uintptr(target.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(target)
	return GestureType(__r)
}

func QScroller_UngrabGesture(target *QObject) {
	purego.SyscallN(addr_QScroller_ungrabGesture, uintptr(target.UnsafePointer()))
	runtime.KeepAlive(target)
}

func QScroller_ActiveScrollers() []*QScroller {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QScroller_activeScrollers, uintptr(unsafe.Pointer(&__n)))
	__out := make([]*QScroller, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = bindQScroller(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QScroller) Target() *QObject {
	__r, _, _ := purego.SyscallN(addr_QScroller_target, this.h)
	runtime.KeepAlive(this)
	return bindQObject(__r)
}

func (this *QScroller) State() QScroller__State {
	__r0, _, _ := purego.SyscallN(addr_QScroller_state, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QScroller__State(__r)
}

func (this *QScroller) HandleInput(input QScroller__Input, position *QPointF) bool {
	__r0, _, _ := purego.SyscallN(addr_QScroller_handleInput, this.h, uintptr(input), uintptr(position.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(position)
	return __r != 0
}

func (this *QScroller) Stop() {
	purego.SyscallN(addr_QScroller_stop, this.h)
	runtime.KeepAlive(this)
}

func (this *QScroller) Velocity() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QScroller_velocity, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QScroller) FinalPosition() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QScroller_finalPosition, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QScroller) PixelPerMeter() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QScroller_pixelPerMeter, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QScroller) ScrollerProperties() *QScrollerProperties {
	__r, _, _ := purego.SyscallN(addr_QScroller_scrollerProperties, this.h)
	runtime.KeepAlive(this)
	return newownedQScrollerProperties(__r)
}

func (this *QScroller) SetSnapPositionsX(positions []float64) {
	positions_arr := make([]float64, len(positions))
	for __i, __e := range positions {
		positions_arr[__i] = __e
	}
	var positions_p uintptr
	if len(positions_arr) > 0 {
		positions_p = uintptr(unsafe.Pointer(&positions_arr[0]))
	}
	purego.SyscallN(addr_QScroller_setSnapPositionsX, this.h, positions_p, uintptr(len(positions)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(positions_arr)
}

func (this *QScroller) SetSnapPositionsX2(first float64, interval float64) {
	pg_QScroller_setSnapPositionsX2(this.h, first, interval)
	runtime.KeepAlive(this)
}

func (this *QScroller) SetSnapPositionsY(positions []float64) {
	positions_arr := make([]float64, len(positions))
	for __i, __e := range positions {
		positions_arr[__i] = __e
	}
	var positions_p uintptr
	if len(positions_arr) > 0 {
		positions_p = uintptr(unsafe.Pointer(&positions_arr[0]))
	}
	purego.SyscallN(addr_QScroller_setSnapPositionsY, this.h, positions_p, uintptr(len(positions)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(positions_arr)
}

func (this *QScroller) SetSnapPositionsY2(first float64, interval float64) {
	pg_QScroller_setSnapPositionsY2(this.h, first, interval)
	runtime.KeepAlive(this)
}

func (this *QScroller) SetScrollerProperties(prop *QScrollerProperties) {
	purego.SyscallN(addr_QScroller_setScrollerProperties, this.h, uintptr(prop.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(prop)
}

func (this *QScroller) ScrollTo(pos *QPointF) {
	purego.SyscallN(addr_QScroller_scrollTo, this.h, uintptr(pos.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QScroller) ScrollTo2(pos *QPointF, scrollTime int) {
	purego.SyscallN(addr_QScroller_scrollTo2, this.h, uintptr(pos.UnsafePointer()), uintptr(scrollTime))
	runtime.KeepAlive(this)
	runtime.KeepAlive(pos)
}

func (this *QScroller) EnsureVisible(rect *QRectF, xmargin float64, ymargin float64) {
	pg_QScroller_ensureVisible(this.h, uintptr(rect.UnsafePointer()), xmargin, ymargin)
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QScroller) EnsureVisible2(rect *QRectF, xmargin float64, ymargin float64, scrollTime int) {
	pg_QScroller_ensureVisible2(this.h, uintptr(rect.UnsafePointer()), xmargin, ymargin, int32(scrollTime))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rect)
}

func (this *QScroller) ResendPrepareEvent() {
	purego.SyscallN(addr_QScroller_resendPrepareEvent, this.h)
	runtime.KeepAlive(this)
}

func (this *QScroller) StateChanged(newstate QScroller__State) {
	purego.SyscallN(addr_QScroller_stateChanged, this.h, uintptr(newstate))
	runtime.KeepAlive(this)
}

func (this *QScroller) ScrollerPropertiesChanged(arg1 *QScrollerProperties) {
	purego.SyscallN(addr_QScroller_scrollerPropertiesChanged, this.h, uintptr(arg1.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(arg1)
}

func QScroller_Connect(sender *QObject, signal string, receiver *QObject, member string) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QScroller_connect, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QScroller_Connect2(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QScroller_connect2, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func QScroller_Tr2(s string, c string) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QScroller_tr2, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QScroller_Tr3(s string, c string, n int) string {
	s_b := append([]byte(s), 0)
	c_b := append([]byte(c), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QScroller_tr3, uintptr(unsafe.Pointer(&s_b[0])), uintptr(unsafe.Pointer(&c_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(s_b)
	runtime.KeepAlive(c_b)
	return pgStr(__r, __n)
}

func QScroller_GrabGesture2(target *QObject, gestureType QScroller__ScrollerGestureType) GestureType {
	__r0, _, _ := purego.SyscallN(addr_QScroller_grabGesture2, uintptr(target.UnsafePointer()), uintptr(gestureType))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(target)
	return GestureType(__r)
}

func (this *QScroller) HandleInput2(input QScroller__Input, position *QPointF, timestamp int64) bool {
	__r0, _, _ := purego.SyscallN(addr_QScroller_handleInput2, this.h, uintptr(input), uintptr(position.UnsafePointer()), uintptr(timestamp))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(position)
	return __r != 0
}

func QScroller_Connect3(sender *QObject, signal string, receiver *QObject, member string, arg5 ConnectionType) *QMetaObject__Connection {
	signal_b := append([]byte(signal), 0)
	member_b := append([]byte(member), 0)
	__r, _, _ := purego.SyscallN(addr_QScroller_connect3, uintptr(sender.UnsafePointer()), uintptr(unsafe.Pointer(&signal_b[0])), uintptr(receiver.UnsafePointer()), uintptr(unsafe.Pointer(&member_b[0])), uintptr(arg5))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal_b)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(member_b)
	return newownedQMetaObject__Connection(__r)
}

func QScroller_Connect4(sender *QObject, signal *QMetaMethod, receiver *QObject, method *QMetaMethod, type_ ConnectionType) *QMetaObject__Connection {
	__r, _, _ := purego.SyscallN(addr_QScroller_connect4, uintptr(sender.UnsafePointer()), uintptr(signal.UnsafePointer()), uintptr(receiver.UnsafePointer()), uintptr(method.UnsafePointer()), uintptr(type_))
	runtime.KeepAlive(sender)
	runtime.KeepAlive(signal)
	runtime.KeepAlive(receiver)
	runtime.KeepAlive(method)
	return newownedQMetaObject__Connection(__r)
}

func (this *QScroller) OnStateChanged(fn func(QScroller__State)) int64 {
	id := registerCallback(func(newstate int32) { fn(QScroller__State(newstate)) })
	pg_QScroller_connect_stateChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScroller) OnScrollerPropertiesChanged(fn func(*QScrollerProperties)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQScrollerProperties(unsafe.Pointer(arg1))) })
	pg_QScroller_connect_scrollerPropertiesChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScroller) OnDestroyed(fn func()) int64 {
	id := registerCallback(func() { fn() })
	pg_QScroller_connect_destroyed(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScroller) OnDestroyedWithQObject(fn func(*QObject)) int64 {
	id := registerCallback(func(arg1 uintptr) { fn(UnsafeNewQObject(unsafe.Pointer(arg1))) })
	pg_QScroller_connect_destroyed2(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScroller) OnObjectNameChanged(fn func(string)) int64 {
	id := registerCallback(func(objectName uintptr, objectName_len int64) { fn(pgStrNoFree(objectName, objectName_len)) })
	pg_QScroller_connect_objectNameChanged(this.h, id)
	runtime.KeepAlive(this)
	return id
}

func (this *QScroller) Disconnect(handle int64) {
	pg_QScroller__disconnect(handle)
	unregisterCallback(handle)
}

func (this *QScroller) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	this.UnsafeRelease()
}
