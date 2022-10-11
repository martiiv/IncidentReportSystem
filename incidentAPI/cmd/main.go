package main

import (
	"fmt"
	databasefunctions "incidentAPI/databaseFunctions"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	databasefunctions.EstablishConnection()
	r := mux.NewRouter()

	//Group endpoint
	//Todo add function to endpoints
	r.Path("/groups").Queries("id", "{id}") //PUT
	r.Path("/groups")                       //GET, POST

	//Log endpoint
	//Todo add function to endpoints
	r.Path("/incident").Queries("id", "{id}") //PUT
	r.Path("/incident")                       //GET, POST

	//System Manager endpoint
	//Todo add function to endpoints
	r.Path("/manager").Queries("id", "{id}") //PUT
	r.Path("/manager")                       //GET, POST, DELETE

	//Warning Receiver endpoint
	//Todo add function to endpoints
	r.Path("/receiver").Queries("id", "{id}") //PUT
	r.Path("/receiver")                       //POST, DELETE

	http.Handle("/", r)
	err := http.ListenAndServe(getPort(), nil)
	if err != nil {
		fmt.Printf(err.Error())
	}
}

/*
Function used for setting the port for the application
We use localhost 8080 for testing
Takes no parameters
Returns the port the software is listening on
*/
func getPort() string {
	var port = os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	return ":" + port
}
