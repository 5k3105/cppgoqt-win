package qt

type QLibrary__LoadHint int64

const (
	QLibrary__ResolveAllSymbolsHint QLibrary__LoadHint = 1
	QLibrary__ExportExternalSymbolsHint QLibrary__LoadHint = 2
	QLibrary__LoadArchiveMemberHint QLibrary__LoadHint = 4
	QLibrary__PreventUnloadHint QLibrary__LoadHint = 8
	QLibrary__DeepBindHint QLibrary__LoadHint = 16
)

type QLibrary__LoadStatusTag int64

const (
	QLibrary__NotLoaded QLibrary__LoadStatusTag = 0
	QLibrary__Loaded QLibrary__LoadStatusTag = 1
)

