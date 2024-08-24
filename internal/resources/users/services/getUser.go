package usersservices

import (
	"errors"

	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/entities"
	customErrors "github.com/eduahcb/hub_api_go/pkg/errors"
	"gorm.io/gorm"
)

func GetUser(userId uint, db *database.Database) error {
	var user entities.User

	err := db.Client.Preload("Module").First(&user, userId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return &customErrors.NotFound{Message: "User not found"}
		}
		return &customErrors.InternalServerError{}
	}

	return nil
}
