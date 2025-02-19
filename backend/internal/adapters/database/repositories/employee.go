package repositories

import (
	"thirawoot/in2forest_shop_backend/internal/domain"
	portout "thirawoot/in2forest_shop_backend/internal/ports/port_out"

	"gorm.io/gorm"
)

type EmployeeRepositoryImpl struct {
	db *gorm.DB
}

func NewEmployeeRepository(db *gorm.DB) portout.EmployeeRepository {
	return &EmployeeRepositoryImpl{
		db: db,
	}
}

func (r *EmployeeRepositoryImpl) Create(data *domain.Employee) (*domain.Employee, error) {
	emp := domain.Employee{
		Email:    data.Email,
		Password: data.Password,
		Name:     data.Name,
		Mobile:   data.Mobile,

		EmployeeRoleID: data.EmployeeRoleID,
	}

	result := r.db.Create(&emp)

	if result.Error != nil {
		return nil, result.Error
	}

	return &emp, nil
}
