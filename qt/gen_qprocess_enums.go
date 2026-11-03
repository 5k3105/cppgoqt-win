package qt

type QProcess__ProcessError int64

const (
	QProcess__FailedToStart QProcess__ProcessError = 0
	QProcess__Crashed QProcess__ProcessError = 1
	QProcess__Timedout QProcess__ProcessError = 2
	QProcess__ReadError QProcess__ProcessError = 3
	QProcess__WriteError QProcess__ProcessError = 4
	QProcess__UnknownError QProcess__ProcessError = 5
)

type QProcess__ProcessState int64

const (
	QProcess__NotRunning QProcess__ProcessState = 0
	QProcess__Starting QProcess__ProcessState = 1
	QProcess__Running QProcess__ProcessState = 2
)

type QProcess__ProcessChannel int64

const (
	QProcess__StandardOutput QProcess__ProcessChannel = 0
	QProcess__StandardError QProcess__ProcessChannel = 1
)

type QProcess__ProcessChannelMode int64

const (
	QProcess__SeparateChannels QProcess__ProcessChannelMode = 0
	QProcess__MergedChannels QProcess__ProcessChannelMode = 1
	QProcess__ForwardedChannels QProcess__ProcessChannelMode = 2
	QProcess__ForwardedOutputChannel QProcess__ProcessChannelMode = 3
	QProcess__ForwardedErrorChannel QProcess__ProcessChannelMode = 4
)

type QProcess__InputChannelMode int64

const (
	QProcess__ManagedInputChannel QProcess__InputChannelMode = 0
	QProcess__ForwardedInputChannel QProcess__InputChannelMode = 1
)

type QProcess__ExitStatus int64

const (
	QProcess__NormalExit QProcess__ExitStatus = 0
	QProcess__CrashExit QProcess__ExitStatus = 1
)

