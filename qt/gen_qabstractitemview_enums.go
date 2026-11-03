package qt

type QAbstractItemView__SelectionMode int64

const (
	QAbstractItemView__NoSelection QAbstractItemView__SelectionMode = 0
	QAbstractItemView__SingleSelection QAbstractItemView__SelectionMode = 1
	QAbstractItemView__MultiSelection QAbstractItemView__SelectionMode = 2
	QAbstractItemView__ExtendedSelection QAbstractItemView__SelectionMode = 3
	QAbstractItemView__ContiguousSelection QAbstractItemView__SelectionMode = 4
)

type QAbstractItemView__SelectionBehavior int64

const (
	QAbstractItemView__SelectItems QAbstractItemView__SelectionBehavior = 0
	QAbstractItemView__SelectRows QAbstractItemView__SelectionBehavior = 1
	QAbstractItemView__SelectColumns QAbstractItemView__SelectionBehavior = 2
)

type QAbstractItemView__ScrollHint int64

const (
	QAbstractItemView__EnsureVisible QAbstractItemView__ScrollHint = 0
	QAbstractItemView__PositionAtTop QAbstractItemView__ScrollHint = 1
	QAbstractItemView__PositionAtBottom QAbstractItemView__ScrollHint = 2
	QAbstractItemView__PositionAtCenter QAbstractItemView__ScrollHint = 3
)

type QAbstractItemView__EditTrigger int64

const (
	QAbstractItemView__NoEditTriggers QAbstractItemView__EditTrigger = 0
	QAbstractItemView__CurrentChanged QAbstractItemView__EditTrigger = 1
	QAbstractItemView__DoubleClicked QAbstractItemView__EditTrigger = 2
	QAbstractItemView__SelectedClicked QAbstractItemView__EditTrigger = 4
	QAbstractItemView__EditKeyPressed QAbstractItemView__EditTrigger = 8
	QAbstractItemView__AnyKeyPressed QAbstractItemView__EditTrigger = 16
	QAbstractItemView__AllEditTriggers QAbstractItemView__EditTrigger = 31
)

type QAbstractItemView__ScrollMode int64

const (
	QAbstractItemView__ScrollPerItem QAbstractItemView__ScrollMode = 0
	QAbstractItemView__ScrollPerPixel QAbstractItemView__ScrollMode = 1
)

type QAbstractItemView__DragDropMode int64

const (
	QAbstractItemView__NoDragDrop QAbstractItemView__DragDropMode = 0
	QAbstractItemView__DragOnly QAbstractItemView__DragDropMode = 1
	QAbstractItemView__DropOnly QAbstractItemView__DragDropMode = 2
	QAbstractItemView__DragDrop QAbstractItemView__DragDropMode = 3
	QAbstractItemView__InternalMove QAbstractItemView__DragDropMode = 4
)

type QAbstractItemView__CursorAction int64

const (
	QAbstractItemView__MoveUp QAbstractItemView__CursorAction = 0
	QAbstractItemView__MoveDown QAbstractItemView__CursorAction = 1
	QAbstractItemView__MoveLeft QAbstractItemView__CursorAction = 2
	QAbstractItemView__MoveRight QAbstractItemView__CursorAction = 3
	QAbstractItemView__MoveHome QAbstractItemView__CursorAction = 4
	QAbstractItemView__MoveEnd QAbstractItemView__CursorAction = 5
	QAbstractItemView__MovePageUp QAbstractItemView__CursorAction = 6
	QAbstractItemView__MovePageDown QAbstractItemView__CursorAction = 7
	QAbstractItemView__MoveNext QAbstractItemView__CursorAction = 8
	QAbstractItemView__MovePrevious QAbstractItemView__CursorAction = 9
)

type QAbstractItemView__State int64

const (
	QAbstractItemView__NoState QAbstractItemView__State = 0
	QAbstractItemView__DraggingState QAbstractItemView__State = 1
	QAbstractItemView__DragSelectingState QAbstractItemView__State = 2
	QAbstractItemView__EditingState QAbstractItemView__State = 3
	QAbstractItemView__ExpandingState QAbstractItemView__State = 4
	QAbstractItemView__CollapsingState QAbstractItemView__State = 5
	QAbstractItemView__AnimatingState QAbstractItemView__State = 6
)

type QAbstractItemView__DropIndicatorPosition int64

const (
	QAbstractItemView__OnItem QAbstractItemView__DropIndicatorPosition = 0
	QAbstractItemView__AboveItem QAbstractItemView__DropIndicatorPosition = 1
	QAbstractItemView__BelowItem QAbstractItemView__DropIndicatorPosition = 2
	QAbstractItemView__OnViewport QAbstractItemView__DropIndicatorPosition = 3
)

