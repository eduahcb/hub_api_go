package techs

import (
	"github.com/eduahcb/hub_api_go/internal/entities"
	"github.com/eduahcb/hub_api_go/pkg/validate"
)

type TechRequest struct {
	Name    string `json:"name" validate:"required"`
	LevelId uint   `json:"level_id" validate:"required"`
}

func (t *TechRequest) Validate() ([]validate.ValidateErrors, error) {
	return validate.ValidateRequestBody(t)
}

func (t *TechRequest) ToEntity() entities.Tech {
	return entities.Tech{
		Name:    t.Name,
		LevelID: t.LevelId,
	}
}
