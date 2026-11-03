package qt

type QDateTimeEdit__Section int64

const (
	QDateTimeEdit__NoSection QDateTimeEdit__Section = 0
	QDateTimeEdit__AmPmSection QDateTimeEdit__Section = 1
	QDateTimeEdit__MSecSection QDateTimeEdit__Section = 2
	QDateTimeEdit__SecondSection QDateTimeEdit__Section = 4
	QDateTimeEdit__MinuteSection QDateTimeEdit__Section = 8
	QDateTimeEdit__HourSection QDateTimeEdit__Section = 16
	QDateTimeEdit__DaySection QDateTimeEdit__Section = 256
	QDateTimeEdit__MonthSection QDateTimeEdit__Section = 512
	QDateTimeEdit__YearSection QDateTimeEdit__Section = 1024
	QDateTimeEdit__TimeSections_Mask QDateTimeEdit__Section = 31
	QDateTimeEdit__DateSections_Mask QDateTimeEdit__Section = 1792
)

