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

func Test_getAllSystemManager(t *testing.T) {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/manager").HandlerFunc(endpoints.HandleSystemManagerRequest)
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("Getting all managers", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("/manager").
			Expect(t).
			Status(http.StatusOK).
			End()
	})
}

func Test_getOneSystemManager(t *testing.T) {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/manager").HandlerFunc(endpoints.HandleSystemManagerRequest).Queries("id", "{id}")
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("Getting manager with id 1", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("/manager").
			Query("id", "1").
			Expect(t).
			Body(`{"id": 1, "userName": "OdaManager","company": "IncidentCorp","credential": "1"}`).
			Status(http.StatusOK).
			End()
	})
}

func Test_createSystemManager(t *testing.T) {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/manager").HandlerFunc(endpoints.HandleSystemManagerRequest)
	ts := httptest.NewServer(r)

	defer ts.Close()
	t.Run("Creating new test manager", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Post("/manager").
			Body(`{"userName": "TestManagerAPITEST","company":"IncidentCorp","email":"testManager@gmail.com","password": "1241erreth23e23r1231"}`).
			Expect(t).
			Status(http.StatusCreated).
			End()
	})

}

/*
func Test_updateSystemManager(t *testing.T) {
	? not implemented yet
}
*/

func Test_deleteSystemManager(t *testing.T) {
	r := mux.NewRouter()
	r.Path("/manager").HandlerFunc(endpoints.HandleSystemManagerRequest)
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("Deleting system manager with name TestManagerAPITEST", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Delete("/manager").
			Body(`[{
		"id": "",
		"name" : "TestManagerAPITEST"
		}]`).
			Expect(t).
			Status(http.StatusOK).
			End()
	})
}
