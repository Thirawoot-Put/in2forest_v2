package repositories

import (
	"errors"
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
	var emp model.Employee

	err := r.db.Preload("EmployeeRole").Where("email = ?", email).First(&emp).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	foundEmp := mapper.ToEmployeeDomain(emp)

	return &foundEmp, nil
}

func (r *EmployeeRepositoryImpl) Find(id uint) (*domain.Employee, error) {
	var empModel model.Employee

	err := r.db.Preload("EmployeeRole").Where("id = ?", id).First(&empModel).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}

		return nil, err
	}

	empDoamin := mapper.ToEmployeeDomain(empModel)
	return &empDoamin, nil
}

func (r *EmployeeRepositoryImpl) Update(id uint, data *domain.Employee) (int64, error) {
	empModel := mapper.ToEmployeeModel(*data)

	result := r.db.Model(&model.Employee{}).Where("id = ?", id).Updates(empModel)
	if result.Error != nil {
		return 0, result.Error
	}

	return result.RowsAffected, nil
}
