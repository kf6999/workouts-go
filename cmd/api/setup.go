package main

import (
	"fmt"
	"net/http"
	"workout.kenfan.org/internal/data"
)

func (app *application) createSetupHandler(w http.ResponseWriter, r *http.Request) {
	var input struct {
		MesocycleNum     int `json:"mesocycleNum"`
		CurrentMesocycle int `json:"currentMesocycle"`
		WeekNum          int `json:"weekNum"`
		CurrentWeek      int `json:"currentWeek"`
		DayNum           int `json:"dayNum"`
		WorkoutNum       int `json:"workoutNum"`
	}

	err := app.readJSON(w, r, &input)
	if err != nil {
		app.errorResponse(w, r, http.StatusBadRequest, err.Error())
		return
	}

	setup := &data.Setup{
		MesocycleNum:     input.MesocycleNum,
		CurrentMesocycle: input.CurrentMesocycle,
		WeekNum:          input.WeekNum,
		CurrentWeek:      input.CurrentWeek,
		DayNum:           input.DayNum,
		WorkoutNum:       input.WorkoutNum,
	}
	// Add validation

	headers := make(http.Header)
	headers.Set("Location", fmt.Sprintf("/workouts/%d", setup.MesocycleNum))

	err = app.models.Workouts.InsertSetup(setup)
	if err != nil {
		app.errorResponse(w, r, http.StatusInternalServerError, err.Error())
		return
	}

	err = app.writeJSON(w, http.StatusAccepted, envelope{"setup": setup}, headers)
	if err != nil {
		app.serverErrorResponse(w, r, err)
		return

	}
}
