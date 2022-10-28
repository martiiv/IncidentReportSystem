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
				"name": "TestGroupAPITEST",
				"info": "Group for testing the api using apitestf"
			}`).
			Expect(t).
			Status(http.StatusCreated).
			End()
	})
}

/*
func Test_updateReceiverGroup(t *testing.T) {
	! PUT Request not implemented for this endpoint
}
*/

func Test_deleteReceiverGroup(t *testing.T) {
	r := mux.NewRouter()
	r.Path("/groups").HandlerFunc(endpoints.HandleReceivingGroup)
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("Deleting test group incident with name TestGroupAPITEST", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Delete("/groups").
			Body(`[{
		"id": "",
		"name" : "TestGroupAPITEST"
		}]`).
			Expect(t).
			Status(http.StatusOK).
			End()
	})
}
