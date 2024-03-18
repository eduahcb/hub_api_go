package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eduahcb/hub_api_go/config"
	"github.com/eduahcb/hub_api_go/internal/entities"
	"github.com/eduahcb/hub_api_go/internal/router"
)

func main() {
	entities := []interface{}{
		&entities.User{},
		&entities.Tech{},
		&entities.Module{},
		&entities.Level{},
	}

	config.SetEntities(entities)

	err := config.Init()
	if err != nil {
		log.Fatal("something went wrong!!! ", err)
	}

	r := router.Init()

	fmt.Printf("server is running on port: %s", config.Envs.Port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%s", config.Envs.Port), r))
}
