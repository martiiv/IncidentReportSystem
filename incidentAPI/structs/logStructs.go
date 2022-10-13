package structs

// Struct method for getting incidents Endpoint /incident Method GET
type GetAllIncidents struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Context        string `json:"context"`
	Information    string `json:"information"`
	ReceivingGroup string `json:"recieveingGroup"`
	Countermeasure string `json:"countermeasure"`
	SystemManager  int    `json:"systemManager"`
}

// Struct method for creating incidents endpoint: /incident Method POST
type CreateIncident struct {
	Id             int    `json:"id"`
	Name           string `json:"name"`
	Context        string `json:"context"`
	Information    string `json:"information"`
	ReceivingGroup string `json:"receivingGroup"`
	Countermeasure string `json:"countermeasure"`
	SystemManager  int    `json:"systemManager"`
}

type SendIndividualIncident struct {
	Name           string `json:"name"`
	Context        string `json:"context"`
	Information    string `json:"information"`
	Receiver       string `json:"receiver"`
	Countermeasure string `json:"countermeasure"`
}

// Struct method for updating countermeasures for an incident endpoint /incident?id=EX Method PUT
type UpdateCountermeasure struct {
	Countermeasure string `json:"countermeasure"`
}
