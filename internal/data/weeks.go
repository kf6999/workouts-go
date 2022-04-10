package data

import (
	"database/sql"
	"errors"
	"workout.kenfan.org/internal/validator"
)

type Week struct {
	ID           int64  `json:"id"`
	Created_at   string `json:"created_at"`
	WeekNum      int32  `json:"weekNum"`
	Mesocycle_id int64  `json:"mesocycle_id"`
}

func (w WorkoutModel) InsertWeek(week *Week) error {
	query := `INSERT INTO week (weeknum, mesocycle_id) values ($1,$2) returning id,created_at`
	args := []interface{}{week.WeekNum, week.Mesocycle_id}

	return w.DB.QueryRow(query, args...).Scan(&week.ID, &week.Created_at)
}

func (w WorkoutModel) GetWeek(id int64) (*Week, error) {
	query := `SELECT id,created_at,weeknum,mesocycle_id FROM week WHERE id = $1`
	// Struct to hold result
	var week Week
	// Execute query with QueryRow, pass in ID then scan response into mesocycle struct
	err := w.DB.QueryRow(query, id).Scan(&week.ID, &week.Created_at, &week.WeekNum, &week.Mesocycle_id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &week, nil
}

func (w WorkoutModel) UpdateWeek(week *Week) error {
	query := `UPDATE week SET weeknum = $1 WHERE id = $2 returning id,created_at,mesocycle_id`

	args := []interface{}{week.WeekNum, week.ID}

	return w.DB.QueryRow(query, args...).Scan(&week.ID, &week.Created_at, &week.Mesocycle_id)
}

func (w WorkoutModel) DeleteWeek(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `DELETE FROM week WHERE id = $1`

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

func ValidateWeek(v *validator.Validator, week *Week) {
	v.Check(week.WeekNum >= 0, "weekNum", "must be greater than 0")
	v.Check(week.Mesocycle_id >= 0, "mesocycle_id", "must be greater than 0")
}
