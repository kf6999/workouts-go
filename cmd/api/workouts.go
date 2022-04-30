package main

import (
	"errors"
	"fmt"
	"net/http"
	"workout.kenfan.org/internal/data"
	"workout.kenfan.org/internal/validator"
)

func (app *application) createWorkoutsHandler(w http.ResponseWriter, r *http.Request) {
	// anonymous struct to hold information in HTTP request ot be target decode destination

	var input struct {
		ExerciseName   string  `json:"exerciseName"`
		SetCount       int32   `json:"setCount"`
		Weight         int32   `json:"weight"`
		RepGoal        int32   `json:"repGoal"`
		RepResults     []int32 `json:"repResults"`
		SorenessRating int32   `json:"sorenessRating"`
		PumpRating     int32   `json:"pumpRating"`
		Day_id         int32   `json:"day_id"`
	}

	// json.Decoder instance to read from request body and decode into input struct
	// Pass pointer to input struct

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}
	// Copy values from input struct to workout struct
	workout := &data.Workout{
		ExerciseName:   input.ExerciseName,
		SetCount:       input.SetCount,
		Weight:         input.Weight,
		RepGoal:        input.RepGoal,
		RepResults:     input.RepResults,
		SorenessRating: input.SorenessRating,
		PumpRating:     input.PumpRating,
		Day_id:         input.Day_id,
	}
	// New validator instance
	v := validator.New()

	// Call the ValidateWorkout() function and return response if check fails
	if data.ValidateWorkout(v, workout); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Call Insert() method on workout model to create record in database
	err = app.models.Workouts.Insert(workout)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
	// Include location header to inform client which URL to find created record
	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/workouts/%d", workout.ID))

	// Write JSON response with 201 status code, workout data in response and location header
	err = app.writeJSON(w, http.StatusAccepted, envelope{"workout": workout}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}
}

func (app *application) showWorkoutsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Call Get() method to fetch data for specific workout
	// Use errors.Is() to check if data is returned
	workout, err := app.models.Workouts.Get(id)
	if err != nil {
		if errors.Is(err, data.ErrRecordNotFound) {
			app.notFoundResponse(w, r)
			return
		}
		app.serverErrorResponse(w, r, err)
		return
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"workout": workout}, nil)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}
}

func (app *application) updateWorkoutsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract workout ID from URL
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Get existing workout from database, send 404 if not found
	workout, err := app.models.Workouts.Get(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}

	// Declare input struct to hold expected values in HTTP request
	var input struct {
		ExerciseName   *string `json:"exerciseName"`
		Weight         *int32  `json:"weight"`
		RepResults     []int32 `json:"repResults"`
		SorenessRating *int32  `json:"sorenessRating"`
		PumpRating     *int32  `json:"pumpRating"`
	}

	// json.Decoder instance to read from request body and decode into input struct
	err = app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	// Copy values from input struct to workout struct
	if input.ExerciseName != nil {
		workout.ExerciseName = *input.ExerciseName
	}
	if input.Weight != nil {
		workout.Weight = *input.Weight
	}
	if input.RepResults != nil {
		workout.RepResults = input.RepResults
	}
	if input.SorenessRating != nil {
		workout.SorenessRating = *input.SorenessRating
	}
	if input.PumpRating != nil {
		workout.PumpRating = *input.PumpRating
	}

	// New validator instance
	v := validator.New()

	if data.ValidateWorkout(v, workout); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	// Pass updated workout record to Update() method to update record in database
	err = app.models.Workouts.Update(workout)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return
	}

	// Write JSON response with 200 status code, workout data in response
	err = app.writeJSON(w, http.StatusOK, envelope{"workout": workout}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) deleteWorkoutsHandler(w http.ResponseWriter, r *http.Request) {
	// Extract workout ID from URL
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}
	// Call Delete() method to delete record from database
	err = app.models.Workouts.Delete(id)
	if err != nil {
		switch {
		case errors.Is(err, data.ErrRecordNotFound):
			app.notFoundResponse(w, r)
		default:
			app.serverErrorResponse(w, r, err)
		}
		return
	}
	// Write JSON response with 204 status code
	err = app.writeJSON(w, http.StatusOK, envelope{"message": "workout successfully deleted"}, nil)
	if err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
