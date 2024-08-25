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

//	@Summary		Sigin up de usuário
//	@Description	Usuário se cadastra
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Param			request	body	users.SignupRequest	true	"Requisição para se cadastrar"
//	@Router			/signup [post]
//	@Header			204	{string}	token	"token"
//	@Success		204
func Signup(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var singupRequest users.SignupRequest

		if err := utils.ParseJSON(r, &singupRequest); err != nil {
			responses.BadRequest(w, &customErrors.BadRequest{Message: "Invalid body"})
		}

		errors, err := singupRequest.Validate()
		if err != nil {
			responses.ErrorJSON(w, http.StatusBadRequest, &customErrors.BadRequest{}, errors)
			return
		}

		user := singupRequest.ToEntity()
		token, err := usersservices.CreateUser(user, db)
		if err != nil {
			switch err.(type) {
			case *customErrors.NotFound:
				responses.NotFound(w, err)
				return
			case *customErrors.BadRequest:
				responses.BadRequest(w, err)
			default:
				responses.InternalServerError(w, err)
			}
		}

		w.Header().Set("Authorization", token)
		w.WriteHeader(http.StatusNoContent)
	}
}
