package qt

type QElapsedTimer__ClockType int64

const (
	QElapsedTimer__SystemTime QElapsedTimer__ClockType = 0
	QElapsedTimer__MonotonicClock QElapsedTimer__ClockType = 1
	QElapsedTimer__TickCounter QElapsedTimer__ClockType = 2
	QElapsedTimer__MachAbsoluteTime QElapsedTimer__ClockType = 3
	QElapsedTimer__PerformanceCounter QElapsedTimer__ClockType = 4
)

