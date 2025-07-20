package routes

import (
	
	"github.com/ByteNirush/fem_project/internal/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux{
	r := chi.NewRouter()
	
	r.Get("/health", app.HealthCheck)
	r.Get("/workouts/{id}", app.WorkoutHandler.HandlerGetWorkoutByID)

	r.Post("/workouts", app.WorkoutHandler.HandleCreateWorkout)
	r.Put("/workouts/{id}", app.WorkoutHandler.HandleUpdateWorkoutByID)
	r.Delete("/workouts/{id}", app.WorkoutHandler.HandleDeleteWorkoutByID)

	r.Post("/users/register", app.UserHandler.HandleRegisterUser)
	r.Post("/tokens/auth", app.TokenHandler.HandleCreateToken)
	return r
}