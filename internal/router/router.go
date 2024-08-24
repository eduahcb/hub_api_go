package router

import (
	"github.com/eduahcb/hub_api_go/config"
	"github.com/gorilla/mux"
)

func Init() *mux.Router {
	r := mux.NewRouter()
	client := config.GetDBClient()
  redis := config.GetRedisClient()

	initRoutes(r, client, redis)

	return r
}
