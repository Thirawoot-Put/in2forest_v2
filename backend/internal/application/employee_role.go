package application

import (
	"thirawoot/in2forest_shop_backend/internal/domain"
	"thirawoot/in2forest_shop_backend/internal/dto"
	portin "thirawoot/in2forest_shop_backend/internal/ports/port_in"
	portout "thirawoot/in2forest_shop_backend/internal/ports/port_out"
	"thirawoot/in2forest_shop_backend/internal/utils/constants"
	"thirawoot/in2forest_shop_backend/internal/utils/response"
)

type EmployeeRoleServiceImpl struct {
	repo portout.EmployeeRoleRepository
}

func NewEmployeeRoleService(repository portout.EmployeeRoleRepository) portin.EmployeeRoleService {
	return &EmployeeRoleServiceImpl{
		repo: repository,
	}
}

func (r *EmployeeRoleServiceImpl) Create(data dto.EmployeeRoleCreate) response.Response {
	role := domain.EmployeeRoleCreate{Role: data.Role}

	id, err := r.repo.Create(&role)
	if err != nil {
		return response.Error("failed to create role", constants.StatusCode.BadRequest)
	}

	return response.Success("create success", constants.StatusCode.Created, id)
}

func (r *EmployeeRoleServiceImpl) Delete(id uint) response.Response {
	err := r.repo.SoftDelete(id)
	if err != nil {
		return response.Error("failed to delete role", constants.StatusCode.ServerError)
	}

	return response.Success("delete success", constants.StatusCode.Ok, nil)
}
