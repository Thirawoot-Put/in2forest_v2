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

	resData := map[string]uint{"id": *id}
	return response.Success("create success", constants.StatusCode.Created, resData)
}

func (r *EmployeeRoleServiceImpl) Delete(id uint) response.Response {
	affected, err := r.repo.HardDelete(id)
	if err != nil {
		return response.Error("failed to delete role", constants.StatusCode.ServerError)
	}

	if affected < 1 {
		return response.Error("invalid role", constants.StatusCode.NotFound)
	}

	resData := map[string]int64{"rowAffected": affected}

	return response.Success("delete success", constants.StatusCode.Ok, resData)
}
