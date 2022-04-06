package data

import (
	"time"
	"workout.kenfan.org/internal/validator"
)

type Workout struct {
	ID           int64     `json:"id"`
	CreatedAt    time.Time `json:"-"`
	Exercise     string    `json:"exercise"`
	ExerciseType string    `json:"exercise_type,omitempty"`
	Sets         int32     `json:"sets,omitempty"`
	Weight       int32     `json:"weight,omitempty"`
	RepGoal      int32     `json:"repGoal,omitempty"`
	Reps         []int32   `json:"reps,omitempty"`
	Soreness     int32     `json:"soreness,omitempty"`
	Pump         int32     `json:"pump,omitempty"`
	Version      int32     `json:"version"`
}

func ValidateWorkout(v *validator.Validator, workout *Workout) {
	v.Check(workout.Exercise != "", "exercise", "must be provided")
	v.Check(len(workout.Exercise) <= 500, "exercise", "must not be more than 500 bytes long")
	v.Check(workout.ExerciseType != "", "exercise type", "must be provided")
	v.Check(len(workout.ExerciseType) <= 500, "exercise type", "must not be more than 500 bytes long")

	v.Check(workout.Sets >= 0, "sets", "must be greater than or equal to 0")
	v.Check(workout.Weight >= 0, "weight", "must be greater than or equal to 0")

	v.Check(workout.RepGoal >= 0, "rep goal", "must be greater than or equal to 0")
	v.Check(workout.RepGoal <= 3, "rep goal", "must be less than or equal to 3")
	v.Check(workout.Reps != nil, "rep", "must be provided")

	v.Check(workout.Soreness >= -1, "soreness", "must be greater than or equal to -1")
	v.Check(workout.Soreness <= 1, "soreness", "must be less than or equal to 1")
	v.Check(workout.Pump >= 0, "pump", "must be less than or equal to 0")
	v.Check(workout.Pump <= 1, "pump", "must be less than or equal to 1")
}
