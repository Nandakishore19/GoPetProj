package repositories

import (
	// "fmt"
	"petProj/src/models"

	"gorm.io/gorm"
)
// TODO: Get daily reports
type ReportsRepository struct {
	db *gorm.DB
}

type ReportsRepositoryInterface interface {
	GetReports() ([]models.DailyAttendanceLog, error)
}

func NewReportsRepository(db *gorm.DB) *ReportsRepository {
	return &ReportsRepository{
		db:db,
	}
}

// func (repo *ReportsRepository) GetWeeklyReports() ([]models.DailyAttendanceLog, error) {
// 	var reports []models.DailyAttendanceLog
// 	fmt.Println("reports", reports)
// 	err := repo.db.Select("employee_id", "date", "first_logged_at", "last_logged_at").Find(&reports).Error
// 	if err!= nil{
// 		fmt.Println("asdlka")
// 	}

// }