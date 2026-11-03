package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QColorSpace_new uintptr
	addr_QColorSpace_new2 uintptr
	addr_QColorSpace_new3 uintptr
	addr_QColorSpace_new4 uintptr
	addr_QColorSpace_new5 uintptr
	pg_QColorSpace_new6 func(int32, float32) uintptr
	addr_QColorSpace_new7 uintptr
	addr_QColorSpace_new8 uintptr
	addr_QColorSpace_new9 uintptr
	addr_QColorSpace_new10 uintptr
	addr_QColorSpace_new11 uintptr
	addr_QColorSpace_new12 uintptr
	pg_QColorSpace_new13 func(uintptr, int32, float32) uintptr
	pg_QColorSpace_new14 func(int32, int32, float32) uintptr
	pg_QColorSpace_new15 func(uintptr, uintptr, uintptr, uintptr, int32, float32) uintptr
	pg_QColorSpace_new16 func(uintptr, int32, float32) uintptr
	addr_QColorSpace_OperatorAssign uintptr
	addr_QColorSpace_swap uintptr
	addr_QColorSpace_primaries uintptr
	addr_QColorSpace_transferFunction uintptr
	pg_QColorSpace_gamma func(uintptr) float32
	addr_QColorSpace_description uintptr
	addr_QColorSpace_setDescription uintptr
	addr_QColorSpace_setTransferFunction uintptr
	addr_QColorSpace_setTransferFunction2 uintptr
	addr_QColorSpace_setTransferFunctions uintptr
	addr_QColorSpace_withTransferFunction uintptr
	addr_QColorSpace_withTransferFunction2 uintptr
	addr_QColorSpace_withTransferFunctions uintptr
	addr_QColorSpace_setPrimaries uintptr
	addr_QColorSpace_setPrimaries2 uintptr
	addr_QColorSpace_setWhitePoint uintptr
	addr_QColorSpace_whitePoint uintptr
	addr_QColorSpace_setPrimaryPoints uintptr
	addr_QColorSpace_primaryPoints uintptr
	addr_QColorSpace_transformModel uintptr
	addr_QColorSpace_colorModel uintptr
	addr_QColorSpace_detach uintptr
	addr_QColorSpace_isValid uintptr
	addr_QColorSpace_isValidTarget uintptr
	addr_QColorSpace_fromIccProfile uintptr
	addr_QColorSpace_iccProfile uintptr
	addr_QColorSpace_transformationToColorSpace uintptr
	addr_QColorSpace_ToQVariant uintptr
	pg_QColorSpace_setTransferFunction3 func(uintptr, int32, float32)
	pg_QColorSpace_withTransferFunction3 func(uintptr, int32, float32) uintptr
	addr_QColorSpace_OperatorEqual uintptr
	addr_QColorSpace_OperatorNotEqual uintptr
	pg_QColorSpace_delete func(uintptr)
)

func init() {
	symaddr(&addr_QColorSpace_new, "QColorSpace_new")
	symaddr(&addr_QColorSpace_new2, "QColorSpace_new2")
	symaddr(&addr_QColorSpace_new3, "QColorSpace_new3")
	symaddr(&addr_QColorSpace_new4, "QColorSpace_new4")
	symaddr(&addr_QColorSpace_new5, "QColorSpace_new5")
	reg(&pg_QColorSpace_new6, "QColorSpace_new6")
	symaddr(&addr_QColorSpace_new7, "QColorSpace_new7")
	symaddr(&addr_QColorSpace_new8, "QColorSpace_new8")
	symaddr(&addr_QColorSpace_new9, "QColorSpace_new9")
	symaddr(&addr_QColorSpace_new10, "QColorSpace_new10")
	symaddr(&addr_QColorSpace_new11, "QColorSpace_new11")
	symaddr(&addr_QColorSpace_new12, "QColorSpace_new12")
	reg(&pg_QColorSpace_new13, "QColorSpace_new13")
	reg(&pg_QColorSpace_new14, "QColorSpace_new14")
	reg(&pg_QColorSpace_new15, "QColorSpace_new15")
	reg(&pg_QColorSpace_new16, "QColorSpace_new16")
	symaddr(&addr_QColorSpace_OperatorAssign, "QColorSpace_OperatorAssign")
	symaddr(&addr_QColorSpace_swap, "QColorSpace_swap")
	symaddr(&addr_QColorSpace_primaries, "QColorSpace_primaries")
	symaddr(&addr_QColorSpace_transferFunction, "QColorSpace_transferFunction")
	reg(&pg_QColorSpace_gamma, "QColorSpace_gamma")
	symaddr(&addr_QColorSpace_description, "QColorSpace_description")
	symaddr(&addr_QColorSpace_setDescription, "QColorSpace_setDescription")
	symaddr(&addr_QColorSpace_setTransferFunction, "QColorSpace_setTransferFunction")
	symaddr(&addr_QColorSpace_setTransferFunction2, "QColorSpace_setTransferFunction2")
	symaddr(&addr_QColorSpace_setTransferFunctions, "QColorSpace_setTransferFunctions")
	symaddr(&addr_QColorSpace_withTransferFunction, "QColorSpace_withTransferFunction")
	symaddr(&addr_QColorSpace_withTransferFunction2, "QColorSpace_withTransferFunction2")
	symaddr(&addr_QColorSpace_withTransferFunctions, "QColorSpace_withTransferFunctions")
	symaddr(&addr_QColorSpace_setPrimaries, "QColorSpace_setPrimaries")
	symaddr(&addr_QColorSpace_setPrimaries2, "QColorSpace_setPrimaries2")
	symaddr(&addr_QColorSpace_setWhitePoint, "QColorSpace_setWhitePoint")
	symaddr(&addr_QColorSpace_whitePoint, "QColorSpace_whitePoint")
	symaddr(&addr_QColorSpace_setPrimaryPoints, "QColorSpace_setPrimaryPoints")
	symaddr(&addr_QColorSpace_primaryPoints, "QColorSpace_primaryPoints")
	symaddr(&addr_QColorSpace_transformModel, "QColorSpace_transformModel")
	symaddr(&addr_QColorSpace_colorModel, "QColorSpace_colorModel")
	symaddr(&addr_QColorSpace_detach, "QColorSpace_detach")
	symaddr(&addr_QColorSpace_isValid, "QColorSpace_isValid")
	symaddr(&addr_QColorSpace_isValidTarget, "QColorSpace_isValidTarget")
	symaddr(&addr_QColorSpace_fromIccProfile, "QColorSpace_fromIccProfile")
	symaddr(&addr_QColorSpace_iccProfile, "QColorSpace_iccProfile")
	symaddr(&addr_QColorSpace_transformationToColorSpace, "QColorSpace_transformationToColorSpace")
	symaddr(&addr_QColorSpace_ToQVariant, "QColorSpace_ToQVariant")
	reg(&pg_QColorSpace_setTransferFunction3, "QColorSpace_setTransferFunction3")
	reg(&pg_QColorSpace_withTransferFunction3, "QColorSpace_withTransferFunction3")
	symaddr(&addr_QColorSpace_OperatorEqual, "QColorSpace_OperatorEqual")
	symaddr(&addr_QColorSpace_OperatorNotEqual, "QColorSpace_OperatorNotEqual")
	reg(&pg_QColorSpace_delete, "QColorSpace_delete")
}

type QColorSpace struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QColorSpace) UnsafeGoID() int64     { return this.vid }
func (this *QColorSpace) UnsafeSetGoID(id int64) { this.vid = id }

func bindQColorSpace(h uintptr) *QColorSpace {
	if h == 0 {
		return nil
	}
	return &QColorSpace{h: h}
}

func newownedQColorSpace(h uintptr) *QColorSpace {
	o := bindQColorSpace(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QColorSpace).Delete)
	}
	return o
}

func (this *QColorSpace) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QColorSpace) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQColorSpace(h unsafe.Pointer) *QColorSpace { return bindQColorSpace(uintptr(h)) }
func UnsafeNewQColorSpaceOwned(h unsafe.Pointer) *QColorSpace { return newownedQColorSpace(uintptr(h)) }

func (this *QColorSpace) UnsafeRelease() { this.h = 0 }

func NewQColorSpace() *QColorSpace {
	__r, _, _ := purego.SyscallN(addr_QColorSpace_new)
	__o := newownedQColorSpace(__r)
	return __o
}

func NewQColorSpace2(namedColorSpace QColorSpace__NamedColorSpace) *QColorSpace {
	__r, _, _ := purego.SyscallN(addr_QColorSpace_new2, uintptr(namedColorSpace))
	__o := newownedQColorSpace(__r)
	return __o
}

func NewQColorSpace3(whitePoint *QPointF, transferFunction QColorSpace__TransferFunction) *QColorSpace {
	__r, _, _ := purego.SyscallN(addr_QColorSpace_new3, uintptr(whitePoint.UnsafePointer()), uintptr(transferFunction))
	__o := newownedQColorSpace(__r)
	runtime.KeepAlive(whitePoint)
	return __o
}

func NewQColorSpace4(whitePoint *QPointF, transferFunctionTable []uint16) *QColorSpace {
	transferFunctionTable_arr := make([]int16, len(transferFunctionTable))
	for __i, __e := range transferFunctionTable {
		transferFunctionTable_arr[__i] = int16(__e)
	}
	var transferFunctionTable_p uintptr
	if len(transferFunctionTable_arr) > 0 {
		transferFunctionTable_p = uintptr(unsafe.Pointer(&transferFunctionTable_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QColorSpace_new4, uintptr(whitePoint.UnsafePointer()), transferFunctionTable_p, uintptr(len(transferFunctionTable)))
	__o := newownedQColorSpace(__r)
	runtime.KeepAlive(whitePoint)
	runtime.KeepAlive(transferFunctionTable_arr)
	return __o
}

func NewQColorSpace5(primaries QColorSpace__Primaries, transferFunction QColorSpace__TransferFunction) *QColorSpace {
	__r, _, _ := purego.SyscallN(addr_QColorSpace_new5, uintptr(primaries), uintptr(transferFunction))
	__o := newownedQColorSpace(__r)
	return __o
}

func NewQColorSpace6(primaries QColorSpace__Primaries, gamma float32) *QColorSpace {
	__o := newownedQColorSpace(pg_QColorSpace_new6(int32(primaries), gamma))
	return __o
}

func NewQColorSpace7(primaries QColorSpace__Primaries, transferFunctionTable []uint16) *QColorSpace {
	transferFunctionTable_arr := make([]int16, len(transferFunctionTable))
	for __i, __e := range transferFunctionTable {
		transferFunctionTable_arr[__i] = int16(__e)
	}
	var transferFunctionTable_p uintptr
	if len(transferFunctionTable_arr) > 0 {
		transferFunctionTable_p = uintptr(unsafe.Pointer(&transferFunctionTable_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QColorSpace_new7, uintptr(primaries), transferFunctionTable_p, uintptr(len(transferFunctionTable)))
	__o := newownedQColorSpace(__r)
	runtime.KeepAlive(transferFunctionTable_arr)
	return __o
}

func NewQColorSpace8(whitePoint *QPointF, redPoint *QPointF, greenPoint *QPointF, bluePoint *QPointF, transferFunction QColorSpace__TransferFunction) *QColorSpace {
	__r, _, _ := purego.SyscallN(addr_QColorSpace_new8, uintptr(whitePoint.UnsafePointer()), uintptr(redPoint.UnsafePointer()), uintptr(greenPoint.UnsafePointer()), uintptr(bluePoint.UnsafePointer()), uintptr(transferFunction))
	__o := newownedQColorSpace(__r)
	runtime.KeepAlive(whitePoint)
	runtime.KeepAlive(redPoint)
	runtime.KeepAlive(greenPoint)
	runtime.KeepAlive(bluePoint)
	return __o
}

func NewQColorSpace9(primaryPoints *QColorSpace__PrimaryPoints, transferFunction QColorSpace__TransferFunction) *QColorSpace {
	__r, _, _ := purego.SyscallN(addr_QColorSpace_new9, uintptr(primaryPoints.UnsafePointer()), uintptr(transferFunction))
	__o := newownedQColorSpace(__r)
	runtime.KeepAlive(primaryPoints)
	return __o
}

func NewQColorSpace10(whitePoint *QPointF, redPoint *QPointF, greenPoint *QPointF, bluePoint *QPointF, transferFunctionTable []uint16) *QColorSpace {
	transferFunctionTable_arr := make([]int16, len(transferFunctionTable))
	for __i, __e := range transferFunctionTable {
		transferFunctionTable_arr[__i] = int16(__e)
	}
	var transferFunctionTable_p uintptr
	if len(transferFunctionTable_arr) > 0 {
		transferFunctionTable_p = uintptr(unsafe.Pointer(&transferFunctionTable_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QColorSpace_new10, uintptr(whitePoint.UnsafePointer()), uintptr(redPoint.UnsafePointer()), uintptr(greenPoint.UnsafePointer()), uintptr(bluePoint.UnsafePointer()), transferFunctionTable_p, uintptr(len(transferFunctionTable)))
	__o := newownedQColorSpace(__r)
	runtime.KeepAlive(whitePoint)
	runtime.KeepAlive(redPoint)
	runtime.KeepAlive(greenPoint)
	runtime.KeepAlive(bluePoint)
	runtime.KeepAlive(transferFunctionTable_arr)
	return __o
}

func NewQColorSpace11(whitePoint *QPointF, redPoint *QPointF, greenPoint *QPointF, bluePoint *QPointF, redTransferFunctionTable []uint16, greenTransferFunctionTable []uint16, blueTransferFunctionTable []uint16) *QColorSpace {
	redTransferFunctionTable_arr := make([]int16, len(redTransferFunctionTable))
	for __i, __e := range redTransferFunctionTable {
		redTransferFunctionTable_arr[__i] = int16(__e)
	}
	var redTransferFunctionTable_p uintptr
	if len(redTransferFunctionTable_arr) > 0 {
		redTransferFunctionTable_p = uintptr(unsafe.Pointer(&redTransferFunctionTable_arr[0]))
	}
	greenTransferFunctionTable_arr := make([]int16, len(greenTransferFunctionTable))
	for __i, __e := range greenTransferFunctionTable {
		greenTransferFunctionTable_arr[__i] = int16(__e)
	}
	var greenTransferFunctionTable_p uintptr
	if len(greenTransferFunctionTable_arr) > 0 {
		greenTransferFunctionTable_p = uintptr(unsafe.Pointer(&greenTransferFunctionTable_arr[0]))
	}
	blueTransferFunctionTable_arr := make([]int16, len(blueTransferFunctionTable))
	for __i, __e := range blueTransferFunctionTable {
		blueTransferFunctionTable_arr[__i] = int16(__e)
	}
	var blueTransferFunctionTable_p uintptr
	if len(blueTransferFunctionTable_arr) > 0 {
		blueTransferFunctionTable_p = uintptr(unsafe.Pointer(&blueTransferFunctionTable_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QColorSpace_new11, uintptr(whitePoint.UnsafePointer()), uintptr(redPoint.UnsafePointer()), uintptr(greenPoint.UnsafePointer()), uintptr(bluePoint.UnsafePointer()), redTransferFunctionTable_p, uintptr(len(redTransferFunctionTable)), greenTransferFunctionTable_p, uintptr(len(greenTransferFunctionTable)), blueTransferFunctionTable_p, uintptr(len(blueTransferFunctionTable)))
	__o := newownedQColorSpace(__r)
	runtime.KeepAlive(whitePoint)
	runtime.KeepAlive(redPoint)
	runtime.KeepAlive(greenPoint)
	runtime.KeepAlive(bluePoint)
	runtime.KeepAlive(redTransferFunctionTable_arr)
	runtime.KeepAlive(greenTransferFunctionTable_arr)
	runtime.KeepAlive(blueTransferFunctionTable_arr)
	return __o
}

func NewQColorSpace12(colorSpace *QColorSpace) *QColorSpace {
	__r, _, _ := purego.SyscallN(addr_QColorSpace_new12, uintptr(colorSpace.UnsafePointer()))
	__o := newownedQColorSpace(__r)
	runtime.KeepAlive(colorSpace)
	return __o
}

func NewQColorSpace13(whitePoint *QPointF, transferFunction QColorSpace__TransferFunction, gamma float32) *QColorSpace {
	__o := newownedQColorSpace(pg_QColorSpace_new13(uintptr(whitePoint.UnsafePointer()), int32(transferFunction), gamma))
	runtime.KeepAlive(whitePoint)
	return __o
}

func NewQColorSpace14(primaries QColorSpace__Primaries, transferFunction QColorSpace__TransferFunction, gamma float32) *QColorSpace {
	__o := newownedQColorSpace(pg_QColorSpace_new14(int32(primaries), int32(transferFunction), gamma))
	return __o
}

func NewQColorSpace15(whitePoint *QPointF, redPoint *QPointF, greenPoint *QPointF, bluePoint *QPointF, transferFunction QColorSpace__TransferFunction, gamma float32) *QColorSpace {
	__o := newownedQColorSpace(pg_QColorSpace_new15(uintptr(whitePoint.UnsafePointer()), uintptr(redPoint.UnsafePointer()), uintptr(greenPoint.UnsafePointer()), uintptr(bluePoint.UnsafePointer()), int32(transferFunction), gamma))
	runtime.KeepAlive(whitePoint)
	runtime.KeepAlive(redPoint)
	runtime.KeepAlive(greenPoint)
	runtime.KeepAlive(bluePoint)
	return __o
}

func NewQColorSpace16(primaryPoints *QColorSpace__PrimaryPoints, transferFunction QColorSpace__TransferFunction, gamma float32) *QColorSpace {
	__o := newownedQColorSpace(pg_QColorSpace_new16(uintptr(primaryPoints.UnsafePointer()), int32(transferFunction), gamma))
	runtime.KeepAlive(primaryPoints)
	return __o
}

func (this *QColorSpace) OperatorAssign(colorSpace *QColorSpace) *QColorSpace {
	__r, _, _ := purego.SyscallN(addr_QColorSpace_OperatorAssign, this.h, uintptr(colorSpace.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colorSpace)
	return bindQColorSpace(__r)
}

func (this *QColorSpace) Swap(colorSpace *QColorSpace) {
	purego.SyscallN(addr_QColorSpace_swap, this.h, uintptr(colorSpace.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colorSpace)
}

func (this *QColorSpace) Primaries() QColorSpace__Primaries {
	__r0, _, _ := purego.SyscallN(addr_QColorSpace_primaries, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QColorSpace__Primaries(__r)
}

func (this *QColorSpace) TransferFunction() QColorSpace__TransferFunction {
	__r0, _, _ := purego.SyscallN(addr_QColorSpace_transferFunction, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QColorSpace__TransferFunction(__r)
}

func (this *QColorSpace) Gamma() float32 {
	__r := pg_QColorSpace_gamma(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QColorSpace) Description() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QColorSpace_description, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QColorSpace) SetDescription(description string) {
	var description_p uintptr
	if len(description) > 0 {
		description_p = uintptr(unsafe.Pointer(unsafe.StringData(description)))
	}
	purego.SyscallN(addr_QColorSpace_setDescription, this.h, description_p, uintptr(len(description)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(description)
}

func (this *QColorSpace) SetTransferFunction(transferFunction QColorSpace__TransferFunction) {
	purego.SyscallN(addr_QColorSpace_setTransferFunction, this.h, uintptr(transferFunction))
	runtime.KeepAlive(this)
}

func (this *QColorSpace) SetTransferFunctionWithTransferFunctionTable(transferFunctionTable []uint16) {
	transferFunctionTable_arr := make([]int16, len(transferFunctionTable))
	for __i, __e := range transferFunctionTable {
		transferFunctionTable_arr[__i] = int16(__e)
	}
	var transferFunctionTable_p uintptr
	if len(transferFunctionTable_arr) > 0 {
		transferFunctionTable_p = uintptr(unsafe.Pointer(&transferFunctionTable_arr[0]))
	}
	purego.SyscallN(addr_QColorSpace_setTransferFunction2, this.h, transferFunctionTable_p, uintptr(len(transferFunctionTable)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transferFunctionTable_arr)
}

func (this *QColorSpace) SetTransferFunctions(redTransferFunctionTable []uint16, greenTransferFunctionTable []uint16, blueTransferFunctionTable []uint16) {
	redTransferFunctionTable_arr := make([]int16, len(redTransferFunctionTable))
	for __i, __e := range redTransferFunctionTable {
		redTransferFunctionTable_arr[__i] = int16(__e)
	}
	var redTransferFunctionTable_p uintptr
	if len(redTransferFunctionTable_arr) > 0 {
		redTransferFunctionTable_p = uintptr(unsafe.Pointer(&redTransferFunctionTable_arr[0]))
	}
	greenTransferFunctionTable_arr := make([]int16, len(greenTransferFunctionTable))
	for __i, __e := range greenTransferFunctionTable {
		greenTransferFunctionTable_arr[__i] = int16(__e)
	}
	var greenTransferFunctionTable_p uintptr
	if len(greenTransferFunctionTable_arr) > 0 {
		greenTransferFunctionTable_p = uintptr(unsafe.Pointer(&greenTransferFunctionTable_arr[0]))
	}
	blueTransferFunctionTable_arr := make([]int16, len(blueTransferFunctionTable))
	for __i, __e := range blueTransferFunctionTable {
		blueTransferFunctionTable_arr[__i] = int16(__e)
	}
	var blueTransferFunctionTable_p uintptr
	if len(blueTransferFunctionTable_arr) > 0 {
		blueTransferFunctionTable_p = uintptr(unsafe.Pointer(&blueTransferFunctionTable_arr[0]))
	}
	purego.SyscallN(addr_QColorSpace_setTransferFunctions, this.h, redTransferFunctionTable_p, uintptr(len(redTransferFunctionTable)), greenTransferFunctionTable_p, uintptr(len(greenTransferFunctionTable)), blueTransferFunctionTable_p, uintptr(len(blueTransferFunctionTable)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(redTransferFunctionTable_arr)
	runtime.KeepAlive(greenTransferFunctionTable_arr)
	runtime.KeepAlive(blueTransferFunctionTable_arr)
}

func (this *QColorSpace) WithTransferFunction(transferFunction QColorSpace__TransferFunction) *QColorSpace {
	__r, _, _ := purego.SyscallN(addr_QColorSpace_withTransferFunction, this.h, uintptr(transferFunction))
	runtime.KeepAlive(this)
	return newownedQColorSpace(__r)
}

func (this *QColorSpace) WithTransferFunctionWithTransferFunctionTable(transferFunctionTable []uint16) *QColorSpace {
	transferFunctionTable_arr := make([]int16, len(transferFunctionTable))
	for __i, __e := range transferFunctionTable {
		transferFunctionTable_arr[__i] = int16(__e)
	}
	var transferFunctionTable_p uintptr
	if len(transferFunctionTable_arr) > 0 {
		transferFunctionTable_p = uintptr(unsafe.Pointer(&transferFunctionTable_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QColorSpace_withTransferFunction2, this.h, transferFunctionTable_p, uintptr(len(transferFunctionTable)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(transferFunctionTable_arr)
	return newownedQColorSpace(__r)
}

func (this *QColorSpace) WithTransferFunctions(redTransferFunctionTable []uint16, greenTransferFunctionTable []uint16, blueTransferFunctionTable []uint16) *QColorSpace {
	redTransferFunctionTable_arr := make([]int16, len(redTransferFunctionTable))
	for __i, __e := range redTransferFunctionTable {
		redTransferFunctionTable_arr[__i] = int16(__e)
	}
	var redTransferFunctionTable_p uintptr
	if len(redTransferFunctionTable_arr) > 0 {
		redTransferFunctionTable_p = uintptr(unsafe.Pointer(&redTransferFunctionTable_arr[0]))
	}
	greenTransferFunctionTable_arr := make([]int16, len(greenTransferFunctionTable))
	for __i, __e := range greenTransferFunctionTable {
		greenTransferFunctionTable_arr[__i] = int16(__e)
	}
	var greenTransferFunctionTable_p uintptr
	if len(greenTransferFunctionTable_arr) > 0 {
		greenTransferFunctionTable_p = uintptr(unsafe.Pointer(&greenTransferFunctionTable_arr[0]))
	}
	blueTransferFunctionTable_arr := make([]int16, len(blueTransferFunctionTable))
	for __i, __e := range blueTransferFunctionTable {
		blueTransferFunctionTable_arr[__i] = int16(__e)
	}
	var blueTransferFunctionTable_p uintptr
	if len(blueTransferFunctionTable_arr) > 0 {
		blueTransferFunctionTable_p = uintptr(unsafe.Pointer(&blueTransferFunctionTable_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QColorSpace_withTransferFunctions, this.h, redTransferFunctionTable_p, uintptr(len(redTransferFunctionTable)), greenTransferFunctionTable_p, uintptr(len(greenTransferFunctionTable)), blueTransferFunctionTable_p, uintptr(len(blueTransferFunctionTable)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(redTransferFunctionTable_arr)
	runtime.KeepAlive(greenTransferFunctionTable_arr)
	runtime.KeepAlive(blueTransferFunctionTable_arr)
	return newownedQColorSpace(__r)
}

func (this *QColorSpace) SetPrimaries(primariesId QColorSpace__Primaries) {
	purego.SyscallN(addr_QColorSpace_setPrimaries, this.h, uintptr(primariesId))
	runtime.KeepAlive(this)
}

func (this *QColorSpace) SetPrimaries2(whitePoint *QPointF, redPoint *QPointF, greenPoint *QPointF, bluePoint *QPointF) {
	purego.SyscallN(addr_QColorSpace_setPrimaries2, this.h, uintptr(whitePoint.UnsafePointer()), uintptr(redPoint.UnsafePointer()), uintptr(greenPoint.UnsafePointer()), uintptr(bluePoint.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(whitePoint)
	runtime.KeepAlive(redPoint)
	runtime.KeepAlive(greenPoint)
	runtime.KeepAlive(bluePoint)
}

func (this *QColorSpace) SetWhitePoint(whitePoint *QPointF) {
	purego.SyscallN(addr_QColorSpace_setWhitePoint, this.h, uintptr(whitePoint.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(whitePoint)
}

func (this *QColorSpace) WhitePoint() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QColorSpace_whitePoint, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QColorSpace) SetPrimaryPoints(primaryPoints *QColorSpace__PrimaryPoints) {
	purego.SyscallN(addr_QColorSpace_setPrimaryPoints, this.h, uintptr(primaryPoints.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(primaryPoints)
}

func (this *QColorSpace) PrimaryPoints() *QColorSpace__PrimaryPoints {
	__r, _, _ := purego.SyscallN(addr_QColorSpace_primaryPoints, this.h)
	runtime.KeepAlive(this)
	return newownedQColorSpace__PrimaryPoints(__r)
}

func (this *QColorSpace) TransformModel() QColorSpace__TransformModel {
	__r0, _, _ := purego.SyscallN(addr_QColorSpace_transformModel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QColorSpace__TransformModel(__r)
}

func (this *QColorSpace) ColorModel() QColorSpace__ColorModel {
	__r0, _, _ := purego.SyscallN(addr_QColorSpace_colorModel, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QColorSpace__ColorModel(__r)
}

func (this *QColorSpace) Detach() {
	purego.SyscallN(addr_QColorSpace_detach, this.h)
	runtime.KeepAlive(this)
}

func (this *QColorSpace) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QColorSpace_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QColorSpace) IsValidTarget() bool {
	__r0, _, _ := purego.SyscallN(addr_QColorSpace_isValidTarget, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func QColorSpace_FromIccProfile(iccProfile *QByteArray) *QColorSpace {
	__r, _, _ := purego.SyscallN(addr_QColorSpace_fromIccProfile, uintptr(iccProfile.UnsafePointer()))
	runtime.KeepAlive(iccProfile)
	return newownedQColorSpace(__r)
}

func (this *QColorSpace) IccProfile() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QColorSpace_iccProfile, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QColorSpace) TransformationToColorSpace(colorspace *QColorSpace) *QColorTransform {
	__r, _, _ := purego.SyscallN(addr_QColorSpace_transformationToColorSpace, this.h, uintptr(colorspace.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colorspace)
	return newownedQColorTransform(__r)
}

func (this *QColorSpace) ToQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QColorSpace_ToQVariant, this.h)
	runtime.KeepAlive(this)
	return newownedQVariant(__r)
}

func (this *QColorSpace) SetTransferFunction2(transferFunction QColorSpace__TransferFunction, gamma float32) {
	pg_QColorSpace_setTransferFunction3(this.h, int32(transferFunction), gamma)
	runtime.KeepAlive(this)
}

func (this *QColorSpace) WithTransferFunction2(transferFunction QColorSpace__TransferFunction, gamma float32) *QColorSpace {
	__r := pg_QColorSpace_withTransferFunction3(this.h, int32(transferFunction), gamma)
	runtime.KeepAlive(this)
	return newownedQColorSpace(__r)
}

func (this *QColorSpace) OperatorEqual(colorSpace2 *QColorSpace) bool {
	__r0, _, _ := purego.SyscallN(addr_QColorSpace_OperatorEqual, this.h, uintptr(colorSpace2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colorSpace2)
	return __r != 0
}

func (this *QColorSpace) OperatorNotEqual(colorSpace2 *QColorSpace) bool {
	__r0, _, _ := purego.SyscallN(addr_QColorSpace_OperatorNotEqual, this.h, uintptr(colorSpace2.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(colorSpace2)
	return __r != 0
}

func (this *QColorSpace) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QColorSpace_delete(this.h)
	}
	this.UnsafeRelease()
}
