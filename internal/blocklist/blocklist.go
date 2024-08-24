package blocklist

import (
	"context"
	"time"

	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/pkg/security"
)

func Add(db *database.Database, token string, duration time.Duration) {
  ctx := context.Background()

  tokenHash := security.GenerateTokenHash(token)
  
  err := db.Redis.Set(ctx, string(tokenHash), "", duration).Err()

  if(err != nil) {
    panic(err)
  }
}

func ContainsTokenHash(db *database.Database, token string) bool {
  ctx := context.Background()

  tokenHash := security.GenerateTokenHash(token)
  
  result, _ := db.Redis.Exists(ctx, string(tokenHash)).Result()

  return result == 1
}
