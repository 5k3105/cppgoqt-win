package qt

type QEventLoop__ProcessEventsFlag int64

const (
	QEventLoop__AllEvents QEventLoop__ProcessEventsFlag = 0
	QEventLoop__ExcludeUserInputEvents QEventLoop__ProcessEventsFlag = 1
	QEventLoop__ExcludeSocketNotifiers QEventLoop__ProcessEventsFlag = 2
	QEventLoop__WaitForMoreEvents QEventLoop__ProcessEventsFlag = 4
	QEventLoop__X11ExcludeTimers QEventLoop__ProcessEventsFlag = 8
	QEventLoop__EventLoopExec QEventLoop__ProcessEventsFlag = 32
	QEventLoop__DialogExec QEventLoop__ProcessEventsFlag = 64
	QEventLoop__ApplicationExec QEventLoop__ProcessEventsFlag = 128
)

