package structs

type PDCountermeasure struct {
	Tag         string `json:"tag"`
	Description string `json:"description"`
}

type DeleteCountermeasure struct {
	Tag     string `json:"tag"`
	Cascade string `json:"cascade"`
}
