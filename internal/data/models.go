package data

import (
	"database/sql"
	"errors"
)

// ErrRecordNotFound Custom ErrRecordNotFound returned from Get() if record not found
var (
	ErrRecordNotFound = errors.New("record not found")
)

// Models Models struct to wrap WorkoutModel
type Models struct {
	Workouts WorkoutModel
}

// NewModels method to return a Model struct with initialized WorkoutModel
func NewModels(db *sql.DB) Models {
	return Models{
		Workouts: WorkoutModel{
			DB: db,
		},
	}
}
