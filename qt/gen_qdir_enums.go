package qt

type QDir__Filter int64

const (
	QDir__Dirs QDir__Filter = 1
	QDir__Files QDir__Filter = 2
	QDir__Drives QDir__Filter = 4
	QDir__NoSymLinks QDir__Filter = 8
	QDir__AllEntries QDir__Filter = 7
	QDir__TypeMask QDir__Filter = 15
	QDir__Readable QDir__Filter = 16
	QDir__Writable QDir__Filter = 32
	QDir__Executable QDir__Filter = 64
	QDir__PermissionMask QDir__Filter = 112
	QDir__Modified QDir__Filter = 128
	QDir__Hidden QDir__Filter = 256
	QDir__System QDir__Filter = 512
	QDir__AccessMask QDir__Filter = 1008
	QDir__AllDirs QDir__Filter = 1024
	QDir__CaseSensitive QDir__Filter = 2048
	QDir__NoDot QDir__Filter = 8192
	QDir__NoDotDot QDir__Filter = 16384
	QDir__NoDotAndDotDot QDir__Filter = 24576
	QDir__NoFilter QDir__Filter = -1
)

type QDir__SortFlag int64

const (
	QDir__Name QDir__SortFlag = 0
	QDir__Time QDir__SortFlag = 1
	QDir__Size QDir__SortFlag = 2
	QDir__Unsorted QDir__SortFlag = 3
	QDir__SortByMask QDir__SortFlag = 3
	QDir__DirsFirst QDir__SortFlag = 4
	QDir__Reversed QDir__SortFlag = 8
	QDir__IgnoreCase QDir__SortFlag = 16
	QDir__DirsLast QDir__SortFlag = 32
	QDir__LocaleAware QDir__SortFlag = 64
	QDir__Type QDir__SortFlag = 128
	QDir__NoSort QDir__SortFlag = -1
)

