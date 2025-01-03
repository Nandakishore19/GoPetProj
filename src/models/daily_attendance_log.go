package models


type DailyAttendanceLog struct {
	AttendanceID uint `gorm:"primaryKey;autoIncrement"`
	EmployeeID  int
	Date         string
	FirstLoggedAt string
	LastLoggedAt string

	// Relations
	// Employee Employee `gorm:"foreignKey:EmployeeID;constraint:OnDelete:CASCADE"`
}