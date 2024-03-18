package techsservices

import (
	"errors"

	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/entities"
	"github.com/eduahcb/hub_api_go/internal/resources/techs"
	customErrors "github.com/eduahcb/hub_api_go/pkg/errors"
	"gorm.io/gorm"
)

func GetById(techId uint64, db *database.Database) (techs.TechResponse, error) {
	tech := entities.Tech{}
	var techResponse techs.TechResponse

	err := db.Client.Preload("Level").First(&tech, techId).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return techResponse, &customErrors.NotFound{Message: "Tech not found"}
		}
		return techResponse, &customErrors.InternalServerError{}
	}

	techResponse = techs.TechResponse{
		ID:   tech.ID,
		Name: tech.Name,
		Level: techs.LevelResponse{
			ID:   tech.Level.ID,
			Name: tech.Level.Name,
		},
	}

	return techResponse, nil
}
