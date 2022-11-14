package main

import (
	"fmt"
	databasefunctions "incidentAPI/databaseFunctions"
	"incidentAPI/endpoints"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	databasefunctions.EstablishConnection()

	//Login
	r.Path("/login").HandlerFunc(endpoints.LoginSystemManager)

	//Group endpoint
	r.Path("/groups").HandlerFunc(endpoints.HandleReceivingGroup).Queries("id", "{id}") //PUT
	r.Path("/groups").HandlerFunc(endpoints.HandleReceivingGroup)                       //GET, POST

	//Log endpoint
	r.Path("/incident").HandlerFunc(endpoints.HandleIncidentRequest).Queries("id", "{id}")   //GET PUT
	r.Path("/incident").HandlerFunc(endpoints.HandleIncidentRequest).Queries("tag", "{tag}") //GET PUT
	r.Path("/incident").HandlerFunc(endpoints.HandleIncidentRequest)

	//System Manager endpoint
	r.Path("/manager").HandlerFunc(endpoints.HandleSystemManagerRequest).Queries("id", "{id}") //PUT
	r.Path("/manager").HandlerFunc(endpoints.HandleSystemManagerRequest)                       //GET, POST, DELETE

	//Warning Receiver endpoint
	r.Path("/receiver").HandlerFunc(endpoints.HandleRequestWarningReceiver).Queries("id", "{id}") //PUT
	r.Path("/receiver").HandlerFunc(endpoints.HandleRequestWarningReceiver)                       //POST, DELETE
	http.Handle("/", r)

	fmt.Print("Listening on port:" + getPort())

	err := http.ListenAndServe(getPort(), nil)
	if err != nil {
		fmt.Printf("%v", err.Error())
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
