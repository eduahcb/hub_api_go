package techshandlers

import (
	"net/http"
	"strconv"

	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/resources/techs"
	techsservices "github.com/eduahcb/hub_api_go/internal/resources/techs/services"
	customErrors "github.com/eduahcb/hub_api_go/pkg/errors"
	"github.com/eduahcb/hub_api_go/pkg/responses"
	"github.com/gorilla/mux"
)

type GetByIdResponse struct {
	Tech techs.TechResponse `json:"tech"`
}

// @Summary		Lista tecnologia por id
// @Description	lista uma única tecnologia por id
// @Tags			Techs
// @Accept			json
// @Produce		json
// @Router		/techs/{id} [get]
// @Param			id	path	int	true	"tech id"
// @Security		BearerAuth
// @Success		200	{array}	techs.TechResponse
func GetById(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		techId, _ := strconv.ParseUint(params["id"], 10, 64)

		tech, err := techsservices.GetById(techId, db)
		if err != nil {
			switch err.(type) {
			case *customErrors.NotFound:
				responses.NotFound(w, err)
				return
			default:
				responses.InternalServerError(w, err)
				return
			}
		}

		responses.OK(w, GetByIdResponse{
			Tech: tech,
		})
	}
}
