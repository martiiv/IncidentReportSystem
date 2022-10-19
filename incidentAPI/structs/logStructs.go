package structs

// Struct method for getting incidents Endpoint /incident Method GET
type GetIncident struct {
	IncidentId     *int    `json:"id"`
	Tag            *string `json:"tag"`
	Name           *string `json:"name"`
	Description    *string `json:"description"`
	Company        *string `json:"company"`
	Credential     *string `json:"credential"`
	ReceivingGroup *string `json:"receivingGroup"`
	Countermeasure *string `json:"countermeasure"`
	Sendbymanager  *string `json:"sendbymanager"`
	Date           *string `json:"date"`
}

// Struct method for creating incidents endpoint: /incident Method POST
type CreateIncident struct {
	Tag            string  `json:"tag"`
	Name           string  `json:"name"`
	Description    string  `json:"description"`
	Company        string  `json:"company"`
	Credential     *string `json:"credential"`
	ReceivingGroup string  `json:"receivingGroup"`
	Countermeasure string  `json:"countermeasure"`
	Sendbymanager  string  `json:"sendbymanager"`
	Date           string  `json:"date"`
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
	IncidentId     int    `json:"incidentId"`
	Countermeasure string `json:"countermeasure"`
}
