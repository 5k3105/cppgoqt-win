package qt

type QMetaMethod__Access int64

const (
	QMetaMethod__Private QMetaMethod__Access = 0
	QMetaMethod__Protected QMetaMethod__Access = 1
	QMetaMethod__Public QMetaMethod__Access = 2
)

type QMetaMethod__MethodType int64

const (
	QMetaMethod__Method QMetaMethod__MethodType = 0
	QMetaMethod__Signal QMetaMethod__MethodType = 1
	QMetaMethod__Slot QMetaMethod__MethodType = 2
	QMetaMethod__Constructor QMetaMethod__MethodType = 3
)

type QMetaMethod__Attributes int64

const (
	QMetaMethod__Compatibility QMetaMethod__Attributes = 1
	QMetaMethod__Cloned QMetaMethod__Attributes = 2
	QMetaMethod__Scriptable QMetaMethod__Attributes = 4
)

