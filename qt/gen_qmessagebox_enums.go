package qt

type QMessageBox__Option int64

const (
	QMessageBox__Option__DontUseNativeDialog QMessageBox__Option = 1
)

type QMessageBox__Icon int64

const (
	QMessageBox__NoIcon QMessageBox__Icon = 0
	QMessageBox__Information QMessageBox__Icon = 1
	QMessageBox__Warning QMessageBox__Icon = 2
	QMessageBox__Critical QMessageBox__Icon = 3
	QMessageBox__Question QMessageBox__Icon = 4
)

type QMessageBox__ButtonRole int64

const (
	QMessageBox__InvalidRole QMessageBox__ButtonRole = -1
	QMessageBox__AcceptRole QMessageBox__ButtonRole = 0
	QMessageBox__RejectRole QMessageBox__ButtonRole = 1
	QMessageBox__DestructiveRole QMessageBox__ButtonRole = 2
	QMessageBox__ActionRole QMessageBox__ButtonRole = 3
	QMessageBox__HelpRole QMessageBox__ButtonRole = 4
	QMessageBox__YesRole QMessageBox__ButtonRole = 5
	QMessageBox__NoRole QMessageBox__ButtonRole = 6
	QMessageBox__ResetRole QMessageBox__ButtonRole = 7
	QMessageBox__ApplyRole QMessageBox__ButtonRole = 8
	QMessageBox__NRoles QMessageBox__ButtonRole = 9
)

type QMessageBox__StandardButton int64

const (
	QMessageBox__NoButton QMessageBox__StandardButton = 0
	QMessageBox__Ok QMessageBox__StandardButton = 1024
	QMessageBox__Save QMessageBox__StandardButton = 2048
	QMessageBox__SaveAll QMessageBox__StandardButton = 4096
	QMessageBox__Open QMessageBox__StandardButton = 8192
	QMessageBox__Yes QMessageBox__StandardButton = 16384
	QMessageBox__YesToAll QMessageBox__StandardButton = 32768
	QMessageBox__No QMessageBox__StandardButton = 65536
	QMessageBox__NoToAll QMessageBox__StandardButton = 131072
	QMessageBox__Abort QMessageBox__StandardButton = 262144
	QMessageBox__Retry QMessageBox__StandardButton = 524288
	QMessageBox__Ignore QMessageBox__StandardButton = 1048576
	QMessageBox__Close QMessageBox__StandardButton = 2097152
	QMessageBox__Cancel QMessageBox__StandardButton = 4194304
	QMessageBox__Discard QMessageBox__StandardButton = 8388608
	QMessageBox__Help QMessageBox__StandardButton = 16777216
	QMessageBox__Apply QMessageBox__StandardButton = 33554432
	QMessageBox__Reset QMessageBox__StandardButton = 67108864
	QMessageBox__RestoreDefaults QMessageBox__StandardButton = 134217728
	QMessageBox__FirstButton QMessageBox__StandardButton = 1024
	QMessageBox__LastButton QMessageBox__StandardButton = 134217728
	QMessageBox__YesAll QMessageBox__StandardButton = 32768
	QMessageBox__NoAll QMessageBox__StandardButton = 131072
	QMessageBox__Default QMessageBox__StandardButton = 256
	QMessageBox__Escape QMessageBox__StandardButton = 512
	QMessageBox__FlagMask QMessageBox__StandardButton = 768
	QMessageBox__ButtonMask QMessageBox__StandardButton = -769
)

