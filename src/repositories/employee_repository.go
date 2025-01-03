package repositories

import (
	"gorm.io/gorm"
	"petProj/src/models"
)

type EmployeeRepository struct {
	db *gorm.DB
}

type EmployeeRepositoryInterface interface {
	GetEmployees()([]models.Employee, error)
	AddEmployee(id int, email,  macAddress, name string)error
}

func NewEmployeeRepository(db *gorm.DB) *EmployeeRepository {
	return &EmployeeRepository{
		db:db,
	}
}

// GetEmployees
func (repo *EmployeeRepository) GetEmployees(mac_address []string)([]models.Employee, error) {
	var employees []models.Employee
	// if err := repo.db.Where("mac_address IN ?", mac_address).Find(&employees).Error; err != nil {
	// 	return nil, err
	// }
	err := repo.db.Where("mac_address IN ?", mac_address).Find(&employees).Error
	if err != nil {
		return nil, err
	}
	return employees, nil
}

// AddEmployee
func (repo *EmployeeRepository) AddEmployee(id int, email, macAddress, name string) error {
	person := models.Employee{
		EmployeeID:  id,
		Email:      email,
		MacAddress: macAddress,
		Name:       name,
	}
	return repo.db.Create(&person).Error
}