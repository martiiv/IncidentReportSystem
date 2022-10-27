package structs

// Method for creating system managers /manager Method POST
type CreateSystemManager struct {
	UserName string `json:"userName"`
	Company  string `json:"company"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

// Method for getting a system manager /manager?id=92 Method GET
type GetSystemManager struct {
	Id         int    `json:"id"`
	UserName   string `json:"userName"`
	Company    string `json:"company"`
	Credential string `json:"credential"`
}

// Method for deleting system managers /manager
type DeleteSystemManager []struct {
	Id int `json:"id"`
}
