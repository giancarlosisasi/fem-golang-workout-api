package routes

import (
	"fm-api-project/internal/app"

	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux {
	r := chi.NewRouter()

	r.Group(func(r chi.Router) {
		r.Use(app.Middleware.Authenticate)

		r.Get("/workouts", app.Middleware.RequireUSer(
			app.WorkoutHandler.HandleGetAllWorkouts,
		))
		r.Get("/workouts/{id}", app.Middleware.RequireUSer(
			app.WorkoutHandler.HandleGetWorkoutById,
		))

		r.Post("/workouts", app.Middleware.RequireUSer(app.WorkoutHandler.HandleCreateWorkout))
		r.Put("/workouts/{id}", app.Middleware.RequireUSer(app.WorkoutHandler.HandleUpdateWorkoutByID))
		r.Delete("/workouts/{id}", app.Middleware.RequireUSer(app.WorkoutHandler.HandleDeleteWorkoutById))
	})

	r.Get("/health", app.HealthCheck)
	r.Post("/users", app.UserHandler.HandleRegisterUser)
	r.Post("/tokens/authentication", app.TokenHandler.HandleCreateToken)

	return r
}
