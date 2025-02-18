package application

import (
	"thirawoot/in2forest_shop_backend/internal/dto"
	portin "thirawoot/in2forest_shop_backend/internal/ports/port_in"
)

type AuthAppImpl struct {
	empApp     portin.EmployeeApp
	empRoleApp portin.EmployeeRoleApp
}

func NewAuthApp(
	empApp portin.EmployeeApp,
	empRoleApp portin.EmployeeRoleApp,
) portin.AuthApp {
	return &AuthAppImpl{
		empApp:     empApp,
		empRoleApp: empRoleApp,
	}
}

func (a *AuthAppImpl) RegisterAdmin(data dto.Employee) (string, error) {
	role, err := a.empRoleApp.FindByRole("ADMIN")
	if err != nil {
		return "", err
	}

	return "token", nil
}
