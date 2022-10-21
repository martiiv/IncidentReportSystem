package main

import (
	"fmt"
	"incidentAPI/communication"
	databasefunctions "incidentAPI/databaseFunctions"
	"incidentAPI/endpoints"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	databasefunctions.EstablishConnection()
	//Group endpoint
	//Todo add function to endpoints
	r.Path("/groups").Queries("id", "{id}") //PUT
	r.Path("/groups")                       //GET, POST

	//Log endpoint
	//Todo add function to endpoints
	r.Path("/incident").HandlerFunc(endpoints.HandleIncidentRequest).Queries("id", "{id}") //GET PUT
	r.Path("/incident").HandlerFunc(endpoints.HandleIncidentRequest)

	// Send email
	//r.Path("/incident/sendMail/").HandlerFunc(communication.SendMail)
	r.HandleFunc("/incident/sendMail/", communication.SendMail)

	//System Manager endpoint
	//Todo add function to endpoints
	r.Path("/manager").HandlerFunc(endpoints.HandleSystemManagerRequest).Queries("id", "{id}") //PUT
	r.Path("/manager").HandlerFunc(endpoints.HandleSystemManagerRequest)                       //GET, POST, DELETE

	//Warning Receiver endpoint
	//Todo add function to endpoints
	r.Path("/receiver").Queries("id", "{id}")                               //PUT
	r.Path("/receiver").HandlerFunc(endpoints.HandleRequestWarningReceiver) //POST, DELETE
	http.Handle("/", r)

	fmt.Print("Listening on port:" + getPort())

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
