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
