package application

import (
	"thirawoot/in2forest_shop_backend/internal/domain"
	"thirawoot/in2forest_shop_backend/internal/dto"
	portin "thirawoot/in2forest_shop_backend/internal/ports/port_in"
	portout "thirawoot/in2forest_shop_backend/internal/ports/port_out"
	"thirawoot/in2forest_shop_backend/internal/utils/constants"
	"thirawoot/in2forest_shop_backend/internal/utils/mapper"
)

type EmployeeRoleAppImpl struct {
	repo portout.EmployeeRoleRepository
}

func NewEmployeeRoleApp(repository portout.EmployeeRoleRepository) portin.EmployeeRoleApp {
	return &EmployeeRoleAppImpl{
		repo: repository,
	}
}

func (a *EmployeeRoleAppImpl) FindByRole(role string) (*dto.EmployeeRole, error) {
	foundRole := a.repo.FindByRole(role)
	if foundRole == nil {
		return nil, ErrNotFound
	}

	mapRole := mapper.ToEmployeeRoleDto(*foundRole)

	return &mapRole, nil
}

func (a *EmployeeRoleAppImpl) Create(data dto.EmployeeRoleCreate) (*dto.EmployeeRole, error) {
	foundRole, _ := a.FindByRole(data.Role)
	if foundRole != nil {
		return nil, ErrAlreadyUse
	}

	role := domain.EmployeeRole{Role: data.Role}

	newRole, err := a.repo.Create(&role)
	if err != nil {
		return nil, NewAppErr("FAILED_TO_CREATE", constants.Code.ServerError, err)
	}

	return &dto.EmployeeRole{ID: newRole.ID, Role: newRole.Role}, nil
}

func (a *EmployeeRoleAppImpl) Delete(id uint) (map[string]int64, error) {
	affected, err := a.repo.SoftDelete(id)
	mapRowAffected := map[string]int64{"rowAffected": affected}

	if err != nil {
		return mapRowAffected, NewAppErr("FAILD_TO_DELETE", constants.Code.ServerError, err)
	}

	if affected < 1 {
		return mapRowAffected, ErrInvalidInternalData
	}

	return mapRowAffected, nil
}

func (a *EmployeeRoleAppImpl) Update(id uint, data dto.EmployeeRoleCreate) (map[string]int64, error) {
	_, err := a.Find(id)
	mapRowAffected := map[string]int64{"rowAffected": 0}
	if err != nil {
		return mapRowAffected, err
	}

	affected, err := a.repo.Update(id, (*domain.EmployeeRoleCreate)(&data))
	mapRowAffected["rowAffected"] = affected
	if err != nil {
		return mapRowAffected, NewAppErr("FAILED_TO_UPDATE", constants.Code.ServerError, err)
	}

	return mapRowAffected, nil
}

func (a *EmployeeRoleAppImpl) Find(id uint) (*dto.EmployeeRole, error) {
	foundRole := a.repo.Find(id)
	if foundRole == nil {
		return nil, ErrNotFound
	}

	resData := dto.EmployeeRole{ID: foundRole.ID, Role: foundRole.Role}

	return &resData, nil
}

func (a *EmployeeRoleAppImpl) FindAll() ([]dto.EmployeeRole, error) {
	foundRoles := a.repo.FindAll()
	roles := make([]dto.EmployeeRole, len(foundRoles))

	if len(foundRoles) == 0 {
		return nil, ErrNotFound
	}

	for i, role := range foundRoles {
		roles[i] = mapper.ToEmployeeRoleDto(role)
	}

	return roles, nil
}
