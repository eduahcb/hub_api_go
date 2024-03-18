package router

import (
	"net/http"

	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/middlewares"
	techshandlers "github.com/eduahcb/hub_api_go/internal/resources/techs/handlers"
	usershandlers "github.com/eduahcb/hub_api_go/internal/resources/users/handlers"
	"github.com/eduahcb/hub_api_go/pkg/responses"
	"github.com/gorilla/mux"
	"gorm.io/gorm"
)

func initRoutes(r *mux.Router, client *gorm.DB) {
	api := r.PathPrefix("/api/v1").Subrouter()

	db := &database.Database{
		Client: client,
	}

	api.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		type response struct {
			Message string
		}

		responses.OK(w, response{Message: "welcome to hub api!!!"})
	}).Methods(http.MethodGet)

	// public routes
	api.HandleFunc("/signin", usershandlers.Signin(db)).Methods(http.MethodPost)
	api.HandleFunc("/signup", usershandlers.Signup(db)).Methods(http.MethodPost)

	// private routes
	api.HandleFunc("/techs", middlewares.Authentication(techshandlers.GetAll(db))).Methods(http.MethodGet)
	api.HandleFunc("/techs", middlewares.Authentication(techshandlers.Create(db))).Methods(http.MethodPost)
	api.HandleFunc("/techs/{id}", middlewares.Authentication(techshandlers.GetById(db))).Methods(http.MethodGet)
	api.HandleFunc("/techs/{id}", middlewares.Authentication(techshandlers.Delete(db))).Methods(http.MethodDelete)
	api.HandleFunc("/techs/{id}", middlewares.Authentication(techshandlers.Update(db))).Methods(http.MethodPut)
}
