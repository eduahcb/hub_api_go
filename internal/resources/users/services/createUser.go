package usersservices

import (
	"errors"
	"time"

	"github.com/eduahcb/hub_api_go/config"
	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/entities"
	customErrors "github.com/eduahcb/hub_api_go/pkg/errors"
	"github.com/eduahcb/hub_api_go/pkg/security"
	"gorm.io/gorm"
)

func CreateUser(user entities.User, db *database.Database) (string, error) {
	err := db.Client.First(&entities.Module{}, user.ModuleID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return "", &customErrors.NotFound{Message: "Module not found"}
		}

		return "", &customErrors.InternalServerError{}
	}

	err = db.Client.First(&user, "email = ?", user.Email).Error

	if err == nil {
		return "", &customErrors.BadRequest{Message: "Email already exists"}
	}

	err = db.Client.Create(&user).Error

	if err != nil {
		return "", &customErrors.InternalServerError{}
	}

	expirationTime := time.Now().Add(time.Minute * 15).Unix()
	token, _ := security.Token(user.ID, expirationTime, config.Envs.SecretKey)

	return token, nil
}
