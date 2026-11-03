package qt

type QFrame__Shape int64

const (
	QFrame__NoFrame QFrame__Shape = 0
	QFrame__Box QFrame__Shape = 1
	QFrame__Panel QFrame__Shape = 2
	QFrame__WinPanel QFrame__Shape = 3
	QFrame__HLine QFrame__Shape = 4
	QFrame__VLine QFrame__Shape = 5
	QFrame__StyledPanel QFrame__Shape = 6
)

type QFrame__Shadow int64

const (
	QFrame__Plain QFrame__Shadow = 16
	QFrame__Raised QFrame__Shadow = 32
	QFrame__Sunken QFrame__Shadow = 48
)

type QFrame__StyleMask int64

const (
	QFrame__Shadow_Mask QFrame__StyleMask = 240
	QFrame__Shape_Mask QFrame__StyleMask = 15
)

