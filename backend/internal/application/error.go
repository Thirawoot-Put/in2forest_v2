package application

import (
	"fmt"
	"thirawoot/in2forest_shop_backend/internal/utils/constants"
)

type AppErr struct {
	Message string `json:"message"`
	Code    int    `json:"statusCode"`
}

func (e *AppErr) Error() string {
	return e.Message
}

func NewAppErr(message string, code int, err error) *AppErr {
	if err != nil {
		combinMsg := fmt.Sprintf("%s: %s", message, err.Error())
		return &AppErr{
			Message: combinMsg,
			Code:    code,
		}
	}
	return &AppErr{
		Message: message,
		Code:    code,
	}
}

var (
	ErrAlreadyUse          = NewAppErr("ALREADY_IN_USE", constants.Code.BadRequest, nil)
	ErrInvalidInternalData = NewAppErr("INVALID_INTERNAL_DATA", constants.Code.BadRequest, nil)
	ErrNotFound            = NewAppErr("NOT_FOUND", constants.Code.NotFound, nil)
	ErrUserNotFound        = NewAppErr("USER_NOT_FOUND", constants.Code.NotFound, nil)
	ErrPasswordNotMatch    = NewAppErr("PASSWORD_NOT_MATCH", constants.Code.BadRequest, nil)
)
