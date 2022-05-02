package data

type Setup struct {
	MesocycleNum     int `json:"mesocycleNum"`
	CurrentMesocycle int `json:"currentMesocycle"`
	WeekNum          int `json:"weekNum"`
	CurrentWeek      int `json:"currentWeek"`
	DayNum           int `json:"dayNum"`
	WorkoutNum       int `json:"workoutNum"`
}

func (w WorkoutModel) InsertSetup(setup *Setup) error {
	queryMesocycle := `insert into mesocycle(mesocyclenum) select generate_series(1,$1) returning id`

	//queryWeek := `insert into week(weeknum,mesocycle_id) select generate_series(1,$1), ($2) returning id`

	//queryDay := `insert into day(daynum,week_id) select generate_series(1,$1), ($2) returning id`

	args := []interface{}{setup.MesocycleNum}
	//argsWeek := []interface{}{setup.WeekNum, setup.CurrentMesocycle}

	//w.DB.QueryRow(queryWeek, argsWeek...).Scan(&setup.WeekNum)
	for i := 0; i < setup.CurrentMesocycle; i++ {
		queryWeek := `insert into week(weeknum,mesocycle_id) select generate_series(1,$1), ($2) returning id`
		argsWeek := []interface{}{setup.WeekNum, setup.CurrentMesocycle}
		w.DB.QueryRow(queryWeek, argsWeek...).Scan(&setup.WeekNum)
	}
	return w.DB.QueryRow(queryMesocycle, args...).Scan(&setup.MesocycleNum)

}
