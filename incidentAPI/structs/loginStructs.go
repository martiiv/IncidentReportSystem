package structs

type LoginStructs struct {
	Email    string `json:"userName"`
	Password string `json:"password"`
}

type LoggedIn struct {
	Email    string `json:"email"`
	UserName string `json:"userName"`
	CiD      string `json:"cid"`
	Company  string `json:"company"`
}
