package validation

import (
	"encoding/json"

	validator "github.com/go-ozzo/ozzo-validation"
	"github.com/go-ozzo/ozzo-validation/is"
	"github.com/vnnyx/gas-api/exception"
	"github.com/vnnyx/gas-api/model"
)

func CreateUSerValidation(request model.UserCreateRequest) {
	err := validator.ValidateStruct(&request,
		validator.Field(&request.Username, validator.Required),
		validator.Field(&request.Email, validator.Required, is.Email),
		validator.Field(&request.Handphone, validator.Required, is.Digit),
		validator.Field(&request.Password, validator.Required),
		validator.Field(&request.PasswordConfirmation, validator.Required),
	)
	if err != nil {
		b, _ := json.Marshal(err)
		err = exception.ValidationError{
			Message: string(b),
		}
		exception.PanicIfNeeded(err)
	}
}
