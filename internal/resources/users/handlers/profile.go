package usershandlers

import (
	"net/http"

	_ "github.com/eduahcb/hub_api_go/docs"
	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/entities"
	"github.com/eduahcb/hub_api_go/internal/resources/users"
	"github.com/eduahcb/hub_api_go/pkg/responses"
)

//	@Summary		Informações do usuário
//	@Description	Traz o informações essências do usuário
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Router			/me [get]
//	@Security		BearerAuth
//	@Success		200	{object}	users.ProfileResponse
func Profile(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("userId").(uint)
		user := entities.User{}

		db.Client.Preload("Module").First(&user, userId)

		profile := users.ProfileResponse{
			ID:      user.ID,
			Name:    user.Name,
			Contact: user.Contact,
			Bio:     user.Bio,
			Module: users.ModuleResponse{
				ID:   user.Module.ID,
				Name: user.Module.Name,
			},
		}

		responses.OK(w, profile)
	}
}
