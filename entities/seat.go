package entities

type Seat struct {
	ID      int `json:"id"`
	StageID int `json:"stage_id"`
	Number  int `json:"number"`
}
