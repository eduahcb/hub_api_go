package usersservices

import (
	"errors"
	"strings"

	"github.com/eduahcb/hub_api_go/config"
	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/entities"
	customErrors "github.com/eduahcb/hub_api_go/pkg/errors"
	"github.com/eduahcb/hub_api_go/pkg/security"
	"github.com/eduahcb/hub_api_go/pkg/utils"
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

	err = db.Client.Create(&user).Error

	if err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return "", &customErrors.BadRequest{Message: "Email already exists"}
		}
		return "", &customErrors.InternalServerError{}
	}

	expirationTime := utils.ExpirationTime(config.Envs.ExpirationTime) 
	token, _ := security.Token(user.ID, expirationTime, config.Envs.SecretKey)

	return token, nil
}
