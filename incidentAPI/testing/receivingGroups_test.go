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

func Test_getAllReceiverGroups(t *testing.T) {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/groups").HandlerFunc(endpoints.HandleReceivingGroup)
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("Getting all warning receiver", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("/groups").
			Expect(t).
			Status(http.StatusOK).
			End()
	})
}

func Test_getOneReceiverGroup(t *testing.T) {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/groups").HandlerFunc(endpoints.HandleReceivingGroup).Queries("id", "{id}")
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("Getting receiver group with id 1", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("/groups").
			Query("id", "1").
			Expect(t).
			Body(`{"id": 1, "name": "Human Resources", "info": "Group for everyone working in the HR departments in the company"}`).
			Status(http.StatusOK).
			End()
	})
}

func Test_createReceiverGroup(t *testing.T) {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/groups").HandlerFunc(endpoints.HandleReceivingGroup)
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("Creating new test receiver group", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Post("/groups").
			Body(`{
				"name": "Test Group",
				"info": "Group for testing the api"
			}`).
			Expect(t).
			Status(http.StatusCreated).
			End()
	})
}

func Test_updateReceiverGroup(t *testing.T) { //! Not working cuz not implemented
	databasefunctions.EstablishConnection() //? Not implemented
	r := mux.NewRouter()
	r.Path("/groups").HandlerFunc(endpoints.HandleIncidentRequest)
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
			Handler(r).Put("/incident").
			Body(`{
				"incidentId": 1,
				"countermeasure" : "Contact janitor, Fix door"
	}`).
			Expect(t).
			Status(http.StatusOK).
			End()
	})

}

func Test_deleteReceiverGroup(t *testing.T) {
	//? Not implemented
}
