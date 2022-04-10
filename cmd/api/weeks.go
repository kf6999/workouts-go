package main

import (
	"errors"
	"fmt"
	"net/http"
	"workout.kenfan.org/internal/data"
	"workout.kenfan.org/internal/validator"
)

func (app *application) createWeekHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		WeekNum      int32 `json:"weekNum"`
		Mesocycle_id int64 `json:"mesocycle_id"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	week := &data.Week{
		WeekNum:      input.WeekNum,
		Mesocycle_id: input.Mesocycle_id,
	}

	v := validator.New()

	if data.ValidateWeek(v, week); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Workouts.InsertWeek(week)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/week/%d", week.ID))

	err = app.writeJSON(w, http.StatusAccepted, envelope{"week": week}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) showWeekHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	week, err := app.models.Workouts.GetWeek(id)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"week": week}, nil)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}
}

func (app *application) updateWeekHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	week, err := app.models.Workouts.GetWeek(id)
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
		WeekNum int32 `json:"weekNum"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	week.WeekNum = input.WeekNum

	v := validator.New()

	if data.ValidateWeek(v, week); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Workouts.UpdateWeek(week)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"week": week}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteWeekHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Workouts.DeleteWeek(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "week successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
