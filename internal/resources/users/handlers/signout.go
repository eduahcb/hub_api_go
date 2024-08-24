package usershandlers

import (
	"net/http"
	"time"

	"github.com/eduahcb/hub_api_go/internal/blocklist"
	"github.com/eduahcb/hub_api_go/internal/database"
)

func Signout(db *database.Database) http.HandlerFunc {
  return func(w http.ResponseWriter, r *http.Request) {
    token := r.Context().Value("token").(string)
    expirationTime := r.Context().Value("expirationTime").(time.Duration)
  
    blocklist.Add(db, token, expirationTime)

    w.WriteHeader(http.StatusNoContent)
  } 
}
