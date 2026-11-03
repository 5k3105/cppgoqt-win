package qt

type QSystemSemaphore__AccessMode int64

const (
	QSystemSemaphore__Open QSystemSemaphore__AccessMode = 0
	QSystemSemaphore__Create QSystemSemaphore__AccessMode = 1
)

type QSystemSemaphore__SystemSemaphoreError int64

const (
	QSystemSemaphore__NoError QSystemSemaphore__SystemSemaphoreError = 0
	QSystemSemaphore__PermissionDenied QSystemSemaphore__SystemSemaphoreError = 1
	QSystemSemaphore__KeyError QSystemSemaphore__SystemSemaphoreError = 2
	QSystemSemaphore__AlreadyExists QSystemSemaphore__SystemSemaphoreError = 3
	QSystemSemaphore__NotFound QSystemSemaphore__SystemSemaphoreError = 4
	QSystemSemaphore__OutOfResources QSystemSemaphore__SystemSemaphoreError = 5
	QSystemSemaphore__UnknownError QSystemSemaphore__SystemSemaphoreError = 6
)

