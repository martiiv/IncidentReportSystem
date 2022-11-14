package testing

import (
	databasefunctions "incidentAPI/databaseFunctions"
	"incidentAPI/endpoints"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/steinfletcher/apitest"
)

func Test_getAllIncidents(t *testing.T) {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/incident").HandlerFunc(endpoints.HandleIncidentRequest)
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("Getting all incidents", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("/incident").
			Expect(t).
			Status(http.StatusOK).
			End()
	})
}

func Test_getOneIncident(t *testing.T) {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/incident").HandlerFunc(endpoints.HandleIncidentRequest).Queries("id", "{id}")
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("Getting incident with id 2", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("/incident").
			Query("id", "2").
			Expect(t).
			Body(`{"id":2,"tag":"Phishing","name":"Hack attack!","description":"An email from an unknown party has sent out a malicious email containing malware!","company":"IncidentCorp","receivingGroup":"2","countermeasure":"Do not open email, Block sender ","sendbymanager":"OdaManager","date":"2022-10-18 11:49:55"}`).
			Status(http.StatusOK).
			End()
	})
}

func Test_createIncident(t *testing.T) {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/incident").HandlerFunc(endpoints.HandleIncidentRequest)
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("Creating new test incident", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Post("/incident").
			Body(`{ 
			"tag": "Test",
			"name":  "TestIncidentAPITEST",
			"description": "I am testing an incident",
			"company": "IncidentCorp",
			"receivingGroup": "Marketing",
			"countermeasure": "Send help",
			"sendByManager": "OdaManager"
		}`).
			Expect(t).
			Status(http.StatusCreated).Body("New incident added with name: TestIncidentAPITEST").
			End()

		stmt, _ := databasefunctions.Db.Prepare("DELETE FROM `Incident` ORDER BY `IncidentId` desc limit 1;")
		_, _ = stmt.Exec()
	})
}

func Test_updateCountermeasures(t *testing.T) {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/incident").HandlerFunc(endpoints.HandleIncidentRequest)
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("Updating incident with id 1", func(t *testing.T) {
		apitest.New().
			Handler(r).Put("/incident").
			Body(`{
		"incidentId": 1,
		"countermeasure" : "Updating countermeasures ins TestIncident"
		}`).
			Expect(t).
			Status(http.StatusOK).
			End()
	})

	t.Run("Reverting the countermeasures", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Put("/incident").
			Body(`{
				"incidentId": 1,
				"countermeasure" : "Contact janitor, Fix door"
	}`).
			Expect(t).
			Status(http.StatusOK).
			End()
	})

}

func Test_deleteIncident(t *testing.T) {
	r := mux.NewRouter()
	r.Path("/incident").HandlerFunc(endpoints.HandleIncidentRequest)
	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("Deleting testIncident incident with name Broken door", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Delete("/incident").
			Body(`[{
		"incidentId": "45",
		"incidentName" : "Broken door"
		}]`).
			Expect(t).
			Status(http.StatusOK).
			End()
	})

	t.Run("Recreating the incident we deleted", func(t *testing.T) {
		apitest.New().Handler(r).
			Post("/incident").
			Body(`{ 
			"tag": "Hazard",
			"name":  "Broken door",
			"description": "A door has been broken in office five",
			"company": "IncidentCorp",
			"receivingGroup": "Human Resources",
			"countermeasure": "Contact janitor, Fix door",
			"sendByManager": "OdaManager"
		}`).
			Expect(t).
			Status(http.StatusCreated).Body("New incident added with name: Broken door").
			End()
	})
}
