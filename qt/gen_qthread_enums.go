package qt

type QThread__Priority int64

const (
	QThread__IdlePriority QThread__Priority = 0
	QThread__LowestPriority QThread__Priority = 1
	QThread__LowPriority QThread__Priority = 2
	QThread__NormalPriority QThread__Priority = 3
	QThread__HighPriority QThread__Priority = 4
	QThread__HighestPriority QThread__Priority = 5
	QThread__TimeCriticalPriority QThread__Priority = 6
	QThread__InheritPriority QThread__Priority = 7
)

type QThread__QualityOfService int64

const (
	QThread__QualityOfService__Auto QThread__QualityOfService = 0
	QThread__QualityOfService__High QThread__QualityOfService = 1
	QThread__QualityOfService__Eco QThread__QualityOfService = 2
)

