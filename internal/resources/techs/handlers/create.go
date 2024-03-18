package techshandlers

import (
	"net/http"

	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/resources/techs"
	techsservices "github.com/eduahcb/hub_api_go/internal/resources/techs/services"
	customErrors "github.com/eduahcb/hub_api_go/pkg/errors"
	"github.com/eduahcb/hub_api_go/pkg/responses"
	"github.com/eduahcb/hub_api_go/pkg/utils"
)

type CreateResponse struct {
	Tech techs.TechResponse `json:"tech"`
}

func Create(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("userId").(uint)

		var techRequest techs.TechRequest

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
    tech.UserID = userId

		techResponse, err := techsservices.Create(tech, db)
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

		responses.JSON(w, http.StatusCreated, CreateResponse{
			Tech: techResponse,
		})
	}
}
