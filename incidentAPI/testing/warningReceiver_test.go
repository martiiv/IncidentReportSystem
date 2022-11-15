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

func Test_getAllReceivers(t *testing.T) {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/receiver").HandlerFunc(endpoints.HandleRequestWarningReceiver)
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("Getting all warning receiver", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("/receiver").
			Expect(t).
			Status(http.StatusOK).
			End()
	})
}

func Test_getOneReceiver(t *testing.T) {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/receiver").HandlerFunc(endpoints.HandleRequestWarningReceiver).Queries("id", "{id}")
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("Getting warning receiver with id 2", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("/receiver").
			Query("id", "2").
			Expect(t).
			Body(`{"id":2,"name":"Ulrik","phoneNumber":"78590153","company":"IncidentCorp","receiverGroup":"Development","receiverEmail":"UlrikUtvikler@gmail.com"}`).
			Status(http.StatusOK).
			End()
	})
}

func Test_createReceiver(t *testing.T) {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/receiver").HandlerFunc(endpoints.HandleRequestWarningReceiver)
	ts := httptest.NewServer(r)

	defer ts.Close()
	t.Run("Creating new test warning receiver", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Post("/receiver").
			Body(`{
				"name":"TestReceiverAPITEST",
				"phoneNumber":"12345678",
				"company":"IncidentCorp",
				"receiverGroup":"Marketing",
				"receiverEmail":"APITEST@gmail.com"
			}`).
			Expect(t).
			Status(http.StatusCreated).
			End()
	})

}

func Test_deleteReceiver(t *testing.T) {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/receiver").HandlerFunc(endpoints.HandleRequestWarningReceiver)
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("Deleting the test warning receiver", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Delete("/receiver").
			Body(`[{
				"id": "",
				"email":"APITEST@gmail.com"
			}]`).
			Expect(t).
			Status(http.StatusOK).
			End()
	})
}
