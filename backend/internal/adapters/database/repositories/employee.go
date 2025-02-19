package repositories

import (
	"errors"
	"fmt"
	"thirawoot/in2forest_shop_backend/internal/domain"
	"thirawoot/in2forest_shop_backend/internal/infras/database/model"
	portout "thirawoot/in2forest_shop_backend/internal/ports/port_out"
	"thirawoot/in2forest_shop_backend/internal/utils/mapper"

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
	emp := mapper.ToEmployeeModel(*data)

	result := r.db.Create(&emp)

	if result.Error != nil {
		return nil, result.Error
	}

	newEmp := mapper.ToEmployeeDomain(emp)

	return &newEmp, nil
}

func (r *EmployeeRepositoryImpl) FindByEmail(email string) (*domain.Employee, error) {
	emp := model.Employee{}

	result := r.db.Joins("JOIN employee_roles ON employee_roles.id = employees.employee_role_id").First(&emp, "email = ?", email)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, result.Error
	}

	fmt.Println("---> emp: ", emp)
	foundEmp := mapper.ToEmployeeDomain(emp)
	fmt.Println("---> foundEmp: ", foundEmp)

	return &foundEmp, nil
}
