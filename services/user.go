package services

import (
	"fmt"
	"runtime/debug"

	"gitlab.sudovi.me/erp/core-ms-api/data"
	"gitlab.sudovi.me/erp/core-ms-api/dto"
	"gitlab.sudovi.me/erp/core-ms-api/errors"

	"github.com/oykos-development-hub/celeritas"
	up "github.com/upper/db/v4"
)

type userServiceImpl struct {
	App  *celeritas.Celeritas
	repo data.User
}

func NewUserServiceImpl(app *celeritas.Celeritas, repo data.User) UserService {
	return &userServiceImpl{
		App:  app,
		repo: repo,
	}
}

func (h *userServiceImpl) CreateUser(userInput dto.UserRegistrationDTO) (*dto.UserResponseDTO, error) {
	_, err := h.repo.GetByEmail(userInput.Email)
	if err == nil {
		return nil, errors.ErrUserEmailExists
	}

	u := userInput.ToUser()

	id, err := h.repo.Insert(*u)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrInternalServer
	}

	u, err = u.Get(id)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToUserResponseDTO(*u)

	return response, nil
}

func (h *userServiceImpl) UpdateUser(userId int, userInput dto.UserUpdateDTO) (*dto.UserResponseDTO, error) {
	u, err := h.repo.Get(userId)
	if err != nil {
		return nil, errors.ErrNotFound
	}

	if userInput.Email != nil && *userInput.Email != u.Email {
		foundUser, _ := h.repo.GetByEmail(*userInput.Email)
		if foundUser != nil {
			return nil, errors.ErrUserEmailExists
		}
	}

	userInput.ToUser(u)
	err = h.repo.Update(*u)
	if err != nil {
		return nil, errors.ErrInternalServer
	}

	response := dto.ToUserResponseDTO(*u)

	return response, nil
}

func (h *userServiceImpl) GetUser(userId int) (*dto.UserResponseDTO, error) {
	u, err := h.repo.Get(userId)
	if err != nil {
		h.App.InfoLog.Println(string(debug.Stack()))
		h.App.ErrorLog.Println(err)
		return nil, errors.ErrNotFound
	}
	response := dto.ToUserResponseDTO(*u)

	return response, nil
}

func (h *userServiceImpl) GetUserList(data dto.GetUserListDTO) ([]dto.UserResponseDTO, *uint64, error) {
	var conditionAndExp *up.AndExpr
	if data.IsActive != nil {
		conditionAndExp = &up.AndExpr{}
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"active": data.IsActive})
	}
	if data.Email != nil && *data.Email != "" {
		if conditionAndExp == nil {
			conditionAndExp = &up.AndExpr{}
		}
		likeCondition := fmt.Sprintf("%%%s%%", *data.Email)
		conditionAndExp = up.And(conditionAndExp, &up.Cond{"email ILIKE": likeCondition})
	}
	u, total, err := h.repo.GetAll(data.Page, data.Size, conditionAndExp)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return nil, nil, errors.ErrInternalServer
	}
	response := dto.ToUsersResponseDTO(u)

	return response, total, nil
}

func (h *userServiceImpl) DeleteUser(id int) error {
	err := h.repo.Delete(id)
	if err != nil {
		h.App.ErrorLog.Println(err)
		return errors.ErrInternalServer
	}

	return nil
}
