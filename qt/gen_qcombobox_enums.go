package qt

type QComboBox__InsertPolicy int64

const (
	QComboBox__NoInsert QComboBox__InsertPolicy = 0
	QComboBox__InsertAtTop QComboBox__InsertPolicy = 1
	QComboBox__InsertAtCurrent QComboBox__InsertPolicy = 2
	QComboBox__InsertAtBottom QComboBox__InsertPolicy = 3
	QComboBox__InsertAfterCurrent QComboBox__InsertPolicy = 4
	QComboBox__InsertBeforeCurrent QComboBox__InsertPolicy = 5
	QComboBox__InsertAlphabetically QComboBox__InsertPolicy = 6
)

type QComboBox__SizeAdjustPolicy int64

const (
	QComboBox__AdjustToContents QComboBox__SizeAdjustPolicy = 0
	QComboBox__AdjustToContentsOnFirstShow QComboBox__SizeAdjustPolicy = 1
	QComboBox__AdjustToMinimumContentsLengthWithIcon QComboBox__SizeAdjustPolicy = 2
)

type QComboBox__LabelDrawingMode int64

const (
	QComboBox__LabelDrawingMode__UseStyle QComboBox__LabelDrawingMode = 0
	QComboBox__LabelDrawingMode__UseDelegate QComboBox__LabelDrawingMode = 1
)

