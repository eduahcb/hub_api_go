package users

import (
	"github.com/eduahcb/hub_api_go/pkg/validate"
)

type SigninRequest struct {
	Email    string `json:"email" validate:"required,email"`
	Password string `json:"password" validate:"required"`
}

func (s *SigninRequest) Validate() ([]validate.ValidateErrors, error) {
	return validate.ValidateRequestBody(s)
}
