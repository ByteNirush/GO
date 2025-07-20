package app

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ByteNirush/fem_project/internal/api"
	"github.com/ByteNirush/fem_project/internal/middleware"
	"github.com/ByteNirush/fem_project/internal/store"
	"github.com/ByteNirush/fem_project/migrations"
)

type Application struct {
	Logger 			*log.Logger
	WorkoutHandler 	*api.WorkoutHandler
	UserHandler 	*api.UserHandler
	TokenHandler 	*api.TokenHandler
	Middleware 		middleware.UserMiddleware
	DB 				*sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	err = store.MigrateFS(pgDB, migrations.FS, ".")
	if err != nil {
		panic(err)
	}

	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)
 
	workoutStore := store.NewPostgresWorkoutStore(pgDB)
	userStore := store.NewPostgresUserStore(pgDB)
	tokenStore := store.NewPostgresTokenStore(pgDB)
	middlewareHandler := middleware.UserMiddleware{UserStore: userStore}


	// our handlers will go here
	workoutHandler := api.NewWorkoutHandler(workoutStore, logger)
	userHandler := api.NewUserHandler(userStore, logger)
	tokenHandler := api.NewTokenHandler(tokenStore, userStore, logger)
	
	app := &Application{
		Logger : 		logger,
		WorkoutHandler: workoutHandler,
		UserHandler: 	userHandler,
		TokenHandler: 	tokenHandler,
		Middleware: 	middlewareHandler,
		DB: 			pgDB,
	}

	return  app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}
