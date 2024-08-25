package router

import (
	"net/http"

	"github.com/eduahcb/hub_api_go/internal/database"
	"github.com/eduahcb/hub_api_go/internal/middlewares"
	techshandlers "github.com/eduahcb/hub_api_go/internal/resources/techs/handlers"
	usershandlers "github.com/eduahcb/hub_api_go/internal/resources/users/handlers"
	"github.com/eduahcb/hub_api_go/pkg/responses"
	"github.com/gorilla/mux"
	"github.com/redis/go-redis/v9"
	"gorm.io/gorm"

	httpSwagger "github.com/swaggo/http-swagger/v2"
)

func initRoutes(r *mux.Router, client *gorm.DB, rdb *redis.Client) {
	api := r.PathPrefix("/api/v1").Subrouter()

	db := &database.Database{
		Client: client,
		Redis:  rdb,
	}

	api.PathPrefix("/swagger").Handler(httpSwagger.WrapHandler).Methods(http.MethodGet)

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

	// profile
	api.HandleFunc("/me", middlewares.Authentication(db, usershandlers.Profile(db))).Methods(http.MethodGet)

	// signout
	api.HandleFunc("/signout", middlewares.Authentication(db, usershandlers.Signout(db))).Methods(http.MethodGet)

	//techs
	api.HandleFunc("/techs", middlewares.Authentication(db, techshandlers.GetAll(db))).Methods(http.MethodGet)
	api.HandleFunc("/techs", middlewares.Authentication(db, techshandlers.Create(db))).Methods(http.MethodPost)
	api.HandleFunc("/techs/{id}", middlewares.Authentication(db, techshandlers.GetById(db))).Methods(http.MethodGet)
	api.HandleFunc("/techs/{id}", middlewares.Authentication(db, techshandlers.Delete(db))).Methods(http.MethodDelete)
	api.HandleFunc("/techs/{id}", middlewares.Authentication(db, techshandlers.Update(db))).Methods(http.MethodPut)
}
