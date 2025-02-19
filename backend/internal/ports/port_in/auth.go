package portin

import "thirawoot/in2forest_shop_backend/internal/dto"

type AuthEmployeeApp interface {
	RegisterAdmin(data dto.Employee) (*dto.AuthResponse, error)
	LoginEmployee(data dto.AuthLogin) (*dto.AuthResponse, error)
}
