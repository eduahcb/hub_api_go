package moduleshandlers

import (
	"net/http"

	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/resources/modules"
	modulesservices "github.com/eduahcb/hub_api_go/internal/resources/modules/services"
	"github.com/eduahcb/hub_api_go/pkg/responses"
)

type GetAllResponse struct {
	Modules []modules.ModuleResponse `json:"modules"`
}

// @Summary		Lista todas os módulos
// @Description	lista todas as módulos disponiveis
// @Tags			Modules
// @Accept			json
// @Produce		json
// @Router			/modules [get]
// @Success		200	{array}	modules.ModuleResponse
func GetAll(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		modulesResponse, err := modulesservices.GeAll(db)
		if err != nil {
			responses.InternalServerError(w, err)
		}

		responses.OK(w, GetAllResponse{
			Modules: modulesResponse,
		})
	}
}
