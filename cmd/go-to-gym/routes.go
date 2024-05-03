package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() http.Handler {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)

	router.HandlerFunc(http.MethodGet, "/v1/workouts", app.requirePermission("workouts:read", app.listWorkoutsHandler))
	router.HandlerFunc(http.MethodPost, "/v1/workouts", app.requirePermission("workouts:write", app.createWorkoutHandler))
	router.HandlerFunc(http.MethodGet, "/v1/workouts/:id", app.requirePermission("workouts:read", app.showWorkoutHandler))
	//router.HandlerFunc(http.MethodGet, "/v1/workouts/:id", app.showWorkoutHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/workouts/:id", app.requirePermission("workouts:write", app.updateWorkoutHandler))
	//router.HandlerFunc(http.MethodPatch, "/v1/workouts/:id", app.updateWorkoutHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/workouts/:id", app.requirePermission("workouts:write", app.deleteWorkoutHandler))

	router.HandlerFunc(http.MethodPost, "/v1/users", app.registerUserHandler)
	router.HandlerFunc(http.MethodPut, "/v1/users/activated", app.activateUserHandler)

	router.HandlerFunc(http.MethodPost, "/v1/tokens/authentication", app.createAuthenticationTokenHandler)

	return app.recoverPanic(app.rateLimit(app.authenticate(router)))
}
