package repositories

import (
	"thirawoot/in2forest_shop_backend/internal/domain"
	portout "thirawoot/in2forest_shop_backend/internal/ports/port_out"

	"gorm.io/gorm"
)

type EmployeeRoleRepositoryImpl struct {
	db *gorm.DB
}

func NewEmployeeRoleRepository(db *gorm.DB) portout.EmployeeRoleRepository {
	return &EmployeeRoleRepositoryImpl{
		db: db,
	}
}

func (r *EmployeeRoleRepositoryImpl) Create(data *domain.EmployeeRoleCreate) (*uint, error) {
	role := domain.EmployeeRole{Role: data.Role}

	result := r.db.Create(&role)

	if result.Error != nil {
		return nil, result.Error
	}

	return &role.ID, nil
}
