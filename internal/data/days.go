package data

import (
	"database/sql"
	"errors"
	"workout.kenfan.org/internal/validator"
)

type Day struct {
	ID         int64  `json:"id"`
	Created_at string `json:"created_at"`
	DayNum     int32  `json:"day_num"`
	Week_id    int64  `json:"week_id"`
}

func (w WorkoutModel) InsertDay(day *Day) error {
	query := `INSERT INTO day (daynum, week_id) VALUES ($1, $2) RETURNING id, created_at`
	args := []interface{}{day.DayNum, day.Week_id}

	return w.DB.QueryRow(query, args...).Scan(&day.ID, &day.Created_at)
}

func (w WorkoutModel) GetDay(id int64) (*Day, error) {
	query := `SELECT id, created_at, daynum, week_id FROM day WHERE id = $1`

	var day Day

	err := w.DB.QueryRow(query, id).Scan(&day.ID, &day.Created_at, &day.DayNum, &day.Week_id)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &day, nil
}

func (w WorkoutModel) UpdateDay(day *Day) error {
	query := `UPDATE day SET daynum = $1 WHERE id = $2 returning  id, created_at, week_id`
	args := []interface{}{day.DayNum, day.ID}

	return w.DB.QueryRow(query, args...).Scan(&day.ID, &day.Created_at, &day.Week_id)
}

func (w WorkoutModel) DeleteDay(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `DELETE FROM day WHERE id = $1`

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

func ValidateDay(v *validator.Validator, day *Day) {
	v.Check(day.DayNum >= 0, "dayNum", "must be greater than 0")
	v.Check(day.Week_id >= 0, "week_id", "must be greater than 0")
}
