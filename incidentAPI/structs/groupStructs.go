package structs

// Add group struct, endpoint: /groups POST
type AddGroup struct {
	Name string `json:"name"`
	Info string `json:"info"`
}

// Fetch a group, enpoint: /groups GET
type GetGroup struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info string `json:"info"`
}

type DeleteGroup []struct {
	Id int `json:"id"`
}
