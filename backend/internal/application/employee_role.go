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

func (s *EmployeeRoleServiceImpl) findByRole(data dto.EmployeeRole) *domain.EmployeeRole {
	role := s.repo.FindByRole(data.Role)

	return role
}

func (s *EmployeeRoleServiceImpl) Create(data dto.EmployeeRoleCreate) response.Response {
	role := domain.EmployeeRoleCreate{Role: data.Role}

	foundRole := s.findByRole(dto.EmployeeRole{Role: role.Role})
	if foundRole != nil {
		return response.Error("THIS_ROLE_ALREADY_EXIST", constants.StatusCode.Conflict)
	}

	id, err := s.repo.Create(&role)
	if err != nil {
		return response.Error("FAILED_TO_CREATE_ROLE", constants.StatusCode.ServerError)
	}

	resData := map[string]uint{"id": *id}
	return response.Success("create success", constants.StatusCode.Created, resData)
}

func (s *EmployeeRoleServiceImpl) Delete(id uint) response.Response {
	affected, err := s.repo.SoftDelete(id)
	if err != nil {
		return response.Error("FAILED_TO_DELETE_ROLE", constants.StatusCode.ServerError)
	}

	if affected < 1 {
		return response.Error("INVALID_ROLE", constants.StatusCode.NotFound)
	}

	resData := map[string]int64{"rowAffected": affected}

	return response.Success("delete success", constants.StatusCode.Ok, resData)
}
