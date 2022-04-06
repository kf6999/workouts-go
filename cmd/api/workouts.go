package main

import (
	"fmt"
	"net/http"
	"time"
	"workout.kenfan.org/internal/data"
	"workout.kenfan.org/internal/validator"
)

func (app *application) createWorkoutsHandler(w http.ResponseWriter, r *http.Request) {
	// anonymous struct to hold information in HTTP request ot be target decode destination

	var input struct {
		Exercise     string  `json:"exercise"`
		ExerciseType string  `json:"exerciseType"`
		Sets         int32   `json:"sets"`
		Weight       int32   `json:"weight"`
		RepGoal      int32   `json:"repGoal"`
		Reps         []int32 `json:"reps"`
		Soreness     int32   `json:"soreness"`
		Pump         int32   `json:"pump"`
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
		Exercise:     input.Exercise,
		ExerciseType: input.ExerciseType,
		Sets:         input.Sets,
		Weight:       input.Weight,
		RepGoal:      input.RepGoal,
		Reps:         input.Reps,
		Soreness:     input.Soreness,
		Pump:         input.Pump,
	}
	// New validator instance
	v := validator.New()

	// Call the ValidateWorkout() function and return response if check fails
	if data.ValidateWorkout(v, workout); !v.Valid() {
		app.failedValidationResponse(w, r, v.Errors)
		return
	}

	fmt.Fprintf(w, "%+v\n", input)
}

func (app *application) showWorkoutsHandler(w http.ResponseWriter, r *http.Request) {
	id, err := app.readIDParam(r)
	if err != nil {
		app.notFoundResponse(w, r)
		return
	}

	workout := data.Workout{
		ID:           id,
		CreatedAt:    time.Now(),
		Exercise:     "Bench Press",
		ExerciseType: "Straight Sets",
		Sets:         3,
		Weight:       200,
		RepGoal:      3,
		Reps:         []int32{1, 2, 3},
		Soreness:     0,
		Pump:         0,
		Version:      1,
	}

	err = app.writeJSON(w, http.StatusOK, envelope{"workout": workout}, nil)
	if err != nil {
		app.badRequestResponse(w, r, err)
	}
	fmt.Fprintf(w, "show workout %d\n", id)

}
