package portout

import "thirawoot/in2forest_shop_backend/internal/domain"

type EmployeeRoleRepository interface {
	Create(data *domain.EmployeeRoleCreate) (*uint, error)
	FindByRole(role string) *domain.EmployeeRole
	Find(id uint) *domain.EmployeeRole
	SoftDelete(id uint) (int64, error)
	HardDelete(id uint) (int64, error)
	Update(id uint, data *domain.EmployeeRoleCreate) (int64, error)
}
