package qt

type QSystemTrayIcon__ActivationReason int64

const (
	QSystemTrayIcon__Unknown QSystemTrayIcon__ActivationReason = 0
	QSystemTrayIcon__Context QSystemTrayIcon__ActivationReason = 1
	QSystemTrayIcon__DoubleClick QSystemTrayIcon__ActivationReason = 2
	QSystemTrayIcon__Trigger QSystemTrayIcon__ActivationReason = 3
	QSystemTrayIcon__MiddleClick QSystemTrayIcon__ActivationReason = 4
)

type QSystemTrayIcon__MessageIcon int64

const (
	QSystemTrayIcon__NoIcon QSystemTrayIcon__MessageIcon = 0
	QSystemTrayIcon__Information QSystemTrayIcon__MessageIcon = 1
	QSystemTrayIcon__Warning QSystemTrayIcon__MessageIcon = 2
	QSystemTrayIcon__Critical QSystemTrayIcon__MessageIcon = 3
)

