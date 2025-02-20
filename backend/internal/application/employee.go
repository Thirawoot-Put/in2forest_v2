package application

import (
	"thirawoot/in2forest_shop_backend/internal/dto"
	portin "thirawoot/in2forest_shop_backend/internal/ports/port_in"
	portout "thirawoot/in2forest_shop_backend/internal/ports/port_out"
	"thirawoot/in2forest_shop_backend/internal/utils/constants"
	"thirawoot/in2forest_shop_backend/internal/utils/mapper"
	"thirawoot/in2forest_shop_backend/pkg"
)

type EmployeeAppImpl struct {
	repo portout.EmployeeRepository
}

func NewEmployeeApp(repository portout.EmployeeRepository) portin.EmployeeApp {
	return &EmployeeAppImpl{
		repo: repository,
	}
}

func (a *EmployeeAppImpl) Create(data dto.Employee) (*dto.Employee, error) {
	hash, err := pkg.HashPwd(data.Password)
	if err != nil {
		return nil, NewAppErr("FAILED_HASH_PASSWORD", constants.Code.ServerError, err)
	}

	data.Password = hash
	domainEmp := mapper.DtoToEmployeeDomain(data)

	result, err := a.repo.Create(&domainEmp)
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

func (a *EmployeeAppImpl) FindByEmail(email string) (*dto.Employee, error) {
	emp, err := a.repo.FindByEmail(email)

	if err != nil {
		return nil, NewAppErr("FAILED_TO_FETCH", constants.Code.ServerError, err)
	} else if emp == nil {
		return nil, ErrUserNotFound
	}

	dtoEmp := mapper.DaminToEmployeeDto(*emp)
	return &dtoEmp, nil
}

func (a *EmployeeAppImpl) Find(id uint) (*dto.EmployeeResponse, error) {
	emp, err := a.repo.Find(id)
	if err != nil {
		return nil, NewAppErr("FAILED_TO_FETCH", constants.Code.ServerError, err)
	} else if emp == nil {
		return nil, ErrUserNotFound
	}

	dtoEmp := mapper.DaminToEmployeeResponseDto(*emp)
	return &dtoEmp, nil
}

func (a *EmployeeAppImpl) Update(id uint, data dto.Employee) (*dto.EmployeeResponse, error) {
	emp, err := a.repo.Find(id)
	if err != nil {
		return nil, NewAppErr("FAILED_TO_FETCH", constants.Code.ServerError, err)
	} else if emp == nil {
		return nil, ErrUserNotFound
	}

	empDomain := mapper.DtoToEmployeeDomain(data)
	updated, err := a.repo.Update(id, &empDomain)
	if err != nil {
		return nil, NewAppErr("FAILED_TO_UPDATE", constants.Code.ServerError, err)
	}

	empDto := mapper.DaminToEmployeeResponseDto(*updated)
	return &empDto, nil
}
