package levelshandlers

import (
	"net/http"

	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/resources/levels"
	levelsservices "github.com/eduahcb/hub_api_go/internal/resources/levels/services"
	"github.com/eduahcb/hub_api_go/pkg/responses"
)

type GetAllResponse struct {
	Levels []levels.LevelResponse `json:"levels"`
}

// @Summary		Lista todos os levels
// @Description	lista todos os levels disponiveis
// @Tags			Levels
// @Accept			json
// @Produce		json
// @Router			/levels [get]
// @Success		200	{array}	levels.LevelResponse
func GetAll(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		levelResponse, err := levelsservices.GeAll(db)
		if err != nil {
			responses.InternalServerError(w, err)
		}

		responses.OK(w, GetAllResponse{
			Levels: levelResponse,
		})
	}
}
