package middlewares

import (
	"context"
	"net/http"
	"strings"

	"github.com/eduahcb/hub_api_go/config"
	"github.com/eduahcb/hub_api_go/internal/blocklist"
	"github.com/eduahcb/hub_api_go/internal/database"
	customErrors "github.com/eduahcb/hub_api_go/pkg/errors"
	"github.com/eduahcb/hub_api_go/pkg/responses"
	"github.com/eduahcb/hub_api_go/pkg/security"
)

func Authentication(db *database.Database ,next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		authorization := r.Header.Get("Authorization")

		if authorization == "" {
			responses.Unauthorized(w, &customErrors.Unauthorized{Message: "The authorization header was not provided"})
      return
		}

		if !strings.HasPrefix(authorization, "Bearer ") {
			responses.Unauthorized(w, &customErrors.Unauthorized{Message: "The authorization header does not contain a Bearer token"})
      return
		}

    token := strings.Replace(authorization, "Bearer ", "", 1)

    tokenExists := blocklist.ContainsTokenHash(db, token)

    if(tokenExists) {
      responses.Unauthorized(w, &customErrors.Unauthorized{Message: "Invalid token by logout"})
      return
    }

    userId, expirationTime, err  := security.ValidateToken(token, config.Envs.SecretKey)
    
    if err != nil {
      responses.Unauthorized(w, err)
      return
    }
    
    ctx := r.Context()
    ctx = context.WithValue(ctx, "userId", userId)
    ctx = context.WithValue(ctx, "token", token)
    ctx = context.WithValue(ctx, "expirationTime", expirationTime)

		next(w, r.WithContext(ctx))
	}
}
