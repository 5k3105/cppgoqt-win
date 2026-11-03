package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QTextOption_new uintptr
	addr_QTextOption_new2 uintptr
	addr_QTextOption_new3 uintptr
	addr_QTextOption_OperatorAssign uintptr
	addr_QTextOption_setAlignment uintptr
	addr_QTextOption_alignment uintptr
	addr_QTextOption_setTextDirection uintptr
	addr_QTextOption_textDirection uintptr
	addr_QTextOption_setWrapMode uintptr
	addr_QTextOption_wrapMode uintptr
	addr_QTextOption_setFlags uintptr
	addr_QTextOption_flags uintptr
	pg_QTextOption_setTabStopDistance func(uintptr, float64)
	pg_QTextOption_tabStopDistance func(uintptr) float64
	addr_QTextOption_setTabArray uintptr
	addr_QTextOption_tabArray uintptr
	addr_QTextOption_setTabs uintptr
	addr_QTextOption_tabs uintptr
	addr_QTextOption_setUseDesignMetrics uintptr
	addr_QTextOption_useDesignMetrics uintptr
	pg_QTextOption_delete func(uintptr)
)

func init() {
	symaddr(&addr_QTextOption_new, "QTextOption_new")
	symaddr(&addr_QTextOption_new2, "QTextOption_new2")
	symaddr(&addr_QTextOption_new3, "QTextOption_new3")
	symaddr(&addr_QTextOption_OperatorAssign, "QTextOption_OperatorAssign")
	symaddr(&addr_QTextOption_setAlignment, "QTextOption_setAlignment")
	symaddr(&addr_QTextOption_alignment, "QTextOption_alignment")
	symaddr(&addr_QTextOption_setTextDirection, "QTextOption_setTextDirection")
	symaddr(&addr_QTextOption_textDirection, "QTextOption_textDirection")
	symaddr(&addr_QTextOption_setWrapMode, "QTextOption_setWrapMode")
	symaddr(&addr_QTextOption_wrapMode, "QTextOption_wrapMode")
	symaddr(&addr_QTextOption_setFlags, "QTextOption_setFlags")
	symaddr(&addr_QTextOption_flags, "QTextOption_flags")
	reg(&pg_QTextOption_setTabStopDistance, "QTextOption_setTabStopDistance")
	reg(&pg_QTextOption_tabStopDistance, "QTextOption_tabStopDistance")
	symaddr(&addr_QTextOption_setTabArray, "QTextOption_setTabArray")
	symaddr(&addr_QTextOption_tabArray, "QTextOption_tabArray")
	symaddr(&addr_QTextOption_setTabs, "QTextOption_setTabs")
	symaddr(&addr_QTextOption_tabs, "QTextOption_tabs")
	symaddr(&addr_QTextOption_setUseDesignMetrics, "QTextOption_setUseDesignMetrics")
	symaddr(&addr_QTextOption_useDesignMetrics, "QTextOption_useDesignMetrics")
	reg(&pg_QTextOption_delete, "QTextOption_delete")
}

type QTextOption struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QTextOption) UnsafeGoID() int64     { return this.vid }
func (this *QTextOption) UnsafeSetGoID(id int64) { this.vid = id }

func bindQTextOption(h uintptr) *QTextOption {
	if h == 0 {
		return nil
	}
	return &QTextOption{h: h}
}

func newownedQTextOption(h uintptr) *QTextOption {
	o := bindQTextOption(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QTextOption).Delete)
	}
	return o
}

func (this *QTextOption) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QTextOption) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQTextOption(h unsafe.Pointer) *QTextOption { return bindQTextOption(uintptr(h)) }
func UnsafeNewQTextOptionOwned(h unsafe.Pointer) *QTextOption { return newownedQTextOption(uintptr(h)) }

func (this *QTextOption) UnsafeRelease() { this.h = 0 }

func NewQTextOption() *QTextOption {
	__r, _, _ := purego.SyscallN(addr_QTextOption_new)
	__o := newownedQTextOption(__r)
	return __o
}

func NewQTextOption2(alignment AlignmentFlag) *QTextOption {
	__r, _, _ := purego.SyscallN(addr_QTextOption_new2, uintptr(alignment))
	__o := newownedQTextOption(__r)
	return __o
}

func NewQTextOption3(o *QTextOption) *QTextOption {
	__r, _, _ := purego.SyscallN(addr_QTextOption_new3, uintptr(o.UnsafePointer()))
	__o := newownedQTextOption(__r)
	runtime.KeepAlive(o)
	return __o
}

func (this *QTextOption) OperatorAssign(o *QTextOption) *QTextOption {
	__r, _, _ := purego.SyscallN(addr_QTextOption_OperatorAssign, this.h, uintptr(o.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(o)
	return bindQTextOption(__r)
}

func (this *QTextOption) SetAlignment(alignment AlignmentFlag) {
	purego.SyscallN(addr_QTextOption_setAlignment, this.h, uintptr(alignment))
	runtime.KeepAlive(this)
}

func (this *QTextOption) Alignment() AlignmentFlag {
	__r0, _, _ := purego.SyscallN(addr_QTextOption_alignment, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return AlignmentFlag(__r)
}

func (this *QTextOption) SetTextDirection(aDirection LayoutDirection) {
	purego.SyscallN(addr_QTextOption_setTextDirection, this.h, uintptr(aDirection))
	runtime.KeepAlive(this)
}

func (this *QTextOption) TextDirection() LayoutDirection {
	__r0, _, _ := purego.SyscallN(addr_QTextOption_textDirection, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return LayoutDirection(__r)
}

func (this *QTextOption) SetWrapMode(wrap QTextOption__WrapMode) {
	purego.SyscallN(addr_QTextOption_setWrapMode, this.h, uintptr(wrap))
	runtime.KeepAlive(this)
}

func (this *QTextOption) WrapMode() QTextOption__WrapMode {
	__r0, _, _ := purego.SyscallN(addr_QTextOption_wrapMode, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextOption__WrapMode(__r)
}

func (this *QTextOption) SetFlags(flags QTextOption__Flag) {
	purego.SyscallN(addr_QTextOption_setFlags, this.h, uintptr(flags))
	runtime.KeepAlive(this)
}

func (this *QTextOption) Flags() QTextOption__Flag {
	__r0, _, _ := purego.SyscallN(addr_QTextOption_flags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QTextOption__Flag(__r)
}

func (this *QTextOption) SetTabStopDistance(tabStopDistance float64) {
	pg_QTextOption_setTabStopDistance(this.h, tabStopDistance)
	runtime.KeepAlive(this)
}

func (this *QTextOption) TabStopDistance() float64 {
	__r := pg_QTextOption_tabStopDistance(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QTextOption) SetTabArray(tabStops []float64) {
	tabStops_arr := make([]float64, len(tabStops))
	for __i, __e := range tabStops {
		tabStops_arr[__i] = __e
	}
	var tabStops_p uintptr
	if len(tabStops_arr) > 0 {
		tabStops_p = uintptr(unsafe.Pointer(&tabStops_arr[0]))
	}
	purego.SyscallN(addr_QTextOption_setTabArray, this.h, tabStops_p, uintptr(len(tabStops)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tabStops_arr)
}

func (this *QTextOption) TabArray() []float64 {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextOption_tabArray, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]float64, __n)
	__sl := unsafe.Slice((*float64)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = __sl[__i]
	}
	pgFree(__r)
	return __out
}

func (this *QTextOption) SetTabs(tabStops []*QTextOption__Tab) {
	tabStops_arr := make([]uintptr, len(tabStops))
	for __i, __e := range tabStops {
		tabStops_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var tabStops_p uintptr
	if len(tabStops_arr) > 0 {
		tabStops_p = uintptr(unsafe.Pointer(&tabStops_arr[0]))
	}
	purego.SyscallN(addr_QTextOption_setTabs, this.h, tabStops_p, uintptr(len(tabStops)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(tabStops)
	runtime.KeepAlive(tabStops_arr)
}

func (this *QTextOption) Tabs() []*QTextOption__Tab {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QTextOption_tabs, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QTextOption__Tab, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQTextOption__Tab(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QTextOption) SetUseDesignMetrics(b bool) {
	purego.SyscallN(addr_QTextOption_setUseDesignMetrics, this.h, uintptr(b2i(b)))
	runtime.KeepAlive(this)
}

func (this *QTextOption) UseDesignMetrics() bool {
	__r0, _, _ := purego.SyscallN(addr_QTextOption_useDesignMetrics, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QTextOption) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QTextOption_delete(this.h)
	}
	this.UnsafeRelease()
}
