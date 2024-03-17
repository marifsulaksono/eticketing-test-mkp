package entities

import "time"

type Showtime struct {
	ID        int       `json:"id"`
	BranchID  int       `json:"branch_id"`
	StageID   int       `json:"stage_id"`
	MovieID   int       `json:"movie_id"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

type ShowtimeResponse struct {
	ID        int       `json:"id"`
	BranchID  int       `json:"-"`
	Branch    Branch    `json:"branch"`
	StageID   int       `json:"-"`
	Stage     Stage     `json:"stage"`
	MovieID   int       `json:"-"`
	Movie     Movie     `json:"movie"`
	StartTime time.Time `json:"start_time"`
	EndTime   time.Time `json:"end_time"`
}

func (ShowtimeResponse) TableName() string {
	return "showtimes"
}
