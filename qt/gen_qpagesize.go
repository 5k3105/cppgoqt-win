package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QPageSize_new uintptr
	addr_QPageSize_new2 uintptr
	addr_QPageSize_new3 uintptr
	addr_QPageSize_new4 uintptr
	addr_QPageSize_new5 uintptr
	addr_QPageSize_new6 uintptr
	addr_QPageSize_new7 uintptr
	addr_QPageSize_new8 uintptr
	addr_QPageSize_new9 uintptr
	addr_QPageSize_OperatorAssign uintptr
	addr_QPageSize_swap uintptr
	addr_QPageSize_isEquivalentTo uintptr
	addr_QPageSize_isValid uintptr
	addr_QPageSize_key uintptr
	addr_QPageSize_name uintptr
	addr_QPageSize_id uintptr
	addr_QPageSize_windowsId uintptr
	addr_QPageSize_definitionSize uintptr
	addr_QPageSize_definitionUnits uintptr
	addr_QPageSize_size uintptr
	addr_QPageSize_sizePoints uintptr
	addr_QPageSize_sizePixels uintptr
	addr_QPageSize_rect uintptr
	addr_QPageSize_rectPoints uintptr
	addr_QPageSize_rectPixels uintptr
	addr_QPageSize_key2 uintptr
	addr_QPageSize_name2 uintptr
	addr_QPageSize_id2 uintptr
	addr_QPageSize_id3 uintptr
	addr_QPageSize_id4 uintptr
	addr_QPageSize_windowsId2 uintptr
	addr_QPageSize_definitionSize2 uintptr
	addr_QPageSize_definitionUnits2 uintptr
	addr_QPageSize_size2 uintptr
	addr_QPageSize_sizePoints2 uintptr
	addr_QPageSize_sizePixels2 uintptr
	addr_QPageSize_id5 uintptr
	addr_QPageSize_id6 uintptr
	addr_QPageSize_OperatorEqual uintptr
	addr_QPageSize_OperatorNotEqual uintptr
	pg_QPageSize_delete func(uintptr)
)

func init() {
	symaddr(&addr_QPageSize_new, "QPageSize_new")
	symaddr(&addr_QPageSize_new2, "QPageSize_new2")
	symaddr(&addr_QPageSize_new3, "QPageSize_new3")
	symaddr(&addr_QPageSize_new4, "QPageSize_new4")
	symaddr(&addr_QPageSize_new5, "QPageSize_new5")
	symaddr(&addr_QPageSize_new6, "QPageSize_new6")
	symaddr(&addr_QPageSize_new7, "QPageSize_new7")
	symaddr(&addr_QPageSize_new8, "QPageSize_new8")
	symaddr(&addr_QPageSize_new9, "QPageSize_new9")
	symaddr(&addr_QPageSize_OperatorAssign, "QPageSize_OperatorAssign")
	symaddr(&addr_QPageSize_swap, "QPageSize_swap")
	symaddr(&addr_QPageSize_isEquivalentTo, "QPageSize_isEquivalentTo")
	symaddr(&addr_QPageSize_isValid, "QPageSize_isValid")
	symaddr(&addr_QPageSize_key, "QPageSize_key")
	symaddr(&addr_QPageSize_name, "QPageSize_name")
	symaddr(&addr_QPageSize_id, "QPageSize_id")
	symaddr(&addr_QPageSize_windowsId, "QPageSize_windowsId")
	symaddr(&addr_QPageSize_definitionSize, "QPageSize_definitionSize")
	symaddr(&addr_QPageSize_definitionUnits, "QPageSize_definitionUnits")
	symaddr(&addr_QPageSize_size, "QPageSize_size")
	symaddr(&addr_QPageSize_sizePoints, "QPageSize_sizePoints")
	symaddr(&addr_QPageSize_sizePixels, "QPageSize_sizePixels")
	symaddr(&addr_QPageSize_rect, "QPageSize_rect")
	symaddr(&addr_QPageSize_rectPoints, "QPageSize_rectPoints")
	symaddr(&addr_QPageSize_rectPixels, "QPageSize_rectPixels")
	symaddr(&addr_QPageSize_key2, "QPageSize_key2")
	symaddr(&addr_QPageSize_name2, "QPageSize_name2")
	symaddr(&addr_QPageSize_id2, "QPageSize_id2")
	symaddr(&addr_QPageSize_id3, "QPageSize_id3")
	symaddr(&addr_QPageSize_id4, "QPageSize_id4")
	symaddr(&addr_QPageSize_windowsId2, "QPageSize_windowsId2")
	symaddr(&addr_QPageSize_definitionSize2, "QPageSize_definitionSize2")
	symaddr(&addr_QPageSize_definitionUnits2, "QPageSize_definitionUnits2")
	symaddr(&addr_QPageSize_size2, "QPageSize_size2")
	symaddr(&addr_QPageSize_sizePoints2, "QPageSize_sizePoints2")
	symaddr(&addr_QPageSize_sizePixels2, "QPageSize_sizePixels2")
	symaddr(&addr_QPageSize_id5, "QPageSize_id5")
	symaddr(&addr_QPageSize_id6, "QPageSize_id6")
	symaddr(&addr_QPageSize_OperatorEqual, "QPageSize_OperatorEqual")
	symaddr(&addr_QPageSize_OperatorNotEqual, "QPageSize_OperatorNotEqual")
	reg(&pg_QPageSize_delete, "QPageSize_delete")
}

type QPageSize struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QPageSize) UnsafeGoID() int64     { return this.vid }
func (this *QPageSize) UnsafeSetGoID(id int64) { this.vid = id }

func bindQPageSize(h uintptr) *QPageSize {
	if h == 0 {
		return nil
	}
	return &QPageSize{h: h}
}

func newownedQPageSize(h uintptr) *QPageSize {
	o := bindQPageSize(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QPageSize).Delete)
	}
	return o
}

func (this *QPageSize) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QPageSize) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQPageSize(h unsafe.Pointer) *QPageSize { return bindQPageSize(uintptr(h)) }
func UnsafeNewQPageSizeOwned(h unsafe.Pointer) *QPageSize { return newownedQPageSize(uintptr(h)) }

func (this *QPageSize) UnsafeRelease() { this.h = 0 }

func NewQPageSize() *QPageSize {
	__r, _, _ := purego.SyscallN(addr_QPageSize_new)
	__o := newownedQPageSize(__r)
	return __o
}

func NewQPageSize2(pageSizeId QPageSize__PageSizeId) *QPageSize {
	__r, _, _ := purego.SyscallN(addr_QPageSize_new2, uintptr(pageSizeId))
	__o := newownedQPageSize(__r)
	return __o
}

func NewQPageSize3(pointSize *QSize) *QPageSize {
	__r, _, _ := purego.SyscallN(addr_QPageSize_new3, uintptr(pointSize.UnsafePointer()))
	__o := newownedQPageSize(__r)
	runtime.KeepAlive(pointSize)
	return __o
}

func NewQPageSize4(size *QSizeF, units QPageSize__Unit) *QPageSize {
	__r, _, _ := purego.SyscallN(addr_QPageSize_new4, uintptr(size.UnsafePointer()), uintptr(units))
	__o := newownedQPageSize(__r)
	runtime.KeepAlive(size)
	return __o
}

func NewQPageSize5(other *QPageSize) *QPageSize {
	__r, _, _ := purego.SyscallN(addr_QPageSize_new5, uintptr(other.UnsafePointer()))
	__o := newownedQPageSize(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQPageSize6(pointSize *QSize, name string) *QPageSize {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QPageSize_new6, uintptr(pointSize.UnsafePointer()), name_p, uintptr(len(name)))
	__o := newownedQPageSize(__r)
	runtime.KeepAlive(pointSize)
	runtime.KeepAlive(name)
	return __o
}

func NewQPageSize7(pointSize *QSize, name string, matchPolicy QPageSize__SizeMatchPolicy) *QPageSize {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QPageSize_new7, uintptr(pointSize.UnsafePointer()), name_p, uintptr(len(name)), uintptr(matchPolicy))
	__o := newownedQPageSize(__r)
	runtime.KeepAlive(pointSize)
	runtime.KeepAlive(name)
	return __o
}

func NewQPageSize8(size *QSizeF, units QPageSize__Unit, name string) *QPageSize {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QPageSize_new8, uintptr(size.UnsafePointer()), uintptr(units), name_p, uintptr(len(name)))
	__o := newownedQPageSize(__r)
	runtime.KeepAlive(size)
	runtime.KeepAlive(name)
	return __o
}

func NewQPageSize9(size *QSizeF, units QPageSize__Unit, name string, matchPolicy QPageSize__SizeMatchPolicy) *QPageSize {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QPageSize_new9, uintptr(size.UnsafePointer()), uintptr(units), name_p, uintptr(len(name)), uintptr(matchPolicy))
	__o := newownedQPageSize(__r)
	runtime.KeepAlive(size)
	runtime.KeepAlive(name)
	return __o
}

func (this *QPageSize) OperatorAssign(other *QPageSize) *QPageSize {
	__r, _, _ := purego.SyscallN(addr_QPageSize_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQPageSize(__r)
}

func (this *QPageSize) Swap(other *QPageSize) {
	purego.SyscallN(addr_QPageSize_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QPageSize) IsEquivalentTo(other *QPageSize) bool {
	__r0, _, _ := purego.SyscallN(addr_QPageSize_isEquivalentTo, this.h, uintptr(other.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return __r != 0
}

func (this *QPageSize) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QPageSize_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QPageSize) Key() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPageSize_key, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QPageSize) Name() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPageSize_name, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QPageSize) Id() QPageSize__PageSizeId {
	__r0, _, _ := purego.SyscallN(addr_QPageSize_id, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPageSize__PageSizeId(__r)
}

func (this *QPageSize) WindowsId() int {
	__r0, _, _ := purego.SyscallN(addr_QPageSize_windowsId, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QPageSize) DefinitionSize() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QPageSize_definitionSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QPageSize) DefinitionUnits() QPageSize__Unit {
	__r0, _, _ := purego.SyscallN(addr_QPageSize_definitionUnits, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QPageSize__Unit(__r)
}

func (this *QPageSize) Size(units QPageSize__Unit) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QPageSize_size, this.h, uintptr(units))
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QPageSize) SizePoints() *QSize {
	__r, _, _ := purego.SyscallN(addr_QPageSize_sizePoints, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QPageSize) SizePixels(resolution int) *QSize {
	__r, _, _ := purego.SyscallN(addr_QPageSize_sizePixels, this.h, uintptr(resolution))
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QPageSize) Rect(units QPageSize__Unit) *QRectF {
	__r, _, _ := purego.SyscallN(addr_QPageSize_rect, this.h, uintptr(units))
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QPageSize) RectPoints() *QRect {
	__r, _, _ := purego.SyscallN(addr_QPageSize_rectPoints, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QPageSize) RectPixels(resolution int) *QRect {
	__r, _, _ := purego.SyscallN(addr_QPageSize_rectPixels, this.h, uintptr(resolution))
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func QPageSize_Key(pageSizeId QPageSize__PageSizeId) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPageSize_key2, uintptr(pageSizeId), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QPageSize_Name(pageSizeId QPageSize__PageSizeId) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QPageSize_name2, uintptr(pageSizeId), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QPageSize_Id(pointSize *QSize) QPageSize__PageSizeId {
	__r0, _, _ := purego.SyscallN(addr_QPageSize_id2, uintptr(pointSize.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(pointSize)
	return QPageSize__PageSizeId(__r)
}

func QPageSize_Id2(size *QSizeF, units QPageSize__Unit) QPageSize__PageSizeId {
	__r0, _, _ := purego.SyscallN(addr_QPageSize_id3, uintptr(size.UnsafePointer()), uintptr(units))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(size)
	return QPageSize__PageSizeId(__r)
}

func QPageSize_IdWithWindowsId(windowsId int) QPageSize__PageSizeId {
	__r0, _, _ := purego.SyscallN(addr_QPageSize_id4, uintptr(windowsId))
	__r := int32(uint32(__r0))
	return QPageSize__PageSizeId(__r)
}

func QPageSize_WindowsId(pageSizeId QPageSize__PageSizeId) int {
	__r0, _, _ := purego.SyscallN(addr_QPageSize_windowsId2, uintptr(pageSizeId))
	__r := int32(uint32(__r0))
	return int(__r)
}

func QPageSize_DefinitionSize(pageSizeId QPageSize__PageSizeId) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QPageSize_definitionSize2, uintptr(pageSizeId))
	return newownedQSizeF(__r)
}

func QPageSize_DefinitionUnits(pageSizeId QPageSize__PageSizeId) QPageSize__Unit {
	__r0, _, _ := purego.SyscallN(addr_QPageSize_definitionUnits2, uintptr(pageSizeId))
	__r := int32(uint32(__r0))
	return QPageSize__Unit(__r)
}

func QPageSize_Size(pageSizeId QPageSize__PageSizeId, units QPageSize__Unit) *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QPageSize_size2, uintptr(pageSizeId), uintptr(units))
	return newownedQSizeF(__r)
}

func QPageSize_SizePoints(pageSizeId QPageSize__PageSizeId) *QSize {
	__r, _, _ := purego.SyscallN(addr_QPageSize_sizePoints2, uintptr(pageSizeId))
	return newownedQSize(__r)
}

func QPageSize_SizePixels(pageSizeId QPageSize__PageSizeId, resolution int) *QSize {
	__r, _, _ := purego.SyscallN(addr_QPageSize_sizePixels2, uintptr(pageSizeId), uintptr(resolution))
	return newownedQSize(__r)
}

func QPageSize_Id3(pointSize *QSize, matchPolicy QPageSize__SizeMatchPolicy) QPageSize__PageSizeId {
	__r0, _, _ := purego.SyscallN(addr_QPageSize_id5, uintptr(pointSize.UnsafePointer()), uintptr(matchPolicy))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(pointSize)
	return QPageSize__PageSizeId(__r)
}

func QPageSize_Id4(size *QSizeF, units QPageSize__Unit, matchPolicy QPageSize__SizeMatchPolicy) QPageSize__PageSizeId {
	__r0, _, _ := purego.SyscallN(addr_QPageSize_id6, uintptr(size.UnsafePointer()), uintptr(units), uintptr(matchPolicy))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(size)
	return QPageSize__PageSizeId(__r)
}

func (this *QPageSize) OperatorEqual(rhs *QPageSize) bool {
	__r0, _, _ := purego.SyscallN(addr_QPageSize_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPageSize) OperatorNotEqual(rhs *QPageSize) bool {
	__r0, _, _ := purego.SyscallN(addr_QPageSize_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QPageSize) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QPageSize_delete(this.h)
	}
	this.UnsafeRelease()
}
