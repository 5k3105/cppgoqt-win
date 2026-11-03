package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

// purego: unsupported shapes skipped (map returns, >15-arg / >9-arg-callback, exotic): NewQVariant17, NewQVariant22

var (
	addr_QVariant_new uintptr
	addr_QVariant_new2 uintptr
	addr_QVariant_new3 uintptr
	addr_QVariant_new4 uintptr
	addr_QVariant_new5 uintptr
	addr_QVariant_new6 uintptr
	addr_QVariant_new7 uintptr
	addr_QVariant_new8 uintptr
	pg_QVariant_new9 func(float64) uintptr
	pg_QVariant_new10 func(float32) uintptr
	addr_QVariant_new11 uintptr
	addr_QVariant_new12 uintptr
	addr_QVariant_new13 uintptr
	addr_QVariant_new14 uintptr
	addr_QVariant_new15 uintptr
	addr_QVariant_new16 uintptr
	addr_QVariant_new18 uintptr
	addr_QVariant_new19 uintptr
	addr_QVariant_new20 uintptr
	addr_QVariant_new21 uintptr
	addr_QVariant_new23 uintptr
	addr_QVariant_new24 uintptr
	addr_QVariant_new25 uintptr
	addr_QVariant_new26 uintptr
	addr_QVariant_new27 uintptr
	addr_QVariant_new28 uintptr
	addr_QVariant_new29 uintptr
	addr_QVariant_new30 uintptr
	addr_QVariant_new31 uintptr
	addr_QVariant_new32 uintptr
	addr_QVariant_new33 uintptr
	addr_QVariant_new34 uintptr
	addr_QVariant_new35 uintptr
	addr_QVariant_new36 uintptr
	addr_QVariant_new37 uintptr
	addr_QVariant_new38 uintptr
	addr_QVariant_new39 uintptr
	addr_QVariant_new40 uintptr
	addr_QVariant_new41 uintptr
	addr_QVariant_new42 uintptr
	addr_QVariant_new43 uintptr
	addr_QVariant_OperatorAssign uintptr
	addr_QVariant_swap uintptr
	addr_QVariant_userType uintptr
	addr_QVariant_typeId uintptr
	addr_QVariant_typeName uintptr
	addr_QVariant_metaType uintptr
	addr_QVariant_canConvert uintptr
	addr_QVariant_convert uintptr
	addr_QVariant_canView uintptr
	addr_QVariant_canConvert2 uintptr
	addr_QVariant_convert2 uintptr
	addr_QVariant_isValid uintptr
	addr_QVariant_isNull uintptr
	addr_QVariant_clear uintptr
	addr_QVariant_detach uintptr
	addr_QVariant_isDetached uintptr
	addr_QVariant_toInt uintptr
	addr_QVariant_toUInt uintptr
	addr_QVariant_toLongLong uintptr
	addr_QVariant_toULongLong uintptr
	addr_QVariant_toBool uintptr
	pg_QVariant_toDouble func(uintptr) float64
	pg_QVariant_toFloat func(uintptr) float32
	pg_QVariant_toReal func(uintptr) float64
	addr_QVariant_toByteArray uintptr
	addr_QVariant_toBitArray uintptr
	addr_QVariant_toString uintptr
	addr_QVariant_toStringList uintptr
	addr_QVariant_toChar uintptr
	addr_QVariant_toDate uintptr
	addr_QVariant_toTime uintptr
	addr_QVariant_toDateTime uintptr
	addr_QVariant_toList uintptr
	addr_QVariant_toPoint uintptr
	addr_QVariant_toPointF uintptr
	addr_QVariant_toRect uintptr
	addr_QVariant_toSize uintptr
	addr_QVariant_toSizeF uintptr
	addr_QVariant_toLine uintptr
	addr_QVariant_toLineF uintptr
	addr_QVariant_toRectF uintptr
	addr_QVariant_toLocale uintptr
	addr_QVariant_toRegularExpression uintptr
	addr_QVariant_toEasingCurve uintptr
	addr_QVariant_toUuid uintptr
	addr_QVariant_toUrl uintptr
	addr_QVariant_toJsonValue uintptr
	addr_QVariant_toJsonObject uintptr
	addr_QVariant_toJsonArray uintptr
	addr_QVariant_toJsonDocument uintptr
	addr_QVariant_toModelIndex uintptr
	addr_QVariant_toPersistentModelIndex uintptr
	addr_QVariant_load uintptr
	addr_QVariant_save uintptr
	addr_QVariant_type uintptr
	addr_QVariant_typeToName uintptr
	addr_QVariant_nameToType uintptr
	addr_QVariant_setValue uintptr
	addr_QVariant_fromMetaType uintptr
	addr_QVariant_compare uintptr
	addr_QVariant_toInt2 uintptr
	addr_QVariant_toUInt2 uintptr
	addr_QVariant_toLongLong2 uintptr
	addr_QVariant_toULongLong2 uintptr
	pg_QVariant_toDouble2 func(uintptr, uintptr) float64
	pg_QVariant_toFloat2 func(uintptr, uintptr) float32
	pg_QVariant_toReal2 func(uintptr, uintptr) float64
	addr_QVariant_OperatorEqual uintptr
	addr_QVariant_OperatorNotEqual uintptr
	pg_QVariant_delete func(uintptr)
)

func init() {
	symaddr(&addr_QVariant_new, "QVariant_new")
	symaddr(&addr_QVariant_new2, "QVariant_new2")
	symaddr(&addr_QVariant_new3, "QVariant_new3")
	symaddr(&addr_QVariant_new4, "QVariant_new4")
	symaddr(&addr_QVariant_new5, "QVariant_new5")
	symaddr(&addr_QVariant_new6, "QVariant_new6")
	symaddr(&addr_QVariant_new7, "QVariant_new7")
	symaddr(&addr_QVariant_new8, "QVariant_new8")
	reg(&pg_QVariant_new9, "QVariant_new9")
	reg(&pg_QVariant_new10, "QVariant_new10")
	symaddr(&addr_QVariant_new11, "QVariant_new11")
	symaddr(&addr_QVariant_new12, "QVariant_new12")
	symaddr(&addr_QVariant_new13, "QVariant_new13")
	symaddr(&addr_QVariant_new14, "QVariant_new14")
	symaddr(&addr_QVariant_new15, "QVariant_new15")
	symaddr(&addr_QVariant_new16, "QVariant_new16")
	symaddr(&addr_QVariant_new18, "QVariant_new18")
	symaddr(&addr_QVariant_new19, "QVariant_new19")
	symaddr(&addr_QVariant_new20, "QVariant_new20")
	symaddr(&addr_QVariant_new21, "QVariant_new21")
	symaddr(&addr_QVariant_new23, "QVariant_new23")
	symaddr(&addr_QVariant_new24, "QVariant_new24")
	symaddr(&addr_QVariant_new25, "QVariant_new25")
	symaddr(&addr_QVariant_new26, "QVariant_new26")
	symaddr(&addr_QVariant_new27, "QVariant_new27")
	symaddr(&addr_QVariant_new28, "QVariant_new28")
	symaddr(&addr_QVariant_new29, "QVariant_new29")
	symaddr(&addr_QVariant_new30, "QVariant_new30")
	symaddr(&addr_QVariant_new31, "QVariant_new31")
	symaddr(&addr_QVariant_new32, "QVariant_new32")
	symaddr(&addr_QVariant_new33, "QVariant_new33")
	symaddr(&addr_QVariant_new34, "QVariant_new34")
	symaddr(&addr_QVariant_new35, "QVariant_new35")
	symaddr(&addr_QVariant_new36, "QVariant_new36")
	symaddr(&addr_QVariant_new37, "QVariant_new37")
	symaddr(&addr_QVariant_new38, "QVariant_new38")
	symaddr(&addr_QVariant_new39, "QVariant_new39")
	symaddr(&addr_QVariant_new40, "QVariant_new40")
	symaddr(&addr_QVariant_new41, "QVariant_new41")
	symaddr(&addr_QVariant_new42, "QVariant_new42")
	symaddr(&addr_QVariant_new43, "QVariant_new43")
	symaddr(&addr_QVariant_OperatorAssign, "QVariant_OperatorAssign")
	symaddr(&addr_QVariant_swap, "QVariant_swap")
	symaddr(&addr_QVariant_userType, "QVariant_userType")
	symaddr(&addr_QVariant_typeId, "QVariant_typeId")
	symaddr(&addr_QVariant_typeName, "QVariant_typeName")
	symaddr(&addr_QVariant_metaType, "QVariant_metaType")
	symaddr(&addr_QVariant_canConvert, "QVariant_canConvert")
	symaddr(&addr_QVariant_convert, "QVariant_convert")
	symaddr(&addr_QVariant_canView, "QVariant_canView")
	symaddr(&addr_QVariant_canConvert2, "QVariant_canConvert2")
	symaddr(&addr_QVariant_convert2, "QVariant_convert2")
	symaddr(&addr_QVariant_isValid, "QVariant_isValid")
	symaddr(&addr_QVariant_isNull, "QVariant_isNull")
	symaddr(&addr_QVariant_clear, "QVariant_clear")
	symaddr(&addr_QVariant_detach, "QVariant_detach")
	symaddr(&addr_QVariant_isDetached, "QVariant_isDetached")
	symaddr(&addr_QVariant_toInt, "QVariant_toInt")
	symaddr(&addr_QVariant_toUInt, "QVariant_toUInt")
	symaddr(&addr_QVariant_toLongLong, "QVariant_toLongLong")
	symaddr(&addr_QVariant_toULongLong, "QVariant_toULongLong")
	symaddr(&addr_QVariant_toBool, "QVariant_toBool")
	reg(&pg_QVariant_toDouble, "QVariant_toDouble")
	reg(&pg_QVariant_toFloat, "QVariant_toFloat")
	reg(&pg_QVariant_toReal, "QVariant_toReal")
	symaddr(&addr_QVariant_toByteArray, "QVariant_toByteArray")
	symaddr(&addr_QVariant_toBitArray, "QVariant_toBitArray")
	symaddr(&addr_QVariant_toString, "QVariant_toString")
	symaddr(&addr_QVariant_toStringList, "QVariant_toStringList")
	symaddr(&addr_QVariant_toChar, "QVariant_toChar")
	symaddr(&addr_QVariant_toDate, "QVariant_toDate")
	symaddr(&addr_QVariant_toTime, "QVariant_toTime")
	symaddr(&addr_QVariant_toDateTime, "QVariant_toDateTime")
	symaddr(&addr_QVariant_toList, "QVariant_toList")
	symaddr(&addr_QVariant_toPoint, "QVariant_toPoint")
	symaddr(&addr_QVariant_toPointF, "QVariant_toPointF")
	symaddr(&addr_QVariant_toRect, "QVariant_toRect")
	symaddr(&addr_QVariant_toSize, "QVariant_toSize")
	symaddr(&addr_QVariant_toSizeF, "QVariant_toSizeF")
	symaddr(&addr_QVariant_toLine, "QVariant_toLine")
	symaddr(&addr_QVariant_toLineF, "QVariant_toLineF")
	symaddr(&addr_QVariant_toRectF, "QVariant_toRectF")
	symaddr(&addr_QVariant_toLocale, "QVariant_toLocale")
	symaddr(&addr_QVariant_toRegularExpression, "QVariant_toRegularExpression")
	symaddr(&addr_QVariant_toEasingCurve, "QVariant_toEasingCurve")
	symaddr(&addr_QVariant_toUuid, "QVariant_toUuid")
	symaddr(&addr_QVariant_toUrl, "QVariant_toUrl")
	symaddr(&addr_QVariant_toJsonValue, "QVariant_toJsonValue")
	symaddr(&addr_QVariant_toJsonObject, "QVariant_toJsonObject")
	symaddr(&addr_QVariant_toJsonArray, "QVariant_toJsonArray")
	symaddr(&addr_QVariant_toJsonDocument, "QVariant_toJsonDocument")
	symaddr(&addr_QVariant_toModelIndex, "QVariant_toModelIndex")
	symaddr(&addr_QVariant_toPersistentModelIndex, "QVariant_toPersistentModelIndex")
	symaddr(&addr_QVariant_load, "QVariant_load")
	symaddr(&addr_QVariant_save, "QVariant_save")
	symaddr(&addr_QVariant_type, "QVariant_type")
	symaddr(&addr_QVariant_typeToName, "QVariant_typeToName")
	symaddr(&addr_QVariant_nameToType, "QVariant_nameToType")
	symaddr(&addr_QVariant_setValue, "QVariant_setValue")
	symaddr(&addr_QVariant_fromMetaType, "QVariant_fromMetaType")
	symaddr(&addr_QVariant_compare, "QVariant_compare")
	symaddr(&addr_QVariant_toInt2, "QVariant_toInt2")
	symaddr(&addr_QVariant_toUInt2, "QVariant_toUInt2")
	symaddr(&addr_QVariant_toLongLong2, "QVariant_toLongLong2")
	symaddr(&addr_QVariant_toULongLong2, "QVariant_toULongLong2")
	reg(&pg_QVariant_toDouble2, "QVariant_toDouble2")
	reg(&pg_QVariant_toFloat2, "QVariant_toFloat2")
	reg(&pg_QVariant_toReal2, "QVariant_toReal2")
	symaddr(&addr_QVariant_OperatorEqual, "QVariant_OperatorEqual")
	symaddr(&addr_QVariant_OperatorNotEqual, "QVariant_OperatorNotEqual")
	reg(&pg_QVariant_delete, "QVariant_delete")
}

type QVariant struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QVariant) UnsafeGoID() int64     { return this.vid }
func (this *QVariant) UnsafeSetGoID(id int64) { this.vid = id }

func bindQVariant(h uintptr) *QVariant {
	if h == 0 {
		return nil
	}
	return &QVariant{h: h}
}

func newownedQVariant(h uintptr) *QVariant {
	o := bindQVariant(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QVariant).Delete)
	}
	return o
}

func (this *QVariant) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QVariant) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQVariant(h unsafe.Pointer) *QVariant { return bindQVariant(uintptr(h)) }
func UnsafeNewQVariantOwned(h unsafe.Pointer) *QVariant { return newownedQVariant(uintptr(h)) }

func (this *QVariant) UnsafeRelease() { this.h = 0 }

func NewQVariant() *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new)
	__o := newownedQVariant(__r)
	return __o
}

func NewQVariant2(type_ *QMetaType) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new2, uintptr(type_.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(type_)
	return __o
}

func NewQVariant3(other *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new3, uintptr(other.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQVariant4(i int) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new4, uintptr(i))
	__o := newownedQVariant(__r)
	return __o
}

func NewQVariant5(ui uint) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new5, uintptr(ui))
	__o := newownedQVariant(__r)
	return __o
}

func NewQVariant6(ll int64) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new6, uintptr(ll))
	__o := newownedQVariant(__r)
	return __o
}

func NewQVariant7(ull uint64) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new7, uintptr(ull))
	__o := newownedQVariant(__r)
	return __o
}

func NewQVariant8(b bool) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new8, uintptr(b2i(b)))
	__o := newownedQVariant(__r)
	return __o
}

func NewQVariant9(d float64) *QVariant {
	__o := newownedQVariant(pg_QVariant_new9(d))
	return __o
}

func NewQVariant10(f float32) *QVariant {
	__o := newownedQVariant(pg_QVariant_new10(f))
	return __o
}

func NewQVariant11(qchar *QChar) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new11, uintptr(qchar.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(qchar)
	return __o
}

func NewQVariant12(date *QDate) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new12, uintptr(date.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(date)
	return __o
}

func NewQVariant13(time *QTime) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new13, uintptr(time.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(time)
	return __o
}

func NewQVariant14(bitarray *QBitArray) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new14, uintptr(bitarray.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(bitarray)
	return __o
}

func NewQVariant15(bytearray *QByteArray) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new15, uintptr(bytearray.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(bytearray)
	return __o
}

func NewQVariant16(datetime *QDateTime) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new16, uintptr(datetime.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(datetime)
	return __o
}

func NewQVariant18(jsonArray *QJsonArray) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new18, uintptr(jsonArray.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(jsonArray)
	return __o
}

func NewQVariant19(jsonObject *QJsonObject) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new19, uintptr(jsonObject.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(jsonObject)
	return __o
}

func NewQVariant20(list []*QVariant) *QVariant {
	list_arr := make([]uintptr, len(list))
	for __i, __e := range list {
		list_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var list_p uintptr
	if len(list_arr) > 0 {
		list_p = uintptr(unsafe.Pointer(&list_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QVariant_new20, list_p, uintptr(len(list)))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(list)
	runtime.KeepAlive(list_arr)
	return __o
}

func NewQVariant21(locale *QLocale) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new21, uintptr(locale.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(locale)
	return __o
}

func NewQVariant23(re *QRegularExpression) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new23, uintptr(re.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(re)
	return __o
}

func NewQVariant24(string_ string) *QVariant {
	var string__p uintptr
	if len(string_) > 0 {
		string__p = uintptr(unsafe.Pointer(unsafe.StringData(string_)))
	}
	__r, _, _ := purego.SyscallN(addr_QVariant_new24, string__p, uintptr(len(string_)))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(string_)
	return __o
}

func NewQVariant25(stringlist []string) *QVariant {
	stringlist_arr := make([]uintptr, len(stringlist))
	for __i, __e := range stringlist {
		stringlist_arr[__i] = pgCString(__e)
	}
	var stringlist_p uintptr
	if len(stringlist_arr) > 0 {
		stringlist_p = uintptr(unsafe.Pointer(&stringlist_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QVariant_new25, stringlist_p, uintptr(len(stringlist)))
	__o := newownedQVariant(__r)
	for __i := range stringlist_arr {
		pgFree(stringlist_arr[__i])
	}
	runtime.KeepAlive(stringlist_arr)
	return __o
}

func NewQVariant26(url *QUrl) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new26, uintptr(url.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(url)
	return __o
}

func NewQVariant27(jsonValue *QJsonValue) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new27, uintptr(jsonValue.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(jsonValue)
	return __o
}

func NewQVariant28(modelIndex *QModelIndex) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new28, uintptr(modelIndex.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(modelIndex)
	return __o
}

func NewQVariant29(uuid *QUuid) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new29, uintptr(uuid.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(uuid)
	return __o
}

func NewQVariant30(size *QSize) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new30, uintptr(size.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(size)
	return __o
}

func NewQVariant31(size *QSizeF) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new31, uintptr(size.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(size)
	return __o
}

func NewQVariant32(pt *QPoint) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new32, uintptr(pt.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(pt)
	return __o
}

func NewQVariant33(pt *QPointF) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new33, uintptr(pt.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(pt)
	return __o
}

func NewQVariant34(line *QLine) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new34, uintptr(line.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(line)
	return __o
}

func NewQVariant35(line *QLineF) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new35, uintptr(line.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(line)
	return __o
}

func NewQVariant36(rect *QRect) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new36, uintptr(rect.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(rect)
	return __o
}

func NewQVariant37(rect *QRectF) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new37, uintptr(rect.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(rect)
	return __o
}

func NewQVariant38(easing *QEasingCurve) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new38, uintptr(easing.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(easing)
	return __o
}

func NewQVariant39(jsonDocument *QJsonDocument) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new39, uintptr(jsonDocument.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(jsonDocument)
	return __o
}

func NewQVariant40(modelIndex *QPersistentModelIndex) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new40, uintptr(modelIndex.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(modelIndex)
	return __o
}

func NewQVariant41(str string) *QVariant {
	str_b := append([]byte(str), 0)
	__r, _, _ := purego.SyscallN(addr_QVariant_new41, uintptr(unsafe.Pointer(&str_b[0])))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(str_b)
	return __o
}

func NewQVariant42(string_ *QLatin1String) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new42, uintptr(string_.UnsafePointer()))
	__o := newownedQVariant(__r)
	runtime.KeepAlive(string_)
	return __o
}

func NewQVariant43(type_ QVariant__Type) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_new43, uintptr(type_))
	__o := newownedQVariant(__r)
	return __o
}

func (this *QVariant) OperatorAssign(other *QVariant) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQVariant(__r)
}

func (this *QVariant) Swap(other *QVariant) {
	purego.SyscallN(addr_QVariant_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QVariant) UserType() int {
	__r0, _, _ := purego.SyscallN(addr_QVariant_userType, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QVariant) TypeId() int {
	__r0, _, _ := purego.SyscallN(addr_QVariant_typeId, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QVariant) TypeName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QVariant_typeName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QVariant) MetaType() *QMetaType {
	__r, _, _ := purego.SyscallN(addr_QVariant_metaType, this.h)
	runtime.KeepAlive(this)
	return newownedQMetaType(__r)
}

func (this *QVariant) CanConvert(targetType *QMetaType) bool {
	__r0, _, _ := purego.SyscallN(addr_QVariant_canConvert, this.h, uintptr(targetType.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(targetType)
	return __r != 0
}

func (this *QVariant) Convert(type_ *QMetaType) bool {
	__r0, _, _ := purego.SyscallN(addr_QVariant_convert, this.h, uintptr(type_.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(type_)
	return __r != 0
}

func (this *QVariant) CanView(targetType *QMetaType) bool {
	__r0, _, _ := purego.SyscallN(addr_QVariant_canView, this.h, uintptr(targetType.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(targetType)
	return __r != 0
}

func (this *QVariant) CanConvertWithTargetTypeId(targetTypeId int) bool {
	__r0, _, _ := purego.SyscallN(addr_QVariant_canConvert2, this.h, uintptr(targetTypeId))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QVariant) ConvertWithTargetTypeId(targetTypeId int) bool {
	__r0, _, _ := purego.SyscallN(addr_QVariant_convert2, this.h, uintptr(targetTypeId))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QVariant) IsValid() bool {
	__r0, _, _ := purego.SyscallN(addr_QVariant_isValid, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QVariant) IsNull() bool {
	__r0, _, _ := purego.SyscallN(addr_QVariant_isNull, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QVariant) Clear() {
	purego.SyscallN(addr_QVariant_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QVariant) Detach() {
	purego.SyscallN(addr_QVariant_detach, this.h)
	runtime.KeepAlive(this)
}

func (this *QVariant) IsDetached() bool {
	__r0, _, _ := purego.SyscallN(addr_QVariant_isDetached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QVariant) ToInt() int {
	__r0, _, _ := purego.SyscallN(addr_QVariant_toInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r)
}

func (this *QVariant) ToUInt() uint {
	__r0, _, _ := purego.SyscallN(addr_QVariant_toUInt, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r)
}

func (this *QVariant) ToLongLong() int64 {
	__r0, _, _ := purego.SyscallN(addr_QVariant_toLongLong, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVariant) ToULongLong() uint64 {
	__r0, _, _ := purego.SyscallN(addr_QVariant_toULongLong, this.h)
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r)
}

func (this *QVariant) ToBool() bool {
	__r0, _, _ := purego.SyscallN(addr_QVariant_toBool, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QVariant) ToDouble() float64 {
	__r := pg_QVariant_toDouble(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVariant) ToFloat() float32 {
	__r := pg_QVariant_toFloat(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVariant) ToReal() float64 {
	__r := pg_QVariant_toReal(this.h)
	runtime.KeepAlive(this)
	return __r
}

func (this *QVariant) ToByteArray() *QByteArray {
	__r, _, _ := purego.SyscallN(addr_QVariant_toByteArray, this.h)
	runtime.KeepAlive(this)
	return newownedQByteArray(__r)
}

func (this *QVariant) ToBitArray() *QBitArray {
	__r, _, _ := purego.SyscallN(addr_QVariant_toBitArray, this.h)
	runtime.KeepAlive(this)
	return newownedQBitArray(__r)
}

func (this *QVariant) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QVariant_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QVariant) ToStringList() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QVariant_toStringList, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QVariant) ToChar() *QChar {
	__r, _, _ := purego.SyscallN(addr_QVariant_toChar, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QVariant) ToDate() *QDate {
	__r, _, _ := purego.SyscallN(addr_QVariant_toDate, this.h)
	runtime.KeepAlive(this)
	return newownedQDate(__r)
}

func (this *QVariant) ToTime() *QTime {
	__r, _, _ := purego.SyscallN(addr_QVariant_toTime, this.h)
	runtime.KeepAlive(this)
	return newownedQTime(__r)
}

func (this *QVariant) ToDateTime() *QDateTime {
	__r, _, _ := purego.SyscallN(addr_QVariant_toDateTime, this.h)
	runtime.KeepAlive(this)
	return newownedQDateTime(__r)
}

func (this *QVariant) ToList() []*QVariant {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QVariant_toList, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	__out := make([]*QVariant, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = newownedQVariant(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QVariant) ToPoint() *QPoint {
	__r, _, _ := purego.SyscallN(addr_QVariant_toPoint, this.h)
	runtime.KeepAlive(this)
	return newownedQPoint(__r)
}

func (this *QVariant) ToPointF() *QPointF {
	__r, _, _ := purego.SyscallN(addr_QVariant_toPointF, this.h)
	runtime.KeepAlive(this)
	return newownedQPointF(__r)
}

func (this *QVariant) ToRect() *QRect {
	__r, _, _ := purego.SyscallN(addr_QVariant_toRect, this.h)
	runtime.KeepAlive(this)
	return newownedQRect(__r)
}

func (this *QVariant) ToSize() *QSize {
	__r, _, _ := purego.SyscallN(addr_QVariant_toSize, this.h)
	runtime.KeepAlive(this)
	return newownedQSize(__r)
}

func (this *QVariant) ToSizeF() *QSizeF {
	__r, _, _ := purego.SyscallN(addr_QVariant_toSizeF, this.h)
	runtime.KeepAlive(this)
	return newownedQSizeF(__r)
}

func (this *QVariant) ToLine() *QLine {
	__r, _, _ := purego.SyscallN(addr_QVariant_toLine, this.h)
	runtime.KeepAlive(this)
	return newownedQLine(__r)
}

func (this *QVariant) ToLineF() *QLineF {
	__r, _, _ := purego.SyscallN(addr_QVariant_toLineF, this.h)
	runtime.KeepAlive(this)
	return newownedQLineF(__r)
}

func (this *QVariant) ToRectF() *QRectF {
	__r, _, _ := purego.SyscallN(addr_QVariant_toRectF, this.h)
	runtime.KeepAlive(this)
	return newownedQRectF(__r)
}

func (this *QVariant) ToLocale() *QLocale {
	__r, _, _ := purego.SyscallN(addr_QVariant_toLocale, this.h)
	runtime.KeepAlive(this)
	return newownedQLocale(__r)
}

func (this *QVariant) ToRegularExpression() *QRegularExpression {
	__r, _, _ := purego.SyscallN(addr_QVariant_toRegularExpression, this.h)
	runtime.KeepAlive(this)
	return newownedQRegularExpression(__r)
}

func (this *QVariant) ToEasingCurve() *QEasingCurve {
	__r, _, _ := purego.SyscallN(addr_QVariant_toEasingCurve, this.h)
	runtime.KeepAlive(this)
	return newownedQEasingCurve(__r)
}

func (this *QVariant) ToUuid() *QUuid {
	__r, _, _ := purego.SyscallN(addr_QVariant_toUuid, this.h)
	runtime.KeepAlive(this)
	return newownedQUuid(__r)
}

func (this *QVariant) ToUrl() *QUrl {
	__r, _, _ := purego.SyscallN(addr_QVariant_toUrl, this.h)
	runtime.KeepAlive(this)
	return newownedQUrl(__r)
}

func (this *QVariant) ToJsonValue() *QJsonValue {
	__r, _, _ := purego.SyscallN(addr_QVariant_toJsonValue, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonValue(__r)
}

func (this *QVariant) ToJsonObject() *QJsonObject {
	__r, _, _ := purego.SyscallN(addr_QVariant_toJsonObject, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonObject(__r)
}

func (this *QVariant) ToJsonArray() *QJsonArray {
	__r, _, _ := purego.SyscallN(addr_QVariant_toJsonArray, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonArray(__r)
}

func (this *QVariant) ToJsonDocument() *QJsonDocument {
	__r, _, _ := purego.SyscallN(addr_QVariant_toJsonDocument, this.h)
	runtime.KeepAlive(this)
	return newownedQJsonDocument(__r)
}

func (this *QVariant) ToModelIndex() *QModelIndex {
	__r, _, _ := purego.SyscallN(addr_QVariant_toModelIndex, this.h)
	runtime.KeepAlive(this)
	return newownedQModelIndex(__r)
}

func (this *QVariant) ToPersistentModelIndex() *QPersistentModelIndex {
	__r, _, _ := purego.SyscallN(addr_QVariant_toPersistentModelIndex, this.h)
	runtime.KeepAlive(this)
	return newownedQPersistentModelIndex(__r)
}

func (this *QVariant) Load(ds *QDataStream) {
	purego.SyscallN(addr_QVariant_load, this.h, uintptr(ds.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ds)
}

func (this *QVariant) Save(ds *QDataStream) {
	purego.SyscallN(addr_QVariant_save, this.h, uintptr(ds.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(ds)
}

func (this *QVariant) Type() QVariant__Type {
	__r0, _, _ := purego.SyscallN(addr_QVariant_type, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QVariant__Type(__r)
}

func QVariant_TypeToName(typeId int) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QVariant_typeToName, uintptr(typeId), uintptr(unsafe.Pointer(&__n)))
	return pgStr(__r, __n)
}

func QVariant_NameToType(name string) QVariant__Type {
	name_b := append([]byte(name), 0)
	__r0, _, _ := purego.SyscallN(addr_QVariant_nameToType, uintptr(unsafe.Pointer(&name_b[0])))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(name_b)
	return QVariant__Type(__r)
}

func (this *QVariant) SetValue(avalue *QVariant) {
	purego.SyscallN(addr_QVariant_setValue, this.h, uintptr(avalue.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(avalue)
}

func QVariant_FromMetaType(type_ *QMetaType) *QVariant {
	__r, _, _ := purego.SyscallN(addr_QVariant_fromMetaType, uintptr(type_.UnsafePointer()))
	runtime.KeepAlive(type_)
	return newownedQVariant(__r)
}

func QVariant_Compare(lhs *QVariant, rhs *QVariant) *QPartialOrdering {
	__r, _, _ := purego.SyscallN(addr_QVariant_compare, uintptr(lhs.UnsafePointer()), uintptr(rhs.UnsafePointer()))
	runtime.KeepAlive(lhs)
	runtime.KeepAlive(rhs)
	return newownedQPartialOrdering(__r)
}

func (this *QVariant) ToIntWithOk() (int, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QVariant_toInt2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return int(__r), ok_out != 0
}

func (this *QVariant) ToUIntWithOk() (uint, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QVariant_toUInt2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return uint(__r), ok_out != 0
}

func (this *QVariant) ToLongLongWithOk() (int64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QVariant_toLongLong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QVariant) ToULongLongWithOk() (uint64, bool) {
	var ok_out int32
	__r0, _, _ := purego.SyscallN(addr_QVariant_toULongLong2, this.h, uintptr(unsafe.Pointer(&ok_out)))
	__r := int64(__r0)
	runtime.KeepAlive(this)
	return uint64(__r), ok_out != 0
}

func (this *QVariant) ToDoubleWithOk() (float64, bool) {
	var ok_out int32
	__r := pg_QVariant_toDouble2(this.h, uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QVariant) ToFloatWithOk() (float32, bool) {
	var ok_out int32
	__r := pg_QVariant_toFloat2(this.h, uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QVariant) ToRealWithOk() (float64, bool) {
	var ok_out int32
	__r := pg_QVariant_toReal2(this.h, uintptr(unsafe.Pointer(&ok_out)))
	runtime.KeepAlive(this)
	return __r, ok_out != 0
}

func (this *QVariant) OperatorEqual(rhs *QVariant) bool {
	__r0, _, _ := purego.SyscallN(addr_QVariant_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QVariant) OperatorNotEqual(rhs *QVariant) bool {
	__r0, _, _ := purego.SyscallN(addr_QVariant_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QVariant) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QVariant_delete(this.h)
	}
	this.UnsafeRelease()
}
