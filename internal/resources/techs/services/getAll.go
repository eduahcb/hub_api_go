package techsservices

import (
	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/entities"
	"github.com/eduahcb/hub_api_go/internal/resources/techs"
	customErrors "github.com/eduahcb/hub_api_go/pkg/errors"
)

func GetAll(userId uint, db *database.Database) ([]techs.TechResponse, error) {
	var technologies []entities.Tech
	techsResponse := []techs.TechResponse{}

	err := db.Client.Preload("Level").Where("user_id = ?", userId).Find(&technologies).Error
	if err != nil {
		return techsResponse, &customErrors.InternalServerError{}
	}

	for _, tech := range technologies {
		techsResponse = append(techsResponse, techs.TechResponse{
			ID:   tech.ID,
			Name: tech.Name,
			Level: techs.LevelResponse{
				ID:   tech.Level.ID,
				Name: tech.Level.Name,
			},
		})
	}

	return techsResponse, nil
}
