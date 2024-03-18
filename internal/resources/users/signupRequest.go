package users

import (
	"github.com/eduahcb/hub_api_go/internal/entities"
	"github.com/eduahcb/hub_api_go/pkg/security"
	"github.com/eduahcb/hub_api_go/pkg/validate"
)

type SignupRequest struct {
	Name            string `json:"name" validate:"required,max=50"`
	Email           string `json:"email" validate:"required,email"`
	Password        string `json:"password" validate:"required,min=8"`
	ConfirmPassword string `json:"confirm_password" validate:"required,eqfield=Password"`
	Bio             string `json:"bio"`
	Contact         string `json:"contact"`
	ModuleId        uint   `json:"module_id" validate:"required"`
}

func (s *SignupRequest) Validate() ([]validate.ValidateErrors, error) {
	return validate.ValidateRequestBody(s)
}

func (s *SignupRequest) ToEntity() entities.User {
	hashPassword, _ := security.CreateHashPassword(s.Password)

	return entities.User{
		Name:     s.Name,
		Email:    s.Email,
		Password: string(hashPassword),
		Bio:      s.Bio,
		Contact:  s.Contact,
		ModuleID: s.ModuleId,
	}
}
