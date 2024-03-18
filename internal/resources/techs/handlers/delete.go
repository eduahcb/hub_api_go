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

type DeleteResponse struct {
	Tech techs.TechResponse `json:"tech"`
}

func Delete(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		params := mux.Vars(r)

		techId, _ := strconv.ParseUint(params["id"], 10, 64)

		tech, err := techsservices.Delete(techId, db)
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

		responses.OK(w, DeleteResponse{
			Tech: tech,
		})
	}
}
