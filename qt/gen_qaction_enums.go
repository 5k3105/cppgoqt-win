package qt

type QAction__MenuRole int64

const (
	QAction__NoRole QAction__MenuRole = 0
	QAction__TextHeuristicRole QAction__MenuRole = 1
	QAction__ApplicationSpecificRole QAction__MenuRole = 2
	QAction__AboutQtRole QAction__MenuRole = 3
	QAction__AboutRole QAction__MenuRole = 4
	QAction__PreferencesRole QAction__MenuRole = 5
	QAction__QuitRole QAction__MenuRole = 6
)

type QAction__Priority int64

const (
	QAction__LowPriority QAction__Priority = 0
	QAction__NormalPriority QAction__Priority = 128
	QAction__HighPriority QAction__Priority = 256
)

type QAction__ActionEvent int64

const (
	QAction__Trigger QAction__ActionEvent = 0
	QAction__Hover QAction__ActionEvent = 1
)

