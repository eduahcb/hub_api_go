package techshandlers

import (
	"net/http"

	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/resources/techs"
	techsservices "github.com/eduahcb/hub_api_go/internal/resources/techs/services"
	"github.com/eduahcb/hub_api_go/pkg/responses"
)

type GetAllResponse struct {
	Techs []techs.TechResponse `json:"techs"`
}

//	@Summary		Lista todas as tecnologias
//	@Description	lista todas as tecnologias do usuário
//	@Tags			Techs
//	@Accept			json
//	@Produce		json
//	@Router			/techs [get]
//	@Security		BearerAuth
//	@Success		200	{array}	techs.TechResponse
func GetAll(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("userId").(uint)

		techsResponse, err := techsservices.GetAll(userId, db)
		if err != nil {
			responses.InternalServerError(w, err)
		}

		responses.OK(w, GetAllResponse{
			Techs: techsResponse,
		})
	}
}
