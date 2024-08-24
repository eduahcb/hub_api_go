package usershandlers

import (
	"net/http"

	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/entities"
	"github.com/eduahcb/hub_api_go/internal/resources/users"
	"github.com/eduahcb/hub_api_go/pkg/responses"
)

func Profile(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		userId := r.Context().Value("userId").(uint)
		user := entities.User{}

		db.Client.Preload("Module").First(&user, userId)
    
    profile := users.ProfileResponse{
      Name: user.Name,
      Contact: user.Contact,
      Bio: user.Bio,
      Module: users.ModuleResponse{
        ID: user.Module.ID,
        Name: user.Module.Name,
      },
    }

		responses.OK(w, profile)
	}
}
