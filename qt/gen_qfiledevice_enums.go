package qt

type QFileDevice__FileError int64

const (
	QFileDevice__NoError QFileDevice__FileError = 0
	QFileDevice__ReadError QFileDevice__FileError = 1
	QFileDevice__WriteError QFileDevice__FileError = 2
	QFileDevice__FatalError QFileDevice__FileError = 3
	QFileDevice__ResourceError QFileDevice__FileError = 4
	QFileDevice__OpenError QFileDevice__FileError = 5
	QFileDevice__AbortError QFileDevice__FileError = 6
	QFileDevice__TimeOutError QFileDevice__FileError = 7
	QFileDevice__UnspecifiedError QFileDevice__FileError = 8
	QFileDevice__RemoveError QFileDevice__FileError = 9
	QFileDevice__RenameError QFileDevice__FileError = 10
	QFileDevice__PositionError QFileDevice__FileError = 11
	QFileDevice__ResizeError QFileDevice__FileError = 12
	QFileDevice__PermissionsError QFileDevice__FileError = 13
	QFileDevice__CopyError QFileDevice__FileError = 14
)

type QFileDevice__FileTime int64

const (
	QFileDevice__FileAccessTime QFileDevice__FileTime = 0
	QFileDevice__FileBirthTime QFileDevice__FileTime = 1
	QFileDevice__FileMetadataChangeTime QFileDevice__FileTime = 2
	QFileDevice__FileModificationTime QFileDevice__FileTime = 3
)

type QFileDevice__Permission int64

const (
	QFileDevice__ReadOwner QFileDevice__Permission = 16384
	QFileDevice__WriteOwner QFileDevice__Permission = 8192
	QFileDevice__ExeOwner QFileDevice__Permission = 4096
	QFileDevice__ReadUser QFileDevice__Permission = 1024
	QFileDevice__WriteUser QFileDevice__Permission = 512
	QFileDevice__ExeUser QFileDevice__Permission = 256
	QFileDevice__ReadGroup QFileDevice__Permission = 64
	QFileDevice__WriteGroup QFileDevice__Permission = 32
	QFileDevice__ExeGroup QFileDevice__Permission = 16
	QFileDevice__ReadOther QFileDevice__Permission = 4
	QFileDevice__WriteOther QFileDevice__Permission = 2
	QFileDevice__ExeOther QFileDevice__Permission = 1
)

type QFileDevice__FileHandleFlag int64

const (
	QFileDevice__AutoCloseHandle QFileDevice__FileHandleFlag = 1
	QFileDevice__DontCloseHandle QFileDevice__FileHandleFlag = 0
)

type QFileDevice__MemoryMapFlag int64

const (
	QFileDevice__NoOptions QFileDevice__MemoryMapFlag = 0
	QFileDevice__MapPrivateOption QFileDevice__MemoryMapFlag = 1
)

