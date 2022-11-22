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

/*
* File countermeasure_test.go
* Will test all functionality related to the predefined countermeasure and tags table in the database
? Last revision Martin Iversen 22.11.2022
*/

func Test_getAllCountermeasures(t *testing.T) {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/countermeasure").HandlerFunc(endpoints.HandlePDC)
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("Getting all countermeasures", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("/countermeasure").
			Expect(t).
			Status(http.StatusOK).
			End()
	})
}

func Test_getOneCountermeasure(t *testing.T) {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/countermeasure").HandlerFunc(endpoints.HandlePDC).Queries("tag", "{tag}")
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("Getting countermeasure with tag Technical Problem ", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Get("/countermeasure").
			Query("tag", "Technical Problem ").
			Expect(t).
			Body(`{"tag":"Technical Problem ","description":"Check that the network is up before doing anything. If the problem is still up technical support will be dispatched."}`).
			Status(http.StatusOK).
			End()
	})
}

func Test_createCountermeasure(t *testing.T) {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/countermeasure").HandlerFunc(endpoints.HandlePDC)
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("Creating new countermeasure with tag Slippery", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Post("/countermeasure").
			Body(`{ 
			"tag":"Slippery",
			"description":"Wait for the floor to dry, Place a warning sign, Notify Janitor"
		}`).
			Expect(t).
			Status(http.StatusCreated).Body("New countermeasure added with tag: Slippery").
			End()

		_, _ = databasefunctions.Db.Exec("DELETE FROM `PredefinedCounterMeasures` ORDER BY `COid` desc limit 1;")
		_, _ = databasefunctions.Db.Exec("DELETE FROM `Tags` ORDER BY `Tid` desc limit 1;")

	})
}

func Test_updateCountermeasure(t *testing.T) {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()
	r.Path("/countermeasure").HandlerFunc(endpoints.HandlePDC)
	ts := httptest.NewServer(r)
	defer ts.Close()
	t.Run("Updating countermeasure with tag DDOS", func(t *testing.T) {
		apitest.New().
			Handler(r).Put("/countermeasure").
			Body(`{
		"tag": "DDOS",
		"description" : "Updating description in DDOS"
		}`).
			Expect(t).
			Status(http.StatusAccepted).Body("Succsessfully updated countermeasure with tag DDOS").
			End()
	})

	t.Run("Reverting the countermeasure", func(t *testing.T) {
		apitest.New().
			Handler(r).
			Put("/countermeasure").
			Body(`{
				"tag": "DDOS",
				"description" : "Try to block the IPs for a short period of time and after that check the logs to ensure where the attack was coming from."
	}`).
			Expect(t).
			Status(http.StatusAccepted).Body("Succsessfully updated countermeasure with tag DDOS").
			End()
	})

}
