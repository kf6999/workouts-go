package main

import (
	"errors"
	"fmt"
	"net/http"
	"workout.kenfan.org/internal/data"
	"workout.kenfan.org/internal/validator"
)

func (app *application) createDayHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		DayNum  int32 `json:"dayNum"`
		Week_id int64 `json:"week_id"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	day := &data.Day{
		DayNum:  input.DayNum,
		Week_id: input.Week_id,
	}

	v := validator.New()

	if data.ValidateDay(v, day); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Workouts.InsertDay(day)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/week/%d", day.ID))

	err = app.writeJSON(w, http.StatusAccepted, envelope{"week": day}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) showDayHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	day, err := app.models.Workouts.GetDay(id)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}
	err = app.writeJSON(w, http.StatusOK, envelope{"day": day}, nil)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}
}

func (app *application) updateDayHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	day, err := app.models.Workouts.GetDay(id)
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
		DayNum int32 `json:"dayNum"`
	}

	err = app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	day.DayNum = input.DayNum

	v := validator.New()

	if data.ValidateDay(v, day); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	err = app.models.Workouts.UpdateDay(day)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"day": day}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteDayHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	err = app.models.Workouts.DeleteDay(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"message": "day successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
