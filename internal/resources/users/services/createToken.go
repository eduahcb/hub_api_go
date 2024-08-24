package usersservices

import (
	"errors"

	"github.com/eduahcb/hub_api_go/config"
	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/entities"
	"github.com/eduahcb/hub_api_go/internal/resources/users"
	customErrors "github.com/eduahcb/hub_api_go/pkg/errors"
	"github.com/eduahcb/hub_api_go/pkg/security"
	"github.com/eduahcb/hub_api_go/pkg/utils"
	"gorm.io/gorm"
)

func CreateToken(signinRequest users.SigninRequest, db *database.Database) (string, error) {
	user := entities.User{}

	err := db.Client.First(&user, "email = ?", signinRequest.Email).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", &customErrors.NotFound{Message: "User not found"}
		}

		return "", &customErrors.InternalServerError{}
	}

	err = security.ComparePasswords(user.Password, signinRequest.Password)

	if err != nil {
		return "", &customErrors.Unauthorized{Message: "Please check your credentials and try again"}
	}

	expiration := utils.ExpirationTime(config.Envs.ExpirationTime)

	token, _ := security.Token(user.ID, expiration, config.Envs.SecretKey)

	return token, nil
}
