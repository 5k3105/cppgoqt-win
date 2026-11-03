package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QUrlQuery_new uintptr
	addr_QUrlQuery_new2 uintptr
	addr_QUrlQuery_new3 uintptr
	addr_QUrlQuery_new4 uintptr
	addr_QUrlQuery_OperatorAssign uintptr
	addr_QUrlQuery_swap uintptr
	addr_QUrlQuery_isEmpty uintptr
	addr_QUrlQuery_isDetached uintptr
	addr_QUrlQuery_clear uintptr
	addr_QUrlQuery_query uintptr
	addr_QUrlQuery_setQuery uintptr
	addr_QUrlQuery_toString uintptr
	addr_QUrlQuery_setQueryDelimiters uintptr
	addr_QUrlQuery_queryValueDelimiter uintptr
	addr_QUrlQuery_queryPairDelimiter uintptr
	addr_QUrlQuery_hasQueryItem uintptr
	addr_QUrlQuery_addQueryItem uintptr
	addr_QUrlQuery_removeQueryItem uintptr
	addr_QUrlQuery_queryItemValue uintptr
	addr_QUrlQuery_allQueryItemValues uintptr
	addr_QUrlQuery_removeAllQueryItems uintptr
	addr_QUrlQuery_defaultQueryValueDelimiter uintptr
	addr_QUrlQuery_defaultQueryPairDelimiter uintptr
	addr_QUrlQuery_query2 uintptr
	addr_QUrlQuery_toString2 uintptr
	addr_QUrlQuery_queryItemValue2 uintptr
	addr_QUrlQuery_allQueryItemValues2 uintptr
	addr_QUrlQuery_OperatorEqual uintptr
	addr_QUrlQuery_OperatorNotEqual uintptr
	pg_QUrlQuery_delete func(uintptr)
)

func init() {
	symaddr(&addr_QUrlQuery_new, "QUrlQuery_new")
	symaddr(&addr_QUrlQuery_new2, "QUrlQuery_new2")
	symaddr(&addr_QUrlQuery_new3, "QUrlQuery_new3")
	symaddr(&addr_QUrlQuery_new4, "QUrlQuery_new4")
	symaddr(&addr_QUrlQuery_OperatorAssign, "QUrlQuery_OperatorAssign")
	symaddr(&addr_QUrlQuery_swap, "QUrlQuery_swap")
	symaddr(&addr_QUrlQuery_isEmpty, "QUrlQuery_isEmpty")
	symaddr(&addr_QUrlQuery_isDetached, "QUrlQuery_isDetached")
	symaddr(&addr_QUrlQuery_clear, "QUrlQuery_clear")
	symaddr(&addr_QUrlQuery_query, "QUrlQuery_query")
	symaddr(&addr_QUrlQuery_setQuery, "QUrlQuery_setQuery")
	symaddr(&addr_QUrlQuery_toString, "QUrlQuery_toString")
	symaddr(&addr_QUrlQuery_setQueryDelimiters, "QUrlQuery_setQueryDelimiters")
	symaddr(&addr_QUrlQuery_queryValueDelimiter, "QUrlQuery_queryValueDelimiter")
	symaddr(&addr_QUrlQuery_queryPairDelimiter, "QUrlQuery_queryPairDelimiter")
	symaddr(&addr_QUrlQuery_hasQueryItem, "QUrlQuery_hasQueryItem")
	symaddr(&addr_QUrlQuery_addQueryItem, "QUrlQuery_addQueryItem")
	symaddr(&addr_QUrlQuery_removeQueryItem, "QUrlQuery_removeQueryItem")
	symaddr(&addr_QUrlQuery_queryItemValue, "QUrlQuery_queryItemValue")
	symaddr(&addr_QUrlQuery_allQueryItemValues, "QUrlQuery_allQueryItemValues")
	symaddr(&addr_QUrlQuery_removeAllQueryItems, "QUrlQuery_removeAllQueryItems")
	symaddr(&addr_QUrlQuery_defaultQueryValueDelimiter, "QUrlQuery_defaultQueryValueDelimiter")
	symaddr(&addr_QUrlQuery_defaultQueryPairDelimiter, "QUrlQuery_defaultQueryPairDelimiter")
	symaddr(&addr_QUrlQuery_query2, "QUrlQuery_query2")
	symaddr(&addr_QUrlQuery_toString2, "QUrlQuery_toString2")
	symaddr(&addr_QUrlQuery_queryItemValue2, "QUrlQuery_queryItemValue2")
	symaddr(&addr_QUrlQuery_allQueryItemValues2, "QUrlQuery_allQueryItemValues2")
	symaddr(&addr_QUrlQuery_OperatorEqual, "QUrlQuery_OperatorEqual")
	symaddr(&addr_QUrlQuery_OperatorNotEqual, "QUrlQuery_OperatorNotEqual")
	reg(&pg_QUrlQuery_delete, "QUrlQuery_delete")
}

type QUrlQuery struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QUrlQuery) UnsafeGoID() int64     { return this.vid }
func (this *QUrlQuery) UnsafeSetGoID(id int64) { this.vid = id }

func bindQUrlQuery(h uintptr) *QUrlQuery {
	if h == 0 {
		return nil
	}
	return &QUrlQuery{h: h}
}

func newownedQUrlQuery(h uintptr) *QUrlQuery {
	o := bindQUrlQuery(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QUrlQuery).Delete)
	}
	return o
}

func (this *QUrlQuery) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QUrlQuery) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQUrlQuery(h unsafe.Pointer) *QUrlQuery { return bindQUrlQuery(uintptr(h)) }
func UnsafeNewQUrlQueryOwned(h unsafe.Pointer) *QUrlQuery { return newownedQUrlQuery(uintptr(h)) }

func (this *QUrlQuery) UnsafeRelease() { this.h = 0 }

func NewQUrlQuery() *QUrlQuery {
	__r, _, _ := purego.SyscallN(addr_QUrlQuery_new)
	__o := newownedQUrlQuery(__r)
	return __o
}

func NewQUrlQuery2(url *QUrl) *QUrlQuery {
	__r, _, _ := purego.SyscallN(addr_QUrlQuery_new2, uintptr(url.UnsafePointer()))
	__o := newownedQUrlQuery(__r)
	runtime.KeepAlive(url)
	return __o
}

func NewQUrlQuery3(queryString string) *QUrlQuery {
	var queryString_p uintptr
	if len(queryString) > 0 {
		queryString_p = uintptr(unsafe.Pointer(unsafe.StringData(queryString)))
	}
	__r, _, _ := purego.SyscallN(addr_QUrlQuery_new3, queryString_p, uintptr(len(queryString)))
	__o := newownedQUrlQuery(__r)
	runtime.KeepAlive(queryString)
	return __o
}

func NewQUrlQuery4(other *QUrlQuery) *QUrlQuery {
	__r, _, _ := purego.SyscallN(addr_QUrlQuery_new4, uintptr(other.UnsafePointer()))
	__o := newownedQUrlQuery(__r)
	runtime.KeepAlive(other)
	return __o
}

func (this *QUrlQuery) OperatorAssign(other *QUrlQuery) *QUrlQuery {
	__r, _, _ := purego.SyscallN(addr_QUrlQuery_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQUrlQuery(__r)
}

func (this *QUrlQuery) Swap(other *QUrlQuery) {
	purego.SyscallN(addr_QUrlQuery_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QUrlQuery) IsEmpty() bool {
	__r0, _, _ := purego.SyscallN(addr_QUrlQuery_isEmpty, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUrlQuery) IsDetached() bool {
	__r0, _, _ := purego.SyscallN(addr_QUrlQuery_isDetached, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return __r != 0
}

func (this *QUrlQuery) Clear() {
	purego.SyscallN(addr_QUrlQuery_clear, this.h)
	runtime.KeepAlive(this)
}

func (this *QUrlQuery) Query() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrlQuery_query, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrlQuery) SetQuery(queryString string) {
	var queryString_p uintptr
	if len(queryString) > 0 {
		queryString_p = uintptr(unsafe.Pointer(unsafe.StringData(queryString)))
	}
	purego.SyscallN(addr_QUrlQuery_setQuery, this.h, queryString_p, uintptr(len(queryString)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(queryString)
}

func (this *QUrlQuery) ToString() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrlQuery_toString, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrlQuery) SetQueryDelimiters(valueDelimiter *QChar, pairDelimiter *QChar) {
	purego.SyscallN(addr_QUrlQuery_setQueryDelimiters, this.h, uintptr(valueDelimiter.UnsafePointer()), uintptr(pairDelimiter.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(valueDelimiter)
	runtime.KeepAlive(pairDelimiter)
}

func (this *QUrlQuery) QueryValueDelimiter() *QChar {
	__r, _, _ := purego.SyscallN(addr_QUrlQuery_queryValueDelimiter, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QUrlQuery) QueryPairDelimiter() *QChar {
	__r, _, _ := purego.SyscallN(addr_QUrlQuery_queryPairDelimiter, this.h)
	runtime.KeepAlive(this)
	return newownedQChar(__r)
}

func (this *QUrlQuery) HasQueryItem(key string) bool {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	__r0, _, _ := purego.SyscallN(addr_QUrlQuery_hasQueryItem, this.h, key_p, uintptr(len(key)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return __r != 0
}

func (this *QUrlQuery) AddQueryItem(key string, value string) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	var value_p uintptr
	if len(value) > 0 {
		value_p = uintptr(unsafe.Pointer(unsafe.StringData(value)))
	}
	purego.SyscallN(addr_QUrlQuery_addQueryItem, this.h, key_p, uintptr(len(key)), value_p, uintptr(len(value)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	runtime.KeepAlive(value)
}

func (this *QUrlQuery) RemoveQueryItem(key string) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	purego.SyscallN(addr_QUrlQuery_removeQueryItem, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func (this *QUrlQuery) QueryItemValue(key string) string {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrlQuery_queryItemValue, this.h, key_p, uintptr(len(key)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return pgStr(__r, __n)
}

func (this *QUrlQuery) AllQueryItemValues(key string) []string {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrlQuery_allQueryItemValues, this.h, key_p, uintptr(len(key)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QUrlQuery) RemoveAllQueryItems(key string) {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	purego.SyscallN(addr_QUrlQuery_removeAllQueryItems, this.h, key_p, uintptr(len(key)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
}

func QUrlQuery_DefaultQueryValueDelimiter() uint16 {
	__r0, _, _ := purego.SyscallN(addr_QUrlQuery_defaultQueryValueDelimiter)
	__r := int16(uint16(__r0))
	return uint16(__r)
}

func QUrlQuery_DefaultQueryPairDelimiter() uint16 {
	__r0, _, _ := purego.SyscallN(addr_QUrlQuery_defaultQueryPairDelimiter)
	__r := int16(uint16(__r0))
	return uint16(__r)
}

func (this *QUrlQuery) QueryWithEncoding(encoding QUrl__ComponentFormattingOption) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrlQuery_query2, this.h, uintptr(encoding), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrlQuery) ToStringWithEncoding(encoding QUrl__ComponentFormattingOption) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrlQuery_toString2, this.h, uintptr(encoding), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QUrlQuery) QueryItemValue2(key string, encoding QUrl__ComponentFormattingOption) string {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrlQuery_queryItemValue2, this.h, key_p, uintptr(len(key)), uintptr(encoding), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	return pgStr(__r, __n)
}

func (this *QUrlQuery) AllQueryItemValues2(key string, encoding QUrl__ComponentFormattingOption) []string {
	var key_p uintptr
	if len(key) > 0 {
		key_p = uintptr(unsafe.Pointer(unsafe.StringData(key)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QUrlQuery_allQueryItemValues2, this.h, key_p, uintptr(len(key)), uintptr(encoding), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(key)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QUrlQuery) OperatorEqual(rhs *QUrlQuery) bool {
	__r0, _, _ := purego.SyscallN(addr_QUrlQuery_OperatorEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QUrlQuery) OperatorNotEqual(rhs *QUrlQuery) bool {
	__r0, _, _ := purego.SyscallN(addr_QUrlQuery_OperatorNotEqual, this.h, uintptr(rhs.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(rhs)
	return __r != 0
}

func (this *QUrlQuery) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QUrlQuery_delete(this.h)
	}
	this.UnsafeRelease()
}
