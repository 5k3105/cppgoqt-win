package qt

type QSharedMemory__AccessMode int64

const (
	QSharedMemory__ReadOnly QSharedMemory__AccessMode = 0
	QSharedMemory__ReadWrite QSharedMemory__AccessMode = 1
)

type QSharedMemory__SharedMemoryError int64

const (
	QSharedMemory__NoError QSharedMemory__SharedMemoryError = 0
	QSharedMemory__PermissionDenied QSharedMemory__SharedMemoryError = 1
	QSharedMemory__InvalidSize QSharedMemory__SharedMemoryError = 2
	QSharedMemory__KeyError QSharedMemory__SharedMemoryError = 3
	QSharedMemory__AlreadyExists QSharedMemory__SharedMemoryError = 4
	QSharedMemory__NotFound QSharedMemory__SharedMemoryError = 5
	QSharedMemory__LockError QSharedMemory__SharedMemoryError = 6
	QSharedMemory__OutOfResources QSharedMemory__SharedMemoryError = 7
	QSharedMemory__UnknownError QSharedMemory__SharedMemoryError = 8
)

