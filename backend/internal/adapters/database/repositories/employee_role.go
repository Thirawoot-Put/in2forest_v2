package repositories

import (
	"errors"
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

func (r *EmployeeRoleRepositoryImpl) Create(data *domain.EmployeeRole) (*domain.EmployeeRole, error) {
	role := domain.EmployeeRole{Name: data.Name}

	result := r.db.Create(&role)

	if result.Error != nil {
		return nil, result.Error
	}

	return &role, nil
}

func (r *EmployeeRoleRepositoryImpl) FindByRole(roleName string) *domain.EmployeeRole {
	foundRole := domain.EmployeeRole{}
	result := r.db.First(&foundRole, domain.EmployeeRole{Name: roleName})

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &foundRole
}

func (r *EmployeeRoleRepositoryImpl) Find(id uint) *domain.EmployeeRole {
	foundRole := domain.EmployeeRole{}
	result := r.db.First(&foundRole, domain.EmployeeRole{ID: id})

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	return &foundRole
}

func (r *EmployeeRoleRepositoryImpl) FindAll() []domain.EmployeeRole {
	var roles []domain.EmployeeRole
	r.db.Find(&roles)

	return roles
}

func (r *EmployeeRoleRepositoryImpl) SoftDelete(id uint) (int64, error) {
	result := r.db.Delete(&domain.EmployeeRole{}, id)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}

	return result.RowsAffected, nil
}

func (r *EmployeeRoleRepositoryImpl) HardDelete(id uint) (int64, error) {
	result := r.db.Unscoped().Delete(&domain.EmployeeRole{}, id)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}

	return result.RowsAffected, nil
}

func (r *EmployeeRoleRepositoryImpl) Update(id uint, data *domain.EmployeeRole) (int64, error) {
	result := r.db.Save(&domain.EmployeeRole{ID: id, Name: data.Name})
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}

	return result.RowsAffected, nil
}
