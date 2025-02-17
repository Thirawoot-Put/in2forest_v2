package portout

import "thirawoot/in2forest_shop_backend/internal/domain"

type EmployeeRole interface {
	Save(data domain.EmployeeRoleCreate) (*uint, error)
}
