package qt

type QAbstractAnimation__Direction int64

const (
	QAbstractAnimation__Forward QAbstractAnimation__Direction = 0
	QAbstractAnimation__Backward QAbstractAnimation__Direction = 1
)

type QAbstractAnimation__State int64

const (
	QAbstractAnimation__Stopped QAbstractAnimation__State = 0
	QAbstractAnimation__Paused QAbstractAnimation__State = 1
	QAbstractAnimation__Running QAbstractAnimation__State = 2
)

type QAbstractAnimation__DeletionPolicy int64

const (
	QAbstractAnimation__KeepWhenStopped QAbstractAnimation__DeletionPolicy = 0
	QAbstractAnimation__DeleteWhenStopped QAbstractAnimation__DeletionPolicy = 1
)

