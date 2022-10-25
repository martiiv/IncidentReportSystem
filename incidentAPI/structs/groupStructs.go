package structs

// Add group struct, endpoint: /groups POST
type CreateReceivingGroup struct {
	Name string `json:"name"`
	Info string `json:"info"`
}

// Fetch a group, enpoint: /groups GET
type GetReceivingGroups struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
	Info string `json:"info"`
}

type DeleteReceivingGroup []struct {
	Id int `json:"id"`
}
