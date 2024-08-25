package usershandlers

import (
	"net/http"
	"time"

	"github.com/eduahcb/hub_api_go/internal/blocklist"
	"github.com/eduahcb/hub_api_go/internal/database"
)

//	@Summary		Sigin out de usuário
//	@Description	Sigin out de usuário no qual
//	@Tags			User
//	@Accept			json
//	@Produce		json
//	@Router			/signout [get]
//	@Security		BearerAuth
//	@Header			204	{string}	token	"token"
//	@Success		204
func Signout(db *database.Database) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.Context().Value("token").(string)
		expirationTime := r.Context().Value("expirationTime").(time.Duration)

		blocklist.Add(db, token, expirationTime)

		w.WriteHeader(http.StatusNoContent)
	}
}
