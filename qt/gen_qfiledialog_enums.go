package qt

type QFileDialog__ViewMode int64

const (
	QFileDialog__Detail QFileDialog__ViewMode = 0
	QFileDialog__List QFileDialog__ViewMode = 1
)

type QFileDialog__FileMode int64

const (
	QFileDialog__AnyFile QFileDialog__FileMode = 0
	QFileDialog__ExistingFile QFileDialog__FileMode = 1
	QFileDialog__Directory QFileDialog__FileMode = 2
	QFileDialog__ExistingFiles QFileDialog__FileMode = 3
)

type QFileDialog__AcceptMode int64

const (
	QFileDialog__AcceptOpen QFileDialog__AcceptMode = 0
	QFileDialog__AcceptSave QFileDialog__AcceptMode = 1
)

type QFileDialog__DialogLabel int64

const (
	QFileDialog__LookIn QFileDialog__DialogLabel = 0
	QFileDialog__FileName QFileDialog__DialogLabel = 1
	QFileDialog__FileType QFileDialog__DialogLabel = 2
	QFileDialog__Accept QFileDialog__DialogLabel = 3
	QFileDialog__Reject QFileDialog__DialogLabel = 4
)

type QFileDialog__Option int64

const (
	QFileDialog__ShowDirsOnly QFileDialog__Option = 1
	QFileDialog__DontResolveSymlinks QFileDialog__Option = 2
	QFileDialog__DontConfirmOverwrite QFileDialog__Option = 4
	QFileDialog__DontUseNativeDialog QFileDialog__Option = 8
	QFileDialog__ReadOnly QFileDialog__Option = 16
	QFileDialog__HideNameFilterDetails QFileDialog__Option = 32
	QFileDialog__DontUseCustomDirectoryIcons QFileDialog__Option = 64
)

