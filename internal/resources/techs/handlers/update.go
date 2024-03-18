package techshandlers

import (
	"net/http"
	"strconv"

	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/resources/techs"
	techsservices "github.com/eduahcb/hub_api_go/internal/resources/techs/services"
	customErrors "github.com/eduahcb/hub_api_go/pkg/errors"
	"github.com/eduahcb/hub_api_go/pkg/responses"
	"github.com/eduahcb/hub_api_go/pkg/utils"
	"github.com/gorilla/mux"
)

type UpdateResponse struct {
	Tech techs.TechResponse `json:"tech"`
}

func Update(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var techRequest techs.TechRequest
		
    params := mux.Vars(r)

		techId, _ := strconv.ParseUint(params["id"], 10, 64)

		if err := utils.ParseJSON(r, &techRequest); err != nil {
			responses.BadRequest(w, &customErrors.BadRequest{Message: "Invalid body"})
			return
		}

		errors, err := techRequest.Validate()
		if err != nil {
			responses.ErrorJSON(w, http.StatusBadRequest, &customErrors.BadRequest{}, errors)
			return
		}

		tech := techRequest.ToEntity()
		tech.ID = uint(techId)

		techResponse, err := techsservices.Update(tech, db)
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

		responses.OK(w, UpdateResponse{
			Tech: techResponse,
		})
	}
}
