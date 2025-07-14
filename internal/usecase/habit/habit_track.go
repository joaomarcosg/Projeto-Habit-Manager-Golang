package habit

import "time"

type TrackHabitReq struct {
	HabitID    int64     `json:"habit_id"`
	StartDate  time.Time `json:"start_date"`
	TargetDate time.Time `json:"target_date"`
}
