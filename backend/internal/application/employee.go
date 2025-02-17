package application

import (
	"thirawoot/in2forest_shop_backend/internal/domain"
	"thirawoot/in2forest_shop_backend/internal/dto"
	portin "thirawoot/in2forest_shop_backend/internal/ports/port_in"
	portout "thirawoot/in2forest_shop_backend/internal/ports/port_out"
	"thirawoot/in2forest_shop_backend/internal/utils/bcrypt"
	"thirawoot/in2forest_shop_backend/internal/utils/constants"
)

type EmployeeAppImpl struct {
	repo portout.EmployeeRepository
}

func NewEmployeeApp(repository portout.EmployeeRepository) portin.EmployeeApp {
	return &EmployeeAppImpl{
		repo: repository,
	}
}

func (s *EmployeeAppImpl) Create(data dto.Employee) (*dto.Employee, error) {
	hash, err := bcrypt.HashPwd(data.Password)
	if err != nil {
		return nil, NewAppErr("FAILED_HASH_PASSWORD", constants.Code.ServerError, err)
	}

	result, err := s.repo.Create(&domain.Employee{
		Email:          data.Email,
		Password:       hash,
		Name:           data.Name,
		Mobile:         data.Mobile,
		EmployeeRoleID: data.EmployeeRoleID,
	})
	if err != nil {
		return nil, NewAppErr("FAILED_TO_CREATE_EMPLOYEE", constants.Code.ServerError, err)
	}

	return &dto.Employee{
		ID:     result.ID,
		Email:  result.Email,
		Name:   result.Name,
		Mobile: result.Mobile,
	}, nil
}
