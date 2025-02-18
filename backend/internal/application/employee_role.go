package application

import (
	"fmt"
	"thirawoot/in2forest_shop_backend/internal/domain"
	"thirawoot/in2forest_shop_backend/internal/dto"
	portin "thirawoot/in2forest_shop_backend/internal/ports/port_in"
	portout "thirawoot/in2forest_shop_backend/internal/ports/port_out"
	"thirawoot/in2forest_shop_backend/internal/utils/constants"
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

func (s *EmployeeRoleServiceImpl) Create(data dto.EmployeeRoleCreate) (*dto.EmployeeRole, error) {
	role := domain.EmployeeRole{Role: data.Role}

	foundRole := s.findByRole(dto.EmployeeRole{Role: role.Role})
	if foundRole != nil {
		return nil, ErrAlreadyUse
	}

	newRole, err := s.repo.Create(&role)
	if err != nil {
		return nil, NewAppErr("FAILED_TO_CREATE", constants.Code.ServerError, err)
	}

	return &dto.EmployeeRole{ID: newRole.ID, Role: newRole.Role}, nil
}

func (s *EmployeeRoleServiceImpl) Delete(id uint) (map[string]int64, error) {
	affected, err := s.repo.SoftDelete(id)
	mapRowAffected := map[string]int64{"rowAffected": affected}

	if err != nil {
		return mapRowAffected, NewAppErr("FAILD_TO_DELETE", constants.Code.ServerError, err)
	}

	if affected < 1 {
		return mapRowAffected, ErrInvalidInternalData
	}

	return mapRowAffected, nil
}

func (s *EmployeeRoleServiceImpl) find(id uint) (*dto.EmployeeRole, error) {
	foundRole := s.repo.Find(id)

	if foundRole == nil {
		return nil, fmt.Errorf("ROLE_NOT_FOUND")
	}

	return &dto.EmployeeRole{ID: foundRole.ID, Role: foundRole.Role}, nil
}

func (s *EmployeeRoleServiceImpl) Find(id uint) (*dto.EmployeeRole, error) {
	foundRole := s.repo.Find(id)
	if foundRole == nil {
		return nil, ErrNotFound
	}

	resData := dto.EmployeeRole{ID: foundRole.ID, Role: foundRole.Role}

	return &resData, nil
}

func (s *EmployeeRoleServiceImpl) Update(id uint, data dto.EmployeeRoleCreate) (map[string]int64, error) {
	_, err := s.Find(id)
	mapRowAffected := map[string]int64{"rowAffected": 0}
	if err != nil {
		return mapRowAffected, err
	}

	affected, err := s.repo.Update(id, (*domain.EmployeeRoleCreate)(&data))
	mapRowAffected["rowAffected"] = affected
	if err != nil {
		return mapRowAffected, NewAppErr("FAILED_TO_UPDATE", constants.Code.ServerError, err)
	}

	return mapRowAffected, nil
}
