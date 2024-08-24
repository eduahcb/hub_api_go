package techsservices

import (
	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/entities"
	"github.com/eduahcb/hub_api_go/internal/resources/techs"
	customErrors "github.com/eduahcb/hub_api_go/pkg/errors"
	"gorm.io/gorm/clause"
)

func Delete(techId uint64, db *database.Database) (techs.TechResponse, error) {
	tech := entities.Tech{}
	var techResponse techs.TechResponse

	result := db.Client.Clauses(clause.Returning{}).Delete(&tech, techId)
	if result.Error != nil {
		return techResponse, &customErrors.InternalServerError{}
	} else if result.RowsAffected == 0 {
		return techResponse, &customErrors.NotFound{Message: "Tech not found"}
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
