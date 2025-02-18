package application

import portin "thirawoot/in2forest_shop_backend/internal/ports/port_in"

type AuthAppImpl struct {
	empApp portin.EmployeeApp
}

func NewAuthApp(empApp portin.EmployeeApp) portin.AuthApp {
	return &AuthAppImpl{
		empApp: empApp,
	}
}

func (a *AuthAppImpl) RegisterAdmin() {

}
