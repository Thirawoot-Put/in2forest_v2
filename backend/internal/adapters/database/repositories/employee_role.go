package repositories

import (
	"errors"
	"thirawoot/in2forest_shop_backend/internal/domain"
	"thirawoot/in2forest_shop_backend/internal/infras/database/model"
	portout "thirawoot/in2forest_shop_backend/internal/ports/port_out"
	"thirawoot/in2forest_shop_backend/internal/utils/mapper"

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
	role := model.EmployeeRole{Name: data.Name}

	err := r.db.Create(&role).Error
	if err != nil {
		return nil, err
	}

	newRole := mapper.ModelToEmployeeRoleDomain(role)
	return &newRole, nil
}

func (r *EmployeeRoleRepositoryImpl) FindByRole(roleName string) *domain.EmployeeRole {
	role := model.EmployeeRole{}
	result := r.db.First(&role, domain.EmployeeRole{Name: roleName})

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	foundRole := mapper.ModelToEmployeeRoleDomain(role)
	return &foundRole
}

func (r *EmployeeRoleRepositoryImpl) Find(id uint) *domain.EmployeeRole {
	role := model.EmployeeRole{}
	result := r.db.First(&role, domain.EmployeeRole{ID: id})

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return nil
	}

	domainRole := mapper.ModelToEmployeeRoleDomain(role)
	return &domainRole
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
	var role model.EmployeeRole

	updateModel := mapper.DomainToEmployeeRoleModel(*data)
	result := r.db.Model(&role).Where("id = ?", id).Updates(updateModel)
	if result.Error != nil {
		return result.RowsAffected, result.Error
	}

	return result.RowsAffected, nil
}
