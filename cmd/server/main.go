package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/eduahcb/hub_api_go/config"
	_ "github.com/eduahcb/hub_api_go/docs"
	"github.com/eduahcb/hub_api_go/internal/entities"
	"github.com/eduahcb/hub_api_go/internal/router"
)

//	@title			Hub API Example
//	@version		1.0
//	@description	Primeira api em go
//	@termsOfService	http://swagger.io/terms/

//	@contact.name	API Support
//	@contact.url	http://www.swagger.io/support
//	@contact.email	support@swagger.io

//	@license.name	Apache 2.0
//	@license.url	http://www.apache.org/licenses/LICENSE-2.0.html

//	@host						localhost:8080
//	@BasePath					/api/v1
//	@securityDefinitions.apikey	BearerAuth
//	@in							header
//	@name						Authorization

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
