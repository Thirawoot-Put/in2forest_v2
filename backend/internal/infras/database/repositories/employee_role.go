package repositories

import (
	"thirawoot/in2forest_shop_backend/internal/domain"
	"thirawoot/in2forest_shop_backend/internal/infras/database/model"
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

func (r *EmployeeRoleRepositoryImpl) SoftDelete(id uint) (int64, error) {
	result := r.db.Delete(&model.EmployeeRole{}, id)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}

	return result.RowsAffected, nil
}

func (r *EmployeeRoleRepositoryImpl) HardDelete(id uint) (int64, error) {
	result := r.db.Unscoped().Delete(&model.EmployeeRole{}, id)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}

	return result.RowsAffected, nil
}
