package entities

type User struct {
	ID          int    `json:"id"`
	Username    string `json:"username"`
	Password    string `json:"password"`
	Email       string `json:"email"`
	Phonenumber string `json:"phonenumber"`
	Firstname   string `json:"firstname"`
	Lastname    string `json:"lastname"`
	Role        string `json:"role"`
}
