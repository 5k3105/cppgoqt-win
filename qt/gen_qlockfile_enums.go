package qt

type QLockFile__LockError int64

const (
	QLockFile__NoError QLockFile__LockError = 0
	QLockFile__LockFailedError QLockFile__LockError = 1
	QLockFile__PermissionError QLockFile__LockError = 2
	QLockFile__UnknownError QLockFile__LockError = 3
)

