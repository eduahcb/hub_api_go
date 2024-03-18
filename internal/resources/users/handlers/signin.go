package usershandlers

import (
	"net/http"

	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/resources/users"
	usersservices "github.com/eduahcb/hub_api_go/internal/resources/users/services"
	customErrors "github.com/eduahcb/hub_api_go/pkg/errors"
	"github.com/eduahcb/hub_api_go/pkg/responses"
	"github.com/eduahcb/hub_api_go/pkg/utils"
)

func Signin(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		signinRequest := users.SigninRequest{}

		if err := utils.ParseJSON(r, &signinRequest); err != nil {
			responses.BadRequest(w, &customErrors.BadRequest{Message: "Invalid body"})
			return
		}

		errors, err := signinRequest.Validate()
		if err != nil {
			responses.ErrorJSON(w, http.StatusBadRequest, &customErrors.BadRequest{}, errors)
			return
		}

		token, err := usersservices.CreateToken(signinRequest, db)
		if err != nil {
			switch err.(type) {
			case *customErrors.NotFound:
				responses.NotFound(w, err)
				return
			case *customErrors.Unauthorized:
				responses.Unauthorized(w, err)
				return
			default:
				responses.InternalServerError(w, err)
				return
			}
		}

		w.Header().Set("Authorization", token)
		w.WriteHeader(http.StatusNoContent)
	}
}
