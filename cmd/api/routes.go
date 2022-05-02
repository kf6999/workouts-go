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

	// Workout Routes
	router.HandlerFunc(http.MethodPost, "/v1/workouts", app.createWorkoutsHandler)
	router.HandlerFunc(http.MethodGet, "/v1/workouts/:id", app.showWorkoutsHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/workouts/:id", app.updateWorkoutsHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/workouts/:id", app.deleteWorkoutsHandler)

	// Mesocycle Routes
	router.HandlerFunc(http.MethodPost, "/v1/mesocycle", app.createMesocycleHandler)
	router.HandlerFunc(http.MethodGet, "/v1/mesocycle/:id", app.showMesocycleHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/mesocycle/:id", app.updateMesocycleHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/mesocycle/:id", app.deleteMesocycleHandler)

	// Week Routes
	router.HandlerFunc(http.MethodPost, "/v1/week", app.createWeekHandler)
	router.HandlerFunc(http.MethodGet, "/v1/week/:id", app.showWeekHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/week/:id", app.updateWeekHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/week/:id", app.deleteWeekHandler)

	// Day Routes
	router.HandlerFunc(http.MethodPost, "/v1/day", app.createDayHandler)
	router.HandlerFunc(http.MethodGet, "/v1/day/:id", app.showDayHandler)
	router.HandlerFunc(http.MethodPatch, "/v1/day/:id", app.updateDayHandler)
	router.HandlerFunc(http.MethodDelete, "/v1/day/:id", app.deleteDayHandler)

	// Setup Routes
	router.HandlerFunc(http.MethodPost, "/v1/setup", app.createSetupHandler)
	return router
}
