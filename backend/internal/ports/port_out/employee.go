package portout

import "thirawoot/in2forest_shop_backend/internal/domain"

type EmployeeRepository interface {
	Create(data *domain.Employee) (*domain.Employee, error)
}
