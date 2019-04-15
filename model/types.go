package model

type User struct {
	ID        string `json:"id"`
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
	Password  string `json:password`
	JobTitle  string `json:"job"`
	Email     string `json:"email"`
	Image     string `json:"image"`
}

type Film struct {
	ID          int
	Title       string
	Description string
	Year        int
	Rate        float32
	Length      int
}

type Customer struct {
	ID        int
	Firstname string
	Lastname  string
	Email     string
}

type Credential struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
