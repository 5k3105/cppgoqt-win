package qt

import (
	"unsafe"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	addr_QCommandLineParser_new uintptr
	addr_QCommandLineParser_tr uintptr
	addr_QCommandLineParser_setSingleDashWordOptionMode uintptr
	addr_QCommandLineParser_setOptionsAfterPositionalArgumentsMode uintptr
	addr_QCommandLineParser_addOption uintptr
	addr_QCommandLineParser_addOptions uintptr
	addr_QCommandLineParser_addVersionOption uintptr
	addr_QCommandLineParser_addHelpOption uintptr
	addr_QCommandLineParser_setApplicationDescription uintptr
	addr_QCommandLineParser_applicationDescription uintptr
	addr_QCommandLineParser_addPositionalArgument uintptr
	addr_QCommandLineParser_clearPositionalArguments uintptr
	addr_QCommandLineParser_process uintptr
	addr_QCommandLineParser_process2 uintptr
	addr_QCommandLineParser_parse uintptr
	addr_QCommandLineParser_errorText uintptr
	addr_QCommandLineParser_isSet uintptr
	addr_QCommandLineParser_value uintptr
	addr_QCommandLineParser_values uintptr
	addr_QCommandLineParser_isSet2 uintptr
	addr_QCommandLineParser_value2 uintptr
	addr_QCommandLineParser_values2 uintptr
	addr_QCommandLineParser_positionalArguments uintptr
	addr_QCommandLineParser_optionNames uintptr
	addr_QCommandLineParser_unknownOptionNames uintptr
	addr_QCommandLineParser_showVersion uintptr
	addr_QCommandLineParser_showHelp uintptr
	addr_QCommandLineParser_helpText uintptr
	addr_QCommandLineParser_showMessageAndExit uintptr
	addr_QCommandLineParser_tr2 uintptr
	addr_QCommandLineParser_tr3 uintptr
	addr_QCommandLineParser_addPositionalArgument2 uintptr
	addr_QCommandLineParser_showHelp2 uintptr
	addr_QCommandLineParser_showMessageAndExit2 uintptr
	pg_QCommandLineParser_delete func(uintptr)
)

func init() {
	symaddr(&addr_QCommandLineParser_new, "QCommandLineParser_new")
	symaddr(&addr_QCommandLineParser_tr, "QCommandLineParser_tr")
	symaddr(&addr_QCommandLineParser_setSingleDashWordOptionMode, "QCommandLineParser_setSingleDashWordOptionMode")
	symaddr(&addr_QCommandLineParser_setOptionsAfterPositionalArgumentsMode, "QCommandLineParser_setOptionsAfterPositionalArgumentsMode")
	symaddr(&addr_QCommandLineParser_addOption, "QCommandLineParser_addOption")
	symaddr(&addr_QCommandLineParser_addOptions, "QCommandLineParser_addOptions")
	symaddr(&addr_QCommandLineParser_addVersionOption, "QCommandLineParser_addVersionOption")
	symaddr(&addr_QCommandLineParser_addHelpOption, "QCommandLineParser_addHelpOption")
	symaddr(&addr_QCommandLineParser_setApplicationDescription, "QCommandLineParser_setApplicationDescription")
	symaddr(&addr_QCommandLineParser_applicationDescription, "QCommandLineParser_applicationDescription")
	symaddr(&addr_QCommandLineParser_addPositionalArgument, "QCommandLineParser_addPositionalArgument")
	symaddr(&addr_QCommandLineParser_clearPositionalArguments, "QCommandLineParser_clearPositionalArguments")
	symaddr(&addr_QCommandLineParser_process, "QCommandLineParser_process")
	symaddr(&addr_QCommandLineParser_process2, "QCommandLineParser_process2")
	symaddr(&addr_QCommandLineParser_parse, "QCommandLineParser_parse")
	symaddr(&addr_QCommandLineParser_errorText, "QCommandLineParser_errorText")
	symaddr(&addr_QCommandLineParser_isSet, "QCommandLineParser_isSet")
	symaddr(&addr_QCommandLineParser_value, "QCommandLineParser_value")
	symaddr(&addr_QCommandLineParser_values, "QCommandLineParser_values")
	symaddr(&addr_QCommandLineParser_isSet2, "QCommandLineParser_isSet2")
	symaddr(&addr_QCommandLineParser_value2, "QCommandLineParser_value2")
	symaddr(&addr_QCommandLineParser_values2, "QCommandLineParser_values2")
	symaddr(&addr_QCommandLineParser_positionalArguments, "QCommandLineParser_positionalArguments")
	symaddr(&addr_QCommandLineParser_optionNames, "QCommandLineParser_optionNames")
	symaddr(&addr_QCommandLineParser_unknownOptionNames, "QCommandLineParser_unknownOptionNames")
	symaddr(&addr_QCommandLineParser_showVersion, "QCommandLineParser_showVersion")
	symaddr(&addr_QCommandLineParser_showHelp, "QCommandLineParser_showHelp")
	symaddr(&addr_QCommandLineParser_helpText, "QCommandLineParser_helpText")
	symaddr(&addr_QCommandLineParser_showMessageAndExit, "QCommandLineParser_showMessageAndExit")
	symaddr(&addr_QCommandLineParser_tr2, "QCommandLineParser_tr2")
	symaddr(&addr_QCommandLineParser_tr3, "QCommandLineParser_tr3")
	symaddr(&addr_QCommandLineParser_addPositionalArgument2, "QCommandLineParser_addPositionalArgument2")
	symaddr(&addr_QCommandLineParser_showHelp2, "QCommandLineParser_showHelp2")
	symaddr(&addr_QCommandLineParser_showMessageAndExit2, "QCommandLineParser_showMessageAndExit2")
	reg(&pg_QCommandLineParser_delete, "QCommandLineParser_delete")
}

type QCommandLineParser struct {
	h   uintptr
	vid int64 // virtual-override id (0 for non-subclass instances)
}

func (this *QCommandLineParser) UnsafeGoID() int64     { return this.vid }
func (this *QCommandLineParser) UnsafeSetGoID(id int64) { this.vid = id }

func bindQCommandLineParser(h uintptr) *QCommandLineParser {
	if h == 0 {
		return nil
	}
	return &QCommandLineParser{h: h}
}

func newownedQCommandLineParser(h uintptr) *QCommandLineParser {
	o := bindQCommandLineParser(h)
	if o != nil {
		runtime.SetFinalizer(o, (*QCommandLineParser).Delete)
	}
	return o
}

func (this *QCommandLineParser) chandle() uintptr {
	if this == nil {
		return 0
	}
	return this.h
}

func (this *QCommandLineParser) UnsafePointer() unsafe.Pointer { return unsafe.Pointer(this.chandle()) }
func UnsafeNewQCommandLineParser(h unsafe.Pointer) *QCommandLineParser { return bindQCommandLineParser(uintptr(h)) }
func UnsafeNewQCommandLineParserOwned(h unsafe.Pointer) *QCommandLineParser { return newownedQCommandLineParser(uintptr(h)) }

func (this *QCommandLineParser) UnsafeRelease() { this.h = 0 }

func NewQCommandLineParser() *QCommandLineParser {
	__r, _, _ := purego.SyscallN(addr_QCommandLineParser_new)
	__o := newownedQCommandLineParser(__r)
	return __o
}

func QCommandLineParser_Tr(sourceText string) string {
	sourceText_b := append([]byte(sourceText), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLineParser_tr, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	return pgStr(__r, __n)
}

func (this *QCommandLineParser) SetSingleDashWordOptionMode(parsingMode QCommandLineParser__SingleDashWordOptionMode) {
	purego.SyscallN(addr_QCommandLineParser_setSingleDashWordOptionMode, this.h, uintptr(parsingMode))
	runtime.KeepAlive(this)
}

func (this *QCommandLineParser) SetOptionsAfterPositionalArgumentsMode(mode QCommandLineParser__OptionsAfterPositionalArgumentsMode) {
	purego.SyscallN(addr_QCommandLineParser_setOptionsAfterPositionalArgumentsMode, this.h, uintptr(mode))
	runtime.KeepAlive(this)
}

func (this *QCommandLineParser) AddOption(commandLineOption *QCommandLineOption) bool {
	__r0, _, _ := purego.SyscallN(addr_QCommandLineParser_addOption, this.h, uintptr(commandLineOption.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(commandLineOption)
	return __r != 0
}

func (this *QCommandLineParser) AddOptions(options []*QCommandLineOption) bool {
	options_arr := make([]uintptr, len(options))
	for __i, __e := range options {
		options_arr[__i] = uintptr(__e.UnsafePointer())
	}
	var options_p uintptr
	if len(options_arr) > 0 {
		options_p = uintptr(unsafe.Pointer(&options_arr[0]))
	}
	__r0, _, _ := purego.SyscallN(addr_QCommandLineParser_addOptions, this.h, options_p, uintptr(len(options)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(options)
	runtime.KeepAlive(options_arr)
	return __r != 0
}

func (this *QCommandLineParser) AddVersionOption() *QCommandLineOption {
	__r, _, _ := purego.SyscallN(addr_QCommandLineParser_addVersionOption, this.h)
	runtime.KeepAlive(this)
	return newownedQCommandLineOption(__r)
}

func (this *QCommandLineParser) AddHelpOption() *QCommandLineOption {
	__r, _, _ := purego.SyscallN(addr_QCommandLineParser_addHelpOption, this.h)
	runtime.KeepAlive(this)
	return newownedQCommandLineOption(__r)
}

func (this *QCommandLineParser) SetApplicationDescription(description string) {
	var description_p uintptr
	if len(description) > 0 {
		description_p = uintptr(unsafe.Pointer(unsafe.StringData(description)))
	}
	purego.SyscallN(addr_QCommandLineParser_setApplicationDescription, this.h, description_p, uintptr(len(description)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(description)
}

func (this *QCommandLineParser) ApplicationDescription() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLineParser_applicationDescription, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCommandLineParser) AddPositionalArgument(name string, description string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var description_p uintptr
	if len(description) > 0 {
		description_p = uintptr(unsafe.Pointer(unsafe.StringData(description)))
	}
	purego.SyscallN(addr_QCommandLineParser_addPositionalArgument, this.h, name_p, uintptr(len(name)), description_p, uintptr(len(description)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	runtime.KeepAlive(description)
}

func (this *QCommandLineParser) ClearPositionalArguments() {
	purego.SyscallN(addr_QCommandLineParser_clearPositionalArguments, this.h)
	runtime.KeepAlive(this)
}

func (this *QCommandLineParser) Process(arguments []string) {
	arguments_arr := make([]uintptr, len(arguments))
	for __i, __e := range arguments {
		arguments_arr[__i] = pgCString(__e)
	}
	var arguments_p uintptr
	if len(arguments_arr) > 0 {
		arguments_p = uintptr(unsafe.Pointer(&arguments_arr[0]))
	}
	purego.SyscallN(addr_QCommandLineParser_process, this.h, arguments_p, uintptr(len(arguments)))
	runtime.KeepAlive(this)
	for __i := range arguments_arr {
		pgFree(arguments_arr[__i])
	}
	runtime.KeepAlive(arguments_arr)
}

func (this *QCommandLineParser) ProcessWithApp(app *QCoreApplication) {
	purego.SyscallN(addr_QCommandLineParser_process2, this.h, uintptr(app.UnsafePointer()))
	runtime.KeepAlive(this)
	runtime.KeepAlive(app)
}

func (this *QCommandLineParser) Parse(arguments []string) bool {
	arguments_arr := make([]uintptr, len(arguments))
	for __i, __e := range arguments {
		arguments_arr[__i] = pgCString(__e)
	}
	var arguments_p uintptr
	if len(arguments_arr) > 0 {
		arguments_p = uintptr(unsafe.Pointer(&arguments_arr[0]))
	}
	__r0, _, _ := purego.SyscallN(addr_QCommandLineParser_parse, this.h, arguments_p, uintptr(len(arguments)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	for __i := range arguments_arr {
		pgFree(arguments_arr[__i])
	}
	runtime.KeepAlive(arguments_arr)
	return __r != 0
}

func (this *QCommandLineParser) ErrorText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLineParser_errorText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func (this *QCommandLineParser) IsSet(name string) bool {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	__r0, _, _ := purego.SyscallN(addr_QCommandLineParser_isSet, this.h, name_p, uintptr(len(name)))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return __r != 0
}

func (this *QCommandLineParser) Value(name string) string {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLineParser_value, this.h, name_p, uintptr(len(name)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	return pgStr(__r, __n)
}

func (this *QCommandLineParser) Values(name string) []string {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLineParser_values, this.h, name_p, uintptr(len(name)), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QCommandLineParser) IsSetWithOption(option *QCommandLineOption) bool {
	__r0, _, _ := purego.SyscallN(addr_QCommandLineParser_isSet2, this.h, uintptr(option.UnsafePointer()))
	__r := int32(uint32(__r0))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	return __r != 0
}

func (this *QCommandLineParser) ValueWithOption(option *QCommandLineOption) string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLineParser_value2, this.h, uintptr(option.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	return pgStr(__r, __n)
}

func (this *QCommandLineParser) ValuesWithOption(option *QCommandLineOption) []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLineParser_values2, this.h, uintptr(option.UnsafePointer()), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(option)
	__out := make([]string, __n)
	__sl := unsafe.Slice((*uintptr)(unsafe.Pointer(__r)), int(__n))
	for __i := range __out {
		__out[__i] = pgGoStr(__sl[__i])
		pgFree(__sl[__i])
	}
	pgFree(__r)
	return __out
}

func (this *QCommandLineParser) PositionalArguments() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLineParser_positionalArguments, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QCommandLineParser) OptionNames() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLineParser_optionNames, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QCommandLineParser) UnknownOptionNames() []string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLineParser_unknownOptionNames, this.h, uintptr(unsafe.Pointer(&__n)))
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

func (this *QCommandLineParser) ShowVersion() {
	purego.SyscallN(addr_QCommandLineParser_showVersion, this.h)
	runtime.KeepAlive(this)
}

func (this *QCommandLineParser) ShowHelp() {
	purego.SyscallN(addr_QCommandLineParser_showHelp, this.h)
	runtime.KeepAlive(this)
}

func (this *QCommandLineParser) HelpText() string {
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLineParser_helpText, this.h, uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(this)
	return pgStr(__r, __n)
}

func QCommandLineParser_ShowMessageAndExit(type_ QCommandLineParser__MessageType, message string) {
	var message_p uintptr
	if len(message) > 0 {
		message_p = uintptr(unsafe.Pointer(unsafe.StringData(message)))
	}
	purego.SyscallN(addr_QCommandLineParser_showMessageAndExit, uintptr(type_), message_p, uintptr(len(message)))
	runtime.KeepAlive(message)
}

func QCommandLineParser_Tr2(sourceText string, disambiguation string) string {
	sourceText_b := append([]byte(sourceText), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLineParser_tr2, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func QCommandLineParser_Tr3(sourceText string, disambiguation string, n int) string {
	sourceText_b := append([]byte(sourceText), 0)
	disambiguation_b := append([]byte(disambiguation), 0)
	var __n int64
	__r, _, _ := purego.SyscallN(addr_QCommandLineParser_tr3, uintptr(unsafe.Pointer(&sourceText_b[0])), uintptr(unsafe.Pointer(&disambiguation_b[0])), uintptr(n), uintptr(unsafe.Pointer(&__n)))
	runtime.KeepAlive(sourceText_b)
	runtime.KeepAlive(disambiguation_b)
	return pgStr(__r, __n)
}

func (this *QCommandLineParser) AddPositionalArgument2(name string, description string, syntax string) {
	var name_p uintptr
	if len(name) > 0 {
		name_p = uintptr(unsafe.Pointer(unsafe.StringData(name)))
	}
	var description_p uintptr
	if len(description) > 0 {
		description_p = uintptr(unsafe.Pointer(unsafe.StringData(description)))
	}
	var syntax_p uintptr
	if len(syntax) > 0 {
		syntax_p = uintptr(unsafe.Pointer(unsafe.StringData(syntax)))
	}
	purego.SyscallN(addr_QCommandLineParser_addPositionalArgument2, this.h, name_p, uintptr(len(name)), description_p, uintptr(len(description)), syntax_p, uintptr(len(syntax)))
	runtime.KeepAlive(this)
	runtime.KeepAlive(name)
	runtime.KeepAlive(description)
	runtime.KeepAlive(syntax)
}

func (this *QCommandLineParser) ShowHelpWithExitCode(exitCode int) {
	purego.SyscallN(addr_QCommandLineParser_showHelp2, this.h, uintptr(exitCode))
	runtime.KeepAlive(this)
}

func QCommandLineParser_ShowMessageAndExit2(type_ QCommandLineParser__MessageType, message string, exitCode int) {
	var message_p uintptr
	if len(message) > 0 {
		message_p = uintptr(unsafe.Pointer(unsafe.StringData(message)))
	}
	purego.SyscallN(addr_QCommandLineParser_showMessageAndExit2, uintptr(type_), message_p, uintptr(len(message)), uintptr(exitCode))
	runtime.KeepAlive(message)
}

func (this *QCommandLineParser) Delete() {
	runtime.SetFinalizer(this, nil)
	if __id := this.UnsafeGoID(); __id != 0 {
		unregisterInstance(__id)
	}
	if this.h != 0 {
		pg_QCommandLineParser_delete(this.h)
	}
	this.UnsafeRelease()
}
