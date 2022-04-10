package data

import (
	"database/sql"
	"errors"
	"github.com/lib/pq"
	"time"
	"workout.kenfan.org/internal/validator"
)

// Workout is a struct that holds the data for a workout
type Workout struct {
	ID             int64     `json:"id"`
	Created_at     time.Time `json:"-"`
	ExerciseName   string    `json:"exerciseName"`
	SetCount       int32     `json:"setCount,omitempty"`
	Weight         int32     `json:"weight,omitempty"`
	RepGoal        int32     `json:"repGoal,omitempty"`
	RepResults     []int32   `json:"repResults,omitempty"`
	SorenessRating int32     `json:"sorenessRating,omitempty"`
	PumpRating     int32     `json:"pumpRating,omitempty"`
	Day_id         int32     `json:"day_id"`
}

type Week struct {
	ID         int64     `json:"id"`
	Created_at time.Time `json:"-"`
	WeekNum    int32     `json:"weekNum"`
}

type Day struct {
	ID         int64     `json:"id"`
	Created_at time.Time `json:"-"`
	DayNum     int32     `json:"dayNum"`
}

type Exercise struct {
	ExeciseName   string    `json:"exerciseName"`
	Created_at    time.Time `json:"-"`
	ExerciseFocus string    `json:"exerciseFocus"`
	VideoURL      string    `json:"videoURL"`
	TenRM         int32     `json:"tenRM"`
}

// WorkoutModel struct type which wraps a sql.DB conection pool.
type WorkoutModel struct {
	DB *sql.DB
}

// Insert Insert() accepts a pointer to Workout struct which has data for new record
func (w WorkoutModel) Insert(workout *Workout) error {
	query := `INSERT INTO workout (exerciseName, setCount, weight, repGoal, repResults,sorenessRating, pumpRating, day_id) VALUES ($1, $2, $3, $4, $5, $6, $7, $8) RETURNING id,created_at`

	// Create args slice containing values for placeholder parameters from workout struct
	args := []interface{}{workout.ExerciseName, workout.SetCount, workout.Weight, workout.RepGoal, pq.Array(workout.RepResults), workout.SorenessRating, workout.PumpRating, workout.Day_id}

	// Use QueryRow() to execute SQL query on connection pool by passing in the args as variadic values
	return w.DB.QueryRow(query, args...).Scan(&workout.ID, &workout.Created_at)
}

func (w WorkoutModel) Get(id int64) (*Workout, error) {
	if id < 1 {
		return nil, ErrRecordNotFound
	}

	query := `SELECT id, created_at, exerciseName, setCount, weight, repGoal, repResults, sorenessRating, pumpRating, day_id FROM workout WHERE id = $1`

	// Workout struct to hold data from query
	var workout Workout

	// Execute query using QueryRow() method by passing in ID value as placeholder parameter
	// then scanning response data into fields of workout struct
	err := w.DB.QueryRow(query, id).Scan(
		&workout.ID,
		&workout.Created_at,
		&workout.ExerciseName,
		&workout.SetCount,
		&workout.Weight,
		&workout.RepGoal,
		pq.Array(&workout.RepResults),
		&workout.SorenessRating,
		&workout.PumpRating,
		&workout.Day_id,
	)
	// If no workout found, Scan() returns sql.ErrNoRows which will return ErrRecordNotFound
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &workout, nil
}

func (w WorkoutModel) Update(workout *Workout) error {
	query := `UPDATE workout SET weight = $1, repResults = $2, sorenessRating = $3, pumpRating = $4 WHERE id = $5 returning day_id`

	// Create args slice containing values for placeholder parameters from workout struct
	args := []interface{}{workout.Weight, pq.Array(workout.RepResults), workout.SorenessRating, workout.PumpRating, workout.ID}

	// Use QueryRow() to execute query by passing in the args slice as variadic values and scanning hte new version value into workout struct
	return w.DB.QueryRow(query, args...).Scan(&workout.Day_id)
}

func (w WorkoutModel) Delete(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}

	query := `DELETE FROM workout WHERE id = $1`

	// Execute query by passing in ID value as placeholder parameter
	result, err := w.DB.Exec(query, id)
	if err != nil {
		return err
	}

	// Check if any rows were affected by query
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}

	// If no rows were affected, return ErrRecordNotFound
	if rowsAffected == 0 {
		return ErrRecordNotFound
	}
	return nil
}

func ValidateWorkout(v *validator.Validator, workout *Workout) {
	v.Check(workout.ExerciseName != "", "exerciseName", "must be provided")
	v.Check(len(workout.ExerciseName) <= 500, "exercise", "must not be more than 500 bytes long")

	v.Check(workout.SetCount >= 0, "sets", "must be greater than or equal to 0")
	v.Check(workout.Weight >= 0, "weight", "must be greater than or equal to 0")

	v.Check(workout.RepGoal >= 0, "rep goal", "must be greater than or equal to 0")
	v.Check(workout.RepGoal <= 3, "rep goal", "must be less than or equal to 3")
	v.Check(workout.RepResults != nil, "rep", "must be provided")

	v.Check(workout.SorenessRating >= -1, "soreness", "must be greater than or equal to -1")
	v.Check(workout.SorenessRating <= 1, "soreness", "must be less than or equal to 1")
	v.Check(workout.PumpRating >= 0, "pump", "must be less than or equal to 0")
	v.Check(workout.PumpRating <= 1, "pump", "must be less than or equal to 1")
}
