package main

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (app *application) routes() *httprouter.Router {
	router := httprouter.New()

	router.NotFound = http.HandlerFunc(app.notFoundResponse)
	router.MethodNotAllowed = http.HandlerFunc(app.methodNotAllowedResponse)

	router.HandlerFunc(http.MethodGet, "/v1/healthcheck", app.healthcheckHandler)
	router.HandlerFunc(http.MethodPost, "/v1/workouts", app.createWorkoutsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/workouts/:id", app.showWorkoutsHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/workouts/:id", app.updateWorkoutsHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/workouts/:id", app.deleteWorkoutsHandler)

	return router
}
