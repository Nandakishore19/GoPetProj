package models

type Employee struct {
	EmployeeID int    `gorm:"primaryKey"`
	Email      string `gorm:"unique"`
	Name       string
	MacAddress string `gorm:"unique"`
	// Defining Relations
	DailyAttendanceLogs []DailyAttendanceLog `gorm:"foreignKey:EmployeeID;rconstraint:OnDelete:CASCADE"`
}