package app

import (
	"fmt"
	"log"
	"os"
	"net/http"
	"database/sql"

	"github.com/ByteNirush/fem_project/internal/api" 
	"github.com/ByteNirush/fem_project/internal/store"
	
)

type Application struct {
	Logger 			*log.Logger
	WorkoutHandler 	*api.WorkoutHandler
	DB 				*sql.DB
}

func NewApplication() (*Application, error) {
	pgDB, err := store.Open()
	if err != nil {
		return nil, err
	}

	logger := log.New(os.Stdout, "", log.Ldate | log.Ltime)

	// our handlers will go here
	workoutHandler := api.NewWorkoutHandler()
	
	app := &Application{
		Logger : 		logger,
		WorkoutHandler: workoutHandler,
		DB: 			pgDB,
	}

	return  app, nil
}

func (a *Application) HealthCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Status is available\n")
}
