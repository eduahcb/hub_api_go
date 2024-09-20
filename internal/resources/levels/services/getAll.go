package levelsservices

import (
	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/entities"
	"github.com/eduahcb/hub_api_go/internal/resources/levels"
	customErrors "github.com/eduahcb/hub_api_go/pkg/errors"
)

func GeAll(db *database.Database) ([]levels.LevelResponse, error) {
	var levelsEntity []entities.Level
	levelsResponse := []levels.LevelResponse{}

	err := db.Client.Find(&levelsEntity).Error

	if err != nil {
		return levelsResponse, &customErrors.InternalServerError{}
	}

	for _, level := range levelsEntity {
		levelsResponse = append(levelsResponse, levels.LevelResponse{
			ID:   level.ID,
			Name: level.Name,
		})
	}

	return levelsResponse, nil
}
