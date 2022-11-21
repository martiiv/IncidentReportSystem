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
	t.Run("Getting incident with id 110", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("/incident").
			Query("id", "110").
			Expect(t).
			Body(`{"id":110,"tag":"Malware","name":"Virus in the print Server","description":"Lorem ipsum dolor sit amet, consectetur adipiscing elit.","company":"NTNU","receivingGroup":"Information Security","countermeasure":"Disconect the affected devices from the network and start checking critical devices such as servers and Administration's machines.","sendbymanager":"aleksaab","date":"2022-11-15 16:06:10","lessonlearned":"test"}`).
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
			"tag": "Phishing",
			"name":  "TestIncidentAPITEST",
			"description": "I am testing an incident",
			"company": "IncidentCorp",
			"receivingGroup": "Development",
			"sendByManager": "OdaManager",
			"lessonlearned": " "
		}`).
			Expect(t).
			Status(http.StatusCreated).Body("New incident added with name: TestIncidentAPITESTMail sent").
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
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/incident").HandlerFunc(endpoints.HandleIncidentRequest)
	ts := httptest.NewServer(r)
	defer ts.Close()

	t.Run("Deleting testIncident incident with name Broken door", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Delete("/incident").
			Body(`[{
		"incidentId": "",
		"incidentName" : "Room 3 is out of the network"
		}]`).
			Expect(t).
			Status(http.StatusOK).
			End()
	})

	t.Run("Recreating the incident we deleted", func(t *testing.T) {
		apitest.New().Handler(r).
			Post("/incident").
			Body(`{ 
			"tag": "Technical Problem ",
			"name":  "Room 3 is out of the network",
			"description": "All of the PCs in room 3 are out of the network.",
			"company": "NTNU",
			"receivingGroup": "Development",
			"sendByManager": "OdaManager",
			"lessonLearned": "IDK lord have mercy"
		}`).
			Expect(t).
			Status(http.StatusCreated).Body("New incident added with name: Room 3 is out of the networkMail sent").
			End()
	})
}
