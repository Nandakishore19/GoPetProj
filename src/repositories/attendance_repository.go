package repositories

import (
	"errors"
	"fmt"
	"petProj/src/models"
	"time"

	"gorm.io/gorm"

)

//LogEmployeeAttendance

type AttendanceRepositoryInterface interface {
	LogEmployeeAttendance(employeeId int, date time.Time)error
}

type AttendanceRepository struct {
	db *gorm.DB
}

func NewAttendanceRepository(db *gorm.DB) *AttendanceRepository {
	return &AttendanceRepository{
		db:db,
	}
}

// TODO: make Emloyee Id and date as composite key, So that there wont be any exception case with duplicate logs
func (repo *AttendanceRepository) LogEmployeeAttendance(employeeId int, date string) error {
	var attendance models.DailyAttendanceLog
	fmt.Println(attendance.EmployeeID)
	err := repo.db.Where(("employee_id = ? AND date = ?"), employeeId, date).First(&attendance).Error

	if errors.Is(err, gorm.ErrRecordNotFound){
		newAttendance := models.DailyAttendanceLog{
			EmployeeID: employeeId,
			Date: date,
			FirstLoggedAt: time.Now().Format("15:04:05"),
			LastLoggedAt: time.Now().Format("15:04:05"),
		}
		if err := repo.db.Create(&newAttendance).Error; err != nil {
			return fmt.Errorf("failed to create new attendance log: %w", err)
		}
	} else {
		fmt.Println("Came here")
		repo.db.Model(&models.DailyAttendanceLog{}).Where("employee_id = ? AND date = ?", employeeId, date).Update("last_logged_at", time.Now().Format("15:04:05"))

		// if err := repo.db.Save(&attendance).Error; err != nil {
		// 	return fmt.Errorf("failed to update attendance log: %w", err)
		// }
	}

	return nil
}