package techsservices

import (
	"errors"

	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/entities"
	"github.com/eduahcb/hub_api_go/internal/resources/techs"
	customErrors "github.com/eduahcb/hub_api_go/pkg/errors"
	"gorm.io/gorm"
)

func Update(tech entities.Tech, db *database.Database) (techs.TechResponse, error) {
	var techResponse techs.TechResponse
  level := entities.Level{}
	
  err := db.Client.First(&level, tech.LevelID).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return techResponse, &customErrors.NotFound{Message: "Level not found"}
		}
		return techResponse, &customErrors.InternalServerError{}
	}

	result := db.Client.Where(tech.ID).Updates(&tech)

	if result.Error != nil {
		return techResponse, &customErrors.InternalServerError{}
	} else if result.RowsAffected == 0 {
		return techResponse, &customErrors.NotFound{Message: "Tech not found"}
	}

	tech.Level = level
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
