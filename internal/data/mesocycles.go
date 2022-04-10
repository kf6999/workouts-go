package data

import (
	"database/sql"
	"errors"
	"time"
	"workout.kenfan.org/internal/validator"
)

type Mesocycle struct {
	ID           int64     `json:"id"`
	Created_at   time.Time `json:"-"`
	MesocycleNum int32     `json:"mesocycleNum"`
}

func (w WorkoutModel) InsertMesocycle(mesocycle *Mesocycle) error {
	query := `INSERT INTO mesocycle (mesocyclenum) values ($1) returning id,created_at`
	args := []interface{}{mesocycle.MesocycleNum}

	return w.DB.QueryRow(query, args...).Scan(&mesocycle.ID, &mesocycle.Created_at)
}

func (w WorkoutModel) GetMesocycle(id int64) (*Mesocycle, error) {
	query := `SELECT id,created_at,mesocyclenum FROM mesocycle WHERE id = $1`
	// Struct to hold result
	var mesocycle Mesocycle
	// Execute query with QueryRow, pass in ID then scan response into mesocycle struct
	err := w.DB.QueryRow(query, id).Scan(&mesocycle.ID, &mesocycle.Created_at, &mesocycle.MesocycleNum)
	if err != nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, ErrRecordNotFound
		default:
			return nil, err
		}
	}

	return &mesocycle, nil
}

func (w WorkoutModel) UpdateMesocycle(mesocycle *Mesocycle) error {
	query := `UPDATE mesocycle SET mesocyclenum = $1 WHERE id = $2 returning id,created_at`

	args := []interface{}{mesocycle.MesocycleNum, mesocycle.ID}

	return w.DB.QueryRow(query, args...).Scan(&mesocycle.ID, &mesocycle.Created_at)
}

func (w WorkoutModel) DeleteMesocycle(id int64) error {
	if id < 1 {
		return ErrRecordNotFound
	}
	query := `DELETE FROM mesocycle WHERE id = $1`

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
func ValidateMesocycle(v *validator.Validator, mesocycle *Mesocycle) {
	v.Check(mesocycle.MesocycleNum >= 0, "mesocycleNum", "must be greater than 0")
}
