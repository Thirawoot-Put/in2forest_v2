package portin

import "thirawoot/in2forest_shop_backend/internal/dto"

type AuthApp interface {
	RegisterAdmin(data dto.Employee) (string, error)
}
