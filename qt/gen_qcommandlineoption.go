package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCommandLineOption_new uintptr
	addr_QCommandLineOption_new2 uintptr
	addr_QCommandLineOption_new3 uintptr
	addr_QCommandLineOption_new4 uintptr
	addr_QCommandLineOption_new5 uintptr
	addr_QCommandLineOption_new6 uintptr
	addr_QCommandLineOption_new7 uintptr
	addr_QCommandLineOption_new8 uintptr
	addr_QCommandLineOption_new9 uintptr
	addr_QCommandLineOption_OperatorAssign uintptr
	addr_QCommandLineOption_swap uintptr
	addr_QCommandLineOption_names uintptr
	addr_QCommandLineOption_setValueName uintptr
	addr_QCommandLineOption_valueName uintptr
	addr_QCommandLineOption_setDescription uintptr
	addr_QCommandLineOption_description uintptr
	addr_QCommandLineOption_setDefaultValue uintptr
	addr_QCommandLineOption_setDefaultValues uintptr
	addr_QCommandLineOption_defaultValues uintptr
	addr_QCommandLineOption_flags uintptr
	addr_QCommandLineOption_setFlags uintptr
	pg_QCommandLineOption_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCommandLineOption_new, "QCommandLineOption_new")
	symaddr(&addr_QCommandLineOption_new2, "QCommandLineOption_new2")
	symaddr(&addr_QCommandLineOption_new3, "QCommandLineOption_new3")
	symaddr(&addr_QCommandLineOption_new4, "QCommandLineOption_new4")
	symaddr(&addr_QCommandLineOption_new5, "QCommandLineOption_new5")
	symaddr(&addr_QCommandLineOption_new6, "QCommandLineOption_new6")
	symaddr(&addr_QCommandLineOption_new7, "QCommandLineOption_new7")
	symaddr(&addr_QCommandLineOption_new8, "QCommandLineOption_new8")
	symaddr(&addr_QCommandLineOption_new9, "QCommandLineOption_new9")
	symaddr(&addr_QCommandLineOption_OperatorAssign, "QCommandLineOption_OperatorAssign")
	symaddr(&addr_QCommandLineOption_swap, "QCommandLineOption_swap")
	symaddr(&addr_QCommandLineOption_names, "QCommandLineOption_names")
	symaddr(&addr_QCommandLineOption_setValueName, "QCommandLineOption_setValueName")
	symaddr(&addr_QCommandLineOption_valueName, "QCommandLineOption_valueName")
	symaddr(&addr_QCommandLineOption_setDescription, "QCommandLineOption_setDescription")
	symaddr(&addr_QCommandLineOption_description, "QCommandLineOption_description")
	symaddr(&addr_QCommandLineOption_setDefaultValue, "QCommandLineOption_setDefaultValue")
	symaddr(&addr_QCommandLineOption_setDefaultValues, "QCommandLineOption_setDefaultValues")
	symaddr(&addr_QCommandLineOption_defaultValues, "QCommandLineOption_defaultValues")
	symaddr(&addr_QCommandLineOption_flags, "QCommandLineOption_flags")
	symaddr(&addr_QCommandLineOption_setFlags, "QCommandLineOption_setFlags")
	reg(&pg_QCommandLineOption_delete, "QCommandLineOption_delete")
}

type QCommandLineOption struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCommandLineOption) UnsafeGoID() int64     { return this.vid }
func (this *QCommandLineOption) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCommandLineOption(h uintptr) *QCommandLineOption {
	if h == 0 {
		return nil
	}
	return &QCommandLineOption{h: h}
}

func newownedQCommandLineOption(h uintptr) *QCommandLineOption {
	o := bindQCommandLineOption(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCommandLineOption).Delete)
	}
	return o
}

func (this *QCommandLineOption) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCommandLineOption) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCommandLineOption(h unsafe.Pointer) *QCommandLineOption { return bindQCommandLineOption(uintptr(h)) }
func UnsafeNewQCommandLineOptionOwned(h unsafe.Pointer) *QCommandLineOption { return newownedQCommandLineOption(uintptr(h)) }

func (this *QCommandLineOption) UnsafeRelease() { this.h = 0 }

func NewQCommandLineOption(name string) *QCommandLineOption {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r, _, _ := purego.SyscallN(addr_QCommandLineOption_new, name_p, uintptr(len(name)))
	__o := newownedQCommandLineOption(__r)
	runtime.KeepAlive(name)
	return __o
}

func NewQCommandLineOption2(names []string) *QCommandLineOption {
	names_arr := make([]uintptr, len(names))
	for __i, __e := range names {
		names_arr[__i] = pgCString(__e)
	}
	var names_p uintptr
	if len(names_arr) > 0 {
		names_p = uintptr(unsafe.Pointer(&names_arr[0]))
	}
	__r, _, _ := purego.SyscallN(addr_QCommandLineOption_new2, names_p, uintptr(len(names)))
	__o := newownedQCommandLineOption(__r)
	for __i := range names_arr {
		pgFree(names_arr[__i])
	}
	runtime.KeepAlive(names_arr)
	return __o
}

func NewQCommandLineOption3(name string, description string) *QCommandLineOption {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var description_p uintptr
	if len(description) > 0 {
		description_p = uintptr(unsafe.Pointer(unsafe.StringData(description)))
	}
	__r, _, _ := purego.SyscallN(addr_QCommandLineOption_new3, name_p, uintptr(len(name)), description_p, uintptr(len(description)))
	__o := newownedQCommandLineOption(__r)
	runtime.KeepAlive(name)
	runtime.KeepAlive(description)
	return __o
}

func NewQCommandLineOption4(names []string, description string) *QCommandLineOption {
	names_arr := make([]uintptr, len(names))
	for __i, __e := range names {
		names_arr[__i] = pgCString(__e)
	}
	var names_p uintptr
	if len(names_arr) > 0 {
		names_p = uintptr(unsafe.Pointer(&names_arr[0]))
	}
	var description_p uintptr
	if len(description) > 0 {
		description_p = uintptr(unsafe.Pointer(unsafe.StringData(description)))
	}
	__r, _, _ := purego.SyscallN(addr_QCommandLineOption_new4, names_p, uintptr(len(names)), description_p, uintptr(len(description)))
	__o := newownedQCommandLineOption(__r)
	for __i := range names_arr {
		pgFree(names_arr[__i])
	}
	runtime.KeepAlive(names_arr)
	runtime.KeepAlive(description)
	return __o
}

func NewQCommandLineOption5(other *QCommandLineOption) *QCommandLineOption {
	__r, _, _ := purego.SyscallN(addr_QCommandLineOption_new5, uintptr(other.UnsafePointer()))
	__o := newownedQCommandLineOption(__r)
	runtime.KeepAlive(other)
	return __o
}

func NewQCommandLineOption6(name string, description string, valueName string) *QCommandLineOption {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var description_p uintptr
	if len(description) > 0 {
		description_p = uintptr(unsafe.Pointer(unsafe.StringData(description)))
	}
	var valueName_p uintptr
	if len(valueName) > 0 {
		valueName_p = uintptr(unsafe.Pointer(unsafe.StringData(valueName)))
	}
	__r, _, _ := purego.SyscallN(addr_QCommandLineOption_new6, name_p, uintptr(len(name)), description_p, uintptr(len(description)), valueName_p, uintptr(len(valueName)))
	__o := newownedQCommandLineOption(__r)
	runtime.KeepAlive(name)
	runtime.KeepAlive(description)
	runtime.KeepAlive(valueName)
	return __o
}

func NewQCommandLineOption7(name string, description string, valueName string, defaultValue string) *QCommandLineOption {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var description_p uintptr
	if len(description) > 0 {
		description_p = uintptr(unsafe.Pointer(unsafe.StringData(description)))
	}
	var valueName_p uintptr
	if len(valueName) > 0 {
		valueName_p = uintptr(unsafe.Pointer(unsafe.StringData(valueName)))
	}
	var defaultValue_p uintptr
	if len(defaultValue) > 0 {
		defaultValue_p = uintptr(unsafe.Pointer(unsafe.StringData(defaultValue)))
	}
	__r, _, _ := purego.SyscallN(addr_QCommandLineOption_new7, name_p, uintptr(len(name)), description_p, uintptr(len(description)), valueName_p, uintptr(len(valueName)), defaultValue_p, uintptr(len(defaultValue)))
	__o := newownedQCommandLineOption(__r)
	runtime.KeepAlive(name)
	runtime.KeepAlive(description)
	runtime.KeepAlive(valueName)
	runtime.KeepAlive(defaultValue)
	return __o
}

func NewQCommandLineOption8(names []string, description string, valueName string) *QCommandLineOption {
	names_arr := make([]uintptr, len(names))
	for __i, __e := range names {
		names_arr[__i] = pgCString(__e)
	}
	var names_p uintptr
	if len(names_arr) > 0 {
		names_p = uintptr(unsafe.Pointer(&names_arr[0]))
	}
	var description_p uintptr
	if len(description) > 0 {
		description_p = uintptr(unsafe.Pointer(unsafe.StringData(description)))
	}
	var valueName_p uintptr
	if len(valueName) > 0 {
		valueName_p = uintptr(unsafe.Pointer(unsafe.StringData(valueName)))
	}
	__r, _, _ := purego.SyscallN(addr_QCommandLineOption_new8, names_p, uintptr(len(names)), description_p, uintptr(len(description)), valueName_p, uintptr(len(valueName)))
	__o := newownedQCommandLineOption(__r)
	for __i := range names_arr {
		pgFree(names_arr[__i])
	}
	runtime.KeepAlive(names_arr)
	runtime.KeepAlive(description)
	runtime.KeepAlive(valueName)
	return __o
}

func NewQCommandLineOption9(names []string, description string, valueName string, defaultValue string) *QCommandLineOption {
	names_arr := make([]uintptr, len(names))
	for __i, __e := range names {
		names_arr[__i] = pgCString(__e)
	}
	var names_p uintptr
	if len(names_arr) > 0 {
		names_p = uintptr(unsafe.Pointer(&names_arr[0]))
	}
	var description_p uintptr
	if len(description) > 0 {
		description_p = uintptr(unsafe.Pointer(unsafe.StringData(description)))
	}
	var valueName_p uintptr
	if len(valueName) > 0 {
		valueName_p = uintptr(unsafe.Pointer(unsafe.StringData(valueName)))
	}
	var defaultValue_p uintptr
	if len(defaultValue) > 0 {
		defaultValue_p = uintptr(unsafe.Pointer(unsafe.StringData(defaultValue)))
	}
	__r, _, _ := purego.SyscallN(addr_QCommandLineOption_new9, names_p, uintptr(len(names)), description_p, uintptr(len(description)), valueName_p, uintptr(len(valueName)), defaultValue_p, uintptr(len(defaultValue)))
	__o := newownedQCommandLineOption(__r)
	for __i := range names_arr {
		pgFree(names_arr[__i])
	}
	runtime.KeepAlive(names_arr)
	runtime.KeepAlive(description)
	runtime.KeepAlive(valueName)
	runtime.KeepAlive(defaultValue)
	return __o
}

func (this *QCommandLineOption) OperatorAssign(other *QCommandLineOption) *QCommandLineOption {
	__r, _, _ := purego.SyscallN(addr_QCommandLineOption_OperatorAssign, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
	return bindQCommandLineOption(__r)
}

func (this *QCommandLineOption) Swap(other *QCommandLineOption) {
	purego.SyscallN(addr_QCommandLineOption_swap, this.h, uintptr(other.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(other)
}

func (this *QCommandLineOption) Names() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLineOption_names, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QCommandLineOption) SetValueName(name string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	purego.SyscallN(addr_QCommandLineOption_setValueName, this.h, name_p, uintptr(len(name)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
}

func (this *QCommandLineOption) ValueName() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLineOption_valueName, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCommandLineOption) SetDescription(description string) {
	var description_p uintptr
	if len(description) > 0 {
		description_p = uintptr(unsafe.Pointer(unsafe.StringData(description)))
	}
	purego.SyscallN(addr_QCommandLineOption_setDescription, this.h, description_p, uintptr(len(description)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(description)
}

func (this *QCommandLineOption) Description() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLineOption_description, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCommandLineOption) SetDefaultValue(defaultValue string) {
	var defaultValue_p uintptr
	if len(defaultValue) > 0 {
		defaultValue_p = uintptr(unsafe.Pointer(unsafe.StringData(defaultValue)))
	}
	purego.SyscallN(addr_QCommandLineOption_setDefaultValue, this.h, defaultValue_p, uintptr(len(defaultValue)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(defaultValue)
}

func (this *QCommandLineOption) SetDefaultValues(defaultValues []string) {
	defaultValues_arr := make([]uintptr, len(defaultValues))
	for __i, __e := range defaultValues {
		defaultValues_arr[__i] = pgCString(__e)
	}
	var defaultValues_p uintptr
	if len(defaultValues_arr) > 0 {
		defaultValues_p = uintptr(unsafe.Pointer(&defaultValues_arr[0]))
	}
	purego.SyscallN(addr_QCommandLineOption_setDefaultValues, this.h, defaultValues_p, uintptr(len(defaultValues)))
	runtime.KeepAlive(this)
	for __i := range defaultValues_arr {
		pgFree(defaultValues_arr[__i])
	}
	runtime.KeepAlive(defaultValues_arr)
}

func (this *QCommandLineOption) DefaultValues() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLineOption_defaultValues, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QCommandLineOption) Flags() QCommandLineOption__Flag {
	__r0, _, _ := purego.SyscallN(addr_QCommandLineOption_flags, this.h)
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	return QCommandLineOption__Flag(__r)
}

func (this *QCommandLineOption) SetFlags(aflags QCommandLineOption__Flag) {
	purego.SyscallN(addr_QCommandLineOption_setFlags, this.h, uintptr(aflags))
	runtime.KeepAlive(this)
}

func (this *QCommandLineOption) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCommandLineOption_delete(this.h)
	}
	this.UnsafeRelease()
}
