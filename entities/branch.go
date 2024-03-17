package entities

type Branch struct {
	ID          int    `json:"id"`
	Name        string `json:"name"`
	Address     string `json:"address"`
	Phonenumber string `json:"phonenumber"`
}
