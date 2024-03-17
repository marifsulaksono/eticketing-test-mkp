package entities

type Stage struct {
	ID       int    `json:"id"`
	BranchID int    `json:"branch_id"`
	Name     string `json:"name"`
	Capacity int    `json:"capacity"`
}
