package main

import (
	"errors"
	"fmt"
	"net/http"
	"workout.kenfan.org/internal/data"
	"workout.kenfan.org/internal/validator"
)

func (app *application) createMesocycleHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		MesocycleNum int32 `json:"mesocycleNum"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	mesocycle := &data.Mesocycle{
		MesocycleNum: input.MesocycleNum,
	}
	v := validator.New()

	if data.ValidateMesocycle(v, mesocycle); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Todo - validate mesocycle

	err = app.models.Workouts.InsertMesocycle(mesocycle)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/mesocycles/%d", mesocycle.ID))

	// Write JSON response with 201 status code, workout data in response and location header
	err = app.writeJSON(w, http.StatusAccepted, envelope{"mesocycle": mesocycle}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) showMesocycleHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	mesocycle, err := app.models.Workouts.GetMesocycle(id)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"mesocycle": mesocycle}, nil)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}
}

func (app *application) updateMesocycleHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	mesocycle, err := app.models.Workouts.GetMesocycle(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	var input struct {
		MesocycleNum int32 `json:"mesocycleNum"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	mesocycle.MesocycleNum = input.MesocycleNum

	v := validator.New()
	if data.ValidateMesocycle(v, mesocycle); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Workouts.UpdateMesocycle(mesocycle)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"mesocycle": mesocycle}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteMesocycleHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Workouts.DeleteMesocycle(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "mesocycle successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
